# Graphics GD [![Go Reference](https://pkg.go.dev/badge/grow.graphics/gd.svg)](https://pkg.go.dev/grow.graphics/gd)

This module provides a safe performant way to work with graphics and game development in Go via the GDExtension
interface of a supported graphics/game engine. So far, Godot 4.3 is the only officially supported engine.

You can support the project and prioritise issues [here](https://buy.stripe.com/4gw14maETbnX3vOcMM)

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
	startup.Wait() // wait for the graphics/game engine to start up.
	hello := Label.New()
	hello.AsControl().SetAnchorsPreset(Control.PresetFullRect) // expand the label to take up the whole screen.
	hello.SetHorizontalAlignment(Label.HorizontalAlignmentCenter)
	hello.SetVerticalAlignment(Label.VerticalAlignmentCenter)
	hello.SetText("Hello, World!")
	SceneTree.Add(hello)
}
```

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

## Recommendations

Start with a main.go file, model your project in Go using structs to represent the
world, space or state of your project. Go is an excellent language for textual
representation. Use the `gd` command to launch the Engine's editor when you want to
create visual representation of your structures. The editor is excellent for importing
media, managing assets and designing the visual and spatial aspects  of a project.

**NOTE:** Don't forget to write tests!

## Where Do I Find?

```
* Engine Class           -> graphics.gd/classdb/{ClassName}
* Engine Class Method    -> graphics.gd/classdb/{ClassName}.Instance.{pascal(MethodName)}
* Utility Functions      -> graphics.gd/variant/* and/or use the Go standard library.
* Enum                   -> graphics.gd/classdb/{ClassName}.{EnumName}
* Enum Value             -> graphics.gd/classdb/{ClassName}.{EnumName}{EnumValueName}
* Singletons             -> graphics.gd/classdb/{ClassName}.{pascal(MethodName)} // package-level functions.
```

## Performance
It's feasible to write high performance code using this module, keep to Engine types where possible and avoid
memory to the heap in frequently called functions. `Advanced` instances are available for each class which allow
more fine-grained control over memory allocation.

### Zero Allocations
Benchmarking shows method calls from Go -> Engine do not allocate in practice.

Allocations are currently unavoidable for Script -> Go calls (but not
for class virtual method overrides such as `Ready` or `Process`, which
should be allocation free).

We've got some ideas to reduce allocations for Script -> Go calls, when
arguments fit entirely within registers. TBA.

## Examples
There are a number of examples in the [examples](https://github.com/grow-graphics/eg)
repo. All examples are designed to be run with `gd run` without any additional setup.

## Testing
To run the go tests for this module `cd internal && gd test`.

## Supported Platforms

* Windows
* Linux   (including Steam Deck)
* Mac     (including Apple Silicon)
* Android (including MetaQuest)
* IOS     (should work, untested)

## Known Limitations

* No support for calling Engine class methods that take varargs.
* No support for script extensions.
* 64bit support only.
* No Web Export Yet (we have some ideas on how to work towards this)
* No planned support for proprietary consoles.

## See Also

* [godot-go](https://github.com/godot-go/godot-go) (Another project aiming to support Go + Godot integration)

## Licensing
This project is licensed under an MIT license (the same license as Godot), you can use
it in any manner you can use the Godot engine. If you use this for a commercially successful
project, please consider [financially supporting us](https://buy.stripe.com/4gw14maETbnX3vOcMM).
