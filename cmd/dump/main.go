package main

import (
	"encoding/json"
	"fmt"
	"github.com/pinealctx/solitairor/pkg/sol"
	"github.com/pinealctx/solitairor/pkg/store"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"os"
	"time"
)

var (
	db        *gorm.DB
	tableName string
)

func main() {
	var app = cli.App{
		Name: "generate solitaire to database",
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
				Value: "solitaire_1m",
			},
			&cli.StringFlag{
				Name:  "dbLogLevel",
				Usage: "log database level",
				Value: "none",
			},
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "debug print info",
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

	tableName = c.String("tableName")
	var t1 = time.Now()
	var rs, err = store.GetItemByAverageStepAsc(db, tableName)
	if err != nil {
		return err
	}
	for i := range rs {
		handleRecordItem(&rs[i])
	}
	var t2 = time.Now()
	var dur = t2.Sub(t1)
	var count = len(rs)

	if c.Bool("debug") {
		fmt.Println("count:", count, "duration:", dur, "average:", dur/time.Duration(count))
	}

	return nil
}

func handleRecordItem(item *sol.Record) {
	var ccs, err = sol.ParseCards(item.InitCards)
	if err != nil {
		return
	}
	var cls = sol.ConvertLegacyCards(ccs).ToBytes()
	var cr = sol.CardRecord{
		SearchCount: item.SearchCount,
		Step:        item.AverageStep,
		Min:         item.MinStep,
		Max:         item.MaxStep,
		Diff:        item.DiffStep,
		Road:        item.SolutionCount,
		Cards:       cls,
	}
	var buf []byte
	buf, err = json.Marshal(cr)
	if err != nil {
		return
	}
	fmt.Println(string(buf))
}
