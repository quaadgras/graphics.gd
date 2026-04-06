package startup

import (
	gdunsafe "graphics.gd"
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdreference"
	"graphics.gd/internal/pointers"
	"graphics.gd/internal/ring"
	"graphics.gd/variant/Callable"

	_ "unsafe"
)

//go:linkname keep_reachable_instances_alive graphics.gd/classdb.keep_reachable_instances_alive
func keep_reachable_instances_alive()

func init() {
	gdunsafe.OnEveryFrame(func() {
		Callable.Cycle()
		ring.Main.Flush()
		keep_reachable_instances_alive()
		gdreference.GC(gd.Free)
		pointers.Cycle()
	})
	gdunsafe.OnFinalFrame(func() {
	})
}
