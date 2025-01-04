package GDExtension

import gd "graphics.gd/internal"

// LibraryPath is the path to the shared library that contains the current GD extension.
func LibraryPath() string {
	return gd.Global.GetLibraryPath(gd.Global.ExtensionToken).String()
}
