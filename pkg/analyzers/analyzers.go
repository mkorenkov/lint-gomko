package analyzers

import (
	"go/ast"
	"go/token"
	"strconv"

	"github.com/mkorenkov/lint-gomko/pkg/analyzers/nolinter"
	"github.com/mkorenkov/lint-gomko/pkg/config"
	"github.com/mkorenkov/lint-gomko/pkg/reports"
	"golang.org/x/tools/go/analysis"
)

// InspectFunc core logic of the linter, acts on an ast.Node, traversed depth-first
// returns possible lint violations, that will be further inspected for nolint rules
type InspectFunc func(n ast.Node, importAliases map[string]string, lastPos token.Pos) []reports.Report

// Analyze generates a `run` function for a linter, based on a simple template:
// it does filtering based on configured glob patterns;
// it collects import aliases and sends it as a linter function param
// it collects reports from a linter function and checks them against nolint rules
func Analyze(inspectFunc InspectFunc) func(pass *analysis.Pass) (interface{}, error) {
	return func(pass *analysis.Pass) (interface{}, error) {
		for _, file := range pass.Files {
			skip, err := config.ShouldSkip(pass.Fset.File(file.Pos()).Name())
			if err != nil {
				return nil, err
			}
			if skip {
				continue
			}

			possibleReports := []*reports.Report{}
			importAliases := map[string]string{}

			ast.Inspect(file, func(n ast.Node) bool {
				for _, report := range possibleReports {
					if n != nil && report.Pos < n.Pos() && report.NextTokenPos < n.Pos() {
						report.NextTokenPos = n.Pos()
					}
				}

				// collect import aliases
				if imp, ok := n.(*ast.ImportSpec); ok {
					if imp.Name != nil {
						val, err := strconv.Unquote(imp.Path.Value)
						if err != nil {
							panic(err)
						}
						importAliases[imp.Name.String()] = val
					}
				}

				// collect reports from actual linter
				newReports := inspectFunc(n, importAliases, file.End())
				for _, r := range newReports {
					possibleReports = append(possibleReports, &r) //nolint:appendr
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
}
