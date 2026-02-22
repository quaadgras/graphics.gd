package Object

import (
	"iter"

	gd "graphics.gd/internal"
	"graphics.gd/internal/gdreference"
)

// Any object.
type Any interface {
	gd.IsClass
	AsObject() [1]gd.Object
}

// Set assigns value to the given property. If the property does not exist or the given
// value's type doesn't match, nothing happens.
//
// Note: property must be in snake_case when referring to built-in Godot properties.
func Set(object Any, property string, value any) { //gd:Object.set
	gd.ObjectSet(object.AsObject()[0], gd.NewStringName(property), gd.NewVariant(value))
}

// Get returns the Variant value of the given property. If the property does not exist,
// this method returns null.
//
// Note: property must be in snake_case when referring to built-in Godot properties.
func Get(object Any, property string) any { //gd:Object.get
	return gd.ObjectGet(object.AsObject()[0], gd.NewStringName(property)).Interface()
}

// HasMethod returns true if the given method name exists in the object.
//
// Note: In C#, method must be in snake_case when referring to built-in Godot methods. Prefer using the names exposed in the MethodName class to avoid allocating a new StringName on each call.
func HasMethod(object Any, method string) bool { //gd:Object.has_method
	return gd.ObjectHasMethod(object.AsObject()[0], gd.NewStringName(method))
}

// Call calls the method on the object and returns the result.
func Call(object Any, method string, args ...any) any { //gd:Object.call
	var converted []gd.Variant
	for _, arg := range args {
		converted = append(converted, gd.NewVariant(arg))
	}
	result, err := gd.ObjectCall(object.AsObject()[0], gd.NewStringName(method), converted...)
	if err != nil {
		panic(err)
	}
	return result.Interface()
}

// InstanceIsValid returns true if the given object instance is valid (the reference has not been
// invalidated and the object has not been freed).
func InstanceIsValid(obj Any) bool { //gd:is_instance_valid
	if gdreference.BadObject(obj.AsObject()[0]) {
		return false
	}
	return Instance(obj.AsObject()).ID().Instance() != Nil
}

// GetPropertyList returns a list of all property names in the object.
func GetPropertyList(object Any) []PropertyInfo { //gd:Object.get_property_list
	return gd.ArrayAs[[]PropertyInfo](gd.ObjectGetPropertyList(object.AsObject()[0]))
}

// SetIndex sets the value at the given index in the object. If the index is out of range or the
// value's type doesn't match, nothing happens.
func SetIndex(object Any, index int, value any) { //gd:Object[]=
	gd.ObjectSetIndex(object.AsObject()[0], index, gd.NewVariant(value))
}

// Index returns the Variant value at the given index in the object. If the index is out of range,
func Index(object Any, index int) any { //gd:Object[]
	return gd.ObjectGetIndex(object.AsObject()[0], index).Interface()
}

// Iter returns an iterator over the elements of an Object that implements Iterable.
func Iter(object Any) iter.Seq[any] {
	iter := gd.NewVariant(object).Iterator()
	return func(yield func(any) bool) {
		for iter.Next() {
			if !yield(iter.Value()) {
				return
			}
		}
	}
}

// Aliases returns true if a and b refer to the same object instance.
func Aliases(a, b Any) bool {
	return gdreference.GetObject(a.AsObject()[0]) == gdreference.GetObject(b.AsObject()[0])
}

// Leak prevents an engine reference from being invalidated until [Free] is
// called on it.
//
// The typical use is to replicate GDScript leak-by-default semantics by
// immediately wrapping a newly created object with this function.
//
//	var obj = Object.Leak(Object.New())
func Leak[T Any](obj T) T {
	_, kind := gdreference.AskObject(obj.AsObject()[0])
	switch kind {
	case gdreference.TypePooled:
		gdreference.PinObject(obj.AsObject()[0])
		fallthrough
	case gdreference.TypePinned, gdreference.TypeStatic, gdreference.TypeUnsafe:
		return obj
	default:
		panic("Object.Leak called on a pointer owned by the engine")
	}
}

// Free immediately invalidates an object reference, enabling any resources
// associated with it to be released, any subsequent use of the object may
// result in a panic. May not have any effect if the object is still in use
// by the engine.
//
// Free is safe to call at any time on any object (invalidated or not).
func Free(obj Any) {
	if obj == nil {
		return
	}
	ptr := obj.AsObject()[0]
	if gdreference.BadObject(ptr) {
		return
	}
	raw, kind := gdreference.AskObject(ptr)
	switch kind {
	case gdreference.TypePooled:
		gd.ObjectFree(ptr)
	case gdreference.TypePinned:
		if gd.ExtensionInstanceLookup(raw) == nil {
			gd.ObjectFree(ptr)
		}
	}
}
