//go:build !ios

package main

import "os"

// defaultProject is the project the harness works on when
// GD_HARNESS_PROJECT is unset: the working directory.
func defaultProject() string {
	if project := os.Getenv("GD_HARNESS_PROJECT"); project != "" {
		return project
	}
	project, _ := os.Getwd()
	return project
}
