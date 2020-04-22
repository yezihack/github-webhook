package main

import (
	"go-github-webhook/cmd"
	"log"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
