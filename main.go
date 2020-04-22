package main

import (
	"go-webhook/cmd"
	"log"
)

// go-webhook main
func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}
