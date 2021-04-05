package cmd

import (
	"stock/server"

	"github.com/urfave/cli/v2"
)

// Start
func Start() *cli.Command {
	return &cli.Command{
		Name:        "start",
		Usage:       "Start",
		Description: "Start",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "./config.yml",
				Usage:   "Load configuration from `FILE`",
			},
		},
		Action: func(c *cli.Context) error {
			// database bootstrap
			if err := bootstrap(c); err != nil {
				return err
			}
			// init server
			if err := server.Init(); err != nil {
				return err
			}

			srv := server.NewServer()
			return srv.Start()

			// // init MQTT
			// go server.MQTTInit()

			// srv := server.NewServer()
			// if config.Server.HTTPS {
			// 	return srv.StartTLS()
			// }

			// return srv.Start()
		},
	}
}
