---
title: The Command Line Tool
slug: guide/gd
sidebar:
  order: 2
---

The `gd` command can be installed using `go install`, this command is a drop in replacement for the `go` command for working with `graphics.gd` projects.

```sh
go install graphics.gd/cmd/gd@release
```

Run this command again whenever you update `graphics.gd` (it doesn't currently stay in sync with your project version).

The `gd` command is also fully compatible with standard GDScript-based Godot projects and can be used to initialise cross-platform export configurations
and to easily launch projects on Web, Android and iOS. To use it this way, run it inside the directory where `project.godot` is located.

#### Quick Rundown

```sh
gd       # launch the Godot editor to manage assets and design the UI.
gd run   # run the project in debug mode.
gd build # export the project as a release depending upon GOOS and GOARCH.
gd test  # run your Go tests.
```

#### Why a new command?

The `gd` command takes care of all the build flags and configuration files needed to setup Go + Godot so that things will "just work" in the same way that the `go` command does.

If you are already well versed in Godot and GDExtension, or you want to use your own build tools, you can use the `go` command to build the shared library yourself:

```sh
CC="zig cc" go build -o example.so -buildmode=c-shared
```

#### What does the `gd` command do?

Everything you would have to setup for yourself if you were using `go` directly:

* downloads and installs dependencies for you on supported platforms.
* creates a `graphics` directory to house the godot project.
* creates a blank `project.godot` and `main.tscn`.
* creates the `.gdextension` library and enables it.
* outputs Go binaries into the `graphics` directory.
* sets up a `.gitignore` to ignore these compiled Go libraries.
* creates export presets that export into `releases`.
* sets the correct `zig` flags when cross-compiling.
* uses `vpk` to bundle the project into a self-updating release.
* sets up Go tests to run correctly within the Godot runtime.
* builds projects for Android and iOS without their SDKs installed.
