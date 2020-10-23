package appendr

import (
	"go/ast"
	"go/token"

	"github.com/mkorenkov/lint-gomko/pkg/analyzers"
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
	Run:      analyzers.Analyze(run),
	Requires: []*analysis.Analyzer{nolinter.Analyzer},
}

func run(n ast.Node, importAliases map[string]string, lastPos token.Pos) []reports.Report {
	res := []reports.Report{}

	if call, ok := n.(*ast.CallExpr); ok {
		if ident, ok := call.Fun.(*ast.Ident); ok {
			if ident.Name == "append" {
				res = append(res, reports.Report{
					Pos:          ident.Pos(),
					NextTokenPos: lastPos,
					Category:     analyzerName,
					Message:      analyzerMsg,
				})
			}
		}
	}

	return res
}
