package bidichk_test

import (
	"os"
	"path/filepath"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/breml/bidichk/pkg/bidichk"
)

func TestAll(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get wd: %s", err)
	}

	testdata := filepath.Join(filepath.Dir(filepath.Dir(wd)), "testdata")
	analysistest.Run(t, testdata, bidichk.NewAnalyzer(), "simple")
	analysistest.Run(t, testdata, bidichk.NewAnalyzer(), "commenting-out")
	analysistest.Run(t, testdata, bidichk.NewAnalyzer(), "stretched-string")

	configuredAnalyzer := bidichk.NewAnalyzer()
	configuredAnalyzer.Flags.Set("disallowed-runes", "LEFT-TO-RIGHT-ISOLATE")
	analysistest.Run(t, testdata, configuredAnalyzer, "commenting-out-lri-only")

	configuredAnalyzerShortName := bidichk.NewAnalyzer()
	configuredAnalyzerShortName.Flags.Set("disallowed-runes", "LRI")
	analysistest.Run(t, testdata, configuredAnalyzerShortName, "commenting-out-lri-only")
}
