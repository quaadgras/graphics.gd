package Object

import gd "graphics.gd/internal"

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

func CallV(object Any, method string, args ...any) any { //gd:Object.callv
	array := gd.NewArray()
	for _, arg := range args {
		array.PushBack(gd.NewVariant(arg))
	}

	return object.AsObject()[0].Callv(gd.NewStringName(method), array)
}
