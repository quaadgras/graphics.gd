package builder

import (
	"bufio"
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
	"sync"
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

// build compiles the project as an android c-shared library. With testing set it
// builds a `go test` binary (run on-device under the engine via the FirstFrame
// hook in startup_cgo.go) instead of the application.
func (android Android) build(testing bool, args ...string) error {
	var debug_keystore string
	switch runtime.GOOS {
	case "linux":
		debug_keystore = filepath.Join(os.Getenv("HOME"), ".local", "share", "godot", "keystores", "debug.keystore")
	case "windows":
		debug_keystore = filepath.Join(os.Getenv("APPDATA"), "Godot", "keystores", "debug.keystore")
	case "darwin":
		debug_keystore = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "Godot", "keystores", "debug.keystore")
	case "android":
		// Building on-device (e.g. under Termux): the Godot Android Editor
		// app opens the compiled extension directly, so no signing keystore,
		// java stub, or SDK scaffolding is needed — skip straight to the
		// compile.
	default:
		return nil
	}
	if debug_keystore != "" {
		if err := setupHostExportTools(debug_keystore); err != nil {
			return xray.New(err)
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
		var target string
		switch GOARCH {
		case "arm64":
			target = "aarch64-linux-android"
		case "amd64":
			target = "x86_64-linux-android"
		default:
			return fmt.Errorf("gd build: cannot cross-compile android/%v on %v", GOARCH, runtime.GOOS)
		}
		// Stub libraries for `-l` flags naming libraries that only exist
		// on-device: with -nostdlib zig has nothing to resolve -lm or
		// -lpthread against (zig 0.15 ships no bundled libc for android
		// targets), so compile stubs from the bundled sources for the
		// linker to find. See the .c files for why they stay empty.
		buildStub := func(name string) error {
			args := append([]string{"cc", "-target", target, "-shared", "-nostdlib"}, tooling.CGOCFlags()...)
			args = append(args,
				"-Wl,-soname,"+name+".so",
				"-o", filepath.Join(ANDROID_SDK, "usr", "lib", name+".so"),
				filepath.Join(ANDROID_SDK, "usr", "lib", name+".c"),
			)
			if err := exec.Command(zig, args...).Run(); err != nil {
				return fmt.Errorf("build %s stub for %s: %w", name, GOARCH, err)
			}
			return nil
		}
		if err := buildStub("libm"); err != nil {
			return xray.New(err)
		}
		if err := buildStub("libpthread"); err != nil {
			return xray.New(err)
		}
		if GOARCH != "arm64" {
			// The bundled liblog.so (no-op shims the dynamic linker
			// substitutes with the device's real liblog.so at runtime)
			// is prebuilt for aarch64 only; rebuild it from source for
			// other targets.
			if err := buildStub("liblog"); err != nil {
				return xray.New(err)
			}
		}
		if err := os.Setenv("CC", zig+" cc -target "+target+" -nostdlib -I"+ANDROID_SDK+"/usr/include -L"+ANDROID_SDK+"/usr/lib"); err != nil {
			return xray.New(err)
		}
		if err := os.Setenv("GOARCH", GOARCH); err != nil {
			return xray.New(err)
		}
	}
	out := filepath.Join(project.GraphicsDirectory, fmt.Sprintf("libandroid_%v.so", GOARCH))
	if testing {
		return tooling.Go.Action("test", args, append(fastcbFlags("android", ""), "-c", "-ldflags=-checklinkname=0", "-buildmode=c-shared", "-o", out)...)
	}
	return tooling.Go.Action("build", args, append(fastcbFlags("android", ""), "-ldflags=-checklinkname=0", "-buildmode=c-shared", "-o", out)...)
}

// setupHostExportTools prepares everything godot's android export needs on a
// desktop host: a debug signing keystore, a java stub, and a fake Android SDK
// layout pointing at gd-managed tools.
func setupHostExportTools(debug_keystore string) error {
	HOME, err := os.UserHomeDir()
	if err != nil {
		return xray.New(err)
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
	// On-device (Termux) nothing else has created GDPATH/bin yet.
	if err := os.MkdirAll(filepath.Join(GDPATH, "bin"), 0755); err != nil {
		return xray.New(err)
	}
	if err := os.WriteFile(filepath.Join(GDPATH, "bin", "java"+exe), []byte("java stub"), 0755); err != nil {
		return xray.New(err)
	}
	var default_sdk_path string
	switch runtime.GOOS {
	case "linux", "android":
		// android: the exporting editor on-device is a linuxbsd build, so it
		// reads the linux default SDK location under Termux's HOME.
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
		// On-device (Termux) GDPATH holds no desktop adb/apksigner to
		// point at; link the Termux packages when installed, otherwise a
		// stub — the unsigned template export only checks these exist.
		fauxTool := func(name, dest string) error {
			if runtime.GOOS != "android" {
				return os.Symlink(filepath.Join(GDPATH, "bin", name), dest)
			}
			if path, err := exec.LookPath(name); err == nil {
				return os.Symlink(path, dest)
			}
			return os.WriteFile(dest, []byte(name+" stub"), 0755)
		}
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
				if err := fauxTool("adb", filepath.Join(default_sdk_path, "platform-tools", "adb")); err != nil {
					return xray.New(err)
				}
			}
			if runtime.GOOS == "windows" {
				if err := project.CopyFile(filepath.Join(GDPATH, "bin", "apksigner.exe"), filepath.Join(default_sdk_path, "build-tools", "35", "apksigner.bat")); err != nil {
					return xray.New(err)
				}
			} else {
				if err := fauxTool("apksigner", filepath.Join(default_sdk_path, "build-tools", "35", "apksigner")); err != nil {
					return xray.New(err)
				}
			}
		}
		// On-device the faux SDK entries point at Termux packages, which
		// come and go with pkg install/uninstall — refresh them every run so
		// a tool installed after the SDK was first created replaces its
		// stub, and a removed one does not linger as a dangling symlink for
		// godot's export to trip over.
		if runtime.GOOS == "android" {
			for name, dest := range map[string]string{
				"adb":       filepath.Join(default_sdk_path, "platform-tools", "adb"),
				"apksigner": filepath.Join(default_sdk_path, "build-tools", "35", "apksigner"),
			} {
				if err := os.MkdirAll(filepath.Dir(dest), 0755); err != nil {
					return xray.New(err)
				}
				os.Remove(dest)
				if err := fauxTool(name, dest); err != nil {
					return xray.New(err)
				}
			}
		}
	}
	return nil
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
	case "android":
		return android.runOnDevice(args...)
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
	if err := ensureProjectIcon(); err != nil {
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
	// Run the test app with --headless so it doesn't depend on a GPU/render loop:
	// the test scheduler is driven by the engine's per-frame main loop, which
	// crawls on the software GL of a CI emulator. Restore the preset afterwards so
	// a real `gd build` for the same project is unaffected.
	restoreHeadless, err := bakeAndroidHeadless(presetName)
	if err != nil {
		return xray.New(err)
	}
	defer restoreHeadless()
	apkPath := filepath.Join(project.GraphicsDirectory, exportPath)
	if err := os.MkdirAll(filepath.Dir(apkPath), 0755); err != nil {
		return xray.New(err)
	}
	if err := ensureProjectIcon(); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	// Release export: a debug export needs Godot to reach a path the local musl
	// editor (a Go test binary) can't, and release is debuggable-independent
	// since we read results from logcat rather than via run-as.
	if err := tooling.Godot.Exec("--headless", "--export-release", presetName); err != nil {
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
	// The test binary routes its stdout to logcat under the Go runtime's "Go"
	// tag (see startup_android.go); clear the buffer, launch, then scrape it.
	// Resolve the launcher activity and start it with `am start`. monkey's exit
	// code is unreliable for a test app that exits quickly (it reports non-zero
	// when the app it's monitoring goes away). Retry the resolve since right
	// after install the package manager may not have it ready yet.
	var activity string
	for i := 0; i < 12 && activity == ""; i++ {
		out, _ := exec.Command(adb, "shell", "cmd", "package", "resolve-activity", "--brief", "-c", "android.intent.category.LAUNCHER", packageName).Output()
		for _, ln := range strings.Split(strings.TrimSpace(string(out)), "\n") {
			if ln = strings.TrimSpace(ln); strings.HasPrefix(ln, packageName+"/") {
				activity = ln
			}
		}
		if activity == "" {
			time.Sleep(2 * time.Second)
		}
	}
	if activity == "" {
		return xray.New(fmt.Errorf("could not resolve launcher activity for %s", packageName))
	}
	// Stop any prior instance android may still be relaunching so it can't flood
	// the log buffer we are about to clear and read.
	_ = exec.Command(adb, "shell", "am", "force-stop", packageName).Run()
	_ = exec.Command(adb, "logcat", "-G", "16M").Run() // headroom for a suite that relaunches for minutes
	_ = exec.Command(adb, "logcat", "-b", "all", "-c").Run()
	// Follow logcat for the whole run rather than dumping it at the end. The
	// device buffer is small and logd prunes chatty UIDs, so a suite that
	// relaunches for eight minutes loses its earliest output — which is where a
	// crash report is. That cost us the goroutine stack of a real failure: the
	// panic line survived to the end of the run, every frame under it did not.
	// `-v raw` strips the logcat prefix so each line is the raw test output.
	// android relaunches the app after os.Exit, so the suite re-runs; force-stop
	// it on the way out. Watch for the TestMain completion sentinel — a clean
	// logd line that, unlike the piped test output, never interleaves.
	defer func() { _ = exec.Command(adb, "shell", "am", "force-stop", packageName).Run() }()
	logs, stopFollowing, err := followAndroidLog(adb)
	if err != nil {
		return xray.New(err)
	}
	defer stopFollowing()
	if out, err := exec.Command(adb, "shell", "am", "start", "-n", activity).CombinedOutput(); err != nil {
		return xray.New(fmt.Errorf("am start %s: %w\n%s", activity, err, out))
	}
	deadline := time.Now().Add(8 * time.Minute)
	var last string
	for time.Now().Before(deadline) {
		last = logs()
		if code, ok := lastSentinel(last); ok {
			crashed := printAndroidResults(last)
			if code == 0 {
				if crashed {
					// lastSentinel takes the LAST verdict, and android relaunches
					// the app after a crash, so a run that crashed and then passed
					// on the relaunch reported success — hiding the crash above,
					// and with it however often this really happens.
					return fmt.Errorf("gd test: android suite crashed, then passed when android relaunched it (see the crash above)")
				}
				return nil
			}
			return fmt.Errorf("gd test: android suite failed (exit code %d)", code)
		}
		time.Sleep(time.Second)
	}
	printAndroidResults(last)
	return fmt.Errorf("gd test: android suite did not finish within the timeout")
}

// followAndroidLog starts tailing the test app's logcat output and returns a
// function giving everything captured so far, plus one to stop the tail. The
// output is accumulated on this side so it survives the device evicting it.
func followAndroidLog(adb string) (read func() string, stop func(), err error) {
	tail := exec.Command(adb, "logcat", "-s", "Go:E", "-v", "raw")
	out, err := tail.StdoutPipe()
	if err != nil {
		return nil, nil, xray.New(err)
	}
	if err := tail.Start(); err != nil {
		return nil, nil, xray.New(err)
	}
	var (
		mutex sync.Mutex
		lines strings.Builder
	)
	go func() {
		scanner := bufio.NewScanner(out)
		scanner.Buffer(make([]byte, 64*1024), 4*1024*1024) // stack frames can be long
		for scanner.Scan() {
			mutex.Lock()
			lines.WriteString(scanner.Text())
			lines.WriteByte('\n')
			mutex.Unlock()
		}
	}()
	return func() string {
			mutex.Lock()
			defer mutex.Unlock()
			return lines.String()
		}, func() {
			_ = tail.Process.Kill()
			_ = tail.Wait()
		}, nil
}

// printAndroidResults prints each distinct go test result line once, and the
// first crash report in full; it reports whether it saw a crash. The system
// relaunches the app after it exits, so by the time we read the verdict the log
// contains the suite repeated many times. Failure details (the indented
// "foo_test.go:12: ..." assertion lines and panics) are kept, or a --- FAIL
// verdict is impossible to act on. The first panic's crash output (goroutine
// stacks) is printed in full: those lines match none of the result patterns,
// and a panic verdict without its stack is impossible to act on too.
//
// The dying process and the relaunched one write to the same log buffer, so the
// crash report has the next run's output interleaved into it. Ending the report
// at the next result line therefore truncated it to nothing on exactly the runs
// that needed it (the emulator interleaves where a local device does not) —
// result lines are skipped instead, and the report ends at the run sentinel or
// a line budget.
func printAndroidResults(log string) (crashed bool) {
	report, crashed := androidResults(log)
	fmt.Print(report)
	return crashed
}

// androidResults renders what [printAndroidResults] prints, so the parsing can
// be tested without a device.
func androidResults(log string) (report string, crashed bool) {
	const crashReportMaxLines = 120
	var out strings.Builder
	seen := make(map[string]bool)
	crash, crashPrinted, crashLines := false, false, 0
	for _, line := range strings.Split(log, "\n") {
		t := strings.TrimSpace(line)
		if crash {
			switch {
			case strings.HasPrefix(t, "GDTEST_DONE") || crashLines >= crashReportMaxLines:
				crash, crashPrinted = false, true
			case isTestResultLine(t):
				// The relaunched suite, spliced into the report: not crash
				// output, but still a result — fall through and report it once,
				// leaving the crash report open around it.
			default:
				fmt.Fprintln(&out, line)
				crashLines++
				continue
			}
		}
		if !crashPrinted && (strings.HasPrefix(t, "panic:") || strings.HasPrefix(t, "fatal error:")) {
			crash, crashed = true, true
			fmt.Fprintln(&out, t)
			continue
		}
		result := strings.HasPrefix(t, "--- PASS") || strings.HasPrefix(t, "--- FAIL")
		detail := strings.Contains(t, "_test.go:") || strings.HasPrefix(t, "panic:")
		if (result || detail) && !seen[t] {
			seen[t] = true
			fmt.Fprintln(&out, t)
		}
	}
	return out.String(), crashed
}

// isTestResultLine reports whether the line is go test's own progress output,
// as opposed to something a crashing process wrote.
func isTestResultLine(line string) bool {
	for _, prefix := range []string{"=== RUN", "=== PAUSE", "=== CONT", "--- PASS", "--- FAIL", "--- SKIP", "GDTEST_DONE"} {
		if strings.HasPrefix(line, prefix) {
			return true
		}
	}
	return line == "PASS" || line == "FAIL"
}

// bakeAndroidHeadless sets command_line/extra_args="--headless" on the named
// export preset so the exported test app runs without rendering, and returns a
// function that restores the original config (so a normal `gd build` is
// unaffected).
func bakeAndroidHeadless(presetName string) (restore func(), err error) {
	cfgPath := filepath.Join(project.GraphicsDirectory, "export_presets.cfg")
	original, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}
	// Find the [preset.N] whose name matches, then set the cmdline in its
	// [preset.N.options] section.
	lines := strings.Split(string(original), "\n")
	idx, cur := "", ""
	for _, line := range lines {
		s := strings.TrimSpace(line)
		if strings.HasPrefix(s, "[preset.") && !strings.HasSuffix(s, ".options]") {
			cur = strings.TrimSuffix(strings.TrimPrefix(s, "[preset."), "]")
		} else if name, ok := strings.CutPrefix(s, "name="); ok && strings.Trim(name, `"`) == presetName {
			idx = cur
			break
		}
	}
	if idx == "" {
		return nil, fmt.Errorf("preset %q not found in %s", presetName, cfgPath)
	}
	optionsHeader := "[preset." + idx + ".options]"
	inOptions, set := false, false
	for i, line := range lines {
		s := strings.TrimSpace(line)
		if strings.HasPrefix(s, "[") {
			inOptions = s == optionsHeader
			continue
		}
		if inOptions && strings.HasPrefix(s, "command_line/extra_args=") {
			lines[i] = `command_line/extra_args="--headless"`
			set = true
			break
		}
	}
	if !set {
		return nil, fmt.Errorf("command_line/extra_args not found for preset %q", presetName)
	}
	if err := os.WriteFile(cfgPath, []byte(strings.Join(lines, "\n")), 0o644); err != nil {
		return nil, err
	}
	return func() { _ = os.WriteFile(cfgPath, original, 0o644) }, nil
}

// lastSentinel returns the exit code from the last "GDTEST_DONE <code>" line
// emitted by the test binary's TestMain (see internal/main_android_test.go).
func lastSentinel(s string) (code int, ok bool) {
	for _, line := range strings.Split(s, "\n") {
		if rest, found := strings.CutPrefix(strings.TrimSpace(line), "GDTEST_DONE "); found {
			if n, err := strconv.Atoi(strings.TrimSpace(rest)); err == nil {
				code, ok = n, true
			}
		}
	}
	return code, ok
}

func (android Android) BuildMain(...string) error {
	if runtime.GOOS == "android" {
		return android.buildMainOnDevice()
	}
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
	if err := ensureProjectIcon(); err != nil {
		return xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return xray.New(err)
	}
	if err := tooling.Godot.Exec("--headless", "--export-release", presetName); err != nil {
		return xray.New(err)
	}
	return android.packageAab(apkPath)
}

// packageAab converts the exported .apk into an .aab that can be uploaded to
// the Play Store, offering to sign it with a passphrase-derived upload key.
func (android Android) packageAab(apkPath string) error {
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

// buildMainOnDevice exports the android APK on-device (Termux). The static
// musl editor built as tooling by gd's main flow performs the headless
// export; the desktop .aab/Play-Store pipeline is skipped because its java
// tooling (apktool/aapt2/bundletool) does not run on bionic.
func (android Android) buildMainOnDevice() error {
	apkPath, signed, err := android.exportOnDevice()
	if err != nil {
		return xray.New(err)
	}
	if signed {
		fmt.Println("gd: built and debug-signed", apkPath)
	} else {
		fmt.Println("gd: built", apkPath)
		fmt.Println("gd: the APK is unsigned and will not install — run 'pkg install apksigner' and rebuild to have gd debug-sign it")
	}
	if err := setupOnDeviceAabTools(); err != nil {
		return xray.New(err)
	}
	return android.packageAab(apkPath)
}

// setupOnDeviceAabTools points the aab pipeline's toolchain entries at what
// can actually run on bionic: Termux packages for the native pieces (aapt2,
// and a JRE to run the jars) and the stock apktool/bundletool jars from the
// release bucket — the desktop downloads are GraalVM native-image
// compilations of those same jars, and native-image cannot target bionic.
func setupOnDeviceAabTools() error {
	if _, err := exec.LookPath("java"); err != nil {
		if err := termuxPkgInstall("openjdk-17"); err != nil {
			return fmt.Errorf("java is required to build the .aab on-device: %w", err)
		}
	}
	if _, err := exec.LookPath("aapt2"); err != nil {
		if err := termuxPkgInstall("aapt2"); err != nil {
			return fmt.Errorf("aapt2 is required to build the .aab on-device: %w", err)
		}
	}
	aapt2, err := exec.LookPath("aapt2")
	if err != nil {
		return xray.New(err)
	}
	tooling.AndroidAssetPackagingTool.Path = aapt2
	tooling.AndroidPackageKitTool.Name = "apktool.jar"
	tooling.AndroidPackageKitTool.DownloadURL = "https://release.graphics.gd/apktool.jar"
	tooling.AndroidPackageKitTool.Jar = true
	tooling.BundleTool.Name = "bundletool.jar"
	tooling.BundleTool.DownloadURL = "https://release.graphics.gd/bundletool.jar"
	tooling.BundleTool.Jar = true
	return nil
}

// exportOnDevice compiles the extension, exports the APK with the static
// musl editor and debug-signs it. The preset exports the APK unsigned
// (package/signed=false), so signing happens afterwards when a Termux
// apksigner is installed; signed reports whether it was.
func (android Android) exportOnDevice(args ...string) (apkPath string, signed bool, err error) {
	if err := android.Build(args...); err != nil {
		return "", false, xray.New(err)
	}
	HOME, err := os.UserHomeDir()
	if err != nil {
		return "", false, xray.New(err)
	}
	// Install apksigner before anything references it: godot itself signs
	// during the export when the preset asks for it (package/signed defaults
	// to true in godot, and pre-existing presets carry that), and
	// setupHostExportTools links the faux SDK's apksigner entry at whatever
	// is on PATH right now.
	apksigner, signerErr := exec.LookPath("apksigner")
	if signerErr != nil && termuxPkgInstall("apksigner") == nil {
		apksigner, signerErr = exec.LookPath("apksigner")
	}
	// The exporting editor is a linuxbsd build, so the keystore and faux SDK
	// it may validate live at the linux locations under Termux's HOME.
	debug_keystore := filepath.Join(HOME, ".local", "share", "godot", "keystores", "debug.keystore")
	if err := setupHostExportTools(debug_keystore); err != nil {
		return "", false, xray.New(err)
	}
	if err := ensureExportEditorSettings(); err != nil {
		return "", false, xray.New(err)
	}
	GOARCH := "arm64"
	if env := os.Getenv("GOARCH"); env != "" {
		GOARCH = env
	}
	presetName, exportPath, err := pickAndroidPreset(GOARCH)
	if err != nil {
		return "", false, xray.New(err)
	}
	apkPath = filepath.Join(project.GraphicsDirectory, exportPath)
	if err := os.MkdirAll(filepath.Dir(apkPath), 0755); err != nil {
		return "", false, xray.New(err)
	}
	if err := ensureProjectIcon(); err != nil {
		return "", false, xray.New(err)
	}
	if err := os.Chdir(project.GraphicsDirectory); err != nil {
		return "", false, xray.New(err)
	}
	if err := tooling.Godot.Exec("--headless", "--export-release", presetName); err != nil {
		return "", false, xray.New(err)
	}
	// Godot exits 0 even when it printed "Project export failed", so trust
	// the artifact rather than the exit status.
	if _, err := os.Stat(apkPath); err != nil {
		return "", false, fmt.Errorf("gd: godot did not produce %s — see its export errors above", apkPath)
	}
	if signerErr != nil {
		return apkPath, false, nil
	}
	cmd := exec.Command(apksigner,
		"sign", "--ks", debug_keystore,
		"--ks-key-alias", "androiddebugkey", "--ks-pass", "pass:android",
		apkPath,
	)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	if err := cmd.Run(); err != nil {
		return "", false, xray.New(err)
	}
	return apkPath, true, nil
}

// ensureExportEditorSettings seeds the static musl editor's settings with the
// faux java SDK path: godot refuses android exports without a valid one
// ("A valid Java SDK path is required in Editor Settings"). On desktops the
// startup editorSetup plugin fills it in, but that runs on the first editor
// frame, which a headless --export invocation never reaches — and on-device
// the plain `gd` flow opens the editor APK, whose settings are a different
// store entirely, so the musl editor's settings start empty. The path points
// at GDPATH, whose bin/java stub setupHostExportTools has already written —
// the same layout editorSetup configures. A value the user has set themselves
// is left alone.
func ensureExportEditorSettings() error {
	HOME, err := os.UserHomeDir()
	if err != nil {
		return xray.New(err)
	}
	GDPATH := os.Getenv("GDPATH")
	if GDPATH == "" {
		GDPATH = filepath.Join(HOME, "gd")
	}
	config := os.Getenv("XDG_CONFIG_HOME")
	if config == "" {
		config = filepath.Join(HOME, ".config")
	}
	// Godot names the settings file after the minor version, e.g.
	// editor_settings-4.7.tres; the musl editor is always the gd-pinned build.
	minor := tooling.Godot.Version
	if parts := strings.SplitN(minor, ".", 3); len(parts) >= 2 {
		minor = parts[0] + "." + parts[1]
	}
	path := filepath.Join(config, "godot", "editor_settings-"+minor+".tres")
	const key = "export/android/java_sdk_path"
	line := key + ` = "` + GDPATH + `"`
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
			return xray.New(err)
		}
		return os.WriteFile(path, []byte("[gd_resource type=\"EditorSettings\" format=3]\n\n[resource]\n"+line+"\n"), 0644)
	}
	if err != nil {
		return xray.New(err)
	}
	// The editor persists its full defaults on shutdown, so the key is
	// usually already present — as the invalid empty string, which must be
	// replaced. Only a non-empty value counts as user-set.
	lines := strings.Split(string(data), "\n")
	for i, l := range lines {
		if strings.HasPrefix(strings.TrimSpace(l), key) {
			if strings.TrimSpace(l) != key+` = ""` {
				return nil
			}
			lines[i] = line
			return os.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
		}
	}
	// The [resource] section runs to the end of the file, so appending keeps
	// the setting inside it even when the marker line has unexpected spacing.
	text := strings.Replace(string(data), "[resource]\n", "[resource]\n"+line+"\n", 1)
	if text == string(data) {
		text = strings.TrimRight(string(data), "\n") + "\n" + line + "\n"
	}
	return os.WriteFile(path, []byte(text), 0644)
}

// runOnDevice exports, installs and launches the project on the device gd
// itself is running on (Termux). Android offers no silent install to a
// terminal app, so the APK is handed to the system package installer via
// termux-open and the user confirms it on screen; gd watches the package
// manager until the install lands and then launches the app with am. Termux
// cannot read other apps' logcat (READ_LOGS is a privileged permission), so
// engine logs need wireless adb from another machine.
func (android Android) runOnDevice(args ...string) error {
	apkPath, signed, err := android.exportOnDevice(args...)
	if err != nil {
		return xray.New(err)
	}
	if !signed {
		return fmt.Errorf("gd run: %s is unsigned and cannot be installed — install apksigner ('pkg install apksigner') and try again", apkPath)
	}
	GOARCH := "arm64"
	if env := os.Getenv("GOARCH"); env != "" {
		GOARCH = env
	}
	packageName, err := androidPackageName(GOARCH)
	if err != nil {
		return xray.New(err)
	}
	// Termux ships wrappers for the platform tools; fall back to the system
	// binaries when they are not on PATH (same as openAndroidEditorApp).
	systemTool := func(name string) string {
		if path, err := exec.LookPath(name); err == nil {
			return path
		}
		return "/system/bin/" + name
	}
	pm := systemTool("pm")
	// Every successful install moves the package to a fresh /data/app path,
	// so a change in `pm path` — including from empty on a first install —
	// means the user accepted the prompt.
	before, _ := exec.Command(pm, "path", packageName).Output()
	opener, err := exec.LookPath("termux-open")
	if err != nil {
		if termuxPkgInstall("termux-tools") == nil {
			opener, err = exec.LookPath("termux-open")
		}
		if err != nil {
			return fmt.Errorf("gd run: termux-open not found to hand %s to the system package installer — install the termux-tools package, or install the APK manually", apkPath)
		}
	}
	// The Termux content provider refuses to serve files to other apps —
	// including the package installer — until the user opts in, and the
	// installer surfaces that refusal as a bogus "There was a problem
	// parsing the package", so opt in for them (transparently) up front.
	if err := ensureTermuxExternalApps(); err != nil {
		return fmt.Errorf("gd run: Termux does not let other apps read its files, so the package installer cannot receive the APK, and gd could not opt in for you (%w). Opt in with:\n\n\tmkdir -p ~/.termux && echo \"allow-external-apps = true\" >> ~/.termux/termux.properties && termux-reload-settings\n\nand rerun gd run. (This also lets apps you explicitly grant Termux permissions interact with it — see the Termux wiki.)", err)
	}
	if out, err := exec.Command(opener, apkPath).CombinedOutput(); err != nil {
		return xray.New(fmt.Errorf("termux-open %s: %w\n%s", apkPath, err, out))
	}
	fmt.Println("gd: accept the install prompt on the device...")
	fmt.Println("    (no prompt? Termux needs the 'Install unknown apps' permission in Android's")
	fmt.Println("    settings, and android only shows the prompt while Termux is in the foreground)")
	deadline := time.Now().Add(5 * time.Minute)
	installed := false
	lastOpen := time.Now()
	for time.Now().Before(deadline) {
		after, _ := exec.Command(pm, "path", packageName).Output()
		if trimmed := bytes.TrimSpace(after); len(trimmed) > 0 && !bytes.Equal(trimmed, bytes.TrimSpace(before)) {
			installed = true
			break
		}
		// Android only shows the installer while Termux is in the
		// foreground, and silently drops the request otherwise (the export
		// takes minutes, plenty of time to have switched away) — re-hand the
		// APK over periodically so returning to Termux still pops the prompt.
		if time.Since(lastOpen) > 45*time.Second {
			_ = exec.Command(opener, apkPath).Run()
			lastOpen = time.Now()
		}
		time.Sleep(2 * time.Second)
	}
	if !installed {
		fmt.Fprintln(os.Stderr, "gd: could not confirm the install; launching whatever is installed. Usual causes:")
		fmt.Fprintln(os.Stderr, "    - Termux lacks the 'Install unknown apps' permission (Android settings > Apps > Termux)")
		fmt.Fprintln(os.Stderr, "    - the installer refused an update over a build signed with a different key")
		fmt.Fprintln(os.Stderr, "      (e.g. installed over adb from a desktop) — uninstall", packageName, "first")
	}
	// Godot's exported APKs expose com.godot.game.GodotAppLauncher as the
	// (only exported) launcher activity, whatever the applicationId; resolve
	// it from the package manager anyway so a template rename keeps working.
	component := packageName + "/com.godot.game.GodotAppLauncher"
	if resolved, err := exec.Command(systemTool("cmd"), "package", "resolve-activity", "--brief", "-c", "android.intent.category.LAUNCHER", packageName).Output(); err == nil {
		for _, ln := range strings.Split(strings.TrimSpace(string(resolved)), "\n") {
			if ln = strings.TrimSpace(ln); strings.HasPrefix(ln, packageName+"/") {
				component = ln
			}
		}
	}
	out, err := exec.Command(systemTool("am"), "start", "-n", component).CombinedOutput()
	os.Stdout.Write(out)
	// `am start` reports failures like a missing activity on stdout with a
	// zero exit status, so scan the output as well.
	if err != nil || strings.Contains(string(out), "Error") {
		if err == nil {
			err = errors.New(strings.TrimSpace(string(out)))
		}
		return xray.New(err)
	}
	fmt.Println("gd: launched", packageName, "— engine logs are not readable from Termux; use wireless adb logcat from another machine to follow them")
	return nil
}

// termuxPkgInstall installs Termux packages non-interactively, so a first
// `gd run` on-device completes without the user copying setup commands
// around. Quietly does nothing outside Termux (no `pkg` on PATH).
func termuxPkgInstall(packages ...string) error {
	pkg, err := exec.LookPath("pkg")
	if err != nil {
		return err
	}
	fmt.Println("gd: installing", strings.Join(packages, " "), "(pkg install)...")
	cmd := exec.Command(pkg, append([]string{"install", "-y"}, packages...)...)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr
	return cmd.Run()
}

// ensureTermuxExternalApps opts Termux into serving files to other apps when
// the user has not decided either way, telling them what changed and why.
// The property also lets apps the user explicitly grants Termux permissions
// interact with it, which is why the change is announced rather than silent.
func ensureTermuxExternalApps() error {
	if termuxAllowsExternalApps() {
		return nil
	}
	HOME, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	dir := filepath.Join(HOME, ".termux")
	if err := os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	file, err := os.OpenFile(filepath.Join(dir, "termux.properties"), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	// java.util.Properties takes the last occurrence, so appending wins over
	// the commented-out line in the stock template.
	_, err = file.WriteString("\n# added by gd: lets the system package installer read exported APKs out of Termux\nallow-external-apps = true\n")
	if closeErr := file.Close(); err == nil {
		err = closeErr
	}
	if err != nil {
		return err
	}
	fmt.Println("gd: enabled allow-external-apps in ~/.termux/termux.properties so the package")
	fmt.Println("    installer can read the exported APK (see the Termux wiki to learn more)")
	reload, err := exec.LookPath("termux-reload-settings")
	if err != nil {
		return fmt.Errorf("termux-reload-settings not found to apply the change")
	}
	return exec.Command(reload).Run()
}

// termuxAllowsExternalApps reports whether termux.properties sets
// allow-external-apps = true (the stock template carries the line commented
// out). Without it the Termux content provider throws on openFile and the
// system installer shows a generic parse error.
func termuxAllowsExternalApps() bool {
	HOME, err := os.UserHomeDir()
	if err != nil {
		return true // cannot tell; let the flow proceed
	}
	for _, path := range []string{
		filepath.Join(HOME, ".termux", "termux.properties"),
		filepath.Join(HOME, ".config", "termux", "termux.properties"),
	} {
		data, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			key, value, ok := strings.Cut(line, "=")
			if ok && strings.TrimSpace(key) == "allow-external-apps" && strings.TrimSpace(value) == "true" {
				return true
			}
		}
	}
	return false
}

// androidPackageName resolves the applicationId the exported APK will carry
// from the preset's package/unique_name — on-device there is no aapt2 to
// dump it from the APK itself.
func androidPackageName(GOARCH string) (string, error) {
	preset, err := androidPresetFor(GOARCH)
	if err != nil {
		return "", xray.New(err)
	}
	name := preset.uniqueName
	if name == "" || strings.Contains(name, "$") {
		return "", fmt.Errorf("gd run: set a literal package/unique_name for preset %q in graphics/export_presets.cfg (found %q)", preset.name, name)
	}
	return name, nil
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
	preset, err := androidPresetFor(GOARCH)
	if err != nil {
		return "", "", err
	}
	return preset.name, preset.exportPath, nil
}

// androidPresetFor implements [pickAndroidPreset], returning the whole preset.
func androidPresetFor(GOARCH string) (androidPreset, error) {
	abi := "arm64-v8a"
	if GOARCH == "amd64" {
		abi = "x86_64"
	}
	presets, err := loadAndroidPresets()
	if err != nil {
		return androidPreset{}, xray.New(err)
	}
	if want := os.Getenv("GD_ANDROID_PRESET"); want != "" {
		for _, p := range presets {
			if p.name == want {
				return p, nil
			}
		}
		return androidPreset{}, fmt.Errorf("gd build: GD_ANDROID_PRESET=%q not found in graphics/export_presets.cfg", want)
	}
	want := "Android " + abi
	for _, p := range presets {
		if p.name == want {
			return p, nil
		}
	}
	for _, p := range presets {
		if p.platform == "Android" && p.archs[abi] {
			return p, nil
		}
	}
	return androidPreset{}, fmt.Errorf("gd build: no Android preset for %s in graphics/export_presets.cfg", abi)
}

type androidPreset struct {
	name, platform, exportPath, uniqueName string
	archs                                  map[string]bool
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
		case "package/unique_name":
			cur.uniqueName = val
		default:
			if abi, ok := strings.CutPrefix(key, "architectures/"); ok && val == "true" {
				cur.archs[abi] = true
			}
		}
	}
	return out, nil
}
