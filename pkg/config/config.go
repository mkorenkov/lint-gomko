package config

import (
	"os"
	"strings"
	"sync"

	"github.com/gobwas/glob"
)

const ignorePathsEnv = "IGNORE"

var ignorePaths []string

var once sync.Once

// ShouldSkip checks `ignorePathsEnv` environment variable for glob patterns and
// matches those against the given `filename`.
func ShouldSkip(filename string) (bool, error) {
	once.Do(func() {
		ignorePaths = strings.Split(os.Getenv(ignorePathsEnv), ",")
	})
	for _, p := range ignorePaths {
		g, err := glob.Compile(p, '/')
		if err != nil {
			return false, err
		}
		if g.Match(filename) {
			return true, nil
		}
	}
	return false, nil
}
