package elser

import (
	"go/ast"

	"github.com/mkorenkov/lint-gomko/pkg/analyzers/nolinter"
	"github.com/mkorenkov/lint-gomko/pkg/reports"
	"golang.org/x/tools/go/analysis"
)

// Analyzer else is unnecessary, prefer early termination.
var Analyzer = &analysis.Analyzer{
	Name:     "elser",
	Doc:      "finds else statements in the code",
	Run:      run,
	Requires: []*analysis.Analyzer{nolinter.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		possibleReports := []*reports.Report{}

		ast.Inspect(file, func(n ast.Node) bool {
			for _, report := range possibleReports {
				if n != nil && report.Pos < n.Pos() && report.NextTokenPos < n.Pos() {
					report.NextTokenPos = n.Pos()
				}
			}

			if b, ok := n.(*ast.IfStmt); ok {
				if b.Else != nil {
					possibleReports = append(possibleReports, &reports.Report{
						Pos:          b.Else.Pos(),
						NextTokenPos: file.End(),
						Category:     "elser",
						Message:      "else is unnecessary, prefer early termination",
					})
				}
			}
			return true
		})

		for _, report := range possibleReports {
			if !nolinter.IsSupressed(pass, report.Pos, report.NextTokenPos) {
				pass.Report(analysis.Diagnostic{
					Pos:            report.Pos,
					End:            report.NextTokenPos,
					Category:       report.Category,
					Message:        report.Message,
					SuggestedFixes: nil,
				})
			}
		}
	}

	return nil, nil
}
