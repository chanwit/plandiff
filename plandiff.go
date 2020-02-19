package plandiff

import (
	"github.com/chanwit/plandiff/difflib"
)

// GetUnifiedDiff generates a unified diff between the two WKSctl plans.
func GetUnifiedDiff(currentPlan, newPlan string) (string, error) {
	return difflib.GetUnifiedDiffString(difflib.UnifiedDiff{
		A:        difflib.SplitLines(currentPlan),
		B:        difflib.SplitLines(newPlan),
		FromFile: "current",
		ToFile:   "new",
		Context:  3,
	})
}
