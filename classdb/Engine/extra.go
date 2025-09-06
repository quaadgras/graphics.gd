package Engine

import (
	"runtime"
	"sync"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/pointers"
)

// Version returns the version of the Engine.
func Version() string {
	return pointers.New[gd.String](gdextension.Host.Version.String()).String()
}

var (
	print_fn,
	print_rich_fn,
	print_err_fn,
	print_verbose_fn,
	push_error_fn,
	push_warning_fn gdextension.FunctionID
	print_setup = sync.OnceFunc(func() {
		print_fn = gdextension.Host.Builtin.Functions.Name(pointers.Get(gd.NewStringName("print")), 2648703342)
		print_rich_fn = gdextension.Host.Builtin.Functions.Name(pointers.Get(gd.NewStringName("print_rich")), 2648703342)
		print_err_fn = gdextension.Host.Builtin.Functions.Name(pointers.Get(gd.NewStringName("printerr")), 2648703342)
		print_verbose_fn = gdextension.Host.Builtin.Functions.Name(pointers.Get(gd.NewStringName("print_verbose")), 2648703342)
		push_error_fn = gdextension.Host.Builtin.Functions.Name(pointers.Get(gd.NewStringName("push_error")), 2648703342)
		push_warning_fn = gdextension.Host.Builtin.Functions.Name(pointers.Get(gd.NewStringName("push_warning")), 2648703342)
	})
)

// Println converts one or more arguments of any type to string in the best way
// possible and prints them to the console.
func Println(v ...any) { //gd:print
	print_setup()
	var variants []gdextension.Variant
	for _, value := range v {
		variants = append(variants, gdextension.Variant(pointers.Get(gd.NewVariant(value))))
	}
	gdextension.Host.Builtin.Functions.Call(print_fn, nil, gdextension.ShapeVariants(len(variants)), gdextension.CallAccepts[any](&variants[0]))
}

// PrintRich converts one or more arguments of any type to string in the best way possible
// and prints them to the console.
//
// The following BBCode tags are supported: b, i, u, s, indent, code, url, center, right,
// color, bgcolor, fgcolor.
//
// Color tags only support the following named colors: black, red, green, yellow, blue,
// magenta, pink, purple, cyan, white, orange, gray. Hexadecimal color codes are not supported.
//
// URL tags only support URLs wrapped by a URL tag, not URLs with a different title.
//
// When printing to standard output, the supported subset of BBCode is converted to ANSI
// escape codes for the terminal emulator to display. Support for ANSI escape codes varies
// across terminal emulators, especially for italic and strikethrough. In standard output,
// code is represented with faint text but without any font change. Unsupported tags are left
// as-is in standard output.
func PrintRich(v ...any) { //gd:print_rich
	print_setup()
	var variants []gdextension.Variant
	for _, value := range v {
		variants = append(variants, gdextension.Variant(pointers.Get(gd.NewVariant(value))))
	}
	gdextension.Host.Builtin.Functions.Call(print_rich_fn, nil, gdextension.ShapeVariants(len(variants)), gdextension.CallAccepts[any](&variants[0]))
}

// Log prints one or more arguments to strings in the best way possible to standard error line
func Log(v ...any) { //gd:printerr
	print_setup()
	var variants []gdextension.Variant
	for _, value := range v {
		variants = append(variants, gdextension.Variant(pointers.Get(gd.NewVariant(value))))
	}
	gdextension.Host.Builtin.Functions.Call(print_err_fn, nil, gdextension.ShapeVariants(len(variants)), gdextension.CallAccepts[any](&variants[0]))
}

// Logv prints if verbose mode is enabled (OS.is_stdout_verbose returning true), converts one or
// more arguments of any type to string in the best way possible and prints them to the console.
func Logv(v ...any) { //gd:print_verbose
	print_setup()
	var variants []gdextension.Variant
	for _, value := range v {
		variants = append(variants, gdextension.Variant(pointers.Get(gd.NewVariant(value))))
	}
	gdextension.Host.Builtin.Functions.Call(print_verbose_fn, nil, gdextension.ShapeVariants(len(variants)), gdextension.CallAccepts[any](&variants[0]))
}

// Print prints one or more arguments to strings in the best way possible to the OS terminal.
// Unlike print, no newline is automatically added at the end.
func Print(v ...any) { //gd:prints printraw printt
	print_setup()
	var variants []gdextension.Variant
	for _, value := range v {
		variants = append(variants, gdextension.Variant(pointers.Get(gd.NewVariant(value))))
	}
	gdextension.Host.Builtin.Functions.Call(print_fn, nil, gdextension.ShapeVariants(len(variants)), gdextension.CallAccepts[any](&variants[0]))
}

// Raise pushes an error message to Godot's built-in debugger and to the OS terminal.
func Raise(err error) { //gd:push_error
	pc, file, line, ok := runtime.Caller(1)
	if ok {
		name := runtime.FuncForPC(pc).Name()
		gdextension.Host.Log.Error(err.Error(), "", name, file, int32(line), false)
	} else {
		print_setup()
		var variants []gdextension.Variant
		variants = append(variants, gdextension.Variant(pointers.Get(gd.NewVariant(err.Error()))))
		gdextension.Host.Builtin.Functions.Call(push_error_fn, nil, gdextension.ShapeVariants(len(variants)), gdextension.CallAccepts[any](&variants[0]))
	}
}

// RaiseWarning pushes a warning message to Godot's built-in debugger and to the OS terminal.
func RaiseWarning(v ...any) { //gd:push_warning
	print_setup()
	var variants []gdextension.Variant
	for _, value := range v {
		variants = append(variants, gdextension.Variant(pointers.Get(gd.NewVariant(value))))
	}
	gdextension.Host.Builtin.Functions.Call(push_warning_fn, nil, gdextension.ShapeVariants(len(variants)), gdextension.CallAccepts[any](&variants[0]))
}
