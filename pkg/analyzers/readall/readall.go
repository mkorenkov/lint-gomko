package readall

import (
	"fmt"
	"go/ast"
	"go/token"
	"path"

	"github.com/mkorenkov/lint-gomko/pkg/analyzers"
	"github.com/mkorenkov/lint-gomko/pkg/analyzers/nolinter"
	"github.com/mkorenkov/lint-gomko/pkg/reports"
	"golang.org/x/tools/go/analysis"
)

const analyzerName = "readall"
const analyzerMsg = "ioutil.ReadAll is expensive and should be avoided"

// Analyzer ioutil.ReadAll is expensive. This linter will nudge you about ioutil.ReadAll presence in your code.
var Analyzer = &analysis.Analyzer{
	Name:     analyzerName,
	Doc:      "finds ioutil.ReadAll usages",
	Run:      analyzers.Analyze(run),
	Requires: []*analysis.Analyzer{nolinter.Analyzer},
}

func run(n ast.Node, importAliases map[string]string, lastPos token.Pos) []reports.Report {
	res := []reports.Report{}

	if call, ok := n.(*ast.CallExpr); ok {
		if fun, ok := call.Fun.(*ast.SelectorExpr); ok {
			funcName := fun.Sel.Name
			if pkgID, ok := fun.X.(*ast.Ident); ok {
				pkgName := pkgID.Name
				if alias, ok := importAliases[pkgName]; ok {
					pkgName = path.Base(alias)
				}
				if fmt.Sprintf("%s.%s", pkgName, funcName) == "ioutil.ReadAll" {
					res = append(res, reports.Report{
						Pos:          pkgID.Pos(),
						NextTokenPos: lastPos,
						Category:     analyzerName,
						Message:      analyzerMsg,
					})
				}
			}
		}
	}
	return res
}
