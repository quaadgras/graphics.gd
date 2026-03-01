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
	"graphics.gd/variant/String"
)

func editorSetup() {
	// Setup Faux SDKs
	settings := EditorInterface.GetEditorSettings()
	if settings.GetSetting("export/android/java_sdk_path").(String.Unicode).String() == "" {
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
	android_sdk_path := settings.GetSetting("export/android/android_sdk_path").(String.Unicode).String()
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
