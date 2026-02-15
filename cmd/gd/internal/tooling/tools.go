package tooling

var Godot = toolchain{
	Name:          "godot",
	Version:       "4.6",
	VersionFlag:   "--version",
	VersionPrefix: "4.6",
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
	VersionFlag:   "clang --version",
	VersionPrefix: "clang version 21.",
	DownloadURL:   "https://release.graphics.gd/llvm.$(GOOS).$(GOARCH)$(EXT)",
	DownloadEXT:   map[string]string{"linux": "", "windows": ".exe", "darwin": ""},
	RequiredFor:   "linking iOS builds",
}

var Zig = toolchain{
	Name:         "zig",
	Version:      "0.15.2",
	VersionFlag:  "version",
	DownloadHint: "https://ziglang.org/download/",
	DownloadURL:  "https://ziglang.org/download/$(VERSION)/zig-$(ARCH)-$(OS)-$(VERSION)$(EXT)",
	DownloadOS:   map[string]string{"windows": "windows", "darwin": "macos", "linux": "linux"},
	DownloadARCH: map[string]string{"amd64": "x86_64", "arm64": "aarch64"},
	DownloadEXT:  map[string]string{"windows": ".zip", "darwin": ".tar.xz", "linux": ".tar.xz"},
	RequiredFor:  "cross-compiling",
}

var Go = toolchain{
	Name:          "go",
	VersionFlag:   "version",
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
	VersionFlag:   "--help",
	VersionPrefix: "Description:\n  Velopack CLI 0.0.1298,",
	RequiredFor:   "self-updating-bundles",
}

var AndroidPackageSigner = toolchain{
	Name:        "apksigner",
	Version:     "0.9",
	VersionFlag: "--version",
	DownloadURL: "https://release.graphics.gd/apksigner.$(GOOS).$(GOARCH)$(EXT)",
	DownloadEXT: map[string]string{"linux": "", "windows": ".exe", "darwin": ""},
	RequiredFor: "building the .apk",
}

var AndroidDebugBridge = toolchain{
	Name:            "adb",
	Version:         "1.0.41",
	VersionFlag:     "--version",
	VersionPrefix:   "Android Debug Bridge version 1.0.41",
	DownloadURL:     "https://release.graphics.gd/adb.$(GOOS).$(GOARCH)$(EXT)",
	DownloadEXT:     map[string]string{"linux": "", "windows": ".zip", "darwin": ""},
	RequiredFor:     "launching the project on a connected android device",
	DarwinUniversal: true,
}

var UltimatePackerForExecutables = toolchain{
	Name:          "upx",
	Version:       "5.0.2",
	VersionFlag:   "--version",
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
	VersionFlag:   "v",
	DownloadURL:   "https://release.graphics.gd/apktool.$(GOOS).$(GOARCH)$(EXT)",
	DownloadEXT:   map[string]string{"linux": "", "windows": ".exe", "darwin": ""},
	RequiredFor:   "converting the exported .apk into an .aab",
}

var AndroidAssetPackagingTool = toolchain{
	Name:            "aapt2",
	Version:         "2.19-android-13.0.0_r6",
	VersionPrefix:   "Android Asset Packaging Tool (aapt) 2.",
	VersionFlag:     "version",
	DownloadURL:     "https://release.graphics.gd/aapt2.$(GOOS).$(GOARCH)$(EXT)",
	DownloadEXT:     map[string]string{"linux": "", "windows": ".exe", "darwin": ""},
	RequiredFor:     "converting the exported .apk into an .aab",
	DarwinUniversal: true,
}

var BundleTool = toolchain{
	Name:        "bundletool",
	Version:     "1.18.3",
	VersionFlag: "version",
	DownloadURL: "https://release.graphics.gd/bundletool.$(GOOS).$(GOARCH)$(EXT)",
	DownloadEXT: map[string]string{"linux": "", "windows": ".exe", "darwin": ""},
	RequiredFor: "converting the exported .apk into an .aab",
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
