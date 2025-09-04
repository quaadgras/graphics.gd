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
	return nil
}

func (Android) BuildMain(...string) error {
	return nil
}

func (Android) Run(args ...string) error {
	wd, err := os.Getwd()
	if err != nil {
		return xray.New(err)
	}
	adb, err := adb()
	if err != nil {
		return xray.New(err)
	} //  adb shell monkey -p com.example.original -c android.intent.category.LAUNCHER 1; adb logcat --pid=$(adb shell pidof com.example.original) > dump.txt
	cmd := exec.Command(adb, "install", "./releases/android/arm64/"+path.Base(wd)+".apk")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return xray.New(err)
	}
	cmd = exec.Command(adb, "shell", "monkey", "-p", "com.example."+project.AndroidSafePackageName(path.Base(wd)), "-c", "android.intent.category.LAUNCHER", "1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return xray.New(err)
	}
	var pid []byte
	for range 3 {
		pid, err = exec.Command(adb, "shell", "pidof", "com.example."+project.AndroidSafePackageName(path.Base(wd))).Output()
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
	return nil
}
