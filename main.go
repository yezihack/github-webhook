package main

import (
	"fmt"
	"go-webhook/cmd"
	"os"
)

// go-webhook main
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
