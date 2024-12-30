//go:generate go run ./internal/tool/generate
//go:generate go run ./internal/tool/generate/v2
//go:generate go fmt ./...
package gd

import (
	"fmt"
	"math/rand"
	"time"

	gd "graphics.gd/internal"
)

type (
	Variant = gd.Variant

	Bool  = gd.Bool
	Int   = gd.Int
	Float = gd.Float

	Color = gd.Color

	Array      = gd.Array
	Dictionary = gd.Dictionary

	Object = gd.Object

	PackedByteArray    = gd.PackedByteArray
	PackedColorArray   = gd.PackedColorArray
	PackedFloat32Array = gd.PackedFloat32Array
	PackedFloat64Array = gd.PackedFloat64Array
	PackedInt32Array   = gd.PackedInt32Array
	PackedInt64Array   = gd.PackedInt64Array
	PackedStringArray  = gd.PackedStringArray
	PackedVector2Array = gd.PackedVector2Array
	PackedVector3Array = gd.PackedVector3Array
	PackedVector4Array = gd.PackedVector4Array

	Plane = gd.Plane

	Projection = gd.Projection
	Quaternion = gd.Quaternion

	AABB   = gd.AABB
	Rect2  = gd.Rect2
	Rect2i = gd.Rect2i

	RID = gd.RID

	Signal   = gd.Signal
	Callable = gd.Callable

	String     = gd.String
	StringName = gd.StringName
	NodePath   = gd.NodePath

	Basis       = gd.Basis
	Transform2D = gd.Transform2D
	Transform3D = gd.Transform3D

	Vector2  = gd.Vector2
	Vector2i = gd.Vector2i
	Vector3  = gd.Vector3
	Vector3i = gd.Vector3i
	Vector4  = gd.Vector4
	Vector4i = gd.Vector4i
)

// Version returns the version of the Engine.
func Version() gd.Version {
	return gd.Global.GetGodotVersion()
}

func randomize() { //gd:randomize
	rand.Seed(time.Now().UnixNano())
}

func seed(s int) { //gd:seed
	rand.Seed(int64(s))
}

func rand_from_seed(seed int) *rand.Rand { //gd:rand_from_seed
	return rand.New(rand.NewSource(int64(seed)))
}

func weakref(v any) any { return v } //gd:weakref

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

type Error int

const (
	/*Methods that return [enum Error] return [constant OK] when no error occurred.
	  Since [constant OK] has value 0, and all other error constants are positive integers, it can also be used in boolean checks.
	  [b]Example:[/b]
	  [codeblock]
	  var error = method_that_returns_error()
	  if error != OK:
	      printerr("Failure!")

	  # Or, alternatively:
	  if error:
	      printerr("Still failing!")
	  [/codeblock]
	  [b]Note:[/b] Many functions do not return an error code, but will print error messages to standard output.*/
	Ok Error = 0
	/*Generic error.*/
	Failed Error = 1
	/*Unavailable error.*/
	ErrUnavailable Error = 2
	/*Unconfigured error.*/
	ErrUnconfigured Error = 3
	/*Unauthorized error.*/
	ErrUnauthorized Error = 4
	/*Parameter range error.*/
	ErrParameterRangeError Error = 5
	/*Out of memory (OOM) error.*/
	ErrOutOfMemory Error = 6
	/*File: Not found error.*/
	ErrFileNotFound Error = 7
	/*File: Bad drive error.*/
	ErrFileBadDrive Error = 8
	/*File: Bad path error.*/
	ErrFileBadPath Error = 9
	/*File: No permission error.*/
	ErrFileNoPermission Error = 10
	/*File: Already in use error.*/
	ErrFileAlreadyInUse Error = 11
	/*File: Can't open error.*/
	ErrFileCantOpen Error = 12
	/*File: Can't write error.*/
	ErrFileCantWrite Error = 13
	/*File: Can't read error.*/
	ErrFileCantRead Error = 14
	/*File: Unrecognized error.*/
	ErrFileUnrecognized Error = 15
	/*File: Corrupt error.*/
	ErrFileCorrupt Error = 16
	/*File: Missing dependencies error.*/
	ErrFileMissingDependencies Error = 17
	/*File: End of file (EOF) error.*/
	ErrFileEof Error = 18
	/*Can't open error.*/
	ErrCantOpen Error = 19
	/*Can't create error.*/
	ErrCantCreate Error = 20
	/*Query failed error.*/
	ErrQueryFailed Error = 21
	/*Already in use error.*/
	ErrAlreadyInUse Error = 22
	/*Locked error.*/
	ErrLocked Error = 23
	/*Timeout error.*/
	ErrTimeout Error = 24
	/*Can't connect error.*/
	ErrCantConnect Error = 25
	/*Can't resolve error.*/
	ErrCantResolve Error = 26
	/*Connection error.*/
	ErrConnectionError Error = 27
	/*Can't acquire resource error.*/
	ErrCantAcquireResource Error = 28
	/*Can't fork process error.*/
	ErrCantFork Error = 29
	/*Invalid data error.*/
	ErrInvalidData Error = 30
	/*Invalid parameter error.*/
	ErrInvalidParameter Error = 31
	/*Already exists error.*/
	ErrAlreadyExists Error = 32
	/*Does not exist error.*/
	ErrDoesNotExist Error = 33
	/*Database: Read error.*/
	ErrDatabaseCantRead Error = 34
	/*Database: Write error.*/
	ErrDatabaseCantWrite Error = 35
	/*Compilation failed error.*/
	ErrCompilationFailed Error = 36
	/*Method not found error.*/
	ErrMethodNotFound Error = 37
	/*Linking failed error.*/
	ErrLinkFailed Error = 38
	/*Script failed error.*/
	ErrScriptFailed Error = 39
	/*Cycling link (import cycle) error.*/
	ErrCyclicLink Error = 40
	/*Invalid declaration error.*/
	ErrInvalidDeclaration Error = 41
	/*Duplicate symbol error.*/
	ErrDuplicateSymbol Error = 42
	/*Parse error.*/
	ErrParseError Error = 43
	/*Busy error.*/
	ErrBusy Error = 44
	/*Skip error.*/
	ErrSkip Error = 45
	/*Help error. Used internally when passing [code]--version[/code] or [code]--help[/code] as executable options.*/
	ErrHelp Error = 46
	/*Bug error, caused by an implementation issue in the method.
	  [b]Note:[/b] If a built-in method returns this code, please open an issue on [url=https://github.com/godotengine/godot/issues]the GitHub Issue Tracker[/url].*/
	ErrBug Error = 47
	/*Printer on fire error (This is an easter egg, no built-in methods return this error code).*/
	ErrPrinterOnFire Error = 48
)

// Error implements the [error] interface.
func (err Error) Error() string { //gd:error_string
	switch err {
	case Failed:
		return "Failed"
	case ErrUnavailable:
		return "Unavailable"
	case ErrUnconfigured:
		return "Unconfigured"
	case ErrUnauthorized:
		return "Unauthorized"
	case ErrParameterRangeError:
		return "Parameter range error"
	case ErrOutOfMemory:
		return "Out of memory"
	case ErrFileNotFound:
		return "File not found"
	case ErrFileBadDrive:
		return "File bad drive"
	case ErrFileBadPath:
		return "File bad path"
	case ErrFileNoPermission:
		return "File no permission"
	case ErrFileAlreadyInUse:
		return "File already in use"
	case ErrFileCantOpen:
		return "File can't open"
	case ErrFileCantWrite:
		return "File can't write"
	case ErrFileCantRead:
		return "File can't read"
	case ErrFileUnrecognized:
		return "File unrecognized"
	case ErrFileCorrupt:
		return "File corrupt"
	case ErrFileMissingDependencies:
		return "File missing dependencies"
	case ErrFileEof:
		return "File EOF"
	case ErrCantOpen:
		return "Can't open"
	case ErrCantCreate:
		return "Can't create"
	case ErrQueryFailed:
		return "Query failed"
	case ErrAlreadyInUse:
		return "Already in use"
	case ErrLocked:
		return "Locked"
	case ErrTimeout:
		return "Timeout"
	case ErrCantConnect:
		return "Can't connect"
	case ErrCantResolve:
		return "Can't resolve"
	case ErrConnectionError:
		return "Connection error"
	case ErrCantAcquireResource:
		return "Can't acquire resource"
	case ErrCantFork:
		return "Can't fork"
	case ErrInvalidData:
		return "Invalid data"
	case ErrInvalidParameter:
		return "Invalid parameter"
	case ErrAlreadyExists:
		return "Already exists"
	case ErrDoesNotExist:
		return "Does not exist"
	case ErrDatabaseCantRead:
		return "Database can't read"
	case ErrDatabaseCantWrite:
		return "Database can't write"
	case ErrCompilationFailed:
		return "Compilation failed"
	case ErrMethodNotFound:
		return "Method not found"
	case ErrLinkFailed:
		return "Link failed"
	case ErrScriptFailed:
		return "Script failed"
	case ErrCyclicLink:
		return "Cyclic link"
	case ErrInvalidDeclaration:
		return "Invalid declaration"
	case ErrDuplicateSymbol:
		return "Duplicate symbol"
	case ErrParseError:
		return "Parse error"
	case ErrBusy:
		return "Busy"
	case ErrSkip:
		return "Skip"
	case ErrHelp:
		return "Help"
	case ErrBug:
		return "Bug"
	case ErrPrinterOnFire:
		return "Printer on fire"
	default:
		return "Unknown error"
	}
}
