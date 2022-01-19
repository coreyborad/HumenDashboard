package cmd

import (
	"stock/concrete"
	"stock/services"

	"github.com/urfave/cli/v2"
)

func Manually() *cli.Command {
	return &cli.Command{
		Name:        "manually",
		Usage:       "Manually",
		Description: "Manually",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "./config.yml",
				Usage:   "Load configuration from `FILE`",
			},
		},
		Subcommands: []*cli.Command{
			{
				Name:  "stock",
				Usage: "Run the manually commands",
				Subcommands: []*cli.Command{
					{
						Name:  "parser",
						Usage: "Run the manually commands",
						Action: func(c *cli.Context) error {
							if err := bootstrap(c); err != nil {
								return err
							}
							stockServ := services.CreateStockService()
							stockServ.DailyParser()
							return nil
						},
					},
					{
						Name:  "calc",
						Usage: "Run the manually commands",
						Action: func(c *cli.Context) error {
							if err := bootstrap(c); err != nil {
								return err
							}
							stockServ := services.CreateStockService()
							// date := time.Date(2021, time.September, 9, 0, 0, 0, 0, time.UTC)
							// stockServ.Calc([]string{"2330"}, &date)
							stockServ.Calc("2330", nil)
							return nil
						},
					},
					{
						Name:  "history",
						Usage: "Run the manually commands",
						Action: func(c *cli.Context) error {
							if err := bootstrap(c); err != nil {
								return err
							}
							stockConcrete := concrete.CreateStockConcrete()
							stockConcrete.CheckHistory()
							return nil
						},
					},
				},
			},
		},
	}
}
