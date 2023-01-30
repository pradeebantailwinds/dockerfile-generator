package main

import (
	"os"

	"github.com/pradeebantailwinds/dockerfile-generator/cmd/dfg/app"
)

func main() {
	if err := app.Run(); err != nil {
		os.Exit(1)
	}
}
