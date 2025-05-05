package main

import (
	"github.com/Nikita-koder/terraformdescriptioncheck"
	"golang.org/x/tools/go/analysis"
)

func NewAnalyzer() *analysis.Analyzer {
	return terraformdescriptioncheck.Analyzer
}
