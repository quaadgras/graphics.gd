package tooling

var Godot = toolchain{
	Name:          "godot",
	Version:       "4.6.2",
	VersionFlags:  []string{"--version"},
	VersionPrefix: "4.6.",
	DownloadHint:  "https://godotengine.org/download",
	DownloadURL:   "https://github.com/godotengine/godot/releases/download/$(VERSION)-stable/Godot_v$(VERSION)-stable_$(OS).zip",
	DownloadOS:    map[string]string{"windows": "win64.exe", "linux": "linux.$(ARCH)", "darwin": "macos.universal"},
	DownloadARCH:  map[string]string{"amd64": "x86_64", "arm64": "arm64"},
	Unzip:         "Godot_v$(VERSION)-stable_$(OS)",
	IsApp:         true,
	RequiredFor:   "graphics",

	ConvertArguments: map[string]string{
		"-v":       "--verbose",
		"-x":       "",
		"-gcflags": "",
	},
}

var LLVM = toolchain{
	Name:          "llvm",
	Version:       "21.1.8",
	VersionFlags:  []string{"clang", "--version"},
	VersionPrefix: "clang version 21.",
	DownloadURL:   "https://release.graphics.gd/llvm.$(GOOS).$(GOARCH)$(EXT)",
	DownloadEXT:   map[string]string{"linux": "", "windows": ".exe", "darwin": ""},
	RequiredFor:   "linking iOS builds",
}

var Zig = toolchain{
	Name:         "zig",
	Version:      "0.15.2",
	VersionFlags: []string{"version"},
	DownloadHint: "https://ziglang.org/download/",
	DownloadURL:  "https://ziglang.org/download/$(VERSION)/zig-$(ARCH)-$(OS)-$(VERSION)$(EXT)",
	DownloadOS:   map[string]string{"windows": "windows", "darwin": "macos", "linux": "linux"},
	DownloadARCH: map[string]string{"amd64": "x86_64", "arm64": "aarch64"},
	DownloadEXT:  map[string]string{"windows": ".zip", "darwin": ".tar.xz", "linux": ".tar.xz"},
	RequiredFor:  "cross-compiling",
}

var Go = toolchain{
	Name:          "go",
	VersionFlags:  []string{"version"},
	Version:       "1.26.0",
	DownloadHint:  "https://go.dev/dl/",
	VersionPrefix: "go version go1.26.",
	RequiredFor:   "compiling",

	ConvertArguments: map[string]string{
		"--verbose": "-v",
	},
}

var Velopack = toolchain{
	Name:          "vpk",
	Version:       "0.0.1298",
	VersionFlags:  []string{"--help"},
	VersionPrefix: "Description:\n  Velopack CLI 0.0.1298,",
	RequiredFor:   "self-updating-bundles",
}

var AndroidPackageSigner = toolchain{
	Name:         "apksigner",
	Version:      "0.9",
	VersionFlags: []string{"--version"},
	DownloadURL:  "https://release.graphics.gd/apksigner.$(GOOS).$(GOARCH)$(EXT)",
	DownloadEXT:  map[string]string{"linux": "", "windows": ".exe", "darwin": ""},
	RequiredFor:  "building the .apk",
}

var AndroidDebugBridge = toolchain{
	Name:            "adb",
	Version:         "1.0.41",
	VersionFlags:    []string{"--version"},
	VersionPrefix:   "Android Debug Bridge version 1.0.41",
	DownloadURL:     "https://release.graphics.gd/adb.$(GOOS).$(GOARCH)$(EXT)",
	DownloadEXT:     map[string]string{"linux": "", "windows": ".zip", "darwin": ""},
	RequiredFor:     "launching the project on a connected android device",
	DarwinUniversal: true,
}

var UltimatePackerForExecutables = toolchain{
	Name:          "upx",
	Version:       "5.0.2",
	VersionFlags:  []string{"--version"},
	VersionPrefix: "upx 5.0.2",
	DownloadHint:  "https://github.com/upx/upx/releases/latest",
	Downloads: map[string]map[string]string{
		"windows": {
			"amd64": "https://github.com/upx/upx/releases/download/v$(VERSION)/upx-$(VERSION)-win64.zip",
		},
	},
	DownloadURL:  "https://github.com/upx/upx/releases/download/v$(VERSION)/upx-$(VERSION)-$(ARCH)_$(OS).zip",
	DownloadOS:   map[string]string{"linux": "linux"},
	DownloadARCH: map[string]string{"amd64": "amd64", "arm64": "arm64"},
	RequiredFor:  "minifying builds",
}

var AndroidPackageKitTool = toolchain{
	Name:          "apktool",
	Version:       "2.12.1",
	VersionPrefix: "2.12.1-",
	VersionFlags:  []string{"v"},
	DownloadURL:   "https://release.graphics.gd/apktool.$(GOOS).$(GOARCH)$(EXT)",
	DownloadEXT:   map[string]string{"linux": "", "windows": ".exe", "darwin": ""},
	RequiredFor:   "converting the exported .apk into an .aab",
}

var AndroidAssetPackagingTool = toolchain{
	Name:            "aapt2",
	Version:         "2.19-android-13.0.0_r6",
	VersionPrefix:   "Android Asset Packaging Tool (aapt) 2.",
	VersionFlags:    []string{"version"},
	DownloadURL:     "https://release.graphics.gd/aapt2.$(GOOS).$(GOARCH)$(EXT)",
	DownloadEXT:     map[string]string{"linux": "", "windows": ".exe", "darwin": ""},
	RequiredFor:     "converting the exported .apk into an .aab",
	DarwinUniversal: true,
}

var BundleTool = toolchain{
	Name:         "bundletool",
	Version:      "1.18.3",
	VersionFlags: []string{"version"},
	DownloadURL:  "https://release.graphics.gd/bundletool.$(GOOS).$(GOARCH)$(EXT)",
	DownloadEXT:  map[string]string{"linux": "", "windows": ".exe", "darwin": ""},
	RequiredFor:  "converting the exported .apk into an .aab",
}

var Android = toolchain{
	Name:        "android.jar",
	DownloadURL: "https://release.graphics.gd/android.jar",
	RequiredFor: "converting the exported .apk into an .aab",
	IsLibrary:   true,
}

var LibGodotEditor = toolchain{
	Name:        "libgodot.$(GOOS).editor.$(GOARCH).$(EXT)",
	DownloadURL: "https://release.graphics.gd/libgodot.$(GOOS).editor.$(GOARCH).$(EXT)",
	DownloadEXT: map[string]string{"musl": "a", "linux": "a", "windows": "lib", "darwin": "a"},
	RequiredFor: "launching the editor on musl systems",
	IsLibrary:   true,
}

var LibGodot = toolchain{
	Name:        "libgodot.$(GOOS).$(GOARCH).$(EXT)",
	DownloadURL: "https://release.graphics.gd/libgodot.$(GOOS).$(GOARCH).$(EXT)",
	DownloadEXT: map[string]string{"musl": "a", "linux": "a", "windows": "lib", "darwin": "a"},
	RequiredFor: "musl systems & single binaries",
	IsLibrary:   true,
}

var ListDynamicDependencies = toolchain{
	Name:        "ldd",
	RequiredFor: "musl detection",
}

// Java is required at build time only for GOOS=metaquest, to run the R8
// dex compiler against the OpenXR vendor AAR. Auto-download is not
// attempted — distributions vary too much. Users on systems without
// Java will see the install hint instead of a download progress bar.
var Java = toolchain{
	Name:          "java",
	VersionFlags:  []string{"-version"},
	VersionPrefix: "openjdk version ",
	DownloadHint:  "https://adoptium.net/temurin/releases/?version=17",
	RequiredFor:   "running the R8 dex compiler when packaging GOOS=metaquest APKs",
}

// R8 is Google's APK shrinker; we use its bundled D8 entry point as a
// pure-Java jar to compile the OpenXR vendor classes.jar to a dex file
// for injection into the Meta Quest APK. BSD-3-Clause. Pulled directly
// from Maven Central.
var R8 = toolchain{
	Name:        "r8-$(VERSION).jar",
	Version:     "8.5.35",
	DownloadURL: "https://repo1.maven.org/maven2/com/android/tools/r8/$(VERSION)/r8-$(VERSION).jar",
	RequiredFor: "compiling Meta Quest OpenXR vendor classes to dex (GOOS=metaquest)",
	IsLibrary:   true,
}

// OpenXRLoaderAAR ships the Khronos libopenxr_loader.so for arm64-v8a
// plus the standard manifest entries every OpenXR app needs on Android.
// Apache-2.0. Pulled directly from Maven Central.
var OpenXRLoaderAAR = toolchain{
	Name:        "openxr_loader_for_android-$(VERSION).aar",
	Version:     "1.1.53",
	DownloadURL: "https://repo1.maven.org/maven2/org/khronos/openxr/openxr_loader_for_android/$(VERSION)/openxr_loader_for_android-$(VERSION).aar",
	RequiredFor: "providing the Khronos OpenXR loader to Meta Quest builds (GOOS=metaquest)",
	IsLibrary:   true,
}

// OpenXRVendorsMetaAAR is the GodotVR community's plugin AAR that
// registers Meta's OpenXR runtime with the engine at startup. Without
// the classes inside, the OpenXR loader has no path to talk to Horizon
// OS. Apache-2.0. Pulled directly from Maven Central.
var OpenXRVendorsMetaAAR = toolchain{
	Name:        "godot-openxr-vendors-meta-$(VERSION).aar",
	Version:     "4.2.2-stable",
	DownloadURL: "https://repo1.maven.org/maven2/org/godotengine/godot-openxr-vendors-meta/$(VERSION)/godot-openxr-vendors-meta-$(VERSION).aar",
	RequiredFor: "registering Meta's OpenXR runtime with the engine (GOOS=metaquest)",
	IsLibrary:   true,
}
