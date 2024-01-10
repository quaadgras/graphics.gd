# Godot 4.0 Bindings for Go

This package provides an interface for writing Godot 4.3 Extensions in Go.

Previous effort was towards getting this working with Godot 4.0-beta, currently 
being revamped to work with Godot 4.3. So this is not functional at the moment.

## Design Principles

All Godot classes have been exported and can be referred to by their
standard Godot names, for example `gd.Object` is an 
[Object](https://docs.godotengine.org/en/latest/classes/class_object.html) 
reference.

Methods have been renamed to follow Go conventions, so instead of
underscores, methods are named as PascalCase. Keep this in mind when
referring to the Godot documentation.

https://docs.godotengine.org/en/latest/index.html

Methods and functions always use the equivalent Go primitives (where 
possible), so Go strings will be copied over to the Godot representation 
for them.

## Non Goals

* Godot 3.x support.
* Class Inheritance.

## Performance

No measurements have been taken. Do keep in mind that due to CGO there can be 
a higher cost to make calls to small Godot functions (that do little work). 
So try to do small calculations and tight loops in Go without calling out to Godot. 

Additionally, arguments are passed to Godot as pointers so any function calls
will cause their arguments to escape to the heap even though they do not escape. 

(https://github.com/golang/go/issues/27538)


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