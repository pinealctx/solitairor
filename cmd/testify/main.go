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

	fmt.Println("start generate solitaire")
	for i := 0; i < runCount; i++ {
		if generateSolitaire() {
			validCount++
		}
		fmt.Print(".")
	}
	fmt.Println("")

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
	p.InitRoot(sol.NewGameState(cards))
	p.Run()
	if p.SolutionCount() > 0 {
		return true
	}
	return false
}
