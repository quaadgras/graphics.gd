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
