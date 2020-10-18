package nolinter

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestNolints(t *testing.T) {
	testdata := "../../../testdata/nolinter"
	testdata, err := filepath.Abs(testdata)
	if err != nil {
		log.Fatal(err)
	}
	res := analysistest.Run(t, testdata, Analyzer, "")
	require.NotNil(t, res)
	require.NotEmpty(t, res)
	facts, ok := res[0].Result.([]NolintComment)
	require.True(t, ok)
	require.Equal(t, 5, len(facts))
}
