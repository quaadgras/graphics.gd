package gdclass

import (
	gd "graphics.gd/internal"
	"graphics.gd/internal/gdreference"
)

type Locked struct{}

type Object = gdreference.Object
type RefCounted = gd.RefCounted

var classDB = make(map[string]func(gdreference.Object) any)

func Register(name string, constructor func(gdreference.Object) any) {
	classDB[name] = constructor
}

func init() {
	gd.ObjectAs = func(name string, ptr gdreference.Object) any {
		if constructor, ok := classDB[name]; ok {
			return constructor(ptr)
		}
		return ptr
	}
}
