package reports

import "go/token"

// Report captures position of current token vs next token. This allows efficient //nolint comment filtering
type Report struct {
	Pos          token.Pos
	NextTokenPos token.Pos
	Category     string
	Message      string
}
