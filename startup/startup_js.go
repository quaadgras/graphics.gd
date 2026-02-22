package startup

import (
	"iter"
	"slices"
	"testing"

	"graphics.gd/classdb"
	EngineClass "graphics.gd/classdb/Engine"
	"graphics.gd/classdb/SceneTree"
	gd "graphics.gd/internal"
	internal "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Float"
)

var loaded = make(chan struct{})
var hasLoaded bool
var intialized = make(chan struct{})
var shutdown = make(chan struct{})
var main_loop_shutdown = make(chan struct{})

func (engine *engineAsSharedLibrary) Start() {
	<-intialized
	<-loaded
}

func (engine *engineAsSharedLibrary) Scene() {
	<-shutdown
}

func (engine *engineAsSharedLibrary) Rendering() iter.Seq[Float.X] {
	classdb.Register[goMain]()
	<-frame_ready
	return func(yield func(Float.X) bool) {
		frame_ready <- false
		for {
			<-frame_ready // we pause here until the next frame is ready (next Process callback).
			if !yield(dt) {
				frame_ready <- true
				break
			}
			frame_ready <- false
		}
		<-main_loop_shutdown
	}
}

func init() {
	gdextension.On.Engine = gdextension.CallbacksForEngine{
		Init: func(level gdextension.InitializationLevel) {
			gd.Init(level)
			if level == 2 {
				for _, fn := range gd.StartupFunctions {
					fn()
				}
				if testing.Testing() {
					classdb.Register[goSceneTree]()
				}
				close(intialized)
				for _, fn := range gd.PostStartupFunctions {
					fn()
				}
			}
		},
		Exit: func(level gdextension.InitializationLevel) {
			if level == 2 {
				for _, cleanup := range slices.Backward(gd.Cleanups()) {
					cleanup()
				}
				pointers.Cycle()
				pointers.Cycle()
				close(shutdown)
				internal.Linked = false
			}
		},
	}
	gdextension.On.MainLoop.FirstFrame = func() {
		Callable.Cycle()
		if EngineClass.IsEditorHint() {
			editorSetup()
		}
		if !hasLoaded {
			close(loaded)
			hasLoaded = true
		}
	}
}

type goMain struct {
	SceneTree.Extension[goSceneTree] `gd:"GoMainLoop"`
}

func (loop goMain) Initialize() {
	Callable.Cycle()
}

func (loop goMain) PhysicsProcess(delta Float.X) bool {
	return false
}

func (loop goMain) Process(delta Float.X) bool {
	defer Callable.Cycle()
	defer keep_reachable_instances_alive()
	defer pointers.Cycle()
	dt = delta
	frame_ready <- false
	return <-frame_ready
}

func (loop goMain) Finalize() {
	close(main_loop_shutdown)
}
