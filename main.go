package main

import (
	"fmt"
	"github.com/yezihack/go-webhook/cmd"
	"os"
)

// go-webhook main
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}
