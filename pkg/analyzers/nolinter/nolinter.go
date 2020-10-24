package nolinter

import (
	"go/token"
	"reflect"
	"regexp"

	"golang.org/x/tools/go/analysis"
)

// Analyzer required by other analyzers to have access to source code comments matchin //nolint pattern.
var Analyzer = &analysis.Analyzer{
	Name:       "nolinter",
	Doc:        "required by other analyzers to have access to source code comments matchin //nolint pattern",
	Run:        run,
	ResultType: reflect.TypeOf([]NolintComment{}),
}

// NolintComment describes known facts about //nolint comment
type NolintComment struct {
	Comment string
	Pos     token.Pos
	End     token.Pos
}

var nolintRe = regexp.MustCompile(`^\s*//\s*nolint:?`)

func run(pass *analysis.Pass) (interface{}, error) {
	facts := []NolintComment{}
	for _, file := range pass.Files {
		for _, comments := range file.Comments {
			for _, comment := range comments.List {
				if nolintRe.MatchString(comment.Text) {
					facts = append(facts, NolintComment{ //nolint:appendr
						Comment: comment.Text,
						Pos:     comment.Slash,
						End:     comment.End(),
					})
				}
			}
		}
	}
	return facts, nil
}

// IsSupressed checks whether there is a nolint comment on the given line
func IsSupressed(pass *analysis.Pass, pos token.Pos, nextTokenPos token.Pos) bool {
	if comments, ok := pass.ResultOf[Analyzer].([]NolintComment); ok {
		for _, comment := range comments {
			if comment.Pos >= pos && comment.End <= nextTokenPos {
				return true
			}
		}
	}
	return false
}
