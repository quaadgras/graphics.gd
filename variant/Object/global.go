package Object

import (
	"iter"

	gd "graphics.gd/internal"
	"graphics.gd/internal/pointers"
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
	object.AsObject()[0].Set(gd.NewStringName(property), gd.NewVariant(value))
}

// Get returns the Variant value of the given property. If the property does not exist,
// this method returns null.
//
// Note: property must be in snake_case when referring to built-in Godot properties.
func Get(object Any, property string) any { //gd:Object.get
	return object.AsObject()[0].Get(gd.NewStringName(property)).Interface()
}

// HasMethod returns true if the given method name exists in the object.
//
// Note: In C#, method must be in snake_case when referring to built-in Godot methods. Prefer using the names exposed in the MethodName class to avoid allocating a new StringName on each call.
func HasMethod(object Any, method string) bool { //gd:Object.has_method
	return object.AsObject()[0].HasMethod(gd.NewStringName(method))
}

// Call calls the method on the object and returns the result.
func Call(object Any, method string, args ...any) any { //gd:Object.call
	var converted []gd.Variant
	for _, arg := range args {
		converted = append(converted, gd.NewVariant(arg))
	}
	result, err := object.AsObject()[0].Call(gd.NewStringName(method), converted...)
	if err != nil {
		panic(err)
	}
	return result.Interface()
}

// InstanceIsValid returns true if the given object instance is valid (the reference has not been
// invalidated and the object has not been freed).
func InstanceIsValid(obj Any) bool { //gd:is_instance_valid
	if !pointers.Bad(obj.AsObject()[0]) {
		return false
	}
	return Instance(obj.AsObject()).ID().Instance() != Nil
}

// GetPropertyList returns a list of all property names in the object.
func GetPropertyList(object Any) []PropertyInfo { //gd:Object.get_property_list
	return gd.ArrayAs[[]PropertyInfo](object.AsObject()[0].GetPropertyList())
}

// SetIndex sets the value at the given index in the object. If the index is out of range or the
// value's type doesn't match, nothing happens.
func SetIndex(object Any, index int, value any) { //gd:Object[]=
	object.AsObject()[0].SetIndex(index, gd.NewVariant(value))
}

// Index returns the Variant value at the given index in the object. If the index is out of range,
func Index(object Any, index int) any { //gd:Object[]
	return object.AsObject()[0].GetIndex(index).Interface()
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
