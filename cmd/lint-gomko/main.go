package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/mkorenkov/lint-gomko/pkg/analyzers/readall"
)

func main() {
	unitchecker.Main(
		readall.Analyzer,
	)
}
