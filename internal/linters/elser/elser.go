package elser

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// Analyzer else is unnecessary, prefer early termination.
var Analyzer = &analysis.Analyzer{
	Name: "elser",
	Doc:  "finds else statement in the code",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if b, ok := n.(*ast.IfStmt); ok {

				if b.Else != nil {
					pass.Report(analysis.Diagnostic{
						Pos:            b.Else.Pos(),
						End:            0,
						Category:       "elser",
						Message:        "else is unnecessary, prefer early termination.",
						SuggestedFixes: nil,
					})
				}
			}

			return true
		})
	}
	return nil, nil
}
