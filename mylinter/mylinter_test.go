package mylinter_test

import (
	"testing"

	"github.com/0daryo/analizer-playground/mylinter"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, mylinter.Analyzer, "a")
}