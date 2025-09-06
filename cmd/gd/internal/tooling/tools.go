package tooling

var Godot = toolchain{
	Name:          "godot",
	Version:       "4.4.1",
	VersionFlag:   "--version",
	VersionPrefix: "4.4.1.",
	DownloadHint:  "https://godotengine.org/download",
	DownloadURL:   "https://github.com/godotengine/godot/releases/download/$(VERSION)-stable/Godot_v$(VERSION)-stable_$(OS).zip",
	DownloadOS:    map[string]string{"windows": "win64.exe", "linux": "linux.$(ARCH)"},
	DownloadARCH:  map[string]string{"amd64": "x86_64", "arm64": "arm64"},
	Unzip:         "Godot_v$(VERSION)-stable_$(OS)",
	RequiredFor:   "graphics",
}

var Zig = toolchain{
	Name:         "zig",
	Version:      "0.15.1",
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
	Version:       "1.25.0",
	DownloadHint:  "https://go.dev/dl/",
	VersionPrefix: "go version go1.25.",
	RequiredFor:   "compiling",
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
	DownloadURL: "https://release.graphics.gd/apksigner.$(GOOS).$(GOARCH)",
	Installations: map[string]string{
		"linux":   "$(HOME)/Android/Sdk/build-tools/35",
		"windows": "$(HOME)/AppData/Local/Android/Sdk/build-tools/35",
	},
	RequiredFor: "building the .apk",
}

var AndroidDebugBridge = toolchain{
	Name:          "adb",
	Version:       "1.0.41",
	VersionFlag:   "--version",
	VersionPrefix: "Android Debug Bridge version 1.0.41",
	DownloadURL:   "https://release.graphics.gd/adb.$(GOOS).$(GOARCH)",
	Installations: map[string]string{
		"linux":   "$(HOME)/Android/Sdk/platform-tools",
		"windows": "$(HOME)/AppData/Local/Android/Sdk/platform-tools",
	},
	RequiredFor: "launching the project on a connected android device",
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
