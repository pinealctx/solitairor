package main

import (
	"github.com/pinealctx/solitairor/pkg/sol"
	"github.com/pinealctx/solitairor/pkg/store"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"log"
	"os"
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
			&cli.BoolFlag{
				Name:  "debugDB",
				Usage: "log database sql",
			},
			&cli.IntFlag{
				Name:  "maxStackSize",
				Usage: "max stack size",
				Value: 100000,
			},
			&cli.IntFlag{
				Name:  "maxSearchSize",
				Usage: "max search size",
				Value: 1000000,
			},
			&cli.IntFlag{
				Name:  "runCount",
				Usage: "run count",
				Value: 100000,
			},
		},
		Action: runCmd,
	}
	var err = app.Run(os.Args)
	if err != nil {
		log.Println("run command error:", err)
	} else {
		log.Println("run command ok")
	}
}

func runCmd(c *cli.Context) error {
	var db = store.NewDB(
		store.User(c.String("user")),
		store.Password(c.String("password")),
		store.Host(c.String("host")),
		store.Schema(c.String("schema")),
		store.DebugMode(c.Bool("debugDB")),
	)
	for i := 0; i < c.Int("runCount"); i++ {
		generateSolitaire(
			c.Int("maxStackSize"),
			c.Int("maxSearchSize"),
			db,
			c.String("tableName"),
		)
	}
	return nil
}

func generateSolitaire(maxStackSize int, maxSearchSize int, db *gorm.DB, tableName string) {
	var cards = sol.GenQRandCards()
	var p = sol.NewPuzzle(maxStackSize, maxSearchSize)
	p.InitRoot(sol.NewGameState(cards))
	p.Run()
	if p.SolutionCount() > 0 {
		var r = &sol.Record{InitCards: cards.Txt()}
		p.Record(r)
		var err = store.InsertItem(db, tableName, r)
		if err != nil {
			log.Println("insert item error:", err)
		}
	}
}
