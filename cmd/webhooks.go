package cmd

import (
	"github.com/urfave/cli/v2"
	"go-webhook/config"
	"os"
)

var (
	scriptBash string // The script bash
	port       int    // The port
	secret     string // The secret
	quiet      bool   // only print info
)

func init() {
	rootCmd.Flags = []cli.Flag{
		//&cli.StringFlag{Name: "path", Destination:&path, Aliases: []string{"u"}, Value: "hook", Usage: "url path, eg: protocal://hostname[:port]/path/"},
		&cli.StringFlag{Name: "bash", Destination: &scriptBash, Aliases: []string{"b"}, Value: "", Usage: "Execute the script path. eg: /home/hook.sh"},
		&cli.IntFlag{Name: "port", Destination: &port, Aliases: []string{"p"}, Value: 6666, Usage: "http port"},
		&cli.StringFlag{Name: "secret", Destination: &secret, Aliases: []string{"s"}, Value: "", Usage: "github hook secret"},
		&cli.BoolFlag{Name: "quiet", Destination: &quiet, Aliases: []string{"q"}, Value: false, Usage: "quiet operation"},
	}
	// This is processing logic
	rootCmd.Action = action
}

// The start app config
var rootCmd = cli.App{
	Name:      config.Name,
	Usage:     config.Usage,
	Version:   config.Version,
	Authors:   authors,
	Copyright: config.Copyright,
}

// Start execute
func Execute() error {
	return rootCmd.Run(os.Args)
}
