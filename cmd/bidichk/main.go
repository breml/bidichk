package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/breml/bidichk/pkg/bidichk"
)

func main() {
	singlechecker.Main(bidichk.NewAnalyzer())
}
