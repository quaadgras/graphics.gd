package gd

import (
	"iter"

	gd "grow.graphics/gd/internal"
	"runtime.link/mmm"
)

// Iter returns an iterator over the items of the [DirAccess].
func (dir DirAccess) Iter() iter.Seq[gd.String] {
	return func(yield func(gd.String) bool) {
		tmp := gd.NewContext(mmm.API(dir.AsPointer()))
		defer tmp.End()
		dir.ListDirBegin()
		defer dir.ListDirEnd()
		for {
			name := dir.GetNext(tmp)
			if name.Len() == 0 {
				break
			}
			if !yield(name) {
				break
			}
		}
	}
}
