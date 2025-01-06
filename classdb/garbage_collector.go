package classdb

import (
	NodeClass "graphics.gd/classdb/Node"
	SceneTreeClass "graphics.gd/classdb/SceneTree"
	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Float"
)

// goRuntime is injected into the scene tree so that the process function can process
// the frame-based garbage collection routine.
type goRuntime struct {
	Extension[goRuntime, NodeClass.Instance] `gd:"GoRuntime"`
}

func (gr goRuntime) AsNode() NodeClass.Instance { return gr.Super().AsNode() }

func (goRuntime) Process(delta Float.X) {
	Callable.New(func() {
		pointers.Cycle()
	}).CallDeferred()
}

func init() {
	gd.StartupFunctions = append(gd.StartupFunctions, func() {
		Register[goRuntime]()
	})
	gd.PostStartupFunctions = append(gd.PostStartupFunctions, func() {
		SceneTreeClass.Add(new(goRuntime))
	})
}
