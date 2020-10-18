package main

import (
	"github.com/mkorenkov/go-style-linters/internal/linters/elser"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(elser.Analyzer)
}
