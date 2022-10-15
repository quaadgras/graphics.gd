//go:build !generate

package gd

// #include <gdnative_interface.h>
import "C"

import (
	"unsafe"
)

var cOnLoad func() = func() {}

var deferred []func()

//export loadExtension
func loadExtension(p_interface, p_library, p_init unsafe.Pointer) uint8 {
	api = (*cInterface)(p_interface)
	classDB = cClassLibrary(p_library)

	ini := (*cInitialization)(p_init)
	ini.cInit()

	return 1
}

//export initialize
func initialize(userdata unsafe.Pointer, level int64) {
	if level == 2 {
		cOnLoad()

		for _, fn := range deferred {
			fn()
		}
	}
}

//export deinitialize
func deinitialize(userdata unsafe.Pointer, level int64) {}

//export create_instance_func
func create_instance_func(userdata uintptr) uintptr {
	return uintptr(extensionClassDB.get(extensionClassID(userdata)).create())
}

//export free_instance_func
func free_instance_func(userdata, instance uintptr) {
	extensionClassDB.get(extensionClassID(userdata)).delete(loadInstanceID(instance).index)
}

//export get_virtual_index
func get_virtual_index(userdata uintptr, p_name *C.char) uint8 {
	name := C.GoString(p_name)
	method, ok := extensionClassDB.get(extensionClassID(userdata)).lookup(name)
	if !ok {
		return 0
	}
	if method.Index > 255 {
		panic("too many exported methods")
	}
	return uint8(method.Index + 1)
}

//export call_virtual_index
func call_virtual_index(index uint8, instance uintptr, args *unsafe.Pointer, result unsafe.Pointer) {
	id := loadInstanceID(instance)
	extensionClassDB.get(id.class).call(id.index, index, args, result)
}

//export method_call
func method_call(method_userdata uintptr, p_instance uintptr, p_args *unsafe.Pointer, p_argument_count int32, r_return unsafe.Pointer, r_error *cCallError) {
	panic("method_call not implemented")
}

//export method_call_ptr
func method_call_ptr(method_userdata uintptr, p_instance uintptr, p_args *unsafe.Pointer, r_ret unsafe.Pointer) {
	method := exportedMethods[method_userdata-1]
	id := loadInstanceID(p_instance)
	extensionClassDB.get(id.class).call(id.index, uint8(method.Index+1), p_args, r_ret)
}

//export method_args_type
func method_args_type(p_method_userdata uintptr, p_argument int32) cVariantType {
	method := exportedMethods[p_method_userdata-1]
	if p_argument == -1 { // this means the return type.
		if method.Type.NumOut() == 0 {
			return cVariantTypeNil
		}
		return cVariantTypeOf(method.Type.Out(0))
	}
	return cVariantTypeOf(method.Type.In(int(p_argument) + 1)) // +1 because the first argument is the receiver.
}

// dictionary for "constant strings" that should stay in memory for the lifetime of the program.
var dictionary = make(map[string]*C.char)

//export method_args_info
func method_args_info(p_method_userdata uintptr, p_argument int32, r_info *cPropertyInfo) {
	method := exportedMethods[p_method_userdata-1]

	*r_info = cPropertyInfo{} // pedantic in case C doesn't initialize it.

	var name string
	if p_argument == -1 {
		name = method.Result
	} else {
		name = method.Arguments[p_argument]
	}

	if name == "" {
		return
	}

	cname, ok := dictionary[name]
	if !ok {
		cname, ok = cNewString(name)
		if ok {
			panic("cNewString must allocate C memory in this case")
		}
		dictionary[name] = cname
	}

	r_info._type = C.uint32_t(method_args_type(p_method_userdata, p_argument))
	r_info.name = cname
}

//export method_args_meta
func method_args_meta(p_method_userdata uintptr, p_argument int32) cExtensionClassMethodArgumentMetadata {
	return 0
}
