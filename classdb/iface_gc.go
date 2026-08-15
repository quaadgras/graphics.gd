//go:build gc && !go1.27

package classdb

import (
	"reflect"
	"sync/atomic"
	"unsafe"

	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/gdextension"
)

// ifaceWords mirrors the gc compiler's layout of a non-empty interface value:
// a type-word (itab) followed by a data-word. This layout is not guaranteed by
// the language spec, so this file is constrained to the gc compiler and to the
// Go versions it has been validated against — after checking a new release,
// bump the !go1.N constraint here and in iface_portable.go. Any other
// compiler or version falls back to reflect in
// [instanceImplementation.Interface].
type ifaceWords struct {
	tab  unsafe.Pointer
	data unsafe.Pointer
}

func (instance *instanceImplementation) cachedInterface(data unsafe.Pointer) (gdclass.Pointer, bool) {
	// Atomic: cacheInterface can run on a goroutine constructing the
	// instance while an engine callback on the main thread reads the cache.
	// The value is the class's static itab either way, so any published
	// value is correct.
	tab := atomic.LoadPointer(&instance.itab)
	if tab == nil {
		return nil, false
	}
	var iface gdclass.Pointer
	*(*ifaceWords)(unsafe.Pointer(&iface)) = ifaceWords{tab: tab, data: data}
	return iface, true
}

func (instance *instanceImplementation) cacheInterface(iface gdclass.Pointer) {
	atomic.StorePointer(&instance.itab, (*ifaceWords)(unsafe.Pointer(&iface)).tab)
}

// instanceID mints the dispatch word for a fresh instance: the address of
// the user's Go struct, pinned so the engine may hold it and hand it back
// on every callback. Pairs with fastInterface, which rebuilds the receiver
// interface from this word plus the class's static itab without any table
// lookup. The pin is released when the engine frees the instance
// (instanceTable.Del) or when Go takes sole ownership of the object
// (ExtensionInstanceGoOnly), whichever comes first.
func instanceID(instance *instanceImplementation, data reflect.Value) gdextension.ExtensionInstanceID {
	ptr := data.UnsafePointer()
	instance.pinner.Pin(ptr)
	return gdextension.ExtensionInstanceID(uintptr(ptr))
}

// repinInstance re-establishes the dispatch-word pin when ownership of an
// object transfers back to the engine (see ExtensionInstanceGoOnly).
func repinInstance(instance *instanceImplementation) {
	iface := instance.strongInterface()
	if iface == nil {
		return
	}
	instance.pinner.Pin((*ifaceWords)(unsafe.Pointer(&iface)).data)
}

// classTab resolves the static itab of (*T, gdclass.Pointer) at class
// registration time, so virtual dispatch can rebuild the receiver
// interface from (itab, instance word) alone.
func classTab(classType reflect.Type) unsafe.Pointer {
	iface, ok := reflect.TypeAssert[gdclass.Pointer](reflect.New(classType))
	if !ok {
		return nil
	}
	return (*ifaceWords)(unsafe.Pointer(&iface)).tab
}

// fastInterface rebuilds the receiver interface for a virtual dispatch
// from the class's static itab and the instance word (the address of the
// user's struct — see instanceID). This is the hot path: no handle table,
// no weak-pointer resolution, no reflection — two register-sourced words.
func fastInterface(tab unsafe.Pointer, id gdextension.ExtensionInstanceID) (gdclass.Pointer, bool) {
	if tab == nil || id == 0 {
		return nil, false
	}
	var iface gdclass.Pointer
	*(*ifaceWords)(unsafe.Pointer(&iface)) = ifaceWords{tab: tab, data: unsafe.Pointer(uintptr(id))}
	return iface, true
}
