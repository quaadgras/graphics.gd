//go:build ios

package main

import (
	"os"
	"path/filepath"

	"graphics.gd/classdb/OS"
)

// defaultProject on iOS is a directory inside the app sandbox: there is
// no working directory or environment to inherit. Call after
// startup.LoadingScene (it asks the engine).
func defaultProject() string {
	project := filepath.Join(OS.GetUserDataDir(), "project")
	os.MkdirAll(project, 0o755)
	return project
}
