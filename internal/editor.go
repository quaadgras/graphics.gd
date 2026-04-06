package gd

import (
	gdunsafe "graphics.gd"
)

func init() {
	gdunsafe.OnEditorClassDetection(func(classes gdunsafe.PackedArray[gdunsafe.String]) gdunsafe.PackedArray[gdunsafe.String] {
		return classes
	})
}
