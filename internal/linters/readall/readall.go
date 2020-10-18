package readall

import (
	"fmt"
	"go/ast"
	"path"
	"strconv"

	"golang.org/x/tools/go/analysis"
)

// Analyzer ioutil.ReadAll is expensive. This linter will nudge you about ioutil.ReadAll presence in your code.
var Analyzer = &analysis.Analyzer{
	Name: "readall",
	Doc:  "finds ioutil.ReadAll usages",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
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
							pass.Report(analysis.Diagnostic{
								Pos:            pkgID.Pos(),
								End:            0,
								Category:       "readall",
								Message:        "ioutil.ReadAll is expensive and should be avoided",
								SuggestedFixes: nil,
							})
						}
					}
				}
			}
			return true
		})
	}
	return nil, nil
}
