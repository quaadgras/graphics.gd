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

type editorPlugin struct {
	EditorPlugin.Extension[editorPlugin] `gd:"GoEditorPlugin"`
}

func (*editorPlugin) Build() bool {
	gd, err := exec.LookPath("gd")
	if err != nil {
		Engine.Raise(err)
		GOPATH, err := exec.Command("go", "env", "GOPATH").CombinedOutput()
		if err != nil {
			Engine.Raise(err)
			Engine.RaiseWarning("go and gd command not found, cannot build Go source")
			return true
		}
		gd = filepath.Join(string(GOPATH), "bin", "gd")
		if _, err := os.Stat(gd); os.IsNotExist(err) {
			cmd := exec.Command("go", "install", "graphics.gd/cmd/gd@release")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			if err := cmd.Run(); err != nil {
				Engine.Raise(err)
				Engine.RaiseWarning("failed to install the gd command, cannot build Go source")
				return true
			}
		}
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
