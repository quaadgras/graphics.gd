//go:build !generate

package gd

import (
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"strings"

	gdunsafe "graphics.gd"
)

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

func recovery(err any) {
	if traceALL || traceSystem {
		gdunsafe.Log(gdunsafe.LogError, fmt.Sprint(err, "\n", string(debug.Stack())), "", "gdextension.recovery", "err.go", 18, true)
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
		gdunsafe.Log(gdunsafe.LogError, fmt.Sprint(err), "", name, file, int32(line), true)
	}
}
