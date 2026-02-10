package startup

import (
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Callable"

	_ "unsafe"
)

//go:linkname keep_reachable_instances_alive graphics.gd/classdb.keep_reachable_instances_alive
func keep_reachable_instances_alive()

func init() {
	gdextension.On.MainLoop.EveryFrame = func() {
		Callable.Cycle()
		keep_reachable_instances_alive()
		pointers.Cycle()
	}
	gdextension.On.MainLoop.FinalFrame = func() {

	}
}
