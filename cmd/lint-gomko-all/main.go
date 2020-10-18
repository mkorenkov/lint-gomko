package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/mkorenkov/lint-gomko/pkg/analyzers/appendr"
	"github.com/mkorenkov/lint-gomko/pkg/analyzers/elser"
	"github.com/mkorenkov/lint-gomko/pkg/analyzers/readall"
)

func main() {
	unitchecker.Main(
		appendr.Analyzer,
		elser.Analyzer,
		readall.Analyzer,
	)
}
