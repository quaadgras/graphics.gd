package startup

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"graphics.gd/classdb/EditorInterface"
	EngineClass "graphics.gd/classdb/Engine"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/mainthread"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/String"

	_ "unsafe"
)

//go:linkname keep_reachable_instances_alive graphics.gd/classdb.keep_reachable_instances_alive
func keep_reachable_instances_alive()

func init() {
	gdextension.On.MainLoop.FirstFrame = func() {
		Callable.Cycle()
		if EngineClass.IsEditorHint() {
			settings := EditorInterface.GetEditorSettings()
			if settings.GetSetting("export/android/java_sdk_path").(String.Readable).String() == "" {
				GDPATH := os.Getenv("GDPATH")
				if GDPATH == "" {
					GDPATH = filepath.Join(os.Getenv("HOME"), "gd")
				}
				settings.SetSetting("export/android/java_sdk_path", GDPATH)
			}
		}
		if pause_main != nil {
			resume_main()
		} else {
			if !hasLoaded {
				close(loaded)
				hasLoaded = true
			}
		}
	}
	var last_gc = time.Now()
	gdextension.On.MainLoop.EveryFrame = func() {
		Callable.Cycle()
		for range 10 {
			mainthread.Yield()
		}
		if time.Since(last_gc) > time.Second/60 {
			fmt.Println("startup.GC()")
			keep_reachable_instances_alive()
			pointers.Cycle()
			last_gc = time.Now()
		}
	}
	gdextension.On.MainLoop.FinalFrame = func() {

	}
}
