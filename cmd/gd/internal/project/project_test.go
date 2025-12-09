package project

import (
	"os"
	"path/filepath"
	"testing"
)

func TestFindGoModNoInfiniteLoop(t *testing.T) {
	// Create a temporary directory to test
	tmpDir := t.TempDir()
	
	// Test that findGoMod doesn't loop infinitely when no go.mod exists
	result, hasGoMod, err := findGoMod(tmpDir)
	if err != nil {
		t.Fatalf("findGoMod returned error: %v", err)
	}
	
	// Should return false when no go.mod is found
	if hasGoMod {
		t.Error("Expected hasGoMod to be false when no go.mod exists")
	}
	
	// Result should be a valid directory
	if result == "" {
		t.Error("Expected result to be non-empty")
	}
}

func TestFindGoModWithGoMod(t *testing.T) {
	// Create a temporary directory with a go.mod file
	tmpDir := t.TempDir()
	goModPath := filepath.Join(tmpDir, "go.mod")
	
	// Create a go.mod file
	if err := os.WriteFile(goModPath, []byte("module test\n"), 0644); err != nil {
		t.Fatalf("Failed to create go.mod: %v", err)
	}
	
	// Test that findGoMod finds the go.mod
	result, hasGoMod, err := findGoMod(tmpDir)
	if err != nil {
		t.Fatalf("findGoMod returned error: %v", err)
	}
	
	// Should return true when go.mod is found
	if !hasGoMod {
		t.Error("Expected hasGoMod to be true when go.mod exists")
	}
	
	// Result should be the directory with go.mod
	if result != tmpDir {
		t.Errorf("Expected result to be %s, got %s", tmpDir, result)
	}
}

func TestFindGoModWithGraphicsDir(t *testing.T) {
	// Create a temporary directory with a graphics subdirectory
	tmpDir := t.TempDir()
	graphicsDir := filepath.Join(tmpDir, "graphics")
	
	// Create graphics directory
	if err := os.MkdirAll(graphicsDir, 0755); err != nil {
		t.Fatalf("Failed to create graphics directory: %v", err)
	}
	
	// Test that findGoMod finds the graphics directory
	result, hasGoMod, err := findGoMod(tmpDir)
	if err != nil {
		t.Fatalf("findGoMod returned error: %v", err)
	}
	
	// Should return true when graphics directory exists
	if !hasGoMod {
		t.Error("Expected hasGoMod to be true when graphics directory exists")
	}
	
	// Result should be the directory with graphics
	if result != tmpDir {
		t.Errorf("Expected result to be %s, got %s", tmpDir, result)
	}
}

func TestFindGoModNestedDirectory(t *testing.T) {
	// Create a nested directory structure with go.mod at the top
	tmpDir := t.TempDir()
	goModPath := filepath.Join(tmpDir, "go.mod")
	nestedDir := filepath.Join(tmpDir, "subdir", "nested")
	
	// Create go.mod file at the root
	if err := os.WriteFile(goModPath, []byte("module test\n"), 0644); err != nil {
		t.Fatalf("Failed to create go.mod: %v", err)
	}
	
	// Create nested directories
	if err := os.MkdirAll(nestedDir, 0755); err != nil {
		t.Fatalf("Failed to create nested directory: %v", err)
	}
	
	// Test that findGoMod finds go.mod from nested directory
	result, hasGoMod, err := findGoMod(nestedDir)
	if err != nil {
		t.Fatalf("findGoMod returned error: %v", err)
	}
	
	// Should return true when go.mod is found
	if !hasGoMod {
		t.Error("Expected hasGoMod to be true when go.mod exists in parent")
	}
	
	// Result should be the root directory with go.mod
	if result != tmpDir {
		t.Errorf("Expected result to be %s, got %s", tmpDir, result)
	}
}
