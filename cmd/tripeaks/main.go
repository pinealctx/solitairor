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

type TripeakF struct {
	Cards []string `json:"cards"`
	Steps int      `json:"steps"`
	Road  int      `json:"road"`
}

type TripeakObj struct {
	ID    int    `gorm:"column:id"`
	Road  int    `gorm:"column:road"`
	Step  int    `gorm:"column:step"`
	Cards []byte `gorm:"column:cards"`
}

func (t *TripeakF) ToRecord() *TripeakObj {
	var cards = sol.ConvertCardsFromStringList(t.Cards)
	cards = faceDownAllCards(cards)
	var buf = cards.ToBytes()
	return &TripeakObj{
		ID:    int(crc32.ChecksumIEEE(buf)),
		Road:  t.Road,
		Step:  t.Steps,
		Cards: buf,
	}
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
				Value: "tripeaks",
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
	tableName = c.String("tableName")
	for _, jTxt := range cs {
		var s = strings.TrimSpace(jTxt)
		if s != "" {
			var t = &TripeakF{}
			err = json.Unmarshal([]byte(s), t)
			if err != nil {
				return err
			}
			err = store.InsertItem(db, tableName, t.ToRecord())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func faceDownAllCards(cards sol.Cards) sol.Cards {
	for i := 0; i < 52; i++ {
		cards[i].SetFaceDown()
	}
	return cards
}
