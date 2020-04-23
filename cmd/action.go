package cmd

import (
	"github.com/urfave/cli/v2"
	"github.com/yezihack/go-webhook/config"
	"github.com/yezihack/go-webhook/router"
)

// action
var actionHandle = func(c *cli.Context) error {
	if c.NumFlags() == 0 {
		return nil
	}
	cfg, err := config.New(scriptBash, secret, port, quiet)
	if err != nil {
		return err
	}
	// Start http server
	return router.New(cfg)
}
