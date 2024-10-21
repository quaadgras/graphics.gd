package gd

import (
	gd "grow.graphics/gd/internal"
)

type Object = gd.Object
type RefCounted = gd.RefCounted

var classDB = make(map[string]func(gd.Pointer) any)

func Register(name string, constructor func(gd.Pointer) any) {
	classDB[name] = constructor
}

func init() {
	gd.ObjectAs = func(name string, ptr gd.Pointer) any {
		if constructor, ok := classDB[name]; ok {
			return constructor(ptr)
		}
		var obj gd.Object
		obj.SetPointer(ptr)
		return obj
	}
}
