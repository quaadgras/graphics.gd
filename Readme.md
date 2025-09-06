# graphics.gd [![Go Reference](https://pkg.go.dev/badge/graphics.gd.svg)](https://pkg.go.dev/graphics.gd)

A cross platform graphics runtime for Go suitable for building native mobile apps, gdextensions, multimedia applications, games and more!

_Why use graphics.gd?_

* [Write shaders in Go!](./shaders/Readme.md)
* Unlike other native Godot languages, RIDs, callables and dictionary arguments are all distinctly typed.
* A good balance of performance and convenience.
* General purpose pure-Go 'variant' packages, reuse them in any Go project.
* After the first build, recompile quickly, with an experience similar to a scripting language.
* Easily cross-compile for windows/macos/android/linux on any host platform.
* Build and launch native apps on connected Android devices, no Java nor Android SDK needed.

Not just a wrapper! graphics.gd is designed from the ground up to provide a cohesive curated graphics runtime on top of gdextension.

We would love you to take part in our [active discussions](https://github.com/quaadgras/graphics.gd/discussions)
section with any questions, comments or feedback you may have. Show us what you're building!

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

You can help fund the project, motivate development and prioritise issues [here](https://buy.stripe.com/4gw14maETbnX3vOcMM)

## Getting Started
The module includes a drop-in replacement for the go command called `gd` that
makes it easy to work with projects that run within the runtime.
It enables you to start developing a new project from a single `main.go` file,
to install it, make sure that your `$GOPATH/bin` is in your `$PATH` and run:

```sh
	go install graphics.gd/cmd/gd@release
```

Now when you can run `gd run`, `gd test` on the main package in your project's
directory, things will work as expected. The tool will create a "graphics"
subdirectory where you can manage your assets via the Engine's Editor.

Running the command without any arguments will startup the editor.

**NOTE** On linux (and macos if you have brew), `gd` will download the engine for you automatically!
**HINT**  On Windows, you'll want to
[setup CGO](https://github.com/go101/go101/wiki/CGO-Environment-Setup).

If you don't want to use the `gd` command, you can build a shared library with
the `go` command directly:

```sh
go build -o example.so -buildmode=c-shared
```

## Next Steps

Check out the [the.graphics.gd/guide](https://the.graphics.gd/guide) which covers more topics.

## Design Principles

Each graphics class is available as a package under `classdb`. To import the
`Node` class you can import `"graphics.gd/classdb/Node"` There's no inheritance,
so to access a 'super' class, you need to call `Super()` on an Extension 'Class'.
All engine classes have methods to cast to any sub-classes they extend for example
`AsObject()` or `AsNode2D()`.

Methods have been renamed to follow Go conventions, so instead of
underscores, methods are named as PascalCase. Keep this in mind when
referring to the Engine documentation.

https://docs.godotengine.org/en/latest/index.html

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

## Performance
It's feasible to write high performance code using this module, keep to Engine types where possible and avoid
allocating memory on the heap in frequently called functions. `Advanced` instances are available for each class
which allow more fine-grained control over memory allocations.

Benchmarking shows `Advanced` method calls from Go -> Engine do not allocate in practice.

## Examples
There are a number of examples in the [samples](https://github.com/quaadgras/graphics.gd/tree/samples)
branch. All the samples are designed to be run with `gd run` without any additional setup.

## Supported Platforms

* Windows `GOOS=windows gd build`
* Linux   `GOOS=linux gd build`
* Mac     `GOOS=darwin gd build`
* Android `GOOS=android GOARCH=arm64 gd run`
* IOS     (no `gd` tooling support for this target)
* Web     `GOOS=js GOARCH=wasm gd run`

## Platform Restrictions

* Go is not available as a 'scripting' language within the editor.
* 64bit only (arm64 && amd64).
* No support for video game consoles (this may be achieved in the future with WASM).

## Contributing

The best way you can contribute to graphics.gd is to **try it**, this project needs you to find out
what's working and what doesn't, so do please let us know of any trouble that you run into! Any
examples you can contribute are more than welcome.

The next best thing you can do to help is improve the Variant packages, these are general-purpose
packages inspired by the Godot engine's Variant types. Specifically any changes you can make to
optimize functionality and/or improve test coverage of these packages is more than welcome
(such as specialized assembly routines for vector operations, anyone?).

If you enjoy hunting down memory-safety issues, we would appreciate any issues being opened on
this front.

Thirdly, the project needs more tests to ensure that everything is working, the best way you can
guarantee that graphics.gd won't break on you is to contribute tests that cover the functionality
you need!

To run the go tests for graphics.gd, cd into the repo and run `cd internal && gd test`.

Note, we are looking for somebody to create benchmarks for the project.

Lastly, spread the word and let people know about graphics.gd!

## See Also

* [godot-go](https://github.com/godot-go/godot-go) (Another project aiming to support Go + Godot integration)

## Licensing
This project is licensed under an MIT license (the same license as Godot), you can use it in any manner
you can use the Godot engine. If you use this project for any commercially successful products, please
consider [financially supporting us](https://buy.stripe.com/4gw14maETbnX3vOcMM) to show your appreciation!
