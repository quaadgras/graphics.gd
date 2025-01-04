// Package FilePath provides methods and a type for working with slash-separated paths with selectors.
package FilePath

import (
	"path"
	"strings"
)

// String represents a path to a file within a filesystem.
//
// A path is composed of slash-separated (/) names and period-separated (.) extensions.
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
type String string

// WithoutExtension returns the path before the last period character (.) in the path.
func (s String) WithoutExtension() String { //gd:String.get_basename
	return String(strings.TrimSuffix(string(s), path.Ext(string(s))))
}

// IsAbsolute returns true if the path's starting point is explicitly defined.
// Unlike a relative path, an absolute path is represented by a leading slash character (/).
// This method is the opposite of [String.IsRelative].
//
// This includes all paths starting with "res://", "user://", "C:\", "/", etc.
func (s String) IsAbsolute() bool { //gd:String.is_absolute_path
	return path.IsAbs(string(s))
}

// IsRelative returns true if the string is a path, and its starting point is dependent on context.
// The path could begin from the current directory, or the current Node (if the string is derived from
// a NodePath), and may sometimes be prefixed with "./". This method is the opposite of [String.IsAbsolute].
func (s String) IsRelative() bool { //gd:String.is_relative_path
	return !path.IsAbs(string(s))
}
