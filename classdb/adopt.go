package classdb

import (
	"reflect"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdclass"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdreference"
	"graphics.gd/internal/pointers"
)

func init() {
	gd.ExtensionInstanceAdopt = adoptForReload
}

// adoptForReload wraps an existing engine object of the given registered
// class in a fresh Go instance and rebinds the object's extension
// instance to it. It exists for the hot-reload host (see
// graphics.gd/startup with -tags reloads): after a module swap, engine
// objects created by the previous module are handed off to the new one
// so they keep dispatching into live Go code. The Go-side state of the
// adopted instance starts from the class's zero value.
func adoptForReload(class gdextension.ExtensionClassID, raw gdextension.Object) gdextension.ExtensionInstanceID {
	impl := classes.Get(class)
	if impl == nil {
		return 0
	}
	value := impl.Constructor()
	super := (*gdreference.Object)(value.UnsafePointer())
	gdreference.PinObject(super, raw)
	instance := impl.reloadInstance(value, super)
	id := instances.New(instance, value)
	gdextension.Host.Objects.Extension.Setup(raw, pointers.Get(impl.Name), id)
	if keepalive := compile_keepalive(reflect.PointerTo(impl.Type)); keepalive != nil {
		roots.Insert(value, keepalive)
	}
	// The engine owns adopted objects (they were alive before this
	// module existed), so the wrapper is rooted strongly and no
	// free-on-collect cleanup is attached.
	s, _ := reflect.TypeAssert[gdclass.Pointer](value)
	instance.setStrong(s)
	for _, field := range impl.Singletons {
		if singleton, ok := singletons.Lookup(field.Type.Elem()); ok {
			value.Elem().FieldByIndex(field.Index).Set(singleton)
		}
	}
	instance.OnCreate(value)
	return id
}
