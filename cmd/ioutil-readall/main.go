package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/mkorenkov/lint-gomko/pkg/analyzers/readall"
)

func main() {
	singlechecker.Main(readall.Analyzer)
}
