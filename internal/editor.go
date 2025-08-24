package gd

import (
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdmemory"
)

func init() {
	gdextension.On.Editor = gdextension.CallbacksForEditor{
		ClassInUseDetection: func(classes gdextension.PackedArray[gdextension.String], result gdextension.Returns[gdextension.PackedArray[gdextension.String]]) {
			gdmemory.Set(gdextension.Pointer(result), classes)
		},
	}
}
