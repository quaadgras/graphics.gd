package startup

import (
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"slices"
	"strings"

	"graphics.gd/classdb"
	"graphics.gd/classdb/EditorInterface"
	"graphics.gd/classdb/EditorPlugin"
	"graphics.gd/classdb/Engine"
)

func editorSetup() {
	// Setup Faux SDKs. Not on-device: the Godot Android Editor app has no
	// export/android/* editor settings, so GetSetting returns nil there
	// (asserting it crashed the whole editor in the Termux flow), and gd's
	// faux SDKs don't apply on-device anyway.
	if runtime.GOOS == "android" {
		return
	}
	settings := EditorInterface.GetEditorSettings()
	if java_sdk_path, _ := settings.GetSetting("export/android/java_sdk_path").(string); java_sdk_path == "" {
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
	android_sdk_path, _ := settings.GetSetting("export/android/android_sdk_path").(string)
	if runtime.GOOS == "windows" && android_sdk_path == os.Getenv("LOCALAPPDATA")+"/Android/Sdk" {
		settings.SetSetting("export/android/java_sdk_path", filepath.Join(os.Getenv("LOCALAPPDATA"), "Android", "Sdk"))
	}
}

type editorPlugin struct {
	EditorPlugin.Extension[editorPlugin] `gd:"GoEditorPlugin"`
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
