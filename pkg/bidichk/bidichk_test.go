package bidichk_test

import (
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/breml/bidichk/pkg/bidichk"
)

func TestAll(t *testing.T) {
	tt := []struct {
		name         string
		analyzerFunc func() *analysis.Analyzer
		testdataDir  string
	}{
		{
			name:         "simple",
			analyzerFunc: bidichk.NewAnalyzer,
			testdataDir:  "simple",
		},
		{
			name:         "commenting-out",
			analyzerFunc: bidichk.NewAnalyzer,
			testdataDir:  "commenting-out",
		},
		{
			name:         "stretched-string",
			analyzerFunc: bidichk.NewAnalyzer,
			testdataDir:  "stretched-string",
		},
		{
			name: "commenting-out-lri-only",
			analyzerFunc: func() *analysis.Analyzer {
				a := bidichk.NewAnalyzer()
				_ = a.Flags.Set("disallowed-runes", "LEFT-TO-RIGHT-ISOLATE")
				return a
			},
			testdataDir: "commenting-out-lri-only",
		},
		{
			name: "commenting-out-lri-only short name",
			analyzerFunc: func() *analysis.Analyzer {
				a := bidichk.NewAnalyzer()
				_ = a.Flags.Set("disallowed-runes", "LRI")
				return a
			},
			testdataDir: "commenting-out-lri-only",
		},
	}

	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			analysistest.Run(t, testdata, tc.analyzerFunc(), tc.testdataDir)
		})
	}
}
