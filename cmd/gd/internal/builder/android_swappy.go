package builder

import (
	"crypto/sha256"
	"embed"
	"encoding/hex"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"graphics.gd/cmd/gd/internal/gdpaths"
	"graphics.gd/cmd/gd/internal/project"
	"graphics.gd/cmd/gd/internal/tooling"
)

// swappy_src is the source of Google's frame pacing library, which the
// engine's android build expects prebuilt (see bundled/gen-swappy).
//
//go:embed all:bundled/swappy
var swappy_src embed.FS

// swappyABIs are the ABIs the engine's build looks for a Swappy library
// under (platform/android/detect.py), whether or not it is building them.
var swappyABIs = []string{"arm64-v8a", "armeabi-v7a", "x86", "x86_64"}

// swappy builds the frame pacing library for the ABI with zig and installs
// it where the engine's build looks for it, thirdparty/swappy-frame-pacing
// inside of the engine's checkout at src. The build only enables Swappy
// when it finds a library for every ABI, so the ABIs gd does not build
// engines for are given empty archives, nothing ever links them.
func swappy(src, abi, triple string, sdk androidSDK) error {
	zig, err := tooling.Zig.Lookup()
	if err != nil {
		return err
	}
	dir := filepath.Join(gdpaths.Lib, "android", "swappy")
	library := filepath.Join(dir, triple, "libswappy_static.a")
	// the vulkan headers come from the engine, whose ABI they are stable
	// across, so the library is only rebuilt when the source (or zig) is.
	hash := sha256.New()
	fmt.Fprintln(hash, dir, zig)
	err = fs.WalkDir(swappy_src, "bundled/swappy", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		data, err := swappy_src.ReadFile(path)
		fmt.Fprintln(hash, path, len(data))
		hash.Write(data)
		return err
	})
	if err != nil {
		return err
	}
	stamp, sum := library+".txt", hex.EncodeToString(hash.Sum(nil))
	if existing, err := os.ReadFile(stamp); err != nil || string(existing) != sum {
		if err := project.SetupFiles(swappy_src, "bundled/swappy", dir); err != nil {
			return err
		}
		if err := buildSwappy(zig, dir, library, triple, filepath.Join(src, "thirdparty", "vulkan", "include"), sdk); err != nil {
			return err
		}
		if err := os.WriteFile(stamp, []byte(sum), 0644); err != nil {
			return err
		}
	}
	for _, name := range swappyABIs {
		installed := filepath.Join(src, "thirdparty", "swappy-frame-pacing", name, "libswappy_static.a")
		if name == abi {
			if err := copyFile(library, installed); err != nil {
				return err
			}
			continue
		}
		if _, err := os.Stat(installed); err != nil {
			if err := archive(zig, installed, nil); err != nil {
				return err
			}
		}
	}
	return nil
}

// buildSwappy compiles the library's sources for the triple, with the
// flags of its own CMake build, archiving them at library. The Java
// classes it loads on older Android versions are compiled in as the dex
// file the upstream build embeds, under the symbols it looks for.
func buildSwappy(zig, dir, library, triple, vulkan string, sdk androidSDK) error {
	src := filepath.Join(dir, "src")
	sources, err := filepath.Glob(filepath.Join(src, "*.cpp"))
	if err != nil {
		return err
	}
	dex := filepath.Join(dir, "classes_dex.S")
	if err := os.WriteFile(dex, []byte(`	.section .rodata
	.balign 8
	.globl _binary_classes_dex_start
_binary_classes_dex_start:
	.incbin "classes.dex"
	.globl _binary_classes_dex_end
_binary_classes_dex_end:
	.section .note.GNU-stack,"",%progbits
`), 0644); err != nil {
		return err
	}
	sources = append(sources, dex)
	objects := filepath.Join(filepath.Dir(library), "obj")
	if err := os.RemoveAll(objects); err != nil {
		return err
	}
	if err := os.MkdirAll(objects, 0755); err != nil {
		return err
	}
	fmt.Printf("gd: building swappy for %s\n", triple)
	var (
		wg      sync.WaitGroup
		mutex   sync.Mutex
		failure error
		jobs    = make(chan string)
		outputs []string
	)
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for source := range jobs {
				object := filepath.Join(objects, strings.TrimSuffix(filepath.Base(source), filepath.Ext(source))+".o")
				args := []string{"c++", "-target", triple + "." + androidAPI, "-c", source, "-o", object,
					"-std=c++17", "-Wall", "-Wthread-safety", "-D_LIBCPP_ENABLE_THREAD_SAFETY_ANNOTATIONS",
					"-O3", "-fPIC", "-fno-exceptions", "-fno-rtti", "-ffunction-sections", "-fdata-sections", "-g0",
					"-I", src, "-I", vulkan,
				}
				cmd := exec.Command(zig, append(args, tooling.CGOCFlags()...)...)
				cmd.Dir = dir // where classes.dex is, for the .incbin
				cmd.Env = append(os.Environ(), "ZIG_LIBC="+sdk.LibC)
				output, err := cmd.CombinedOutput()
				mutex.Lock()
				if err != nil && failure == nil {
					failure = fmt.Errorf("gd: failed to build swappy: %w\n%s", err, output)
				}
				outputs = append(outputs, object)
				mutex.Unlock()
			}
		}()
	}
	for _, source := range sources {
		jobs <- source
	}
	close(jobs)
	wg.Wait()
	if failure != nil {
		return failure
	}
	os.Remove(library)
	cmd := exec.Command(zig, append([]string{"ar", "rcs", library}, outputs...)...)
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("gd: failed to archive swappy: %w\n%s", err, output)
	}
	return nil
}
