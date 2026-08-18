//go:build ios

package main

import (
	"os"
	"path/filepath"

	"graphics.gd/classdb/OS"

	"graphics.gd/harness/internal/term"
)

// defaultProject on iOS is a directory inside the app sandbox: there is
// no working directory or environment to inherit. Call after
// startup.LoadingScene (it asks the engine).
func defaultProject() string {
	project := filepath.Join(OS.GetUserDataDir(), "project")
	os.MkdirAll(project, 0o755)
	return project
}

// openURL hands a sidestore:// install link from a remote build to
// SideStore on this device, closing the loop: edit here, build there,
// install back here.
func openURL(console *term.Console, url string) {
	console.System("opening SideStore to install: " + url)
	if err := OS.ShellOpen(url); err != nil {
		console.Error("could not open " + url + ": " + err.Error())
	}
}
