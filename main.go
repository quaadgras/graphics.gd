package main

import (
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"runtime.link/api"
	"runtime.link/api/cmdl"
)

var rclone = api.Import[struct {
	Copy func(src, dst string) error `cmdl:"copy --s3-no-check-bucket %v %v"`
}](cmdl.API, "rclone", nil)

var gd = api.Import[struct {
	Build func() error `cmdl:"build"`
}](cmdl.API, "gd", nil)

func main() {
	if len(os.Args) > 1 {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println("error:", err)
			return
		}
		os.Setenv("GOOS", "web")
		filepath.WalkDir(".", func(p string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			if d.Name() == "graphics" {
				example := filepath.Dir(p)
				if os.Args[1] == "copy" {
					if err := os.Chdir(example); err != nil {
						fmt.Println("error:", err)
					}
					if err := gd.Build(); err != nil {
						fmt.Println("error:", err)
						return err
					}
					if err := os.Chdir(wd); err != nil {
						fmt.Println("error:", err)
					}
					if err := rclone.Copy(example+"/releases/js/wasm", "r2:samples-graphics-gd/"+filepath.ToSlash(example)); err != nil {
						fmt.Println("error:", err)
						return err
					}
					if err := rclone.Copy(example+".webp", "r2:samples-graphics-gd/"+path.Dir(filepath.ToSlash(example))); err != nil {
						fmt.Println("error:", err)
						return err
					}
				} else {
					fmt.Println(example)
				}
				return filepath.SkipDir
			}
			return nil
		})
		return
	}
	fs := http.FileServer(http.Dir("."))
	contentTypes := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".json") {
			w.Header().Set("Content-Type", "application/json")
		}

		fs.ServeHTTP(w, r)

	})

	http.ListenAndServe("localhost:8080", contentTypes)
}
