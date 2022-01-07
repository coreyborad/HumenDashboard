package main

import (
	"log"
	"os"
	"stock/cmd"
	"stock/config"

	cli "github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = config.App.Name
	app.Usage = config.App.Usage
	app.Version = config.App.Version
	app.EnableBashCompletion = true
	app.Commands = []*cli.Command{
		cmd.Start(),
		cmd.Schedule(),
		cmd.Manually(),
		// cmd.Migrate(),
		// cmd.Seed(),
		// cmd.Mqtt(),
		// cmd.Schedule(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
