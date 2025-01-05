package DirAccess

import (
	"iter"
)

// Iter returns an iterator over the items of the [DirAccess].
func (dir Instance) Iter() iter.Seq[string] {
	return func(yield func(string) bool) {
		dir.ListDirBegin()
		defer dir.ListDirEnd()
		for {
			name := dir.GetNext()
			if name == "" {
				break
			}
			if !yield(name) {
				break
			}
		}
	}
}
