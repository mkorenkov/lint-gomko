package appendr

import (
	"go/ast"

	"github.com/mkorenkov/lint-gomko/pkg/analyzers/nolinter"
	"github.com/mkorenkov/lint-gomko/pkg/reports"
	"golang.org/x/tools/go/analysis"
)

const analyzerName = "appendr"
const analyzerMsg = "append is not efficient on the heap and is not prone to race conditions"

// Analyzer append is not efficient on the heap and is not prone to race conditions.
var Analyzer = &analysis.Analyzer{
	Name:     analyzerName,
	Doc:      "finds append statements in the code",
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

			if call, ok := n.(*ast.CallExpr); ok {
				if ident, ok := call.Fun.(*ast.Ident); ok {
					if ident.Name == "append" {
						possibleReports = append(possibleReports, &reports.Report{
							Pos:          ident.Pos(),
							NextTokenPos: file.End(),
							Category:     analyzerName,
							Message:      analyzerMsg,
						})
					}
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
