package startup

import (
	"os"
	"os/user"
	"path/filepath"
	"runtime"

	"graphics.gd/classdb/EditorInterface"
	EngineClass "graphics.gd/classdb/Engine"
	"graphics.gd/internal/gdextension"
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
				my, err := user.Current()
				if err == nil {
					HOME := my.HomeDir
					GDPATH := os.Getenv("GDPATH")
					if GDPATH == "" && HOME != "" {
						GDPATH = filepath.Join(HOME, "gd")
					}
					settings.SetSetting("export/android/java_sdk_path", GDPATH)
				}
			}
			// work around godot bug on windows
			android_sdk_path := settings.GetSetting("export/android/android_sdk_path").(String.Readable).String()
			if runtime.GOOS == "windows" && android_sdk_path == os.Getenv("LOCALAPPDATA")+"/Android/Sdk" {
				settings.SetSetting("export/android/java_sdk_path", filepath.Join(os.Getenv("LOCALAPPDATA"), "Android", "Sdk"))
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
	gdextension.On.MainLoop.EveryFrame = func() {
		Callable.Cycle()
		keep_reachable_instances_alive()
		pointers.Cycle()
	}
	gdextension.On.MainLoop.FinalFrame = func() {

	}
}
