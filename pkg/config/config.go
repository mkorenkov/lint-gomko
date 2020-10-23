package config

import (
	"os"
	"path"
	"path/filepath"
	"strings"
	"sync"
)

const ignorePathsEnv = "IGNORE"

var ignorePaths []string

// ShouldSkip checks `ignorePathsEnv` environment variable for glob patterns and
// matches those against the given `filename`.
func ShouldSkip(filename string) (bool, error) {
	var once sync.Once
	once.Do(func() {
		ignorePaths = strings.Split(os.Getenv(ignorePathsEnv), ",")
	})
	for _, p := range ignorePaths {
		d, err := os.Getwd()
		if err != nil {
			return false, err
		}

		res, err := filepath.Glob(p)
		if err != nil {
			return false, err
		}
		for _, matchedFileName := range res {
			if filename == path.Join(d, matchedFileName) {
				return true, nil
			}
		}
	}
	return false, nil
}
