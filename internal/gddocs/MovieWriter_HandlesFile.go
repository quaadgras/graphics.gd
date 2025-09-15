/*
func _handles_file(path):
    # Allows specifying an output file with a `.mkv` file extension (case-insensitive),
    # either in the Project Settings or with the `--write-movie <path>` command line argument.
    return path.get_extension().to_lower() == "mkv"
*/

package main

import (
	"path/filepath"
	"strings"
)

func MovieWriter_HandlesFile() {
	HandlesFile := func(path string) bool {
		return strings.ToLower(filepath.Ext(path)) == ".mkv"
	}
	_ = HandlesFile
}
