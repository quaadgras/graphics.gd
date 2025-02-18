package gd

import (
	"unsafe"
)

//go:linkname MethodBindPointerCall graphics.gd/startup.method_bind_ptrcall
//go:noescape
func MethodBindPointerCall(method MethodBind, obj [1]Object, arg []unsafe.Pointer, ret unsafe.Pointer)
