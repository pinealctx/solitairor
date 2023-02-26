package main

import (
	"fmt"
	"github.com/pinealctx/solitairor/pkg/sol"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	var app = cli.App{
		Name: "parse cards",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "cards",
				Aliases: []string{
					"cs",
				},
				Usage: "cards string",
			},
		},
		Action: parseCardsAction,
	}
	var err = app.Run(os.Args)
	if err != nil {
		fmt.Println("run command error:", err)
	}
}

func parseCardsAction(c *cli.Context) error {
	var ss = c.String("cards")
	fmt.Println("input cards:", ss)
	var ccs, err = sol.ParseCards(ss)
	if err != nil {
		return err
	}
	fmt.Println("output cards:", ccs)
	fmt.Println("output cards:", ccs.ToBytes())
	var cls = sol.ConvertLegacyCards(ccs)
	fmt.Println("client cards:", cls)
	fmt.Println("client cards:", cls.ToBytes())
	return nil
}
