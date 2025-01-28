package Path

import (
	"path/filepath"

	"graphics.gd/variant/String"
)

// ToDirectory semantically represents a path to what is expected to be a directory/folder
// on the filesystem.
type ToDirectory String.Readable

// AsDirectory returns the given path as a directory.
func AsDirectory[T String.Any](path T) ToDirectory {
	return ToDirectory(String.New(path))
}

func (s ToDirectory) String() string               { return String.Readable(s).String() }
func (s ToDirectory) MarshalText() ([]byte, error) { return String.Readable(s).MarshalText() }
func (s *ToDirectory) UnmarshalText(text []byte) error {
	return (*String.Readable)(s).UnmarshalText(text)
}

// Parent returns the parent directory if the directory is valid.
func (s ToFile) Parent() ToDirectory {
	return ToDirectory(String.New(filepath.Dir(s.String())))
}
