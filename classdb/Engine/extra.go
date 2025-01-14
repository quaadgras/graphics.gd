package Engine

import (
	"fmt"

	gd "graphics.gd/internal"
)

type (
	Bool   = gd.Bool
	Int    = gd.Int
	FloatX = gd.Float
)

// Version returns the version of the Engine.
func Version() gd.Version {
	return gd.Global.GetGodotVersion()
}

// Println converts one or more arguments of any type to string in the best way
// possible and prints them to the console.
func Println(v ...any) { //gd:print
	gd.Print(gd.NewVariant(fmt.Sprint(v...)))
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
	gd.PrintRich(gd.NewVariant(fmt.Sprint(v...)))
}

// Log prints one or more arguments to strings in the best way possible to standard error line
func Log(v ...any) { //gd:printerr
	gd.Printerr(gd.NewVariant(fmt.Sprint(v...)))
}

// Logv prints if verbose mode is enabled (OS.is_stdout_verbose returning true), converts one or
// more arguments of any type to string in the best way possible and prints them to the console.
func Logv(v ...any) { //gd:print_verbose
	gd.PrintVerbose(gd.NewVariant(fmt.Sprint(v...)))
}

// Print prints one or more arguments to strings in the best way possible to the OS terminal.
// Unlike print, no newline is automatically added at the end.
func Print(v ...any) { //gd:prints printraw printt
	gd.Print(gd.NewVariant(fmt.Sprint(v...)))
}

// Raise pushes an error message to Godot's built-in debugger and to the OS terminal.
func Raise(err error) { //gd:push_error
	gd.PushError(gd.NewVariant(err.Error()))
}

// RaiseWarning pushes a warning message to Godot's built-in debugger and to the OS terminal.
func RaiseWarning(v ...any) { //gd:push_warning
	gd.PushWarning(gd.NewVariant(fmt.Sprint(v...)))
}
