// Package StringName provides unique strings.
package StringName

import (
	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
)

// This package will be replaced with String-variant package based string interning when Go 1.24 is released
// the goal is to use the new weak pointers to automatically maintain a cross-implementation interning table.

var static = make(map[string]gd.StringName)

func init() {
	gd.StartupFunctions = append(gd.StartupFunctions, func() {
		for s, v := range static {
			str := gd.Global.StringNames.New(s)
			raw, ok := pointers.End(str)
			if ok {
				pointers.Set(v, raw)
			}
		}
	})
}

// Advanced StringName for those who know what they are doing.
type Advanced = gd.StringName

// New returns an engine-optimised StringName for use with Advanced functions.
// Can be initialised ahead of time as a global variable.
func New(s string) gd.StringName {
	if v, ok := static[s]; ok {
		return v
	}
	return gd.Global.StringNames.New(s)
}
