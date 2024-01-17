package gdextension

/*
#include "gdextension_interface.h"

uintptr_t classdb_construct_object(uintptr_t fn, uintptr_t p_classname) {
	return (uintptr_t)((GDExtensionInterfaceClassdbConstructObject)fn)((GDExtensionConstStringNamePtr)p_classname);
}

*/
import "C"

import (
	"grow.graphics/gd"
	internal "grow.graphics/gd/internal"
	"runtime.link/api/call"
	"runtime.link/mmm"
)

func linkCGO(API *internal.API) {
	classdb_construct_object := dlsymGD("classdb_construct_object")
	API.ClassDB.ConstructObject = func(ctx gd.Context, name gd.StringName) gd.Object {
		var frame = call.New()
		var obj gd.Object
		obj.SetPointer(mmm.Make[internal.API, internal.Pointer, uintptr](ctx, ctx.API(), uintptr(C.classdb_construct_object(
			C.uintptr_t(uintptr(classdb_construct_object)),
			C.uintptr_t(call.Arg(frame, name.Pointer()).Uintptr()),
		))))
		frame.Free()
		return obj
	}
}
