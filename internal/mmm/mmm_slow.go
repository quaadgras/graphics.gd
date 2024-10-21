//go:build debug

// Package mmm provides a way to manually manage memory and resource lifetimes with protections against unsafe double-free and use-after-free errors.
package mmm

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"
)

type freeable struct {
	prv *freeable
	nxt *freeable
	api unsafe.Pointer // nil if root
	end func(genericPointer)

	history []history

	// rev highest bit is set if pinned
	// rev second highest bit is set if reference counted.
	rev revision
}

type history struct {
	kind  string
	stack [10]uintptr
}

func (f *freeable) record(kind string) {
	var stack [10]uintptr
	runtime.Callers(2, stack[:])
	f.history = append(f.history, history{
		kind:  kind,
		stack: stack,
	})
}

func crash(ptr *freeable, name string) {
	s := strings.Builder{}
	s.WriteString(name)
	s.WriteString("\n")
	for i := range ptr.history {
		s.WriteString(ptr.history[i].kind)
		s.WriteString("\n")
		for j := range ptr.history[i].stack {
			pc := ptr.history[i].stack[j]
			if pc == 0 {
				break
			}
			fn := runtime.FuncForPC(pc)
			if fn == nil {
				break
			}
			file, line := fn.FileLine(ptr.history[i].stack[j])
			s.WriteString("\t")
			s.WriteString(file)
			s.WriteString(":")
			s.WriteString(fn.Name())
			s.WriteString(":")
			s.WriteString(fmt.Sprint(line))
			s.WriteString("\n")
		}
	}
	panic(s.String())
}
