package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"path/filepath"
	"strings"
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
	fs := http.FileServer(http.Dir("."))
	contentTypes := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".json") {
			w.Header().Set("Content-Type", "application/json")
		}

		fs.ServeHTTP(w, r)

	})

	http.ListenAndServe("localhost:8080", contentTypes)
}
