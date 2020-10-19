package main

import (
	"github.com/yezihack/github-webhook/cmd"
	"log"
)

// go-webhook main
func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	cmd.Execute()
}
