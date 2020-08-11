package main

import (
	"log"
	"os"

	"github.com/npm622/qo/internal/commands"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "bqo",
		HelpName: "go run cmd/cli/main.go",
		Usage:    "a baseball qualifying offer tool",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    commands.ArgURL,
				Aliases: []string{"u"},
				Value:   "https://questionnaire-148920.appspot.com/swe",
				Usage:   "fetch data from here",
			},
			&cli.BoolFlag{
				Name:    commands.ArgFerbose,
				Aliases: []string{"f"},
				Usage:   "include only the numbers (fake-verbose)",
			},
			&cli.BoolFlag{
				Name:    commands.ArgVerbose,
				Aliases: []string{"v"},
				Usage:   "include all the data",
			},
		},
		Action: commands.Action,
		Commands: []*cli.Command{
			&cli.Command{
				Name:    "monte-carlo",
				Aliases: []string{"mc"},
				Usage:   "calculate the qualifying offer many times",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    commands.ArgMonteCarloCount,
						Aliases: []string{"n"},
						Value:   10,
						Usage:   "Number of iterations to calculate qualifying offer",
					},
				},
				Action: commands.MonteCarlo,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
