package startup

import (
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"reflect"
	"runtime"
	"slices"
	"strings"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Engine"
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdreference"
	"graphics.gd/variant/Object"
)

func editorSetup() {
	// Setup Faux SDKs. Not on-device: the Godot Android Editor app has no
	// export/android/* editor settings, so GetSetting returns nil there
	// (asserting it crashed the whole editor in the Termux flow), and gd's
	// faux SDKs don't apply on-device anyway.
	if runtime.GOOS == "android" {
		return
	}
	settings := editorSingleton("get_editor_settings")
	if settings == Object.Nil {
		return
	}
	if java_sdk_path, _ := Object.Call(settings, "get_setting", "export/android/java_sdk_path").(string); java_sdk_path == "" {
		my, err := user.Current()
		if err == nil {
			HOME := my.HomeDir
			GDPATH := os.Getenv("GDPATH")
			if GDPATH == "" && HOME != "" {
				GDPATH = filepath.Join(HOME, "gd")
			}
			Object.Call(settings, "set_setting", "export/android/java_sdk_path", GDPATH)
		}
	}
	// work around godot bug on windows, the default is not a native path.
	android_sdk_path, _ := Object.Call(settings, "get_setting", "export/android/android_sdk_path").(string)
	if runtime.GOOS == "windows" && android_sdk_path == os.Getenv("LOCALAPPDATA")+"/Android/Sdk" {
		Object.Call(settings, "set_setting", "export/android/android_sdk_path", filepath.Join(os.Getenv("LOCALAPPDATA"), "Android", "Sdk"))
	}
}

// editorSingleton calls a getter on the EditorInterface singleton by name, so
// that startup does not import the editor classes into every program. It
// returns Object.Nil outside the editor.
func editorSingleton(getter string) Object.Instance {
	if !Engine.HasSingleton("EditorInterface") {
		return Object.Nil
	}
	result, _ := Object.Call(Engine.GetSingleton("EditorInterface"), getter).(Object.Any)
	if result == nil {
		return Object.Nil
	}
	return Object.Instance(result.AsObject())
}

// editorPluginClass stands in for EditorPlugin.Instance, so that startup can
// extend EditorPlugin without importing the editor classes (and every class
// their methods mention) into every program. classdb names the engine class
// after the element type, recognises an editor plugin by its AsEditorPlugin
// method and resolves virtual methods through Virtual, so this is all of
// EditorPlugin that GoEditorPlugin needs.
type editorPluginClass [1]gdclass.EditorPlugin

func (o editorPluginClass) AsObject() [1]gdreference.Object   { return gdclass.GetEditorPlugin(o[0]) }
func (o editorPluginClass) AsEditorPlugin() editorPluginClass { return o }

func (o editorPluginClass) Virtual(name string) reflect.Value {
	switch name {
	case "_build":
		return reflect.ValueOf(o._build)
	}
	return reflect.Value{}
}

func (editorPluginClass) _build(impl func(ptr gdclass.Receiver) bool) (cb gd.ExtensionClassCallVirtualFunc) {
	return func(class any, p_args, p_back gdextension.Pointer) {
		gd.UnsafeSet(p_back, impl(gdclass.ReceiverOf(class)))
	}
}

type editorPlugin struct {
	gdclass.Extension[editorPlugin, editorPluginClass] `gd:"GoEditorPlugin"`
}

func (*editorPlugin) Build() bool {
	gd, err := exec.LookPath("gd")
	if err != nil {
		return true // no gd, passthrough to usual process.
	}
	cmd := exec.Command(gd)
	environ := os.Environ()
	environ = slices.DeleteFunc(environ, func(env string) bool {
		return strings.HasPrefix(env, "GOOS=")
	})
	cmd.Env = append(environ, "RUNNING_INSIDE_GODOT=1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		Engine.Raise(err)
		return false
	}
	return true
}

func init() {
	classdb.Register[editorPlugin]()
}
