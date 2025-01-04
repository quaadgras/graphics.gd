// Package NodePath provides methods and a type for working with slash-separated paths with colon selectors.
package NodePath

import (
	"path"
	"strings"
)

// String represents a path to an element within a tree.
//
// A path is composed of slash-separated (/) names and colon-separated (:) selectors.
// ".." and "." are special names. They refer to the parent and the current element
// respectively.
//
// The following examples are paths relative to the current element:
//
//	"A"     			  // Points to the direct child A.
//	"A/B"   			  // Points to A's child B.
//	"."     			  // Points to the current element.
//	".."    			  // Points to the parent element.
//	"../C"  			  // Points to the sibling element C.
//	"../.." 		      // Points to the grandparent element.
//	"/root"     		  // Points to the root element.
//	":Position"           // Points to this items's position.
//	":Position:X"         // Points to this items's position in the x axis.
//	"Camera3D:Rotation:Y" // Points to the child Camera3D and its y rotation.
//	"/root:Size:X"        // Points to the root Window and its width.
type String string

// PrefixedWithSelector returns a copy of this path with a period character (.) prefixed,
// transforming it to a pure selector with no items (relative to the current item).
func (s String) PrefixedWithSelector() String { return ":" + s } //gd:NodePath.get_as_property_path

// WithoutSelector returns the path up until the first selector.
func (s String) WithoutSelector() String { //gd:NodePath.get_concatenated_names
	index := strings.LastIndex(string(s), "/")
	if index == -1 {
		return s
	}
	return s[index:]
}

// Selector returns the selector in the path.
func (s String) Selector() String { //gd:NodePath.get_concatenated_subnames
	index := strings.LastIndex(string(s), ".")
	if index == -1 {
		return ""
	}
	return s[index:]
}

// Element returns the element indicated by idx, starting from 0.
func (s String) Element(i int) String { //gd:NodePath.get_name
	return String(strings.Split(string(s), "/")[i])
}

// NumElement returns the number of elements in the path.
func (s String) NumElement() int { //gd:NodePath.get_name_count
	return len(strings.Split(string(s), "/"))
}

// Name returns the sub-selector indicated by idx, starting from 0.
func (s String) Name(i int) String { //gd:NodePath.get_subname
	return String(strings.Split(string(s.Selector()), ".")[i])
}

// NumName returns the number of sub-selectors in the path.
func (s String) NumName() int { //gd:NodePath.get_subname_count
	return len(strings.Split(string(s.Selector()), "."))
}

// Hash returns the 32-bit hash value representing the paths contents.
//
// Note: paths with equal hash values are not guaranteed to be the same,
// as a result of hash collisions. Paths with different hash values are
// guaranteed to be different.
func (s String) Hash() uint32 { //gd:NodePath.hash
	return String.Hash(s)
}

// IsAbsolute returns true if the path's starting point is explicitly defined.
// Unlike a relative path, an absolute path is represented by a leading slash character (/).
// This method is the opposite of [String.IsRelative].
//
// This includes all paths starting with "res://", "user://", "C:\", "/", etc.
func (s String) IsAbsolute() bool { //gd:NodePath.is_absolute
	return path.IsAbs(string(s))
}

// IsEmpty returns true if the paths's length is 0 ("").
func (s String) IsEmpty() bool { //gd:NodePath.is_empty
	return len(s) == 0
}

// StartingFrom returns a slice of the path from the given start index to the end of the path.
func (s String) StartingFrom(start int) String { //gd:NodePath.slice
	return String(string(s)[start:])
}
