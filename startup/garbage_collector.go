package startup

import (
	"graphics.gd/classdb"

	NodeClass "graphics.gd/classdb/Node"
	SceneTreeClass "graphics.gd/classdb/SceneTree"
	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Float"
)

// TODO: switch to mainloop callbacks for Godot 4.5

// goRuntime is injected into the scene tree so that the process function can process
// the frame-based garbage collection routine.
type goRuntime struct {
	NodeClass.Extension[goRuntime] `gd:"GoRuntime"`
}

func (goRuntime) Ready() {
	Callable.Cycle()
}

func (gr goRuntime) ExitTree() {
	gd.NewCallable(func() {
		SceneTreeClass.Add(new(goRuntime)) // we need the GoRuntime to be present in the scene tree
	}).CallDeferred()
}

func (goRuntime) Process(delta Float.X) {
	gd.NewCallable(func() {
		Callable.Cycle()
		pointers.Cycle()
	}).CallDeferred()
}

func init() {
	gd.StartupFunctions = append(gd.StartupFunctions, func() {
		classdb.Register[goRuntime]()
	})
	gd.PostStartupFunctions = append(gd.PostStartupFunctions, func() {
		gd.NewCallable(func() {
			SceneTreeClass.Add(new(goRuntime))
		}).CallDeferred()
	})
}
