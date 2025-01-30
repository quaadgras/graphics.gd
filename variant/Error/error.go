// Package Error provides generic or codes for use as error return values that do not allocate.
package Error

// Code is an error code.
type Code uint8

func New(err error) Code {
	if err == nil {
		return 0
	}
	return Failed
}

// Code constants.
const (
	Failed                  Code = 1
	Unavailable             Code = 2
	Unconfigured            Code = 3
	Unauthorized            Code = 4
	ParametangeCode         Code = 5
	OutOfMemory             Code = 6
	FileNotFound            Code = 7
	FileBadDrive            Code = 8
	FileBadPath             Code = 9
	FileNoPermission        Code = 10
	FileAlreadyInUse        Code = 11
	FileCantOpen            Code = 12
	FileCantWrite           Code = 13
	FileCantRead            Code = 14
	FileUnrecognized        Code = 15
	FileCorrupt             Code = 16
	FileMissingDependencies Code = 17
	FileEof                 Code = 18
	CantOpen                Code = 19
	CantCreate              Code = 20
	QueryFailed             Code = 21
	AlreadyInUse            Code = 22
	Locked                  Code = 23
	Timeout                 Code = 24
	CantConnect             Code = 25
	CantResolve             Code = 26
	ConnectionCode          Code = 27
	CantAcquireResource     Code = 28
	CantFork                Code = 29
	InvalidData             Code = 30
	InvalidParameter        Code = 31
	AlreadyExists           Code = 32
	DoesNotExist            Code = 33
	DatabaseCantRead        Code = 34
	DatabaseCantWrite       Code = 35
	CompilationFailed       Code = 36
	MethodNotFound          Code = 37
	LinkFailed              Code = 38
	ScriptFailed            Code = 39
	CyclicLink              Code = 40
	InvalidDeclaration      Code = 41
	DuplicateSymbol         Code = 42
	ParseCode               Code = 43
	Busy                    Code = 44
	Skip                    Code = 45
	Help                    Code = 46
	Bug                     Code = 47
	PrinterOnFire           Code = 48
)

// Code implements the [error] interface.
func (code Code) Error() string { //gd:or_string
	switch code {
	case Failed:
		return "Failed"
	case Unavailable:
		return "Unavailable"
	case Unconfigured:
		return "Unconfigured"
	case Unauthorized:
		return "Unauthorized"
	case ParametangeCode:
		return "Parameter range or"
	case OutOfMemory:
		return "Out of memory"
	case FileNotFound:
		return "File not found"
	case FileBadDrive:
		return "File bad drive"
	case FileBadPath:
		return "File bad path"
	case FileNoPermission:
		return "File no permission"
	case FileAlreadyInUse:
		return "File already in use"
	case FileCantOpen:
		return "File can't open"
	case FileCantWrite:
		return "File can't write"
	case FileCantRead:
		return "File can't read"
	case FileUnrecognized:
		return "File unrecognized"
	case FileCorrupt:
		return "File corrupt"
	case FileMissingDependencies:
		return "File missing dependencies"
	case FileEof:
		return "File EOF"
	case CantOpen:
		return "Can't open"
	case CantCreate:
		return "Can't create"
	case QueryFailed:
		return "Query failed"
	case AlreadyInUse:
		return "Already in use"
	case Locked:
		return "Locked"
	case Timeout:
		return "Timeout"
	case CantConnect:
		return "Can't connect"
	case CantResolve:
		return "Can't resolve"
	case ConnectionCode:
		return "Connection or"
	case CantAcquireResource:
		return "Can't acquire resource"
	case CantFork:
		return "Can't fork"
	case InvalidData:
		return "Invalid data"
	case InvalidParameter:
		return "Invalid parameter"
	case AlreadyExists:
		return "Already exists"
	case DoesNotExist:
		return "Does not exist"
	case DatabaseCantRead:
		return "Database can't read"
	case DatabaseCantWrite:
		return "Database can't write"
	case CompilationFailed:
		return "Compilation failed"
	case MethodNotFound:
		return "Method not found"
	case LinkFailed:
		return "Link failed"
	case ScriptFailed:
		return "Script failed"
	case CyclicLink:
		return "Cyclic link"
	case InvalidDeclaration:
		return "Invalid declaration"
	case DuplicateSymbol:
		return "Duplicate symbol"
	case ParseCode:
		return "Parse or"
	case Busy:
		return "Busy"
	case Skip:
		return "Skip"
	case Help:
		return "Help"
	case Bug:
		return "Bug"
	case PrinterOnFire:
		return "Printer on fire"
	default:
		return "Unknown or"
	}
}
