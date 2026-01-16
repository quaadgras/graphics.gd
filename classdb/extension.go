package classdb

import (
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdclass"
)

// ExtensionTo is an interface implemented by extensions to the given engine class.
type ExtensionTo[T gd.IsClass] interface {
	Class
	super() T
}

// ExtensionInherits can be used to enable a Go class to inherit another, this means
// they will show up nested in the Node selection menu within the editor. This is an
// advanced technique and any desired helper methods like AsNode, AsParent are not
// included and will need to be defined as methods on T.
//
// The super class S can be accessed by calling Super().
type ExtensionInherits[S gdclass.Interface, T Class] = gdclass.ExtensionInherits[S, T]

// Super returns the parent class of a given extension class.
func Super[T gd.IsClass](class ExtensionTo[T]) T {
	return class.super()
}
