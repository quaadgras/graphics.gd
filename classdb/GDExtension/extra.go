package GDExtension

import (
	gdunsafe "graphics.gd"
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
)

// LibraryPath is the path to the shared library that contains the current GD extension.
func LibraryPath() string {
	return pointers.New[gd.String](gdextension.String{gdextension.Pointer(gdunsafe.LibraryLocation())}).String()
}
