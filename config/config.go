package config

import (
	"errors"
	"github.com/howeyc/gopass"
	"github.com/yezihack/github-webhook/util"
	"os"
	"strconv"
	"strings"
)

var (
	Version     = "v1.4.3"
	Author      = "barry"
	Email       = "freeit@126.com"
	Name        = "github-webhook"
	Copyright   = "@2020 github.com/yezihack/github-webhook"
	Usage       = "This is a github web hooks tools"
	DefaultPort = 2020
)

type Config struct {
	ScriptBash string // script bash path
	Port       int
	Secret     string // secret
	Quiet      bool   // only print info, errors
	Verbose    bool   // print verbose
}

func New(scriptBash, secret string, port int, quiet, verbose bool) (cfg Config, err error) {
	cfg = Config{}
	if scriptBash == "" {
		err = errors.New("The script path is null, Use: --bash value")
		return
	}
	// Check that the file is valid
	if !util.IsFile(scriptBash) {
		err = errors.New("The script path not valid, script path:" + scriptBash)
		return
	}
	cfg.ScriptBash = scriptBash
	if port != 0 {
		if port > 65535 {
			err = errors.New("Invalid value for port, port:" + strconv.Itoa(port))
			return
		}
		cfg.Port = port
	}
	if strings.TrimSpace(secret) == "" {
		var (
			pass []byte
		)
		for {
			pass, err = gopass.GetPasswdPrompt("Enter secret:", true, os.Stdin, os.Stdout)
			if err != nil {
				return
			}
			if pass != nil {
				cfg.Secret = string(pass)
				break
			}
		}
	} else {
		cfg.Secret = secret
	}
	cfg.Quiet = quiet
	cfg.Verbose = verbose
	return
}
