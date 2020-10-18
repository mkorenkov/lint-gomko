package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/mkorenkov/go-style-linters/internal/linters/elser"
	"github.com/mkorenkov/go-style-linters/internal/linters/readall"
)

func main() {
	unitchecker.Main(
		readall.Analyzer,
		elser.Analyzer,
	)
}
