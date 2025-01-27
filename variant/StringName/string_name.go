// Package StringName provides unique strings.
package StringName

import (
	"graphics.gd/variant/String"
)

// This package will be replaced with String-variant package based string interning when Go 1.24 is released
// the goal is to use the new weak pointers to automatically maintain a cross-implementation interning table.

// Advanced StringName for those who know what they are doing.
type Advanced = String.Name

// New returns an engine-optimised StringName for use with Advanced functions.
// Can be initialised ahead of time as a global variable.
func New(s string) String.Name {
	return String.Name(String.New(s))
}
