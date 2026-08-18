package buildkit

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestLinkMachO drives the real graphics.gd iOS game link through
// buildkit's Toolchain — lld.wasm under wazero, from Go, not a shell —
// and checks the output against the native ld64.lld baseline. It proves
// the linker tier is wired in and callable as a library.
//
// Opt-in and environment-specific: it needs the staged modules
// (GD_HARNESS_WASMTC_DIR or the default cache), a GOOS=ios gd build tree
// pointed to by GD_HARNESS_IOS_RELEASE (…/releases/ios), and a native
// ld64.lld on PATH for the baseline. Set GD_HARNESS_LLD_TEST=1 to run.
func TestLinkMachO(t *testing.T) {
	if os.Getenv("GD_HARNESS_LLD_TEST") != "1" {
		t.Skip("set GD_HARNESS_LLD_TEST=1 (needs staged modules + a GOOS=ios release tree in GD_HARNESS_IOS_RELEASE)")
	}
	release := os.Getenv("GD_HARNESS_IOS_RELEASE")
	if release == "" {
		t.Skip("set GD_HARNESS_IOS_RELEASE to a …/releases/ios directory")
	}
	arm64 := filepath.Join(release, "arm64")
	sdk := filepath.Join(release, "sdk")
	app := filepath.Join(arm64, "graphics.gd-harness.app")

	// gd's exact ld64 argument list (cmd/gd/internal/builder/ios.go),
	// with every path made absolute for the guest's / working directory.
	ldArgs := func(out string) []string {
		return []string{
			"-arch", "arm64",
			"-platform_version", "ios", "15.0.0", "15.0.0",
			"-fixup_chains", "-ignore_auto_link",
			"-syslibroot", "/dev/null",
			"-headerpad", "0x1000",
			filepath.Join(arm64, "MoltenVK.xcframework/ios-arm64/libMoltenVK.a"),
			filepath.Join(arm64, "graphics.gd-harness.xcframework/ios-arm64/libgodot.a"),
			filepath.Join(app, "dummy.o"),
			filepath.Join(app, "swift_stubs.o"),
			filepath.Join(arm64, "graphics.gd-harness/dylibs/go.xcframework/ios-arm64/libgo.a"),
			"-o", out,
			"-F", filepath.Join(sdk, "Frameworks"), "-L", filepath.Join(sdk, "lib"),
			"-lSystem", "-lobjc", "-lc++", "-lc++abi", "-lresolv",
			"-framework", "IOSurface", "-framework", "OpenGLES", "-framework", "CoreText",
			"-framework", "CoreGraphics", "-framework", "CoreFoundation", "-framework", "QuartzCore",
			"-framework", "UIKit", "-framework", "Foundation", "-framework", "Metal",
			"-framework", "GameController", "-framework", "CoreMotion", "-framework", "CoreHaptics",
			"-framework", "AVFAudio", "-framework", "AudioToolbox", "-framework", "SwiftUI",
			"-framework", "Security",
			"-lswiftCore", "-lswift_Concurrency", "-lswiftos", "-lswiftCoreFoundation",
			"-lswiftCoreImage", "-lswiftDarwin", "-lswiftDispatch", "-lswiftFoundation",
			"-lswiftMetal", "-lswiftOSLog", "-lswiftObjectiveC", "-lswiftQuartzCore",
			"-lswiftSpatial", "-lswiftUIKit", "-lswiftUniformTypeIdentifiers",
			"-lswiftXPC", "-lswiftsimd",
		}
	}

	ctx := context.Background()
	tc, err := Load(ctx, ModulesDir())
	if err != nil {
		t.Fatal(err)
	}
	defer tc.Close(ctx)
	if !tc.CanLinkMachO() {
		t.Skip("ld64.lld.wasm not staged")
	}

	wasmOut := filepath.Join(t.TempDir(), "game-wasm-lld")
	var log bytes.Buffer
	if err := tc.LinkMachO(ctx, "/", ldArgs(wasmOut), &log); err != nil {
		t.Fatalf("LinkMachO: %v\n%s", err, log.String())
	}
	wasmBin, err := os.ReadFile(wasmOut)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.HasPrefix(wasmBin, []byte{0xcf, 0xfa, 0xed, 0xfe}) { // Mach-O 64 LE magic
		t.Fatalf("output is not a 64-bit Mach-O (starts % x)", wasmBin[:4])
	}
	t.Logf("wasm-lld linked %d bytes", len(wasmBin))

	// Baseline: native ld64.lld from the same args must exist and match
	// except the linker's UUID (and the code-signature hashes cascading
	// from it). If no native linker is present, the Mach-O check above
	// still stands on its own.
	native, err := exec.LookPath("ld64.lld")
	if err != nil {
		t.Skip("no native ld64.lld for the byte-comparison baseline; Mach-O output verified")
	}
	nativeOut := filepath.Join(t.TempDir(), "game-native-lld")
	cmd := exec.Command(native, ldArgs(nativeOut)...)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("native ld64.lld baseline: %v\n%s", err, out)
	}
	nativeBin, err := os.ReadFile(nativeOut)
	if err != nil {
		t.Fatal(err)
	}
	if len(nativeBin) != len(wasmBin) {
		t.Fatalf("size mismatch: native %d, wasm %d", len(nativeBin), len(wasmBin))
	}
	// Count differing bytes; expect only the UUID + cascading signature.
	diffs := 0
	for i := range nativeBin {
		if nativeBin[i] != wasmBin[i] {
			diffs++
		}
	}
	const uuidPlusSignatureCeiling = 20000 // UUID (16B) + ad-hoc code-sig page hashes
	if diffs > uuidPlusSignatureCeiling {
		t.Fatalf("%d differing bytes — more than the expected UUID + signature (%d ceiling)", diffs, uuidPlusSignatureCeiling)
	}
	t.Logf("byte-identical to native ld64.lld except %d bytes (UUID + ad-hoc signature)", diffs)
	if strings.Contains(log.String(), "error") {
		t.Fatalf("linker reported errors:\n%s", log.String())
	}
}
