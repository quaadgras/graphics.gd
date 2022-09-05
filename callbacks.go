//go:build !generate

package gd

import "C"

import (
	"fmt"
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
	if level == 4 {
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
	fmt.Println("lookup", name, ok)
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
