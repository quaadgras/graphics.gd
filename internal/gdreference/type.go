package gdreference

import (
	"graphics.gd/internal/gdextension"
)

type Type int

const (
	TypeUnsafe Type = iota // Raw pointer.
	TypePooled             // main thread owns the pointer.
	TypePinned             // main thread owns the pointer, Kept alive.
	TypeThread             // off-thread owns the pointer.
	TypeBorrow             // Engine owns the pointer.
	TypeStatic             // Allocated at init.
)

func GC(free func(gdextension.Object)) {
	Barrier()
	for chunk := range pool {
		for i := range pool[chunk] {
			obj := &pool[chunk][i]
			if obj.objectID == 0 {
				continue
			}
			if obj.inEngine == 0 {
				if obj.objectID != 0 {
					raw := gdextension.Host.Objects.Lookup(obj.objectID)
					*obj = object{}
					free(raw)
				}
				pool_free = append(pool_free, obj)
			} else if obj.objectID != 0 {
				obj.inEngine = 0
			}
		}
	}
}
