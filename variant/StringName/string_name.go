// Package StringName provides unique strings.
package StringName

import (
	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
)

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

// New returns an engine-optimised StringName for use with Advanced functions.
// Can be initialised ahead of time as a global variable.
func New(s string) gd.StringName {
	if v, ok := static[s]; ok {
		return v
	}
	return gd.Global.StringNames.New(s)
}
