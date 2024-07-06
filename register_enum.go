package gd

import (
	"reflect"

	gd "grow.graphics/gd/internal"
)

type integer interface {
	~int | ~uint | ~int8 | ~int16 | ~int32 | ~int64 | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Enum can be registered, usually within a OnRegister method.
//
// Example:
//
//	func (*MyClass) OnRegister(godot gd.Lifetime) {
//		godot.Register(gd.Enum[MyClass, int]{
//			Name: "MyEnum",
//			Values: map[string]int{
//				"Value1": 1,
//				"Value2": 2,
//			},
//		})
//	}
type Enum[Class gd.IsClass, T integer] struct {

	// Name of the enum as it will appear in Godot.
	Name string

	// Bitfield indicates that the enum is a bitfield.
	Bitfield bool

	// Values maps the name of each value to the
	// integer value it represents.
	Values map[string]T
}

// Register the enum with Godot.
func (enum Enum[Class, T]) Register(godot Lifetime) {
	var classType = reflect.TypeOf([0]Class{}).Elem()

	// Support 'gd' tag for renaming the class within Godot.
	var rename = classType.Name()
	for i := 0; i < classType.NumField(); i++ {
		field := classType.Field(i)
		if field.Name == "Class" {
			if val := field.Tag.Get("gd"); val != "" {
				rename = val
			}
		}
	}
	var className = godot.StringName(rename)

	tmp := gd.NewLifetime(godot.API)
	defer tmp.End()

	enumName := tmp.StringName(enum.Name)

	for name, value := range enum.Values {
		valueName := tmp.StringName(name)
		godot.API.ClassDB.RegisterClassIntegerConstant(godot.API.ExtensionToken, className, enumName, valueName, int64(value), enum.Bitfield)
	}
}
