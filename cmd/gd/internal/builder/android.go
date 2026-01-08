package builder

import (
	"bytes"
	"context"
	"crypto"
	"embed"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"syscall"
	"time"

	"golang.org/x/term"
	"graphics.gd/cmd/gd/internal/cryptic"
	"graphics.gd/cmd/gd/internal/cryptic/certloader"
	"graphics.gd/cmd/gd/internal/cryptic/signjar"
	"graphics.gd/cmd/gd/internal/cryptic/zipslicer"
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

func (Android) Build(args ...string) error {
	HOME, err := os.UserHomeDir()
	if err != nil {
		return xray.New(err)
	}
	var GDPATH = os.Getenv("GDPATH")
	if GDPATH == "" {
		GDPATH = filepath.Join(os.Getenv("HOME"), "gd")
	}
	var default_sdk_path string
	switch runtime.GOOS {
	case "linux":
		default_sdk_path = filepath.Join(HOME, "Android", "Sdk")
	case "windows":
		default_sdk_path = filepath.Join(os.Getenv("LOCALAPPDATA"), "Android", "Sdk")
	case "darwin":
		default_sdk_path = filepath.Join(HOME, "Library", "Android", "Sdk")
	}
	if default_sdk_path != "" {
		if _, err := os.Stat(default_sdk_path); os.IsNotExist(err) {
			if err := os.MkdirAll(filepath.Join(default_sdk_path, "platform-tools"), 0755); err != nil {
				return xray.New(err)
			}
			if err := os.MkdirAll(filepath.Join(default_sdk_path, "build-tools", "35"), 0755); err != nil {
				return xray.New(err)
			}
			var suffix = ""
			if runtime.GOOS == "windows" {
				suffix = ".exe"
			}
			if err := os.Symlink(filepath.Join(GDPATH, "bin", "adb"+suffix), filepath.Join(default_sdk_path, "platform-tools", "adb")); err != nil {
				return xray.New(err)
			}
			suffix = ""
			if runtime.GOOS == "windows" {
				suffix = ".bat"
			}
			if err := os.Symlink(filepath.Join(GDPATH, "bin", "apksigner"), filepath.Join(default_sdk_path, "build-tools", "35", "apksigner"+suffix)); err != nil {
				return xray.New(err)
			}
		}
	}
	if !project.IncludesGo {
		return nil
	}
	var GOARCH = "arm64"
	if goarch := os.Getenv("GOARCH"); goarch != "" {
		GOARCH = goarch
	}
	if runtime.GOOS != "android" || runtime.GOARCH != GOARCH {
		zig, err := tooling.Zig.Lookup()
		if err != nil {
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
			if err := os.Setenv("CC", zig+" cc -target aarch64-linux-android -nostdlib -I"+ANDROID_SDK+"/usr/include -L"+ANDROID_SDK+"/usr/lib"); err != nil {
				return xray.New(err)
			}
			if err := os.Setenv("GOARCH", "arm64"); err != nil {
				return xray.New(err)
			}
		default:
			return fmt.Errorf("gd build: cannot cross-compile android/%v on %v", GOARCH, runtime.GOOS)
		}
	}
	return tooling.Go.Action("build", args, "-ldflags=-checklinkname=0", "-buildmode=c-shared", "-o", filepath.Join(project.GraphicsDirectory, fmt.Sprintf("libandroid_%v.so", GOARCH)))
}

func (android Android) Run(args ...string) error {
	if err := android.Build(args...); err != nil {
		return xray.New(err)
	}
	adb, err := tooling.AndroidDebugBridge.Lookup()
	if err != nil {
		return xray.New(err)
	}
	if _, err := tooling.AndroidPackageSigner.Lookup(); err != nil {
		return xray.New(err)
	}
	if err := os.MkdirAll(filepath.Join(project.ReleasesDirectory, "android", "arm64"), 0755); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	if err := tooling.Godot.Exec("--headless", "--export-debug", "Android"); err != nil {
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
		time.Sleep(time.Second / 3)
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

func (android Android) BuildMain(...string) error {
	if err := android.Build(); err != nil {
		return xray.New(err)
	}
	_, err := tooling.AndroidDebugBridge.Lookup()
	if err != nil {
		return xray.New(err)
	}
	if _, err := tooling.AndroidPackageSigner.Lookup(); err != nil {
		return xray.New(err)
	}
	GDPATH := os.Getenv("GDPATH")
	if GDPATH == "" {
		GDPATH = filepath.Join(os.Getenv("HOME"), "gd")
	}
	if err := os.WriteFile(filepath.Join(GDPATH, "bin", "java"), []byte("java stub"), 0755); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	if err := tooling.Godot.Exec("--headless", "--export-release", "Android"); err != nil {
		return xray.New(err)
	}
	// Now that we have the .apk, we also want an .aab that can be uploaded to the Play Store.
	if err := errors.Join(
		os.RemoveAll(filepath.Join(project.ReleasesDirectory, "android", "decompiled")),
		os.RemoveAll(filepath.Join(project.ReleasesDirectory, "android", "recompiled")),
		os.RemoveAll(filepath.Join(project.ReleasesDirectory, "android", "res.zip")),
		os.RemoveAll(filepath.Join(project.ReleasesDirectory, "android", "base.zip")),
		os.RemoveAll(filepath.Join(project.ReleasesDirectory, "android", "modules.zip")),
		os.RemoveAll(filepath.Join(project.ReleasesDirectory, "android", project.Name+".aab")),
	); err != nil {
		return xray.New(err)
	}
	if err := tooling.AndroidPackageKitTool.Exec("d",
		filepath.Join(project.ReleasesDirectory, "android", "arm64", project.Name+".apk"),
		"-s", "-o",
		filepath.Join(project.ReleasesDirectory, "android", "decompiled"),
		"-f",
	); err != nil {
		return xray.New(err)
	}
	if err := os.Remove(
		filepath.Join(project.ReleasesDirectory, "android", "decompiled", "res", "values-anydpi-v26", "mipmaps.xml"),
	); err != nil {
		return xray.New(err)
	}
	public, err := os.ReadFile(
		filepath.Join(project.ReleasesDirectory, "android", "decompiled", "res", "values", "public.xml"),
	)
	if err != nil {
		return xray.New(err)
	}
	public = bytes.Replace(public, []byte(`<public type="mipmap" name="themed_icon" id="0x7f0a0004" />`), nil, 1)
	if err := os.WriteFile(
		filepath.Join(project.ReleasesDirectory, "android", "decompiled", "res", "values", "public.xml"),
		public,
		0644,
	); err != nil {
		return xray.New(err)
	}
	originalPackageName, err := tooling.AndroidAssetPackagingTool.Output("dump", "packagename",
		filepath.Join(project.ReleasesDirectory, "android", "arm64", project.Name+".apk"),
	)
	if err != nil {
		return xray.New(err)
	}
	// restore intended package name
	manifest, err := os.ReadFile(
		filepath.Join(project.ReleasesDirectory, "android", "decompiled", "AndroidManifest.xml"),
	)
	if err != nil {
		return xray.New(err)
	}
	manifest = bytes.Replace(manifest,
		[]byte(`package="com.godot.game"`),
		[]byte(`package="`+originalPackageName+`"`),
		1,
	)
	manifest = bytes.Replace(manifest,
		[]byte(`android:name="com.godot.game"`),
		[]byte(`android:name="`+originalPackageName+`"`),
		1,
	)
	manifest = bytes.Replace(manifest,
		[]byte(`android:version="\1"`),
		[]byte(`android:version="1"`),
		1,
	)
	if err := os.WriteFile(
		filepath.Join(project.ReleasesDirectory, "android", "decompiled", "AndroidManifest.xml"),
		manifest,
		0644,
	); err != nil {
		return xray.New(err)
	}
	if err := tooling.AndroidAssetPackagingTool.Exec("compile", "--dir",
		filepath.Join(project.ReleasesDirectory, "android", "decompiled", "res"),
		"-o", filepath.Join(project.ReleasesDirectory, "android", "res.zip"),
	); err != nil {
		return xray.New(err)
	}
	android_jar, err := tooling.Android.Lookup()
	if err != nil {
		return xray.New(err)
	}
	export_presets, err := os.ReadFile(filepath.Join(project.GraphicsDirectory, "export_presets.cfg"))
	if err != nil {
		return xray.New(err)
	}
	var version_code_string string
	for line := range bytes.SplitSeq(export_presets, []byte("\n")) {
		if line, ok := bytes.CutPrefix(line, []byte("version/code=")); ok {
			version_code_string = string(bytes.TrimSpace(line))
		}
	}
	if version_code_string == "" {
		version_code_string = "0"
	}
	version_code, err := strconv.Atoi(version_code_string)
	if err != nil {
		return xray.New(err)
	}
	version_code++
	export_presets = bytes.ReplaceAll(export_presets, []byte("version/code="+version_code_string), []byte("version/code="+fmt.Sprint(version_code)))
	if err := os.WriteFile(filepath.Join(project.GraphicsDirectory, "export_presets.cfg"), export_presets, 0644); err != nil {
		return xray.New(err)
	}
	if err := tooling.AndroidAssetPackagingTool.Exec("link", "--proto-format", "-o",
		filepath.Join(project.ReleasesDirectory, "android", "base.zip"),
		"-I", android_jar, "--manifest",
		filepath.Join(project.ReleasesDirectory, "android", "decompiled", "AndroidManifest.xml"),
		"--min-sdk-version", "15", "--target-sdk-version", "35", "--version-code", fmt.Sprint(version_code),
		"--version-name", project.Version, "-R",
		filepath.Join(project.ReleasesDirectory, "android", "res.zip"),
		"--auto-add-overlay",
	); err != nil {
		return xray.New(err)
	}
	if err := tooling.ExtractArchive(filepath.Join(project.ReleasesDirectory, "android", "base.zip"), filepath.Join(project.ReleasesDirectory, "android", "recompiled"), "zip", "", true); err != nil {
		return xray.New(err)
	}
	if err := errors.Join(
		os.MkdirAll(filepath.Join(project.ReleasesDirectory, "android", "recompiled", "dex"), 0755),
		os.MkdirAll(filepath.Join(project.ReleasesDirectory, "android", "recompiled", "manifest"), 0755),
		os.Rename(
			filepath.Join(project.ReleasesDirectory, "android", "recompiled", "AndroidManifest.xml"),
			filepath.Join(project.ReleasesDirectory, "android", "recompiled", "manifest", "AndroidManifest.xml"),
		),
		os.Rename(
			filepath.Join(project.ReleasesDirectory, "android", "decompiled", "assets"),
			filepath.Join(project.ReleasesDirectory, "android", "recompiled", "assets"),
		),
		os.Rename(
			filepath.Join(project.ReleasesDirectory, "android", "decompiled", "lib"),
			filepath.Join(project.ReleasesDirectory, "android", "recompiled", "lib"),
		),
	); err != nil {
		return xray.New(err)
	}
	decompiled, err := os.ReadDir(filepath.Join(project.ReleasesDirectory, "android", "decompiled"))
	if err != nil {
		return xray.New(err)
	}
	for _, dex := range decompiled {
		if path.Ext(dex.Name()) == ".dex" {
			if err := os.Rename(
				filepath.Join(project.ReleasesDirectory, "android", "decompiled", dex.Name()),
				filepath.Join(project.ReleasesDirectory, "android", "recompiled", "dex", dex.Name()),
			); err != nil {
				return xray.New(err)
			}
		}
	}
	if err := tooling.CreateZip(filepath.Join(project.ReleasesDirectory, "android", "recompiled"), filepath.Join(project.ReleasesDirectory, "android", "modules.zip")); err != nil {
		return xray.New(err)
	}
	if err := tooling.BundleTool.Exec("build-bundle", "--modules="+filepath.Join(project.ReleasesDirectory, "android", "modules.zip"),
		"--output="+filepath.Join(project.ReleasesDirectory, "android", project.Name+".aab"),
	); err != nil {
		return xray.New(err)
	}
	if err := errors.Join(
		os.RemoveAll(filepath.Join(project.ReleasesDirectory, "android", "decompiled")),
		os.RemoveAll(filepath.Join(project.ReleasesDirectory, "android", "recompiled")),
		os.Remove(filepath.Join(project.ReleasesDirectory, "android", "res.zip")),
		os.Remove(filepath.Join(project.ReleasesDirectory, "android", "base.zip")),
		os.Remove(filepath.Join(project.ReleasesDirectory, "android", "modules.zip")),
	); err != nil {
		return xray.New(err)
	}
	fmt.Println("\nBuilt Version", project.Name, project.Version, "("+strconv.Itoa(version_code)+")")
	fmt.Println("\nFor the .aab to be elligible for upload to Play Console, gd can sign it with an Upload Key derived from a passphrase.")
	fmt.Println("This means, you don't need to manage any keys and as long as you use Google Play's App Signing and use the same password ")
	fmt.Println("(and project name) each build. If you forget this or change the project's name, you'll need Google to reset the Upload Key.")
	fmt.Println("\nLeave the passphrase blank if you would prefer to manage your own Upload Key / App Signing, ie. keystore/jarsigner.")
	var password string
	for {
		fmt.Print("\nProvide passphrase: ")
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			return nil
		}
		fmt.Println()
		password = string(bytePassword)
		if password == "" {
			fmt.Println("\nsigning skipped, you will need to sign the .aab yourself before uploading to Play Console")
			return nil
		}
		fmt.Print("Confirm passphrase: ")
		bytePassword, err = term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println("\nsigning skipped, you will need to sign the .aab yourself before uploading to Play Console")
			return nil
		}
		fmt.Println()
		if password != string(bytePassword) {
			fmt.Println("passphrases do not match, please try again")
			continue
		}
		break
	}
	key, cert, err := cryptic.DeterministicCertificate(password, project.Name)
	if err != nil {
		return xray.New(err)
	}
	aab, err := os.Open(filepath.Join(project.ReleasesDirectory, "android", project.Name+".aab"))
	if err != nil {
		return xray.New(err)
	}
	r, w := io.Pipe()
	go func() {
		_ = w.CloseWithError(zipslicer.ZipToTar(aab, w))
	}()
	digest, err := signjar.DigestJarStream(r, crypto.SHA256)
	if err != nil {
		return xray.New(err)
	}
	patch, _, err := digest.Sign(context.Background(), &certloader.Certificate{
		Leaf:       cert,
		PrivateKey: key,
	}, "upload", false, false, false)
	if _, err := aab.Seek(0, io.SeekStart); err != nil {
		return xray.New(err)
	}
	if err := patch.Apply(aab, filepath.Join(project.ReleasesDirectory, "android", project.Name+".aab")); err != nil {
		return xray.New(err)
	}
	return nil
}
