package appendr

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

// Analyzer append is not efficient on the heap and is not prone to race conditions.
var Analyzer = &analysis.Analyzer{
	Name: "appendr",
	Doc:  "finds append statements in the code",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			if call, ok := n.(*ast.CallExpr); ok {
				if ident, ok := call.Fun.(*ast.Ident); ok {
					if ident.Name == "append" {
						pass.Report(analysis.Diagnostic{
							Pos:            ident.Pos(),
							End:            0,
							Category:       "appendr",
							Message:        "append is not efficient on the heap and is not prone to race conditions",
							SuggestedFixes: nil,
						})
					}
				}
			}
			return true
		})
	}
	return nil, nil
}
