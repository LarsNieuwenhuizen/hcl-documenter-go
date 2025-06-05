package main

import (
	"github.com/larsnieuwenhuizen/hcl-documenter-go/cmd"
	"os"
)

func main() {
	err := cmd.RootCmd().Execute()
	if err != nil {
		os.Exit(1)
	}
}
