// Package path provides methods and a type for working with slash-separated paths with selectors.
package Path

import (
	"path"
	"path/filepath"

	"graphics.gd/variant/String"
)

// ToFile represents a path to a file within a filesystem.
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
type ToFile String.Readable

// AsFile returns the given path as a directory.
func AsFile[T String.Any](path T) ToFile {
	return ToFile(String.New(path))
}

func (s ToFile) String() string               { return String.Readable(s).String() }
func (s ToFile) MarshalText() ([]byte, error) { return String.Readable(s).MarshalText() }
func (s *ToFile) UnmarshalText(text []byte) error {
	return (*String.Readable)(s).UnmarshalText(text)
}

// FileName returns the file name, including the extension, if the string is a valid file path,
func (s ToFile) FileName() string { //gd:String.get_file
	return filepath.Base(s.String())
}

// WithoutExtension returns the path before the last period character (.) in the path.
func (s ToFile) WithoutExtension() ToFile { //gd:String.get_basename
	return ToFile(String.TrimSuffix(s, path.Ext(s.String())))
}

// IsAbsolute returns true if the path's starting point is explicitly defined.
// Unlike a relative path, an absolute path is represented by a leading slash character (/).
// This method is the opposite of [ToFile.IsRelative].
//
// This includes all paths starting with "res://", "user://", "C:\", "/", etc.
func (s ToFile) IsAbsolute() bool { //gd:String.is_absolute_path
	return path.IsAbs(s.String())
}

// IsRelative returns true if the string is a path, and its starting point is dependent on context.
// The path could begin from the current directory, or the current Node (if the string is derived from
// a NodePath), and may sometimes be prefixed with "./". This method is the opposite of [ToFile.IsAbsolute].
func (s ToFile) IsRelative() bool { //gd:String.is_relative_path
	return !path.IsAbs(s.String())
}

// Directory returns the directory if the string is a valid file path.
func (s ToFile) Directory() ToDirectory { //gd:String.get_base_dir
	return ToDirectory(String.New(filepath.Dir(s.String())))
}

// FileExtension returns the file extension without the leading period (.) if the string is a
// valid file name or path. Otherwise, returns an empty string.
func (s ToFile) FileExtension() String.Readable { //gd:String.get_extension
	ext := filepath.Ext(s.String())
	if ext != "" {
		return String.New(String.Slice(ext, 1, String.Length(ext)))
	}
	return String.New()
}
