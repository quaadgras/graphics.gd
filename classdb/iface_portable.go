//go:build !gc || go1.28

package classdb

import (
	"reflect"
	"sync/atomic"
	"unsafe"

	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/gdextension"
)

// Portable fallback for compilers or Go versions where the interface layout
// assumed by iface_gc.go has not been validated: the itab cache is disabled
// and [instanceImplementation.Interface] always goes through reflect.
func (instance *instanceImplementation) cachedInterface(data unsafe.Pointer) (gdclass.Pointer, bool) {
	return nil, false
}

func (instance *instanceImplementation) cacheInterface(iface gdclass.Pointer) {}

// nextInstanceID mints opaque dispatch words on builds where the interface
// layout is unknown: virtual dispatch cannot rebuild the receiver from the
// word, so every path resolves through the instances table instead and no
// pin is taken.
var nextInstanceID atomic.Uintptr

func instanceID(instance *instanceImplementation, data reflect.Value) gdextension.ExtensionInstanceID {
	return gdextension.ExtensionInstanceID(nextInstanceID.Add(1))
}

func repinInstance(instance *instanceImplementation) {}

func classTab(classType reflect.Type) unsafe.Pointer { return nil }

func fastInterface(tab unsafe.Pointer, id gdextension.ExtensionInstanceID) (gdclass.Pointer, bool) {
	return nil, false
}
