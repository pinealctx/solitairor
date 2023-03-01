package main

import (
	"encoding/json"
	"fmt"
	"github.com/pinealctx/solitairor/pkg/sol"
	"github.com/pinealctx/solitairor/pkg/store"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"hash/crc32"
	"os"
	"strings"
)

var (
	db        *gorm.DB
	tableName string
)

type KlondikeObj struct {
	ID        int    `gorm:"column:id"`
	MinSearch int    `gorm:"column:min_search"`
	Step      int    `gorm:"column:step"`
	Road      int    `gorm:"column:road"`
	Cards     []byte `gorm:"column:cards"`
}

func main() {
	var app = cli.App{
		Name: "store tripeaks cards",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "user",
				Usage: "database user",
				Value: "root",
			},
			&cli.StringFlag{
				Name:  "password",
				Usage: "database password",
				Value: "12345678",
			},
			&cli.StringFlag{
				Name:  "host",
				Usage: "database host",
				Value: "127.0.0.1",
			},
			&cli.StringFlag{
				Name:  "schema",
				Usage: "database schema",
				Value: "testDB",
			},
			&cli.StringFlag{
				Name:  "tableName",
				Usage: "table name",
				Value: "klondike",
			},
			&cli.StringFlag{
				Name:  "dbLogLevel",
				Usage: "log database level",
				Value: "none",
			},
			&cli.StringFlag{
				Name:  "f",
				Usage: "file name",
			},
		},
		Action: runCmd,
	}
	var err = app.Run(os.Args)
	if err != nil {
		fmt.Println("run command error:", err)
	}
}

func runCmd(c *cli.Context) error {
	db = store.NewDB(
		store.User(c.String("user")),
		store.Password(c.String("password")),
		store.Host(c.String("host")),
		store.Schema(c.String("schema")),
		store.LogLevel(c.String("dbLogLevel")),
	)
	var cs, err = sol.ReadFileLine(c.String("f"))
	if err != nil {
		return err
	}
	var mm = make(map[string]*KlondikeObj)
	tableName = c.String("tableName")
	var ccs = make([]*sol.CardRecord, 0, len(cs))
	for _, jTxt := range cs {
		var s = strings.TrimSpace(jTxt)
		if s != "" {
			var rec = &sol.CardRecord{}
			err = json.Unmarshal([]byte(s), rec)
			if err != nil {
				return err
			}
			ccs = append(ccs, rec)
			updateKlondikeMap(mm, rec)
			if err != nil {
				return err
			}
		}
	}
	var tgm = make(map[string]struct{})
	for _, cc := range ccs {
		var fKey = string(cc.Cards)
		var _, ok = tgm[fKey]
		if ok {
			continue
		}
		var fetch = mm[fKey]
		err = store.InsertItem(db, tableName, fetch)
		if err != nil {
			return err
		}
		tgm[fKey] = struct{}{}
	}
	return nil
}

func updateKlondikeMap(mm map[string]*KlondikeObj, rec *sol.CardRecord) {
	var key = string(rec.Cards)
	var exist, ok = mm[key]
	if !ok {
		var cards = sol.MakeCardsFromBytes(rec.Cards)
		cards = faceDownAllCards(cards)
		cards[24].SetFaceUp()
		cards[31].SetFaceUp()
		cards[37].SetFaceUp()
		cards[42].SetFaceUp()
		cards[46].SetFaceUp()
		cards[49].SetFaceUp()
		cards[51].SetFaceUp()
		mm[key] = &KlondikeObj{
			ID:        int(crc32.ChecksumIEEE(rec.Cards)),
			MinSearch: rec.SearchCount,
			Step:      rec.Min,
			Road:      rec.Road,
			Cards:     cards.ToBytes(),
		}
		return
	}
	if rec.SearchCount < exist.MinSearch {
		exist.MinSearch = rec.SearchCount
	}
	if rec.Min < exist.Step {
		exist.Step = rec.Min
	}
	if rec.Road > exist.Road {
		exist.Road = rec.Road
	}
}

func faceDownAllCards(cards sol.Cards) sol.Cards {
	for i := 0; i < 52; i++ {
		cards[i].SetFaceDown()
	}
	return cards
}
