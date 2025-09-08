package main

import (
	"os"

	"github.com/Portfolio-jaime/${{values.app_name}}/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
