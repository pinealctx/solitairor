package main

import (
	"encoding/json"
	"fmt"
	"github.com/pinealctx/solitairor/pkg/sol"
	"github.com/urfave/cli/v2"
	"os"
)

const (
	StackSize = M1
	Count     = K1

	K1     = 1000
	K1Dot2 = 1200
	K1Dot5 = 1500

	K5   = K1 * 5
	K10  = K1 * 10
	K50  = K1 * 50
	K100 = K1 * 100
	K500 = K1 * 500

	M1  = K1 * 1000
	M5  = M1 * 5
	M10 = M1 * 10
	M50 = M1 * 50
)

var (
	searchCountList = []int{
		500,
		700,
		900,

		K1,
		K1Dot2,
		K1Dot5,

		K5,
		K10,
		K50,
		K100,
		K500,

		M1,
		M5,
		M10,
		//M50,
	}
)

func main() {
	var app = cli.App{
		Name: "generate solitaire from ladder",
		Usage: "500 -- about 2% pass\n" +
			"700 -- about 8% pass \n" +
			"900 -- about 16% pass \n" +
			"1k -- about 22% pass\n" +
			"1.2k -- about 31% pass\n" +
			"1.5k -- about 43% pass\n" +
			"5k -- about 50% pass\n" +
			"10k -- about 53% pass\n" +
			"50k -- about 60% pass\n" +
			"100k -- about 63% pass\n" +
			"500k -- about 70% pass\n" +
			"1m -- about 70% pass\n" +
			"5m -- about 80% pass\n" +
			"10m -- about 81% pass\n", // +
		//"50m -- about 81% pass\n",
		Action: runCmd,
	}
	var err = app.Run(os.Args)
	if err != nil {
		fmt.Println("run command error:", err)
	}
}

type BCards struct {
	Cards []byte `json:"cards"`
}

func runCmd(c *cli.Context) error {
	var ccs = make([]sol.Cards, Count)
	for i := 0; i < Count; i++ {
		ccs[i] = sol.GenQRandCards()
	}

	//print init cards
	for i := 0; i < Count; i++ {
		printInitCards(ccs[i])
	}

	for i := range searchCountList {
		fmt.Println("--------", searchCountList[i], "----------")
		for j := 0; j < Count; j++ {
			handleCards(ccs[j], searchCountList[i])
		}
	}
	return nil
}

func printInitCards(cards sol.Cards) {
	var cr = BCards{Cards: cards.ToBytes()}
	var buf, err = json.Marshal(cr)
	if err != nil {
		return
	}
	fmt.Println(string(buf))
}

func handleCards(cards sol.Cards, maxSearchSize int) {
	var p = sol.NewPuzzle(StackSize, maxSearchSize)
	p.InitRoot(sol.NewGameStateFromCards(cards))
	p.Run()
	if p.SolutionCount() > 0 {
		var r = &sol.Record{}
		p.Record(r)
		var cr = sol.CardRecord{
			Step:  r.AverageStep,
			Min:   r.MinStep,
			Max:   r.MaxStep,
			Diff:  r.DiffStep,
			Road:  r.SolutionCount,
			Cards: cards.ToBytes(),
		}
		var buf, err = json.Marshal(cr)
		if err != nil {
			return
		}
		fmt.Println(string(buf))
	}
}
