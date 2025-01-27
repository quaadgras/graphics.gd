// Package Path provides methods and a type for working with slash-separated path strings.
package Path

import (
	"path"
	"slices"

	"graphics.gd/variant/String"
)

// ToNode represents a path to a generic node within a tree.
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
type ToNode String.Readable

// String implements the [fmt.Stringer] interface.
func (s ToNode) String() string { return String.Readable(s).String() }

// PrefixedWithSelector returns a copy of this path with a period character (.) prefixed,
// transforming it to a pure selector with no items (relative to the current item).
func (s ToNode) PrefixedWithSelector() ToNode { return String.Append(ToNode(String.New(":")), s) } //gd:NodePath.get_as_property_path

// WithoutSelector returns the path up until the first selector.
func (s ToNode) WithoutSelector() ToNode { //gd:NodePath.get_concatenated_names
	index := String.FindLast(s, "/")
	if index == -1 {
		return s
	}
	return String.Slice(s, index, String.Length(s))
}

// Selector returns the selector in the path.
func (s ToNode) Selector() String.Readable { //gd:NodePath.get_concatenated_subnames
	index := String.FindLast(s, ".")
	if index == -1 {
		return String.New()
	}
	return String.Readable(String.Slice(s, index, String.Length(s)))
}

// Element returns the element indicated by idx, starting from 0.
func (s ToNode) Element(i int) String.Readable { //gd:NodePath.get_name
	return String.Readable(slices.Collect(String.Splits(s, "/"))[i])
}

// NumElement returns the number of elements in the path.
func (s ToNode) NumElement() int { //gd:NodePath.get_name_count
	return len(slices.Collect(String.Splits(s, "/"))) // FIXME optimise
}

// Name returns the sub-selector indicated by idx, starting from 0.
func (s ToNode) Name(i int) String.Readable { //gd:NodePath.get_subname
	return String.Readable(slices.Collect(String.Splits(s.Selector(), "."))[i])
}

// NumName returns the number of sub-selectors in the path.
func (s ToNode) NumName() int { //gd:NodePath.get_subname_count
	return len(slices.Collect(String.Splits(s.Selector(), "."))) // FIXME optimise
}

// Hash returns the 32-bit hash value representing the paths contents.
//
// Note: paths with equal hash values are not guaranteed to be the same,
// as a result of hash collisions. Paths with different hash values are
// guaranteed to be different.
func (s ToNode) Hash() uint32 { //gd:NodePath.hash
	return ToNode.Hash(s)
}

// IsAbsolute returns true if the path's starting point is explicitly defined.
// Unlike a relative path, an absolute path is represented by a leading slash character (/).
// This method is the opposite of [ToNode.IsRelative].
//
// This includes all paths starting with "res://", "user://", "C:\", "/", etc.
func (s ToNode) IsAbsolute() bool { //gd:NodePath.is_absolute
	return path.IsAbs(s.String())
}

// IsEmpty returns true if the paths's length is 0 ("").
func (s ToNode) IsEmpty() bool { //gd:NodePath.is_empty
	return String.Length(s) == 0
}

// StartingFrom returns a slice of the path from the given start index to the end of the path.
func (s ToNode) StartingFrom(start int) ToNode { //gd:NodePath.slice
	return String.Slice(s, start, String.Length(s))
}
