package cmd

import (
	"stock/config"
	"stock/database"
	"stock/mongodb"

	"github.com/urfave/cli/v2"
)

func bootstrap(c *cli.Context) error {
	// init config
	if err := config.Load(c.String("config")); err != nil {
		return err
	}

	// init database
	if err := database.Init(); err != nil {
		return err
	}

	// init mongodb
	if err := mongodb.NewMongoDB(); err != nil {
		return err
	}

	return nil
}
