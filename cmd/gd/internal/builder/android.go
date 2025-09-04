package builder

import (
	"embed"
	"fmt"
	"go/build"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"

	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/tooling"
	"graphics.gd/cmd/gd/internal/tooling/godot"
	"graphics.gd/cmd/gd/internal/tooling/golang"
	"graphics.gd/cmd/gd/internal/tooling/zig"

	"runtime.link/api/xray"
)

var (
	//go:embed bundled/android
	android_sdk embed.FS
)

type Android struct {
	Graphics string
}

func adb() (string, error) {
	GOPATH := os.Getenv("GOPATH")
	if GOPATH == "" {
		GOPATH = build.Default.GOPATH
	}
	// create dummy java so that godot stops complaining.
	os.WriteFile(filepath.Join(GOPATH, "bin", "java"), []byte("stub java so that godot stops complaining"), 0o664)

	adb, err := exec.LookPath("adb")
	if err == nil {
		return adb, nil
	}
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	gobin := os.Getenv("GOBIN")
	if gobin == "" {
		gobin = filepath.Join(gopath, "bin")
	}
	var arch = runtime.GOARCH
	if runtime.GOOS == "darwin" {
		arch = "universal"
	}
	if _, err = os.Stat(filepath.Join(gobin, "adb")); err == nil {
		return filepath.Join(gobin, "adb"), nil
	}
	var unzip = ""
	if runtime.GOOS == "windows" && runtime.GOARCH == "amd64" {
		unzip = "adb.windows.amd64"
	}
	fmt.Println("gd: downloading https://release.graphics.gd/adb." + runtime.GOOS + "." + arch)
	adb, err = tooling.Download(filepath.Join(gobin, "adb"), unzip, "https://release.graphics.gd/adb."+runtime.GOOS+"."+arch)
	if err != nil {
		return "", fmt.Errorf("unable to download adb for your platform, please ensure adb is installed (and in your path) to launch your app on an Android device: %w", err)
	}
	return adb, nil
}

func (Android) Build(args ...string) error {
	var GOARCH = runtime.GOARCH
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if runtime.GOOS != "android" || runtime.GOARCH != GOARCH {
		if err := zig.Assert(); err != nil {
			return xray.New(err)
		}
		if err := project.SetupFiles(android_sdk, "bundled/android", filepath.Join(project.ReleasesDirectory, "android", "sdk")); err != nil {
			return xray.New(err)
		}
		ANDROID_SDK, err := filepath.Abs(filepath.Join(project.ReleasesDirectory, "android", "sdk"))
		if err != nil {
			return xray.New(err)
		}
		switch GOARCH {
		case "arm64":
			if err := os.Setenv("CC", zig.Executable+" cc -target aarch64-linux-android -nostdlib -I"+ANDROID_SDK+"/usr/include -L"+ANDROID_SDK+"/usr/lib"); err != nil {
				return xray.New(err)
			}
		default:
			return fmt.Errorf("gd build: cannot cross-compile linux %v on %v", GOARCH, runtime.GOOS)
		}
	}
	return golang.Build(args, "-buildmode=c-shared", "-o", filepath.Join(project.GraphicsDirectory, fmt.Sprintf("libandroid_%v.so", GOARCH)))
}

func (android Android) BuildMain(...string) error {
	if err := android.Build(); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	if err := godot.Run("--headless", "--export-release", "Android"); err != nil {
		return xray.New(err)
	}
	return nil
}

func (android Android) Run(args ...string) error {
	if err := android.Build(args...); err != nil {
		return xray.New(err)
	}
	adb, err := adb()
	if err != nil {
		return xray.New(err)
	}
	if err := os.MkdirAll(filepath.Join(project.ReleasesDirectory, "android", "arm64"), 0755); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	if err := godot.Run("--headless", "--export-debug", "Android"); err != nil {
		return xray.New(err)
	}
	//  adb shell monkey -p com.example.original -c android.intent.category.LAUNCHER 1; adb logcat --pid=$(adb shell pidof com.example.original) > dump.txt
	cmd := exec.Command(adb, "install", filepath.Join(project.ReleasesDirectory, "android", "arm64", path.Base(project.Directory)+".apk"))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return xray.New(err)
	}
	cmd = exec.Command(adb, "shell", "monkey", "-p", "com.example."+project.AndroidSafePackageName(path.Base(project.Directory)), "-c", "android.intent.category.LAUNCHER", "1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return xray.New(err)
	}
	var pid []byte
	for range 3 {
		pid, err = exec.Command(adb, "shell", "pidof", "com.example."+project.AndroidSafePackageName(path.Base(project.Directory))).Output()
		if err != nil {
			continue
		}
	}
	if pid == nil {
		return nil
	}
	fmt.Println("PID=", string(pid))
	cmd = exec.Command(adb, "logcat", "--pid="+string(pid[:len(pid)-1]))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return xray.New(err)
	}
	return nil
}

func (Android) Test(args ...string) error {
	return fmt.Errorf("gd test: android not supported")
}
