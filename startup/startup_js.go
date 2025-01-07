package startup

import (
	"syscall/js"

	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
)

func init() {
	GDExtension := js.Global().Get("GDExtension")
	if GDExtension.IsUndefined() {
		GDExtension = js.Global().Get("Object").New()
		js.Global().Set("GDExtension", GDExtension)
	}
	GDExtension.Set("res://library.gdextension", js.FuncOf(func(_ js.Value, args []js.Value) any {
		get_proc_func := args[0]
		linkJS(&gd.Global, func(s string) js.Value {
			return get_proc_func.Invoke(s)
		})
		initialization := args[2]
		initialization.Set("minimum_initialization_level", uint32(gd.GDExtensionInitializationLevelScene))
		initialization.Set("initialize", js.FuncOf(func(_ js.Value, args []js.Value) any {
			level := gd.GDExtensionInitializationLevel(args[0].Int())
			gd.Global.Init(level)
			if level == 2 {
				for _, fn := range gd.StartupFunctions {
					fn()
				}
				for _, fn := range gd.PostStartupFunctions {
					fn()
				}
			}
			return nil
		}))
		initialization.Set("deinitialize", js.FuncOf(func(_ js.Value, args []js.Value) any {
			for _, cleanup := range gd.Cleanups() {
				cleanup()
			}
			pointers.Cycle()
			pointers.Cycle()
			close(done)
			return nil
		}))
		return 0
	}))
}

func linkJS(API *gd.API, dlsym func(string) js.Value) {
	string_name_new := dlsym("string_name_new")
	API.StringNames.New = func(s string) gd.StringName {
		return pointers.New[gd.StringName]([1]uintptr{uintptr(string_name_new.Invoke(s).Float())})
	}
}
