# Godot 4.3 + Go

This module provides a safe and simple way to work with Godot 4.3, in Go.

```
package main

import (
	"fmt"
	"runtime"

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
	godot, classdb, ok := gdextension.Link()
	if !ok {
		return
	}
	gdextension.RegisterClass[HelloWorld](godot, classdb)
}

```

## Design Principles

Godot classes are exported by the `gd` package and can be referred to by 
their standard Godot names, for example `gd.Object` is an 
[Object](https://docs.godotengine.org/en/latest/classes/class_object.html) 
reference.

Methods have been renamed to follow Go conventions, so instead of
underscores, methods are named as PascalCase. Keep this in mind when
referring to the Godot documentation.

https://docs.godotengine.org/en/latest/index.html

## Semi-Automatic Memory Management

Godot types are preferred over Go types, so that allocations can be kept under
control. All values are tied to a [gd.Context] type, which is either:

    (a) a function argument and any values associated with it will be freed
        when the function returns.
    (b) builtin to the class you are extending, and any values associated 
        with it will be freed when the class is destroyed by Godot.

This module aims to offer memory safety for race-free extensions, if you discover
a way to unintentionally do something unsafe (like double free, use-after-free or
a segfault), using methods on types exported by the root `gd` package please open 
an issue. 

## Performance

No serious profiling has been completed, however all Go -> Godot calls are optimised
in a way to avoid almost all allocations. Allocations are currently unavoidable
for all Godot -> Go script calls (but not for virtual overrides, which are essentially
allocation free).

## Example

Run `make` in the example/src directory or manually build a C-shared library:

```sh
cd example/src
make # go build -o ./bin/example.so -buildmode=c-shared
```

**NOTE:** the first time you build a project, it will take a little while to compile.
After this, subsequent builds will be quite a bit faster!

Now open the example project in Godot from a terminal and you will be able to 
see Go printing things to the console.

## Disclaimer

This module is still in development, some functionality may not be available.