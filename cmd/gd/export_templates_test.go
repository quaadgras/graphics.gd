package main

import (
	"archive/zip"
	"io"
	"os"
	"path"
	"strings"
	"testing"
)

// TestTemplateSelectiveExtract checks per-platform selection against the real
// remote .tpz. Gated on GD_TEST_NETWORK because it hits github.com; it only
// reads the central directory and the tiny version.txt, so it does not download
// any of the large per-platform files.
func TestTemplateSelectiveExtract(t *testing.T) {
	if os.Getenv("GD_TEST_NETWORK") == "" {
		t.Skip("set GD_TEST_NETWORK=1 to run (downloads from github.com)")
	}
	const url = "https://github.com/godotengine/godot/releases/download/4.7-stable/Godot_v4.7-stable_export_templates.tpz"

	size, err := remoteSize(url)
	if err != nil {
		t.Fatal(err)
	}
	if size < 100<<20 {
		t.Fatalf("unexpected .tpz size %d", size)
	}
	ra := &blockReaderAt{ra: httpReaderAt{url: url}, size: size, block: 8 << 20}
	zr, err := zip.NewReader(ra, size)
	if err != nil {
		t.Fatal(err)
	}

	spec := templateSpecs["android"]
	var selected []string
	var versionEntry *zip.File
	for _, f := range zr.File {
		if strings.HasSuffix(f.Name, "/") {
			continue
		}
		base := path.Base(f.Name)
		if wantTemplateFile(base, spec) {
			selected = append(selected, base)
			if base == "version.txt" {
				versionEntry = f
			}
		}
		// A few representative foreign files must NOT be selected for android.
		if strings.HasPrefix(base, "windows_") || strings.HasPrefix(base, "web_") {
			if wantTemplateFile(base, spec) {
				t.Errorf("foreign file %q wrongly selected for android", base)
			}
		}
	}

	// The marker and the common files must be in the android selection.
	for _, must := range []string{spec.marker, "version.txt", "icudt_godot.dat"} {
		found := false
		for _, s := range selected {
			if s == must {
				found = true
			}
		}
		if !found {
			t.Errorf("expected %q in android selection, got %v", must, selected)
		}
	}
	if versionEntry == nil {
		t.Fatal("version.txt not found")
	}

	// Extract the tiny version.txt to prove single-entry extraction works.
	rc, err := versionEntry.Open()
	if err != nil {
		t.Fatal(err)
	}
	data, err := io.ReadAll(rc)
	rc.Close()
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), "4.7") {
		t.Errorf("version.txt = %q, want it to mention 4.7", data)
	}

	// Extract icudt_godot.dat (tens of MB) to exercise multi-block reads through
	// the block buffer and zip's CRC check on a large entry.
	var icu *zip.File
	for _, f := range zr.File {
		if path.Base(f.Name) == "icudt_godot.dat" {
			icu = f
		}
	}
	if icu == nil {
		t.Fatal("icudt_godot.dat not found")
	}
	rc2, err := icu.Open()
	if err != nil {
		t.Fatal(err)
	}
	n, err := io.Copy(io.Discard, rc2) // io.Copy completing means CRC validated
	rc2.Close()
	if err != nil {
		t.Fatalf("extracting icudt_godot.dat (%d bytes in): %v", n, err)
	}
	if n != int64(icu.UncompressedSize64) {
		t.Errorf("icudt_godot.dat: read %d bytes, want %d", n, icu.UncompressedSize64)
	}

	t.Logf(".tpz=%d MB; android selection (%d files): %v; version=%q; icudt=%d MB ok",
		size>>20, len(selected), selected, strings.TrimSpace(string(data)), n>>20)
}
