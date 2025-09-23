package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func main() {
	filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.Name() == "graphics" {
			fmt.Println(filepath.Dir(path))
			return filepath.SkipDir
		}
		return nil
	})
}
