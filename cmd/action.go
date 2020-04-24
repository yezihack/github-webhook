package cmd

import (
	"github.com/urfave/cli/v2"
	"github.com/yezihack/github-webhook/config"
	"github.com/yezihack/github-webhook/router"
)

// action
var actionHandle = func(c *cli.Context) error {
	if c.NumFlags() == 0 {
		return nil
	}
	cfg, err := config.New(scriptBash, secret, port, quiet, verbose)
	if err != nil {
		return err
	}
	// Start http server
	return router.New(cfg)
}
