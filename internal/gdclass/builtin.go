package gdclass

import (
	gd "graphics.gd/internal"
)

type Object = gd.Object
type RefCounted = gd.RefCounted

var classDB = make(map[string]func(gd.Object) any)

func Register(name string, constructor func(gd.Object) any) {
	classDB[name] = constructor
}

func init() {
	gd.ObjectAs = func(name string, ptr gd.Object) any {
		if constructor, ok := classDB[name]; ok {
			return constructor(ptr)
		}
		return ptr
	}
}
