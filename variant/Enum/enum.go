// Package Enum provides a way to define enums.
//
//	type MyEnum Enum.Int[struct {
//		One MyEnum `gd:"ONE"`
//		Two MyEnum `gd:"TWO"`
//	}]
//
//	var MyEnums = Enum.Values[MyEnum]()
package Enum

import (
	"errors"
	"reflect"

	"graphics.gd/variant/String"
)

// Int is an enum backed by an increasing integer. T should be a struct
// with fields of the defined type.
type Int[T any] struct {
	methods[T]
}

// Any integer enum should implement this interface.
type Any interface {
	Int() int
	Enum(yield func(string, int) bool)
}

// Pointer to any enum.
type Pointer interface {
	Any

	SetInt(int)
}

type methods[V any] int

func (m methods[V]) Int() int      { return int(m) }
func (m *methods[V]) SetInt(i int) { *m = methods[V](i) }
func (m methods[V]) Enum(yield func(string, int) bool) {
	rtype := reflect.TypeFor[V]()
	for i := 0; i < rtype.NumField(); i++ {
		field := rtype.Field(i)
		name, ok := field.Tag.Lookup("gd")
		if !ok {
			name = String.ToUpper(String.ToSnakeCase(field.Name))
		}
		if !yield(name, i) {
			break
		}
	}
}

// String returns the name of the enum value or an empty string
// if the value is not defined.
func (m methods[V]) String() string {
	for name, value := range m.Enum {
		if value == m.Int() {
			return name
		}
	}
	return ""
}

// MarshalText implements encoding.TextMarshaler.
func (m methods[V]) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (m *methods[V]) UnmarshalText(text []byte) error {
	for name, value := range m.Enum {
		if name == string(text) {
			*m = methods[V](value)
			return nil
		}
	}
	return errors.New("invalid enum value: " + string(text))
}

type isEnum[T any] interface {
	~struct {
		methods[T]
	}
}

// Values returns the available values for an enum. It should be stored and reused inside a global variable.
func Values[T isEnum[V], V any]() V {
	var values V
	rvalue := reflect.ValueOf(&values).Elem()
	for i := 0; i < rvalue.NumField(); i++ {
		rvalue.Field(i).Set(reflect.ValueOf(T(struct{ methods[V] }{methods[V](i)})).Convert(rvalue.Field(i).Type()))
	}
	return values
}
