# Godot 4.0 Bindings for Go

This package provides an interface for writing Godot 4.0 Extensions in Go. Currently requires Godot ~master to be built with an additional PR
https://github.com/godotengine/godot/pull/65018 (aiming to get this merged).

```go
    package main

    import "github.com/readykit/gd"

    func main() {} // ignore this
    func init() {
        gd.Export(
            new(HelloWorld), // export the extension so that it is available in Godot.
        )
    }

    // HelloWorld is an extension that prints "Hello World" after it 
    // has been placed into a scene and is ready.
    type HelloWorld struct {
        gd.Node2D
        gd.Extension
    }

    // Ready implements the Godot '_ready' interface (virtual method).
    // It will run after the node is added to the scene.
    func (HelloWorld) Ready() {
        fmt.Println("Hello World")
    }
```

## Roadmap

* Stability.
* Import Godot Documentation?
* ownership move semantics.
* go:noescape the C calls.

## Design Principles

All Godot classes have been exported and can be referred to by their
standard Godot names, for example `gd.Object` is an 
[Object](https://docs.godotengine.org/en/latest/classes/class_object.html) 
reference.

They can be instantiated with `gd.New[ObjectName](owner)` and/or `gd.New[ObjectName]At(&dst, owner)`.

Methods have been renamed to follow Go conventions, so instead of
underscores, methods are named as PascalCase. Keep this in mind when
referring to the Godot documentation.

https://docs.godotengine.org/en/latest/index.html

Methods and functions always use the equivalent Go types (where possible), 
so Go strings and slices for example will be copied over to the Godot
representation of them.

The goal is for all single-threaded Go operations to be memory-safe and
to have all memory leaks detectable at runtime (they will cause a panic).
This is achieved with an ownership model `gd.BelongsTo` which is 
responsible for tracing the memory ownership of any Godot objects
referenced in Go. 

## Sub Packages

All Godot classes have been organised into a set of sub packages that alias
into the root `gd` package, prefixes of these classes are stripped so that
for example `gd.AudioEffect` can be used as `audio.Effect`. These packages
serve as a more digestable way to explore the Godot API. Unfortunately the
methods of these classes are not available to view in the Go documentation.

This is a limitation of the Go documentation tooling.

https://github.com/golang/go/issues/44905

## Non Goals

* Godot 3.x support.
* Inheritance.

## Performance

No measurements have been taken. Do keep in mind that due to CGO there can be 
a higher cost to make calls to small Godot functions (that do little work). 
So try to do small calculations and tight loops in Go without calling out to Godot. 

Additionally, arguments are passed to Godot as pointers so any function calls
will cause their arguments to escape to the heap even though they do not escape. 

(https://github.com/golang/go/issues/27538)


## Contributing

I'm no C/C++ whiz so you can help by identifying any CGO/unsafe issues. 
Please open a issue if you do find anything *very unsafe*! 

Alternatively if you have any ideas for performance improvements, be sure
to let me know, and/or submit a pull request!

## Example

Run `make` in the example/src directory or manually build a C-shared library:

```sh
cd example/src
make
```

Now open the example project in Godot from a terminal and you will be able to 
see Go printing things to the console.

## Disclaimer

These bindings are not stable, some functions may panic or corrupt memory.
Consider this package to be in a experimental state until further notice.
Godot 4.0 has not been officially released and its API is subject to change,
this means that this package could break at any time!
