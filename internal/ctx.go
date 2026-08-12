//go:build !generate

package gd

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strings"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdmemory"
	"graphics.gd/internal/pointers"
)

// String returns a [String] from a standard UTF8 Go string.
func NewString(s string) String {
	return pointers.New[String](gdextension.Host.Strings.Decode.UTF8(s))
}

// StringName returns a [StringName] from a standard UTF8 Go string.
func NewStringName(s string) StringName {
	return pointers.New[StringName](gdextension.Host.Strings.Intern.UTF8(s))
}

var traceALL = os.Getenv("GOTRACEBACK") == "all" || os.Getenv("GOTRACEBACK") == "1"
var traceSystem = os.Getenv("GOTRACEBACK") == "system"
var traceCrash = os.Getenv("GOTRACEBACK") == "crash"

func Recover() {
	if !traceCrash {
		if err := recover(); err != nil {
			recovery(err)
		}
	}
}

// RecoverCall behaves like [Recover] but is for callbacks that report their
// outcome to the engine through a [gdextension.CallError]. The engine leaves
// that struct uninitialized before calling us (see CallableCustomExtension::call
// and ScriptInstanceExtension::callp), so a recovered panic that returned
// without writing it would have the engine report a bogus second error on top
// of the panic we already logged, such as "Method expected 4 argument(s), but
// called with 4." Report the call as OK; the panic itself is the error.
func RecoverCall(call_error gdextension.Returns[gdextension.CallError]) {
	if !traceCrash {
		if err := recover(); err != nil {
			gdmemory.Set(gdextension.Pointer(call_error), gdextension.CallError{})
			recovery(err)
		}
	}
}

func recovery(err any) {
	if traceALL || traceSystem {
		gdextension.Host.Log.Error(fmt.Sprint(err, "\n", string(debug.Stack())), "", "gdextension.recovery", "err.go", 18, true)
	} else {
		name, file, line := "", "", 0
		var buf [10]uintptr
		for i := range runtime.Callers(0, buf[:]) {
			pc := buf[i]
			if pc == 0 {
				break
			}
			fn := runtime.FuncForPC(pc)
			name = fn.Name()
			if strings.HasPrefix(name, "runtime.") || strings.HasPrefix(name, "graphics.gd") {
				continue
			}
			file, line = fn.FileLine(pc)
			break
		}
		gdextension.Host.Log.Error(fmt.Sprint(err), "", name, file, int32(line), true)
	}
}
