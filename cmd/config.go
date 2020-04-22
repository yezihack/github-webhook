package cmd

import (
	"github.com/urfave/cli/v2"
	"go-webhook/config"
)

// authors info
var authors = []*cli.Author{
	{
		Name:  config.Author,
		Email: config.Email,
	},
}
