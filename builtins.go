package gdunsafe

import (
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector4"
)

// builtin methods that are strictly required for graphics.gd to function.
var builtin = Import[struct {
	Array struct {
		size           BuiltinMethod[Array, struct{}, int64]             `hash:"3173160232"`
		resize         BuiltinMethod[Array, struct{ size int64 }, int64] `hash:"848867239"`
		is_read_only   BuiltinMethod[Array, struct{}, bool]              `hash:"3918633141"`
		make_read_only BuiltinMethod[Array, struct{}, struct{}]          `hash:"3218959716"`
	}
	Callable struct {
		get_method          BuiltinMethod[Callable, struct{}, StringName]           `hash:"1825232092"`
		get_bound_arguments BuiltinMethod[Callable, struct{}, Array]                `hash:"4144163970"`
		get_argument_count  BuiltinMethod[Callable, struct{}, int64]                `hash:"3173160232"`
		callv               BuiltinMethod[Callable, struct{ args Array }, Variant]  `hash:"413578926"`
		bindv               BuiltinMethod[Callable, struct{ args Array }, Callable] `hash:"3564560322"`
		call_deferred       BuiltinMethod[Callable, []Variant, Variant]             `hash:"3286317445"`
	}
	Dictionary struct {
		keys           BuiltinMethod[Dictionary, struct{}, Array]             `hash:"4144163970"`
		has            BuiltinMethod[Dictionary, struct{ key Variant }, bool] `hash:"3680194679"`
		clear          BuiltinMethod[Dictionary, struct{}, struct{}]          `hash:"3218959716"`
		sort           BuiltinMethod[Dictionary, struct{}, struct{}]          `hash:"3218959716"`
		erase          BuiltinMethod[Dictionary, struct{ key Variant }, bool] `hash:"1776646889"`
		hash           BuiltinMethod[Dictionary, struct{}, int64]             `hash:"3173160232"`
		size           BuiltinMethod[Dictionary, struct{}, int64]             `hash:"3173160232"`
		is_read_only   BuiltinMethod[Dictionary, struct{}, bool]              `hash:"3918633141"`
		make_read_only BuiltinMethod[Dictionary, struct{}, struct{}]          `hash:"3218959716"`
	}
	PackedByteArray struct {
		resize    BuiltinMethodMutable[PackedArray[byte], struct{ size int64 }, int64] `hash:"848867239"`
		size      BuiltinMethod[PackedArray[byte], struct{}, int64]                    `hash:"3173160232"`
		duplicate BuiltinMethod[PackedArray[byte], struct{}, PackedArray[byte]]        `hash:"851781288"`
	}
	PackedColorArray struct {
		resize BuiltinMethodMutable[PackedArray[Color.RGBA], struct{ size int64 }, int64] `hash:"848867239"`
		size   BuiltinMethod[PackedArray[Color.RGBA], struct{}, int64]                    `hash:"3173160232"`
	}
	PackedFloat32Array struct {
		resize BuiltinMethodMutable[PackedArray[float32], struct{ size int64 }, int64] `hash:"848867239"`
		size   BuiltinMethod[PackedArray[float32], struct{}, int64]                    `hash:"3173160232"`
	}
	PackedFloat64Array struct {
		resize BuiltinMethodMutable[PackedArray[float64], struct{ size int64 }, int64] `hash:"848867239"`
		size   BuiltinMethod[PackedArray[float64], struct{}, int64]                    `hash:"3173160232"`
	}
	Packedint6432Array struct {
		resize BuiltinMethodMutable[PackedArray[int32], struct{ size int64 }, int64] `hash:"848867239"`
		size   BuiltinMethod[PackedArray[int32], struct{}, int64]                    `hash:"3173160232"`
	}
	PackedStringArray struct {
		resize BuiltinMethodMutable[PackedArray[String], struct{ size int64 }, int64] `hash:"848867239"`
		size   BuiltinMethod[PackedArray[String], struct{}, int64]                    `hash:"3173160232"`
	}
	PackedVector2Array struct {
		resize BuiltinMethodMutable[PackedArray[Vector2.XY], struct{ size int64 }, int64] `hash:"848867239"`
		size   BuiltinMethod[PackedArray[Vector2.XY], struct{}, int64]                    `hash:"3173160232"`
	}
	PackedVector3Array struct {
		resize BuiltinMethodMutable[PackedArray[Vector3.XYZ], struct{ size int64 }, int64] `hash:"848867239"`
		size   BuiltinMethod[PackedArray[Vector3.XYZ], struct{}, int64]                    `hash:"3173160232"`
	}
	PackedVector4Array struct {
		resize BuiltinMethodMutable[PackedArray[Vector4.XYZW], struct{ size int64 }, int64] `hash:"848867239"`
		size   BuiltinMethod[PackedArray[Vector4.XYZW], struct{}, int64]                    `hash:"3173160232"`
	}
	Packedint6464Array struct {
		resize BuiltinMethodMutable[PackedArray[int64], struct{ size int64 }, int64] `hash:"848867239"`
		size   BuiltinMethod[PackedArray[int64], struct{}, int64]                    `hash:"3173160232"`
	}
	Signal struct {
		emit    BuiltinMethod[Signal, []Variant, struct{}] `hash:"3286317445"`
		connect BuiltinMethod[Signal, struct {
			callable Callable
			flags    int64
		}, int64] `hash:"979702392"`
		disconnect      BuiltinMethod[Signal, struct{ callable Callable }, struct{}] `hash:"3470848906"`
		get_name        BuiltinMethod[Signal, struct{}, StringName]                  `hash:"1825232092"`
		get_connections BuiltinMethod[Signal, struct{}, Array]                       `hash:"4144163970"`
		get_object      BuiltinMethod[Signal, struct{}, Object]                      `hash:"4008621732"`
	}
	String struct {
		length     BuiltinMethod[String, struct{}, int64]                    `hash:"3173160232"`
		substr     BuiltinMethod[String, struct{ begin, end int64 }, String] `hash:"787537301"`
		casecmp_to BuiltinMethod[String, struct{ other String }, int64]      `hash:"2920860731"`
	}
	StringName struct {
		length     BuiltinMethod[StringName, struct{}, int64]                    `hash:"3173160232"`
		substr     BuiltinMethod[StringName, struct{ begin, end int64 }, String] `hash:"787537301"`
		casecmp_to BuiltinMethod[StringName, struct{ other String }, int64]      `hash:"2920860731"`
	}
}]()

type ArrayOf[T Any] Array

func (a *ArrayOf[T]) Resize(i int) {
	builtin.Array.resize.Call(Array(*a), struct{ size int64 }{int64(i)})
}
func (a *ArrayOf[T]) Index(i int) T {
	return VariantInto[T](Array(*a).Index(i).Copy(false))
}
func (a *ArrayOf[T]) SetIndex(i int, val T) {
	Array(*a).SetIndex(i, VariantFrom(val))
}
func (a *ArrayOf[T]) Len() int {
	return int(builtin.Array.size.Call(Array(*a), struct{}{}))
}
func (a *ArrayOf[T]) IsReadOnly() bool {
	return builtin.Array.is_read_only.Call(Array(*a), struct{}{})
}
func (a *ArrayOf[T]) MakeReadOnly() {
	builtin.Array.make_read_only.Call(Array(*a), struct{}{})
}

func (a *Array) Resize(i int) {
	builtin.Array.resize.Call(*a, struct{ size int64 }{int64(i)})
}
func (a *Array) Len() int {
	return int(builtin.Array.size.Call(*a, struct{}{}))
}
func (a *Array) IsReadOnly() bool {
	return builtin.Array.is_read_only.Call(*a, struct{}{})
}
func (a *Array) MakeReadOnly() {
	builtin.Array.make_read_only.Call(*a, struct{}{})
}
