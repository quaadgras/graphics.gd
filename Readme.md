# graphics.gd [![Go Reference](https://pkg.go.dev/badge/graphics.gd.svg)](https://pkg.go.dev/graphics.gd)

A cross platform 2D/3D graphics runtime for Go suitable for building native mobile apps, gdextensions,
multimedia applications, games and more.

_Why use graphics.gd?_

* [Write shaders in Go!](./shaders/Readme.md)
* Full compatibility with the [Godot Engine](https://godotengine.org/) ecosystem.
* Unlike C++/C#/GDScript/Rust/Swift, RIDs, Callables and Dictionary arguments are strongly typed.
* Fully documented on [pkg.go.dev](https://pkg.go.dev/graphics.gd), with all code snippets in Go.
* Pure-Go ported `variant` packages, for vector math and more, reuse them in any Go project.
* After an initial build, recompile quickly, with an experience similar to scripting languages.
* Easily cross-compile for windows/macos/android/linux/ios/web on any host platform.
* Neither Java, nor an Android SDK/NDK is needed to build Android apps.
* Neither Xcode nor MacOS is needed to build iOS apps.
* Drop in `gd` command, a cross-platform build tool compatible with `.gd` script projects.

Not just a 1:1 wrapper for gdextension! graphics.gd has been designed from the ground up to 
provide a cohesive and curated experience for using Go on top of Godot + GDExtension.

## Hello World

```go
// This file is all you need to start a project.
// Save it somewhere, install the `gd` command and use `gd run` to get started.
package main

import (
	"graphics.gd/startup"

	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/GUI"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/SceneTree"
)

func main() {
	startup.LoadingScene() // setup the SceneTree and wait until we have access to engine functionality
	hello := Label.New()
	hello.AsControl().SetAnchorsPreset(Control.PresetFullRect) // expand the label to take up the whole screen.
	hello.SetHorizontalAlignment(GUI.HorizontalAlignmentCenter)
	hello.SetVerticalAlignment(GUI.VerticalAlignmentCenter)
	hello.SetText("Hello, World!")
	SceneTree.Add(hello)
	startup.Scene() // starts up the scene and blocks until the engine shuts down.
}
```

## Getting Started
The module includes a drop-in replacement for the go command called `gd` that
makes it easy to work with projects that run within the runtime (including `.gd` 
script projects). It also enables you to start developing a new project starting from
a single  `main.go` file, to install it, make sure that your `$GOPATH/bin` is in your 
`$PATH` and run:

$ `go install graphics.gd/cmd/gd@release`

Now when you can run `gd run`, `gd test` on the main package in your project's
directory, things will work as expected. The tool will create a "graphics"
subdirectory where you can manage your assets via the
[Godot Engine](https://godotengine.org/) editor.

Running the command without any additional arguments will startup the editor.

If you don't want to use the `gd` command, you can also build a shared library with
the standard `go` command (this can be included in an existing 
[Godot Engine](https://godotengine.org/) project):

$ `go build -o example.so -buildmode=c-shared`

## Next Steps

Check out the [the.graphics.gd/guide](https://the.graphics.gd/guide) which covers much, 
much more!

## TLDR

Each engine class is available as a package under `classdb`. To import the
`Node` class you can import `"graphics.gd/classdb/Node"` There's no inheritance,
so to access a 'super' class, you need to call `Super()` on an extension 'class'.
All engine classes have methods to cast to any sub-classes they extend for example
`AsObject()` or `AsNode2D()`.

Methods have been renamed to follow Go conventions, so instead of
underscores, methods are named as PascalCase. Keep this in mind when
referring to Godot documentation.

https://docs.godotengine.org/en/latest/index.html

Optional arguments are omitted by default, convert an `Instance` into either the `Expanded`
or `Advanced` type to use them. ie. for a `Node.Instance` called `node`, it can be converted:

```go
Node.Expanded(node).AddChild(...)
```

## Where Do I Find?
Ctrl+F in the project for a specific `//gd:symbol` to find the matching Go symbol.
```
* Engine Class           -> `//gd:ClassName`
* Engine Class Method    -> `//gd:ClassName.method_name`
* Utility Functions      -> `//gd:utility_function_name`
* Enum                   -> `//gd:ClassName.EnumName`
```
_NOTE_ in order to avoid circular dependencies, a handful of functions have moved packages,
for example `Node.get_tree()` (GDScript) has moved to `SceneTree.Get()` (Go).

## Community & Support

Join the [active discussions](https://github.com/quaadgras/graphics.gd/discussions)
with any questions, comments or feedback you may have. Show us what you're building!

The API surface of the [Godot Engine] is huge, not everything has been translated to
Go optimally, the best thing you can do is to [report] anything that seems 'off' this way
you can reduce the likelihood of being affected by any breaking changes in the future.

*Sponsors receive priority support!* Secure the development of `graphics.gd` and prioritise 
having issues resolved by [sponsoring me](https://github.com/sponsors/Splizard). 

## Performance
It's feasible to write high performance code with `graphics.gd`, keep to variant types where possible 
and avoid allocating memory on the heap in frequently called functions. `Advanced` instances are 
available for each class which allow more fine-grained control over memory allocations.

Benchmarks show that `Advanced` method calls from Go -> Godot don't typically allocate any
memory.

## Examples
There are a number of examples in the [samples](https://github.com/quaadgras/graphics.gd/tree/samples)
branch. All the samples are designed to be run with `gd run` without any additional setup.

## Supported Platforms

* Windows `GOOS=windows gd build`
* Linux   `GOOS=linux gd build`
* MacOS   `GOOS=macos gd build`
* Android `GOOS=android GOARCH=arm64 gd run`
* IOS     `GOOS=ios gd run` (requires [SideStore](https://sidestore.io) on the IOS device)
* Web     `GOOS=web gd run`

## Platform Restrictions

* 64bit only (arm64 && amd64).
* No support for Playstation/Xbox/Switch (achievable in the future with WASI, wasm2c or hitsumabushi).

## Contributing

The best way you can contribute to graphics.gd is to **try it**, this project needs you to find out
what's working and what doesn't, so do please let us know of any trouble that you run into! Any
examples you can contribute are more than welcome.

The next best thing you can do to help is improve the Variant packages, these are general-purpose
packages inspired by the Godot engine's Variant types. Specifically any changes you can make to
optimize functionality and/or improve test coverage of these packages is more than welcome.

If you enjoy hunting down memory-safety issues, we would appreciate this.

`graphics.gd` is looking for someone to create benchmarks to compare this project with `.gd` script 
and/or other gdextension-based alternatives.

The project also needs more tests to ensure that everything is working, the best way you can
guarantee that graphics.gd won't break on you is to contribute tests that cover any functionality
that you need!

To run the go tests for graphics.gd, cd into the repo and run `cd internal && gd test`.

Another great way to contribute, is to write a blog, share a post or let others know about your 
experience with graphics.gd!

## See Also

If you're interested in graphics.gd, you may also wish to explore:

* [godot-go](https://github.com/godot-go/godot-go) (a different Go + Godot project)
* [ebiten](https://github.com/hajimehoshi/ebiten/) (a 2D game engine for Go)
* [g3n](https://github.com/g3n/engine) (a 3D game engine for Go)
* [gdext](https://github.com/godot-rust/gdext) (Rust bindings for Godot 4) 

## Licensing
This project is licensed under an MIT license (the same license as Godot), you can use it in any manner
you can use the Godot engine. If you do use this project for any commercially successful products, please
consider [sponsoring the maintainer](https://github.com/sponsors/Splizard) to show your appreciation!
