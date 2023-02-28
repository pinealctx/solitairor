package main

import (
	"fmt"
	"github.com/pinealctx/solitairor/pkg/store"
	"github.com/urfave/cli/v2"
	"gorm.io/gorm"
	"os"
	"sort"
)

var (
	db *gorm.DB
)

type TripeakObj struct {
	ID    int    `gorm:"column:id"`
	Road  int    `gorm:"column:road"`
	Step  int    `gorm:"column:step"`
	Cards []byte `gorm:"column:cards"`
}

func (x TripeakObj) ToGateObj(d int) *TripeakGate {
	return &TripeakGate{
		ID:         x.ID,
		Difficulty: d,
		Cards:      x.Cards,
		Road:       x.Road,
		Step:       x.Step,
	}
}

type TripeakGate struct {
	ID         int    `gorm:"column:id"`
	Difficulty int    `gorm:"column:difficulty"`
	Cards      []byte `gorm:"column:cards"`
	Road       int    `gorm:"column:road"`
	Step       int    `gorm:"column:step"`
}

type TripeakObjList []TripeakObj

func (x TripeakObjList) Len() int {
	return len(x)
}

func (x TripeakObjList) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}

func (x TripeakObjList) Less(i, j int) bool {
	if x[i].Road > x[j].Road {
		return true
	}
	if x[i].Road == x[j].Road {
		return x[i].Step < x[j].Step
	}
	return false
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
				Name:  "dbLogLevel",
				Usage: "log database level",
				Value: "warn",
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
	var cs TripeakObjList
	var err = db.Table("tripeaks").Find(&cs).Error
	if err != nil {
		return err
	}
	var count = len(cs)
	fmt.Println("get count:", count)
	sort.Sort(cs)

	var pieceLeft = count % 100
	var pieceCount = count / 100

	var cursor int
	var gate *TripeakGate
	for i := 0; i < 100; i++ {
		if i < pieceLeft {
			for j := 0; j < pieceCount+1; j++ {
				gate = cs[cursor].ToGateObj(i + 1)
				cursor++
				err = store.InsertItem(db, "tripeakgate", gate)
				if err != nil {
					return err
				}
			}
		} else {
			for j := 0; j < pieceCount; j++ {
				gate = cs[cursor].ToGateObj(i + 1)
				cursor++
				err = store.InsertItem(db, "tripeakgate", gate)
				if err != nil {
					return err
				}
			}
		}

	}
	return err
}
