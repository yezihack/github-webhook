package cmd

import (
	"github.com/urfave/cli/v2"
	"github.com/yezihack/github-webhook/config"
)

// authors info
var authors = []*cli.Author{
	{
		Name:  config.Author,
		Email: config.Email,
	},
}
