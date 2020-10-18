package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/mkorenkov/go-style-linters/pkg/linters/appendr"
	"github.com/mkorenkov/go-style-linters/pkg/linters/elser"
	"github.com/mkorenkov/go-style-linters/pkg/linters/readall"
)

func main() {
	unitchecker.Main(
		appendr.Analyzer,
		elser.Analyzer,
		readall.Analyzer,
	)
}
