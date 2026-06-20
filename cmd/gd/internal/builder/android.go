package builder

import (
	"bytes"
	"context"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"embed"
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/pavlo-v-chernykh/keystore-go/v4"
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

func (android Android) Build(args ...string) error {
	return android.build(false, args...)
}

// androidTestOutputFile mirrors startup.TestOutputFile: the basename, inside the
// app's user-data dir, that `gd test` output is redirected to on-device. Read
// back here via `adb shell run-as <pkg> cat files/<this>`.
const androidTestOutputFile = "gdtest.log"

// build compiles the project as an android c-shared library. With testing set it
// builds a `go test` binary (run on-device under the engine via the FirstFrame
// hook in startup_cgo.go) instead of the application.
func (android Android) build(testing bool, args ...string) error {
	HOME, err := os.UserHomeDir()
	if err != nil {
		return xray.New(err)
	}
	var debug_keystore string
	switch runtime.GOOS {
	case "linux":
		debug_keystore = filepath.Join(os.Getenv("HOME"), ".local", "share", "godot", "keystores", "debug.keystore")
	case "windows":
		debug_keystore = filepath.Join(os.Getenv("APPDATA"), "Godot", "keystores", "debug.keystore")
	case "darwin":
		debug_keystore = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Godot", "keystores", "debug.keystore")
	default:
		return nil
	}
	if err := os.MkdirAll(filepath.Dir(debug_keystore), 0755); err != nil {
		return xray.New(err)
	}
	if _, err := os.Stat(debug_keystore); os.IsNotExist(err) {
		// Generate RSA private key
		privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
		if err != nil {
			return xray.New(err)
		}

		// Create self-signed certificate
		notBefore := time.Now()
		notAfter := notBefore.Add(10000 * 24 * time.Hour) // 10,000 days
		serialNumber, _ := rand.Int(rand.Reader, new(big.Int).Lsh(big.NewInt(1), 128))

		certTemplate := x509.Certificate{
			SerialNumber: serialNumber,
			Subject: pkix.Name{
				CommonName:   "Android Debug",
				Organization: []string{"Android"},
				Country:      []string{"US"},
			},
			NotBefore: notBefore,
			NotAfter:  notAfter,
			KeyUsage:  x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
			ExtKeyUsage: []x509.ExtKeyUsage{
				x509.ExtKeyUsageServerAuth,
				x509.ExtKeyUsageClientAuth,
			},
			BasicConstraintsValid: true,
		}

		certDER, err := x509.CreateCertificate(rand.Reader, &certTemplate, &certTemplate, &privateKey.PublicKey, privateKey)
		if err != nil {
			return xray.New(err)
		}

		// Encode private key to PKCS#8 (required for JKS)
		privateKeyDER, err := x509.MarshalPKCS8PrivateKey(privateKey)
		if err != nil {
			return xray.New(err)
		}

		// Create keystore
		ks := keystore.New()
		ks.SetPrivateKeyEntry("androiddebugkey", keystore.PrivateKeyEntry{
			PrivateKey:   privateKeyDER,
			CreationTime: time.Now(),
			CertificateChain: []keystore.Certificate{
				{
					Type:    "X.509",
					Content: certDER,
				},
			},
		}, []byte("android"))

		// Write to file
		f, err := os.OpenFile(debug_keystore, os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			return xray.New(err)
		}
		defer f.Close()

		err = ks.Store(f, []byte("android")) // Store password: "android"
		if err != nil {
			return xray.New(err)
		}
	}

	var GDPATH = os.Getenv("GDPATH")
	if GDPATH == "" {
		GDPATH = filepath.Join(HOME, "gd")
	}
	var exe string
	if runtime.GOOS == "windows" {
		exe = ".exe"
	}
	if err := os.WriteFile(filepath.Join(GDPATH, "bin", "java"+exe), []byte("java stub"), 0755); err != nil {
		return xray.New(err)
	}
	var default_sdk_path string
	switch runtime.GOOS {
	case "linux":
		default_sdk_path = filepath.Join(HOME, "Android", "Sdk")
	case "windows":
		default_sdk_path = filepath.Join(os.Getenv("LOCALAPPDATA"), "Android", "Sdk")
		if _, err := tooling.AndroidDebugBridge.Lookup(); err != nil {
			return xray.New(err)
		}
		if _, err := tooling.AndroidPackageSigner.Lookup(); err != nil {
			return xray.New(err)
		}
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
			if runtime.GOOS == "windows" {
				if err := project.CopyFile(filepath.Join(GDPATH, "bin", "AdbWinApi.dll"), filepath.Join(default_sdk_path, "platform-tools", "AdbWinApi.dll")); err != nil {
					return xray.New(err)
				}
				if err := project.CopyFile(filepath.Join(GDPATH, "bin", "AdbWinUsbApi.dll"), filepath.Join(default_sdk_path, "platform-tools", "AdbWinUsbApi.dll")); err != nil {
					return xray.New(err)
				}
				if err := project.CopyFile(filepath.Join(GDPATH, "bin", "adb.exe"), filepath.Join(default_sdk_path, "platform-tools", "adb.exe")); err != nil {
					return xray.New(err)
				}
			} else {
				if err := os.Symlink(filepath.Join(GDPATH, "bin", "adb"), filepath.Join(default_sdk_path, "platform-tools", "adb")); err != nil {
					return xray.New(err)
				}
			}
			if runtime.GOOS == "windows" {
				if err := project.CopyFile(filepath.Join(GDPATH, "bin", "apksigner.exe"), filepath.Join(default_sdk_path, "build-tools", "35", "apksigner.bat")); err != nil {
					return xray.New(err)
				}
			} else {
				if err := os.Symlink(filepath.Join(GDPATH, "bin", "apksigner"), filepath.Join(default_sdk_path, "build-tools", "35", "apksigner")); err != nil {
					return xray.New(err)
				}
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
		case "amd64":
			// The bundled liblog.so stub is aarch64; for x86_64 we
			// compile the same set of no-op shims from liblog.c into
			// a fresh stub the linker can resolve `-llog` against.
			// The dynamic linker substitutes the device's real
			// liblog.so at runtime. -nostdlib is required because
			// zig 0.15 doesn't ship a bundled libc for the
			// x86_64-linux-android target.
			liblog := filepath.Join(ANDROID_SDK, "usr", "lib", "liblog.so")
			liblogSrc := filepath.Join(ANDROID_SDK, "usr", "lib", "liblog.c")
			if err := exec.Command(zig, "cc", "-target", "x86_64-linux-android", "-shared", "-nostdlib",
				"-Wl,-soname,liblog.so", "-o", liblog, liblogSrc).Run(); err != nil {
				return xray.New(fmt.Errorf("build liblog stub for amd64: %w", err))
			}
			if err := os.Setenv("CC", zig+" cc -target x86_64-linux-android -nostdlib -I"+ANDROID_SDK+"/usr/include -L"+ANDROID_SDK+"/usr/lib"); err != nil {
				return xray.New(err)
			}
			if err := os.Setenv("GOARCH", "amd64"); err != nil {
				return xray.New(err)
			}
		default:
			return fmt.Errorf("gd build: cannot cross-compile android/%v on %v", GOARCH, runtime.GOOS)
		}
	}
	out := filepath.Join(project.GraphicsDirectory, fmt.Sprintf("libandroid_%v.so", GOARCH))
	if testing {
		return tooling.Go.Action("test", args, "-c", "-ldflags=-checklinkname=0", "-buildmode=c-shared", "-o", out)
	}
	return tooling.Go.Action("build", args, "-ldflags=-checklinkname=0", "-buildmode=c-shared", "-o", out)
}

func (android Android) Run(args ...string) error {
	var debug_keystore string
	switch runtime.GOOS {
	case "linux":
		debug_keystore = filepath.Join(os.Getenv("HOME"), ".local", "share", "godot", "keystores", "debug.keystore")
	case "windows":
		debug_keystore = filepath.Join(os.Getenv("APPDATA"), "Godot", "keystores", "debug.keystore")
	case "darwin":
		debug_keystore = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Godot", "keystores", "debug.keystore")
	default:
		return nil
	}
	if err := android.Build(args...); err != nil {
		return xray.New(err)
	}
	GOARCH := "arm64"
	if env := os.Getenv("GOARCH"); env != "" {
		GOARCH = env
	}
	adb, err := tooling.AndroidDebugBridge.Lookup()
	if err != nil {
		return xray.New(err)
	}
	if _, err := tooling.AndroidPackageSigner.Lookup(); err != nil {
		return xray.New(err)
	}
	presetName, exportPath, err := pickAndroidPreset(GOARCH)
	if err != nil {
		return xray.New(err)
	}
	apkPath := filepath.Join(project.GraphicsDirectory, exportPath)
	if err := os.MkdirAll(filepath.Dir(apkPath), 0755); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	if err := tooling.Godot.Exec("--headless", "--export-debug", presetName); err != nil {
		return xray.New(err)
	}
	if err := tooling.AndroidPackageSigner.Exec(
		"sign", "--ks", debug_keystore,
		"--ks-key-alias", "androiddebugkey", "--ks-pass", "pass:android",
		apkPath,
	); err != nil {
		return xray.New(err)
	}
	//  adb shell monkey -p com.example.original -c android.intent.category.LAUNCHER 1; adb logcat --pid=$(adb shell pidof com.example.original) > dump.txt
	cmd := exec.Command(adb, "install", apkPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Println("Device not recognized? Make sure developer mode is enabled:")
		fmt.Println("	(go to Settings > About Phone, find the Build Number, and tap it 7 times quickly).")
		fmt.Println("Also make sure to unlock your device and accept any USB debugging prompts!")
		return xray.New(err)
	}
	// Resolve the real package name from the APK manifest instead of
	// reconstructing "com.example.<dir>" — the user may have set a
	// custom package/unique_name in the export preset and the
	// hardcoded form would only match by accident.
	pkgOut, err := tooling.AndroidAssetPackagingTool.Output("dump", "packagename", apkPath)
	if err != nil {
		return xray.New(err)
	}
	packageName := strings.TrimSpace(pkgOut)
	// Clear the log buffer so any post-launch dump only shows this run's output.
	_ = exec.Command(adb, "logcat", "-c").Run()
	cmd = exec.Command(adb, "shell", "monkey", "-p", packageName, "-c", "android.intent.category.LAUNCHER", "1")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return xray.New(err)
	}
	var pid []byte
	for range 10 {
		out, err := exec.Command(adb, "shell", "pidof", packageName).Output()
		if err == nil {
			if trimmed := bytes.TrimSpace(out); len(trimmed) > 0 {
				pid = trimmed
				break
			}
		}
		time.Sleep(time.Second / 3)
	}
	if len(pid) == 0 {
		fmt.Fprintf(os.Stderr, "%s did not start. Recent device error logs:\n", packageName)
		dump := exec.Command(adb, "logcat", "-d", "-t", "200", "*:E")
		dump.Stdout = os.Stderr
		dump.Stderr = os.Stderr
		_ = dump.Run()
		return fmt.Errorf("gd run: %s failed to launch", packageName)
	}
	fmt.Println("PID=", string(pid))
	cmd = exec.Command(adb, "logcat", "--pid="+string(pid))
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return xray.New(err)
	}
	return nil
}

// Test builds the suite as a c-shared android library, deploys it to the
// connected device/emulator, runs it under the engine, and reads the result
// back from the app's user-data dir.
//
// UNVERIFIED: this has not yet run on a real emulator (the dev host has no KVM).
// Known first-run risks to shake out: (1) whether Godot's user:// resolves to
// the app's *internal* files dir (so `run-as ... cat files/...` works) vs an
// external dir; (2) the android export template must be installed
// (`gd build android` once, or AssertExportTemplates); (3) -run/-v passthrough
// is not wired — the whole suite runs (see startup_android.go), like web needed
// the wasm_exec argv patch.
func (android Android) Test(args ...string) error {
	var debug_keystore string
	switch runtime.GOOS {
	case "linux":
		debug_keystore = filepath.Join(os.Getenv("HOME"), ".local", "share", "godot", "keystores", "debug.keystore")
	case "windows":
		debug_keystore = filepath.Join(os.Getenv("APPDATA"), "Godot", "keystores", "debug.keystore")
	case "darwin":
		debug_keystore = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Godot", "keystores", "debug.keystore")
	default:
		return fmt.Errorf("gd test: android not supported on %v", runtime.GOOS)
	}
	// Pass the test flags straight to `go test -c`; they are consumed at compile
	// time and are harmless. The built c-shared lib can't receive them as argv
	// on-device anyway, so startup_android.go resets to a clean -test.v
	// invocation and the whole suite runs (per-test -run/-v passthrough on
	// android is a follow-up).
	if err := android.build(true, args...); err != nil {
		return xray.New(err)
	}
	GOARCH := "arm64"
	if env := os.Getenv("GOARCH"); env != "" {
		GOARCH = env
	}
	adb, err := tooling.AndroidDebugBridge.Lookup()
	if err != nil {
		return xray.New(err)
	}
	if _, err := tooling.AndroidPackageSigner.Lookup(); err != nil {
		return xray.New(err)
	}
	presetName, exportPath, err := pickAndroidPreset(GOARCH)
	if err != nil {
		return xray.New(err)
	}
	apkPath := filepath.Join(project.GraphicsDirectory, exportPath)
	if err := os.MkdirAll(filepath.Dir(apkPath), 0755); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	if err := tooling.Godot.Exec("--headless", "--export-debug", presetName); err != nil {
		return xray.New(err)
	}
	if err := tooling.AndroidPackageSigner.Exec(
		"sign", "--ks", debug_keystore,
		"--ks-key-alias", "androiddebugkey", "--ks-pass", "pass:android",
		apkPath,
	); err != nil {
		return xray.New(err)
	}
	pkgOut, err := tooling.AndroidAssetPackagingTool.Output("dump", "packagename", apkPath)
	if err != nil {
		return xray.New(err)
	}
	packageName := strings.TrimSpace(pkgOut)
	if out, err := exec.Command(adb, "install", "-r", apkPath).CombinedOutput(); err != nil {
		return xray.New(fmt.Errorf("adb install: %w\n%s", err, out))
	}
	resultPath := "files/" + androidTestOutputFile // relative to /data/data/<pkg> (run-as home)
	_ = exec.Command(adb, "shell", "run-as", packageName, "rm", "-f", resultPath).Run()
	if out, err := exec.Command(adb, "shell", "monkey", "-p", packageName, "-c", "android.intent.category.LAUNCHER", "1").CombinedOutput(); err != nil {
		return xray.New(fmt.Errorf("launch %s: %w\n%s", packageName, err, out))
	}
	// Poll the on-device result file until go test prints its terminal PASS/FAIL
	// line, the app dies, or we time out.
	deadline := time.Now().Add(5 * time.Minute)
	var last string
	for time.Now().Before(deadline) {
		out, _ := exec.Command(adb, "shell", "run-as", packageName, "cat", resultPath).Output()
		last = string(out)
		if passed, done := classifyGoTest(last); done {
			fmt.Print(last)
			if passed {
				return nil
			}
			return fmt.Errorf("gd test: android suite failed")
		}
		pid, _ := exec.Command(adb, "shell", "pidof", packageName).Output()
		if len(bytes.TrimSpace(pid)) == 0 && strings.TrimSpace(last) != "" {
			fmt.Print(last)
			return fmt.Errorf("gd test: android app exited before the suite finished")
		}
		time.Sleep(time.Second)
	}
	fmt.Print(last)
	return fmt.Errorf("gd test: android suite timed out after 5m")
}

// classifyGoTest scans go test output for its terminal verdict. done stays false
// until a "PASS"/"FAIL" line appears (go test prints one as its final line).
func classifyGoTest(s string) (passed, done bool) {
	for _, line := range strings.Split(s, "\n") {
		switch strings.TrimSpace(line) {
		case "PASS":
			return true, true
		case "FAIL":
			return false, true
		}
	}
	return false, false
}

func (android Android) BuildMain(...string) error {
	if err := android.Build(); err != nil {
		return xray.New(err)
	}
	GOARCH := "arm64"
	if env := os.Getenv("GOARCH"); env != "" {
		GOARCH = env
	}
	_, err := tooling.AndroidDebugBridge.Lookup()
	if err != nil {
		return xray.New(err)
	}
	if _, err := tooling.AndroidPackageSigner.Lookup(); err != nil {
		return xray.New(err)
	}
	my, err := user.Current()
	if err != nil {
		return xray.New(err)
	}
	HOME := my.HomeDir
	GDPATH := os.Getenv("GDPATH")
	if GDPATH == "" {
		GDPATH = filepath.Join(HOME, "gd")
	}
	var exe string
	if runtime.GOOS == "windows" {
		exe = ".exe"
	}
	if err := os.WriteFile(filepath.Join(GDPATH, "bin", "java"+exe), []byte("java stub"), 0755); err != nil {
		return xray.New(err)
	}
	presetName, exportPath, err := pickAndroidPreset(GOARCH)
	if err != nil {
		return xray.New(err)
	}
	apkPath := filepath.Join(project.GraphicsDirectory, exportPath)
	if err := os.MkdirAll(filepath.Dir(apkPath), 0755); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	if err := tooling.Godot.Exec("--headless", "--export-release", presetName); err != nil {
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
		apkPath,
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
		apkPath,
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
		if filepath.Ext(dex.Name()) == ".dex" {
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

// pickAndroidPreset chooses the Godot export preset for the current
// Android target architecture. Selection order:
//
//  1. GD_ANDROID_PRESET env var (explicit override).
//  2. The default preset for this arch: "Android arm64-v8a" or
//     "Android x86_64".
//  3. The first preset whose platform="Android" has the matching
//     architectures/<abi>=true. Lets users rename or hand-craft.
//
// Returns the preset name (passed to godot --export-*) and the
// project-relative export_path declared by that preset.
func pickAndroidPreset(GOARCH string) (name, exportPath string, err error) {
	abi := "arm64-v8a"
	if GOARCH == "amd64" {
		abi = "x86_64"
	}
	presets, err := loadAndroidPresets()
	if err != nil {
		return "", "", xray.New(err)
	}
	if want := os.Getenv("GD_ANDROID_PRESET"); want != "" {
		for _, p := range presets {
			if p.name == want {
				return p.name, p.exportPath, nil
			}
		}
		return "", "", fmt.Errorf("gd build: GD_ANDROID_PRESET=%q not found in graphics/export_presets.cfg", want)
	}
	want := "Android " + abi
	for _, p := range presets {
		if p.name == want {
			return p.name, p.exportPath, nil
		}
	}
	for _, p := range presets {
		if p.platform == "Android" && p.archs[abi] {
			return p.name, p.exportPath, nil
		}
	}
	return "", "", fmt.Errorf("gd build: no Android preset for %s in graphics/export_presets.cfg", abi)
}

type androidPreset struct {
	name, platform, exportPath string
	archs                      map[string]bool
}

func loadAndroidPresets() ([]androidPreset, error) {
	raw, err := os.ReadFile(filepath.Join(project.GraphicsDirectory, "export_presets.cfg"))
	if err != nil {
		return nil, err
	}
	var (
		out []androidPreset
		cur *androidPreset
	)
	for _, line := range bytes.Split(raw, []byte("\n")) {
		s := strings.TrimSpace(string(line))
		if strings.HasPrefix(s, "[") && strings.HasSuffix(s, "]") {
			if strings.HasPrefix(s, "[preset.") && !strings.HasSuffix(s, ".options]") {
				out = append(out, androidPreset{archs: map[string]bool{}})
				cur = &out[len(out)-1]
			}
			continue
		}
		if cur == nil {
			continue
		}
		key, val, ok := strings.Cut(s, "=")
		if !ok {
			continue
		}
		val = strings.Trim(val, `"`)
		switch key {
		case "name":
			cur.name = val
		case "platform":
			cur.platform = val
		case "export_path":
			cur.exportPath = val
		default:
			if abi, ok := strings.CutPrefix(key, "architectures/"); ok && val == "true" {
				cur.archs[abi] = true
			}
		}
	}
	return out, nil
}

