package main

import (
	"golang.org/x/tools/go/analysis"

	"github.com/0daryo/analizer-playground/mylinter"
)

// cf: https://golangci-lint.run/contributing/new-linters/#how-to-add-a-private-linter-to-golangci-lint

type analyzerPlugin struct{}

// This must be implemented
func (*analyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		mylinter.Analyzer,
	}
}

// This must be defined and named 'AnalyzerPlugin'
var AnalyzerPlugin analyzerPlugin
