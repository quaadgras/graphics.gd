package gdenums

// Error implements the [error] interface.
func (err Error) Error() string {
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
