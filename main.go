package main

import (
	"os"

	"github.com/laupski/open-calculator/command/open-calculator"
)

func main() {
	os.Exit(command.Run(os.Args[1:]))
}