package main

import (
	"github.com/0daryo/analizer-playground/mylinter"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

func main() {
	multichecker.Main(
		inspect.Analyzer,
		mylinter.Analyzer,
	)
}
