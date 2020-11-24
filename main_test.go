package main

import (
	"testing"
)

func TestFilterLineNumber(t *testing.T) {
	result := `root := folder
	if options.ExpandWorkspaceToModule {
		wsRoot, _ := findWorkspaceRoot(ctx, root, s)
		if wsRoot != "" {
			root = wsRoot
		}
	}`
	code := `root := folder
677
	if options.ExpandWorkspaceToModule {
678
		wsRoot, _ := findWorkspaceRoot(ctx, root, s)
682
		if wsRoot != "" {
683
			root = wsRoot
684
		}
685
	}`
	filterRet := filterLineNumber(code)
	if filterRet != result {
		t.Fatalf("filter result %v != result %v", filterRet, result)
	}
}
