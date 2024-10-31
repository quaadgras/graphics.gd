// Package StringName provides unique strings.
package StringName

import gd "grow.graphics/gd/internal"

// New returns an engine-optimised StringName for use with Advanced functions.
func New(s string) gd.StringName {
	return gd.Global.StringNames.New(s)
}
