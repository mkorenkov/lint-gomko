package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/mkorenkov/lint-gomko/pkg/linters/appendr"
	"github.com/mkorenkov/lint-gomko/pkg/linters/elser"
	"github.com/mkorenkov/lint-gomko/pkg/linters/readall"
)

func main() {
	unitchecker.Main(
		appendr.Analyzer,
		elser.Analyzer,
		readall.Analyzer,
	)
}
