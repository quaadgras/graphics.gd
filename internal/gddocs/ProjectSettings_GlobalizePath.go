/*
var path = ""
if OS.has_feature("editor"):
	# Running from an editor binary.
	# `path` will contain the absolute path to `hello.txt` located in the project root.
	path = ProjectSettings.globalize_path("res://hello.txt")
else:
	# Running from an exported project.
	# `path` will contain the absolute path to `hello.txt` next to the executable.
	# This is *not* identical to using `ProjectSettings.globalize_path()` with a `res://` path,
	# but is close enough in spirit.
	path = OS.get_executable_path().get_base_dir().path_join("hello.txt")
*/

package main

import (
	"path/filepath"

	"graphics.gd/classdb/OS"
	"graphics.gd/classdb/ProjectSettings"
)

func ProjectSettings_GlobalizePath() {
	var path = ""
	if OS.HasFeature("editor") {
		// Running from an editor binary.
		// `path` will contain the absolute path to `hello.txt` located in the project root.
		path = ProjectSettings.GlobalizePath("res://hello.txt")
	} else {
		// Running from an exported project.
		// `path` will contain the absolute path to `hello.txt` next to the executable.
		// This is *not* identical to using `ProjectSettings.globalize_path()` with a `res://` path,
		// but is close enough in spirit.
		path = filepath.Join(filepath.Base(OS.GetExecutablePath()), "hello.txt")
	}
	_ = path
}
