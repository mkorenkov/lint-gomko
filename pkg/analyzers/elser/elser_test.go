package elser

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestElser(t *testing.T) {
	testdata := "../../../testdata/elser"
	testdata, err := filepath.Abs(testdata)
	if err != nil {
		log.Fatal(err)
	}
	res := analysistest.Run(t, testdata, Analyzer, "")
	require.NotNil(t, res)
	require.NotEmpty(t, res)
	require.Equal(t, 1, len(res[0].Diagnostics))
	diagRes := res[0].Diagnostics[0]
	require.Equal(t, analyzerName, diagRes.Category)
	require.Equal(t, analyzerMsg, diagRes.Message)
}
