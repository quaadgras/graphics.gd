# Graphics GD [![Go Reference](https://pkg.go.dev/badge/graphics.gd.svg)](https://pkg.go.dev/graphics.gd)

A safe performant way to work with graphics and game development in Go via
a GDExtension host (ie. Godot 4.4).

_Why use graphics.gd?_

* [Write shaders in Go!](./shaders/Readme.md)
* Unlike with other options, RIDs, callables and dictionary arguments are all distinctly typed.
* A good balance of performance and convenience.
* General-purpose pure-Go 'variant' packages, reuse them in any Go project.
* Recompile your code quickly, with a build experience similar to a scripting language.

Not just a wrapper! graphics.gd has been holistically designed and curated from the ground up to provide a cohesive way to interface with the engine.

We would love you to take part in our [active discussions](https://github.com/grow-graphics/gd/discussions)
section with any questions, comments or feedback you may have. Show us what you're building!

```go
// This file is all you need to start a project.
// Save it somewhere, install the `gd` command and use `gd run` to get started.
package main

import (
	"graphics.gd/startup"

	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/SceneTree"
)

func main() {
	startup.LoadingScene() // setup the SceneTree and wait until we have access to engine functionality
	hello := Label.New()
	hello.AsControl().SetAnchorsPreset(Control.PresetFullRect) // expand the label to take up the whole screen.
	hello.SetHorizontalAlignment(Label.HorizontalAlignmentCenter)
	hello.SetVerticalAlignment(Label.VerticalAlignmentCenter)
	hello.SetText("Hello, World!")
	SceneTree.Add(hello)
	startup.Scene() // starts up the scene and blocks until the engine shuts down.
}
```

You can help fund the project, motivate development and prioritise issues [here](https://buy.stripe.com/4gw14maETbnX3vOcMM)

## Getting Started
The module includes a drop-in replacement for the go command called `gd` that
makes it easy to work with projects that run alongside a graphics/game engine.
It enables you to start developing a new project from a single `main.go` file,
to install it, make sure that your `$GOPATH/bin` is in your `$PATH` and run:

```sh
	go install graphics.gd/cmd/gd@master
```

Now when you can run `gd run`, `gd test` on the main package in your project's
directory, things will work as expected. The tool will create a "graphics"
subdirectory where you can manage your assets via the Engine's Editor.

Running the command without any arguments will startup the Engine's Editor.

**NOTE** On linux (and macos if you have brew), `gd` will download an engine for you automatically!
**HINT**  On Windows, you'll want to
[setup CGO](https://github.com/go101/go101/wiki/CGO-Environment-Setup).

If you don't want to use the `gd` command, you can build a shared library with
the `go` command directly:

```sh
go build -o example.so -buildmode=c-shared
```

## Design Principles

Each engine class is available as its own package under `classdb`. To import the
`Node` class you can import `"graphics.gd/classdb/Node"` There's no inheritance,
so to access a 'super' class, you need to call `Super()` on an Extension 'Class'.
All engine classes have methods to cast to any sub-classes they extend for example
`AsObject()` or `AsNode2D()`.

Methods have been renamed to follow Go conventions, so instead of
underscores, methods are named as PascalCase. Keep this in mind when
referring to the Engine documentation.

https://docs.godotengine.org/en/latest/index.html

## Frame-Based Memory Management

Engine references must be 'used' every frame in order to remain alive, otherwise
they will automatically be garbage collected each frame. You shouldn't have to
worry about any memory management as long as you keep Engine references inside an
extension class and don't hold onto references across frames. If you get an
"expired pointer" error, it means either the reference has outlived its frame and
has not been used since or the ownership of the value was transferred to the engine.

The project aims to provide as much memory safety as possible for working with the
Engine, please open an issue if you determine there to be any issues here.

## Recommendations

Start with a main.go file, model your project in Go using structs to represent the
world, space or state of your project. Go is an excellent language for textual
representation. Use the `gd` command to launch the Engine's editor when you want to
create visual representation of your structures. The editor is excellent for importing
media, managing assets and designing the visual and spatial aspects  of a project.

**NOTE:** Don't forget to write tests!

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

## Variant Type Table

| Engine Type        | Convenience Type          | Best Performance Type           |
| ------------------ | ------------------------- | ------------------------------- |
| Variant            | `any`                     | `variant.Any`                   |
| bool               | `bool`                    | `bool`                          |
| int                | `int`                     | `int64`                         |
| float              | `Float.X`                 | `float64`                       |
| String             | `string`                  | `String.Readable`               |
| Vector2            | `Vector2.XY`              | `Vector2.XY`                    |
| Vector2i           | `Vector2i.XY`             | `Vector2i.XY`                   |
| Rect2              | `Rect2.PositionSize`      | `Rect2.PositionSize`            |
| Rect2i             | `Rect2i.PositionSize`     | `Rect2i.PositionSize`           |
| Vector3            | `Vector3.XYZ`             | `Vector3.XYZ`                   |
| Vector3i           | `Vector3i.XYZ`            | `Vector3i.XYZ`                  |
| Transform2D        | `Transform2D`             | `Transform2D`                   |
| Vector4            | `Vector4.XYZW`            | `Vector4.XYZW`                  |
| Vector4i           | `Vector4i.XYZW`           | `Vector4i.XYZW`                 |
| Plane              | `Plane.NormalD`           | `Plane.NormalD`                 |
| Quaternion         | `Quaternion.IJKL`         | `Quaternion.IJKL`               |
| AABB               | `AABB.PositionSize`       | `AABB.PositionSize`             |
| Basis              | `Basis.XYZ`               | `Basis.XYZ`                     |
| Transform3D        | `Transform3D.BasisOrigin` | `Transform3D.BasisOrigin`       |
| Projection         | `Projection.XYZW`         | `Projection.XYZW`               |
| Color              | `Color.RGBA`              | `Color.RGBA`                    |
| StringName         | `string`                  | `String.Name`                   |
| NodePath           | `string`                  | `Path.ToNode`                   |
| Signal             | `chan T`                  | `Signal.Any`                    |
| RID                | `RID.T`                   | `RID.Any`                       |
| Object             | `T.Instance`              | `T.Advanced`                    |
| Callable           | `func(...T) (...T)`       | `Callable.Function`             |
| Dictionary         | `struct/map[T]T`          | `Dictionary.Any`                |
| Array              | `[]T`                     | `Array.Any`                     |
| PackedByteArray    | `[]byte`                  | `Packed.Bytes`                  |
| PackedInt32Array   | `[]int32`                 | `Packed.Array[int32]`           |
| PackedInt64Array   | `[]int64`                 | `Packed.Array[int64]`           |
| PackedFloat32Array | `[]float32`               | `Packed.Array[float32]`         |
| PackedFloat64Array | `[]float64`               | `Packed.Array[float64]`         |
| PackedStringArray  | `[]string`                | `Packed.Strings`                |
| PackedVector2Array | `[]Vector2.XY`            | `Packed.Array[Vector2.XY]`      |
| PackedVector3Array | `[]Vector3.XYZ`           | `Packed.Array[Vector3.XYZ]`     |
| PackedColorArray   | `[]Color.RGBA`            | `Packed.Array[Color.RGBA]`      |
| PackedVector4Array | `[]Vector4.XYZW`          | `Packed.Array[Vector4.XYZW]`    |

## Performance
It's feasible to write high performance code using this module, keep to Engine types where possible and avoid
allocating memory on the heap in frequently called functions. `Advanced` instances are available for each class
which allow more fine-grained control over memory allocations.

Benchmarking shows `Advanced` method calls from Go -> Engine do not allocate in practice.

Allocations are currently unavoidable for any Script -> Go calls (but not
for `Advanced` class virtual method overrides such as `Ready` or `Process`,
which do not allocate in practice).

We've got some ideas to reduce allocations for Script -> Go calls, when
arguments fit entirely within registers. TBA.

## Examples
There are a number of examples in the [examples](https://github.com/grow-graphics/eg)
repo. All examples are designed to be run with `gd run` without any additional setup.

## Supported Platforms

* Windows
* Linux   (including Steam Deck)
* Mac     (including Apple Silicon)
* Android (including MetaQuest)
* IOS     (should work, untested)
* Web     (experimental) `GOOS=js GOARCH=wasm gd run`

## Known Limitations

* No support for Go 'scripts'.
* 64bit only (arm64 && amd64).
* No console support, will likely be achieved in the future with WASM.

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

Lastly, spread the word and let people know about graphics.gd!

## See Also

* [godot-go](https://github.com/godot-go/godot-go) (Another project aiming to support Go + Godot integration)

## Licensing
This project is licensed under an MIT license (the same license as Godot), you can use it in any manner
you can use the Godot engine. If you use this project for any commercially successful products, please
consider [financially supporting us](https://buy.stripe.com/4gw14maETbnX3vOcMM) to show your appreciation!
