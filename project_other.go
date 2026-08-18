//go:build !ios

package main

import (
	"os"

	"graphics.gd/harness/internal/term"
)

// defaultProject is the project the harness works on when
// GD_HARNESS_PROJECT is unset: the working directory.
func defaultProject() string {
	if project := os.Getenv("GD_HARNESS_PROJECT"); project != "" {
		return project
	}
	project, _ := os.Getwd()
	return project
}

// openURL on desktop just surfaces the link; the QR in the build
// output is already scannable from the screen.
func openURL(console *term.Console, url string) {
	console.System("install link: " + url)
}
