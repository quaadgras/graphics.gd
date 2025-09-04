package tooling

var Godot = toolchain{
	Name:          "godot",
	VersionFlag:   "--version",
	VersionPrefix: "4.4.1.",
	DownloadHint:  "https://godotengine.org/download",
	DownloadURL:   "https://github.com/godotengine/godot-builds/releases/download/$(VERSION)-stable/Godot_$(VERSION)-stable_$(OS).zip",
	DownloadOS:    map[string]string{"windows": "win64.exe", "linux": "linux.$(ARCH)"},
	DownloadARCH:  map[string]string{"amd64": "x86_64", "arm64": "arm64"},
	GOBIN:         true,
	Unzip:         "Godot_v$(VERSION)-stable_$(OS)",
	Installation:  "$(GOBIN)",
	RequiredFor:   "graphics",
}

var Zig = toolchain{
	Name:          "zig",
	VersionFlag:   "version",
	VersionEquals: "0.15.1",
	DownloadHint:  "https://ziglang.org/download/",
	DownloadURL:   "https://ziglang.org/builds/zig-$(ARCH)-$(OS)-$(VERSION).zip",
	DownloadOS:    map[string]string{"windows": "windows", "darwin": "macos", "linux": "linux"},
	DownloadARCH:  map[string]string{"amd64": "x86_64", "arm64": "aarch64"},
	Installation:  "$(GOPATH)/zig",
	RequiredFor:   "cross-compiling",
}

var Go = toolchain{
	Name:          "go",
	DownloadHint:  "https://go.dev/dl/",
	VersionPrefix: "go version go1.25.0",
	RequiredFor:   "compiling",
}

var Velopack = toolchain{
	Name:          "vpk",
	VersionFlag:   "--help",
	VersionPrefix: "Description:\n  Velopack CLI 0.0.1298,",
	RequiredFor:   "self-updating-bundles",
}

var AndroidPackageSigner = toolchain{
	Name:          "apksigner",
	VersionFlag:   "--version",
	VersionEquals: "0.9",
	DownloadURL:   "https://release.graphics.gd/apksigner.$(GOOS).$(GOARCH)",
	Installations: map[string]string{
		"linux":   "$(HOME)/Android/Sdk/build-tools/35",
		"windows": "$(HOME)/AppData/Local/Android/Sdk/build-tools/35",
	},
	RequiredFor: "building the .apk",
}

var AndroidDebugBridge = toolchain{
	Name:          "adb",
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
