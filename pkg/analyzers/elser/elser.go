package elser

import (
	"go/ast"
	"go/token"

	"github.com/mkorenkov/lint-gomko/pkg/analyzers"
	"github.com/mkorenkov/lint-gomko/pkg/analyzers/nolinter"
	"github.com/mkorenkov/lint-gomko/pkg/reports"
	"golang.org/x/tools/go/analysis"
)

const analyzerName = "elser"
const analyzerMsg = "else is unnecessary, prefer early termination"

// Analyzer else is unnecessary, prefer early termination.
var Analyzer = &analysis.Analyzer{
	Name:     analyzerName,
	Doc:      "finds else statements in the code",
	Run:      analyzers.Analyze(run),
	Requires: []*analysis.Analyzer{nolinter.Analyzer},
}

func run(n ast.Node, importAliases map[string]string, lastPos token.Pos) []reports.Report {
	res := []reports.Report{}

	if b, ok := n.(*ast.IfStmt); ok {
		if b.Else != nil {
			res = append(res, reports.Report{
				Pos:          b.Else.Pos(),
				NextTokenPos: lastPos,
				Category:     analyzerName,
				Message:      analyzerMsg,
			})
		}
	}

	return res
}
