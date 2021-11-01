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
	analysistest.Run(t, testdata, bidichk.Analyzer, "simple")
	analysistest.Run(t, testdata, bidichk.Analyzer, "commenting-out")
	analysistest.Run(t, testdata, bidichk.Analyzer, "stretched-string")
}
