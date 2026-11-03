package startup

import "os"

// setMainExtension tells an engine that is created by this program (libgodot)
// that the project's library.gdextension is this program's own extension, so
// that the engine registers it under that path, instead of libgodot://main,
// and does not load it a second time from the file's libraries (which a static
// program could not load anyway). The engine unsets it once read, and an engine
// that does not know of it loads the file's libraries as it always has.
func setMainExtension() {
	if isStub {
		return // a Godot project: it has no library.gdextension.
	}
	os.Setenv("LIBGODOT_MAIN_EXTENSION", "res://library.gdextension")
}
