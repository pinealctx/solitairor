package main

import (
	"fmt"
	"github.com/pinealctx/solitairor/pkg/sol"
	"github.com/urfave/cli/v2"
	"os"
	"time"
)

var (
	maxStackSize  int
	maxSearchSize int
)

func main() {
	var app = cli.App{
		Name: "generate solitaire to database",
		Flags: []cli.Flag{
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
				Value: 1000,
			},
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "debug solver",
			},
		},
		Action: runCmd,
	}
	var err = app.Run(os.Args)
	if err != nil {
		fmt.Println("run command error:", err)
	} else {
		fmt.Println("run command ok")
	}
}

func runCmd(c *cli.Context) error {
	var t1 = time.Now()
	maxStackSize = c.Int("maxStackSize")
	maxSearchSize = c.Int("maxSearchSize")
	var runCount = c.Int("runCount")
	var validCount = 0

	var debug = c.Bool("debug")
	fmt.Println("start generate solitaire")
	for i := 0; i < runCount; i++ {
		if generateSolitaire() {
			validCount++
		}
		if debug {
			fmt.Print(".")
		}
	}
	if debug {
		fmt.Println("")
	}

	var t2 = time.Now()
	var dur = t2.Sub(t1)
	fmt.Println("total time:", dur, "average time:", dur/time.Duration(runCount))
	fmt.Println(
		"total count:", runCount, "valid count:", validCount, "pass:", float64(validCount)/float64(runCount))
	return nil
}

func generateSolitaire() bool {
	var cards = sol.GenQRandCards()
	var p = sol.NewPuzzle(maxStackSize, maxSearchSize)
	p.InitRoot(sol.NewGameStateFromCards(cards))
	p.Run()
	return p.SolutionCount() > 0
}
