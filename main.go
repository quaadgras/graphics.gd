package main

import (
	"net/http"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("."))
	contentTypes := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasSuffix(r.URL.Path, ".json") {
			w.Header().Set("Content-Type", "application/json")
		}

		fs.ServeHTTP(w, r)

	})

	http.ListenAndServe("localhost:8080", contentTypes)
}
