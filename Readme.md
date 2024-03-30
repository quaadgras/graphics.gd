# Godot 4.2.1 + Go [![Go Reference](https://pkg.go.dev/badge/grow.graphics/gd.svg)](https://pkg.go.dev/grow.graphics/gd)

This module provides a safe and simple way to work with Godot 4.2.1, in Go via the GDExtension interface.

You can support the project and prioritise issues [here](https://buy.stripe.com/4gw14maETbnX3vOcMM)

```go
package main

import (
	"fmt"

	"grow.graphics/gd"
	"grow.graphics/gd/gdextension"
)

type HelloWorld struct {
	gd.Class[HelloWorld, gd.Node2D]
}

// Ready implements the Godot Node2D _ready interface (virtual function).
func (h *HelloWorld) Ready(gd.Context) {
	fmt.Println("Hello World from Go!")
}

func main() {
	godot, ok := gdextension.Link()
	if !ok {
		return
	}
	gdextension.RegisterClass[HelloWorld](godot)
}

```

## Getting Started
The module includes a drop-in replacement for the go command called `gd` that 
makes it easy to work with projects that run within Godot. It enables you to
start developing a new project from a single main.go file, to install it, make
sure that your `$GOPATH/bin` is in your `$PATH` and run:

```sh
	go install grow.graphics/gd/cmd/gd@master
```

Now when you can run `gd run`, `gd test` on the main package in your project's 
directory, things will work as expected. The tool will create a "graphics" 
subdirectory where you can manage your assets via the Godot Editor. 

Running the command without any arguments will startup the editor. 

**NOTE** On linux, `gd` will download Godot for you automatically!

## Design Principles

Godot classes are exported by the `gd` package and can be referred to by 
their standard Godot names, for example `gd.Object` is an 
[Object](https://docs.godotengine.org/en/latest/classes/class_object.html) 
reference. There's no inheritance, so to access the 'super' class, you need 
to call `Super()` on your custom 'Class'. All Godot classes have methods
to cast to the classes they extend for example `AsObject()` or `AsNode2D()`.

Methods have been renamed to follow Go conventions, so instead of
underscores, methods are named as PascalCase. Keep this in mind when
referring to the Godot documentation.

https://docs.godotengine.org/en/latest/index.html

## Semi-Automatic Memory Management

Godot types are preferred over Go types, in order to keep allocations optional. 
All values are tied to a [gd.Context] type, which is either:

    (a) a function argument and any values associated with it will be freed
        when the function returns.
    (b) builtin to the class you are extending, and any values associated 
        with it will be freed when the class is destroyed by Godot.

This module aims to offer memory safety for race-free extensions, if you discover
a way to unintentionally do something unsafe (like double free, use-after-free or
a segfault), using methods on types exported by the root `gd` package please open 
an issue. 

## Recommendations

Start with a main.go file, model your project in Go using structs to represent the 
world, space or state of your project. Go is an excellent language for textual 
representation. Use the `gd` command to launch the Godot editor when you want to 
create visual representation of your structures. The Godot editor is an excellent 
tool for importing media, managing assets and designing the visual and spatial aspects 
of a project. Don't forget to write tests!

## Where Do I Find?

```
* Godot Class            -> gd.{ClassName}
* Godot Class Method     -> gd.{ClassName}.{pascal(MethodName)}
* Godot Utility Function -> gd.Context.{pascal(UtilityName)} OR gd.{pascal(UtilityName)} (pure)
* Godot Enum             -> gd.{EnumName}
* Godot Enum Value       -> gd.{EnumName}{EnumValueName}
* Godot Singleton        -> gd.{ClassName}(gd.Context) // function returns the singleton, they cannot be stored.
```

## Performance
It's feasible to write high performance code using this module, keep to Godot types where possible and avoid escaping memory to the heap in frequently called functions. 

### Zero Allocations
Benchmarking shows method calls from Go -> Godot do not allocate in practice. 

Allocations are currently unavoidable for GDScript -> Go calls (but not 
for class virtual method overrides such as `Ready` or `Process`, which 
should be allocation free).

We've got some ideas to reduce allocations for GDScript -> Go calls, when
arguments fit entirely within registers. TBA.

## Examples

Run `make` in the example/src directory or manually build a C-shared library:

```sh
cd example/src
make # go build -o ./bin/example.so -buildmode=c-shared
```

**NOTE:** the first time you build a project, it will take a little while to compile.
After this, subsequent builds will be quite a bit faster!

Now open the example project in Godot from a terminal and you will be able to 
see Go printing things to the console.

There is also a [pong](https://github.com/grow-graphics/eg/blob/master/2d/pong/pong.go)
example in the [examples](https://github.com/grow-graphics/eg) repo.

## Testing
To run the go tests for this module `cd internal && gd test`.

## Known Limitations

* No support for indexed properties
* No support for Godot functions with varargs.
* No support for script extensions.
* 64bit support only.
* Not tested on Windows.
* No planned support for Web export or proprietary consoles.
* Android/MetaQuest not working right now (support is planned).

## See Also

* [Our pure Go zero-dependencies port of Godot's Color Class](https://github.com/grow-graphics/uc)
* [Our pure Go zero-dependencies port of Godot's Math Variants](https://github.com/grow-graphics/xy)
* [godot-go](https://github.com/godot-go/godot-go) (Another project aiming to support Go + Godot integration)

## Licensing
This project is licensed under an MIT license (the same license as Godot), you can use 
it in any manner you can use the Godot engine. If you use this for a commercially successful
project, please consider [financially supporting us](https://buy.stripe.com/4gw14maETbnX3vOcMM).
