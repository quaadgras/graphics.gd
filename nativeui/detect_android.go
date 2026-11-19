//go:build android

package nativeui

import (
	"os/exec"
	"strings"

	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/Engine"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

func platformLook(Version) Look { return Android }

func init() { platformDetect = detectAndroid }

// detectAndroid picks One UI on Samsung devices, which draw their own
// widgets rather than Material's, and reads the wallpaper colour.
func detectAndroid(opts *Options) {
	if isOneUI(getprop("ro.product.manufacturer"), getprop("ro.build.version.oneui")) {
		opts.Look = OneUI
	}
	// Android sizes are in dp, 160 to the inch. Godot's screen scale is a
	// rounded density bucket, capped by the window size, so use the DPI.
	if dpi := DisplayServer.ScreenGetDpi(); dpi > 0 {
		opts.Scale = Float.X(dpi) / 160
	}
	// Godot's accent is system_accent1_600, the wallpaper palette's tone 40,
	// which Material derives its scheme from. Samsung's widgets use their
	// own system_primary_light/dark instead.
	if opts.Look == OneUI {
		name := "system_primary_light"
		if opts.Dark {
			name = "system_primary_dark"
		}
		if accent, ok := systemColor(name); ok {
			opts.Accent = accent
		}
	}
}

// systemColor reads one of the framework's @android:color resources
// through Godot's AndroidRuntime plugin and JNI.
func systemColor(name string) (c Color.RGBA, ok bool) {
	defer func() {
		if r := recover(); r != nil {
			Engine.Print("nativeui: reading", name, ":", r)
			ok = false
		}
	}()
	if !Engine.HasSingleton("AndroidRuntime") {
		return c, false
	}
	asObject := func(v any) (Object.Any, bool) { o, ok := v.(Object.Any); return o, ok }
	v := Object.Call(Engine.GetSingleton("AndroidRuntime"), "getApplicationContext")
	ctx, ok := asObject(v)
	if !ok {
		return c, false
	}
	v = Object.Call(ctx, "getResources")
	res, ok := asObject(v)
	if !ok {
		return c, false
	}
	v = Object.Call(res, "getIdentifier", name, "color", "android")
	id := toInt(v)
	if id == 0 {
		return c, false
	}
	argb := uint32(toInt(Object.Call(ctx, "getColor", id)))
	return Color.RGBA{R: Float.X(argb>>16&0xff) / 255, G: Float.X(argb>>8&0xff) / 255, B: Float.X(argb&0xff) / 255, A: 1}, true
}

func toInt(v any) int64 {
	switch n := v.(type) {
	case int64:
		return n
	case int:
		return int64(n)
	case int32:
		return int64(n)
	case float64:
		return int64(n)
	}
	return 0
}

func getprop(name string) string {
	out, err := exec.Command("/system/bin/getprop", name).Output()
	if err != nil {
		return ""
	}
	return strings.TrimSpace(string(out))
}
