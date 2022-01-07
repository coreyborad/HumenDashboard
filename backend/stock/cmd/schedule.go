package cmd

import (
	"stock/schedule"

	"github.com/urfave/cli/v2"
)

func Schedule() *cli.Command {
	return &cli.Command{
		Name:        "schedule",
		Usage:       "Schedule",
		Description: "Schedule",
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
				Name:  "run",
				Usage: "Run the scheduled commands",
				Action: func(c *cli.Context) error {
					if err := bootstrap(c); err != nil {
						return err
					}

					// init schedule
					if err := schedule.Init(); err != nil {
						return err
					}

					scheduler := schedule.GetScheduler()
					defer scheduler.Stop()

					select {}
				},
			},
		},
	}
}
