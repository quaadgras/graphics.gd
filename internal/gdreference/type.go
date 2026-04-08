package gdreference

import (
	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
)

type Type int

const (
	TypeUnsafe Type = iota // Raw pointer.
	TypePooled             // main thread owns the pointer.
	TypePinned             // Go controls the lifetime of the pointer.
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
					raw := gdextension.Object(gdunsafe.ObjectID(obj.objectID).Object())
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
