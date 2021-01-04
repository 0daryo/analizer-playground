package main

import (
	"github.com/0daryo/analizer-playground/mylinter"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() { singlechecker.Main(mylinter.Analyzer) }
