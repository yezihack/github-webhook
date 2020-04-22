package config

import (
	"errors"
	"fmt"
	"github.com/unknwon/com"
	"strings"
)

var (
	Version   = "v1.0"
	Author    = "barry"
	Email     = "freeit@126.com"
	Name      = "go-github-webhook"
	Copyright = "@2020 github.com/yezihack"
	Usage     = "This is a github web hooks tools"
)

type Config struct {
	ScriptBash string // script bash path
	Port       int
	Secret     string // secret
	Quiet      bool   // only print info, errors
}

func New(scriptBash, secret string, port int, quiet bool) (cfg Config, err error) {
	cfg = Config{}
	if scriptBash == "" {
		err = errors.New("the script path is null")
		return
	}
	// Check that the file is valid
	if !com.IsFile(scriptBash) {
		err = fmt.Errorf("the script path not valid")
		return
	}
	cfg.ScriptBash = scriptBash
	if port != 0 {
		if port > 65535 {
			err = errors.New("invalid value for port")
			return
		}
		cfg.Port = port
	}
	if strings.TrimSpace(secret) == "" {
		err = errors.New("the github secret is empty")
		return
	}
	cfg.Secret = secret
	cfg.Quiet = quiet
	return
}
