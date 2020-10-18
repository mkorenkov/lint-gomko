package readall

import (
	"fmt"
	"go/ast"
	"path"
	"strconv"

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
	Run:      run,
	Requires: []*analysis.Analyzer{nolinter.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		possibleReports := []*reports.Report{}
		importAlias := map[string]string{}

		ast.Inspect(file, func(n ast.Node) bool {
			// collect import aliases
			if imp, ok := n.(*ast.ImportSpec); ok {
				if imp.Name != nil {
					val, err := strconv.Unquote(imp.Path.Value)
					if err != nil {
						panic(err)
					}
					importAlias[imp.Name.String()] = val
				}
			}
			// check function calls
			if call, ok := n.(*ast.CallExpr); ok {
				if fun, ok := call.Fun.(*ast.SelectorExpr); ok {
					funcName := fun.Sel.Name
					if pkgID, ok := fun.X.(*ast.Ident); ok {
						pkgName := pkgID.Name
						if alias, ok := importAlias[pkgName]; ok {
							pkgName = path.Base(alias)
						}
						if fmt.Sprintf("%s.%s", pkgName, funcName) == "ioutil.ReadAll" {
							possibleReports = append(possibleReports, &reports.Report{
								Pos:          pkgID.Pos(),
								NextTokenPos: file.End(),
								Category:     analyzerName,
								Message:      analyzerMsg,
							})
						}
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
