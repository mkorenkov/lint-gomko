package main

import (
	"github.com/mkorenkov/go-style-linters/internal/linters/readall"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(readall.Analyzer)
}
