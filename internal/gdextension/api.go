// Package gdextension is the graphics.gd authorative Go representation of the Godot C GDExtension API.
package gdextension

import (
	"unsafe"

	gdunsafe "graphics.gd"
	"graphics.gd/variant/AABB"
	"graphics.gd/variant/Basis"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Plane"
	"graphics.gd/variant/Projection"
	"graphics.gd/variant/Quaternion"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Rect2i"
	"graphics.gd/variant/Transform2D"
	"graphics.gd/variant/Transform3D"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector2i"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector3i"
	"graphics.gd/variant/Vector4"
	"graphics.gd/variant/Vector4i"
)

var Host API

type API struct {
	/*Version struct {
	Major     func() uint32             `gd:"version_major"`
	Minor     func() uint32             `gd:"version_minor"`
	Patch     func() uint32             `gd:"version_patch"`
	Hex       func() uint32             `gd:"version_hex"`
	Status    func() String             `gd:"version_status"`
	Build     func() String             `gd:"version_build"`
	Hash      func() String             `gd:"version_hash"`
	Timestamp func(CallReturns[uint64]) `gd:"version_timestamp"`
	String    func() String             `gd:"version_string"`
	}*/
	/*Library struct {
		Location func() String `gd:"library_location"`
	}*/
	/*Memory struct {
		Malloc func(size int) Pointer               `gd:"memory_malloc"`
		Sizeof func(name StringName) int            `gd:"memory_sizeof"`
		Resize func(addr Pointer, size int) Pointer `gd:"memory_resize"`
		Clear  func(addr Pointer, size int)         `gd:"memory_clear"`

		Load struct {
			Byte   func(addr Pointer) byte   `gd:"memory_load_byte"`
			Uint16 func(addr Pointer) uint16 `gd:"memory_load_u16"`
			Uint32 func(addr Pointer) uint32 `gd:"memory_load_u32"`
		}
		Edit struct {
			Byte   func(addr Pointer, value byte)   `gd:"memory_edit_byte"`
			Uint16 func(addr Pointer, value uint16) `gd:"memory_edit_u16"`
			Uint32 func(addr Pointer, value uint32) `gd:"memory_edit_u32"`
			Uint64 func(addr Pointer, value uint64) `gd:"memory_edit_u64"`

			Bits128 func(addr Pointer, value [2]uint64) `gd:"memory_edit_128"`
			Bits256 func(addr Pointer, value [4]uint64) `gd:"memory_edit_256"`
			Bits512 func(addr Pointer, value [8]uint64) `gd:"memory_edit_512"`
		}
		Free func(addr Pointer) `gd:"memory_free"`
	}*/
	/*Threads struct {
		Main func() bool `gd:"thread_is_main"`
	}*/
	/*Log struct {
		Error   func(text, code, fn, file string, line int32, notify_editor bool) `gd:"log_error"`
		Warning func(text, code, fn, file string, line int32, notify_editor bool) `gd:"log_warning"`
	}*/
	/*Builtin struct {
		Types struct {
			Name            func(vtype VariantType) String                                                                                                                                  `gd:"variant_type_name"`
			Make            func(vtype VariantType, result CallReturns[Variant], arg_count int, args CallAccepts[Variant], err CallReturns[CallError])                                      `gd:"variant_type_make"`
			Call            func(vtype VariantType, static StringName, result CallReturns[Variant], arg_count int, args CallAccepts[Variant], err CallReturns[CallError])                   `gd:"variant_type_call"`
			Convertable     func(vtype VariantType, to VariantType, strict bool) bool                                                                                                       `gd:"variant_type_convertable"`
			SetupArray      func(array Array, vtype VariantType, class_name StringName, script Variant)                                                                                     `gd:"variant_type_setup_array"`
			SetupDictionary func(dict Dictionary, key_type VariantType, key_class_name StringName, key_script Variant, val_type VariantType, val_class_name StringName, val_script Variant) `gd:"variant_type_setup_dictionary"`
			FetchConstant   func(vtype VariantType, constant StringName, result CallReturns[Variant])                                                                                       `gd:"variant_type_fetch_constant"`
			Constructor     func(vtype VariantType, n int) FunctionID                                                                                                                       `gd:"variant_type_unsafe_constructor"`
			Evaluator       func(op VariantOperator, a, b VariantType) FunctionID                                                                                                           `gd:"variant_type_evaluator"`
			Setter          func(vtype VariantType, field StringName) FunctionID                                                                                                            `gd:"variant_type_setter"`
			Getter          func(vtype VariantType, field StringName) FunctionID                                                                                                            `gd:"variant_type_getter"`
			HasProperty     func(vtype VariantType, field StringName) bool                                                                                                                  `gd:"variant_type_has_property"`

			Method func(vtype VariantType, builtin StringName, hash int64) MethodForBuiltinType `gd:"variant_type_builtin_method"`

			Unsafe struct {
				Call func(value CallMutates[any], fn MethodForBuiltinType, result CallReturns[any], shape Shape, args CallAccepts[any]) `gd:"variant_type_unsafe_call"`
				Make func(constructor FunctionID, result CallReturns[any], shape Shape, args CallAccepts[any])                          `gd:"variant_type_unsafe_make"`
				Free func(vtype VariantType, shape Shape, args CallAccepts[any])                                                        `gd:"variant_type_unsafe_free"`
			}
		}
		Functions struct {
			Name func(utility StringName, hash int64) FunctionID                                  `gd:"builtin_name"`
			Call func(fn FunctionID, result CallReturns[any], shape Shape, args CallAccepts[any]) `gd:"builtin_call"`
		}
	}*/
	/*Variants struct {
		Zero func(result CallReturns[Variant])                                                                                                     `gd:"variant_zero"`
		Copy func(v Variant, result CallReturns[Variant])                                                                                          `gd:"variant_copy"`
		Call func(v Variant, method StringName, result CallReturns[Variant], arg_count int, args CallAccepts[Variant], err CallReturns[CallError]) `gd:"variant_call"`
		Eval func(op VariantOperator, a, b Variant, call CallReturns[Variant]) bool                                                                `gd:"variant_eval"`
		Hash func(v Variant, hash CallReturns[int64])                                                                                              `gd:"variant_hash"`
		Bool func(v Variant) bool                                                                                                                  `gd:"variant_bool"`
		Text func(v Variant) String                                                                                                                `gd:"variant_text"`
		Type func(v Variant) VariantType                                                                                                           `gd:"variant_type"`
		Deep struct {
			Copy func(v Variant, result CallReturns[Variant])                    `gd:"variant_deep_copy"`
			Hash func(v Variant, recursion_count int64, hash CallReturns[int64]) `gd:"variant_deep_hash"`
		}
		Set struct {
			Index func(v Variant, key, val Variant) bool                                 `gd:"variant_set_index"`
			Array func(v Variant, idx int, val Variant, err CallReturns[CallError]) bool `gd:"variant_set_array"`
			Field func(v Variant, field StringName, value Variant) bool                  `gd:"variant_set_field"`
		}
		Get struct {
			Index func(v Variant, key Variant, result CallReturns[Variant]) bool                         `gd:"variant_get_index"`
			Array func(v Variant, idx int, result CallReturns[Variant], err CallReturns[CallError]) bool `gd:"variant_get_array"`
			Field func(v Variant, field StringName, result CallReturns[Variant]) bool                    `gd:"variant_get_field"`
		}
		Has struct {
			Index  func(v Variant, index Variant) bool     `gd:"variant_has_index"`
			Method func(v Variant, method StringName) bool `gd:"variant_has_method"`
		}
		Unsafe struct {
			Call func(fn FunctionID, result CallReturns[any], shape Shape, args CallAccepts[any]) `gd:"variant_unsafe_call"`
			Eval func(fn FunctionID, result CallReturns[any], shape Shape, args CallAccepts[any]) `gd:"variant_unsafe_eval"`
			Free func(v Variant)                                                                  `gd:"variant_unsafe_free"`

			MakeNative func(vtype VariantType, v Variant, shape Shape, result CallReturns[any])                 `gd:"variant_unsafe_make_native"`
			FromNative func(vtype VariantType, result CallReturns[Variant], shape Shape, args CallAccepts[any]) `gd:"variant_unsafe_from_native"`

			Set struct {
				Field func(setter FunctionID, shape Shape, args CallAccepts[any])          `gd:"variant_unsafe_set_field"`
				Array func(vtype VariantType, idx int, shape Shape, args CallAccepts[any]) `gd:"variant_unsafe_set_array"`
				Index func(vtype VariantType, shape Shape, args CallAccepts[any])          `gd:"variant_unsafe_set_index"`
			}
			Get struct {
				Field func(getter FunctionID, result CallReturns[any], shape Shape, args CallAccepts[any])          `gd:"variant_unsafe_get_field"`
				Array func(vtype VariantType, idx int, result CallReturns[any], shape Shape, args CallAccepts[any]) `gd:"variant_unsafe_get_array"`
				Index func(vtype VariantType, result CallReturns[any], shape Shape, args CallAccepts[any])          `gd:"variant_unsafe_get_index"`
			}
			InternalPointer func(vtype VariantType, v Variant) Pointer `gd:"variant_unsafe_internal_pointer"`
		}
	}
	Iterators struct {
		Make func(v Variant, result CallReturns[Iterator], err CallReturns[CallError])               `gd:"iterator_make"`
		Next func(v Variant, iter CallMutates[Iterator], err CallReturns[CallError]) bool            `gd:"iterator_next"`
		Load func(v Variant, iter Iterator, result CallReturns[Variant], err CallReturns[CallError]) `gd:"iterator_load"`
	}
	Strings struct {
		Access func(s String, idx int) rune    `gd:"string_access"`
		Resize func(s String, size int) String `gd:"string_resize"`
		Unsafe func(s String) Pointer          `gd:"string_unsafe"`

		Decode struct {
			Latin1 func(s string) String                     `gd:"string_decode_latin1"`
			UTF8   func(s string) String                     `gd:"string_decode_utf8"`
			UTF16  func(s string, little_endian bool) String `gd:"string_decode_utf16"`
			UTF32  func(s string) String                     `gd:"string_decode_utf32"`
			Wide   func(s string) String                     `gd:"string_decode_wide"`
		}
		Encode struct {
			Latin1 func(s String, buf []byte) int `gd:"string_encode_latin1"`
			UTF8   func(s String, buf []byte) int `gd:"string_encode_utf8"`
			UTF16  func(s String, buf []byte) int `gd:"string_encode_utf16"`
			UTF32  func(s String, buf []byte) int `gd:"string_encode_utf32"`
			Wide   func(s String, buf []byte) int `gd:"string_encode_wide"`
		}
		Append struct {
			String func(s String, other String) String `gd:"string_append"`
			Rune   func(s String, char rune) String    `gd:"string_append_rune"`
		}
		Intern struct {
			Latin1 func(s string) StringName `gd:"string_intern_latin1"`
			UTF8   func(s string) StringName `gd:"string_intern_utf8"`
		}
	}*/
	/*Packed struct {
		Bytes struct {
			Unsafe func(p PackedArray[byte]) Pointer       `gd:"packed_byte_array_unsafe"`
			Access func(p PackedArray[byte], idx int) byte `gd:"packed_byte_array_access"`
		}
		Float32s struct {
			Unsafe func(p PackedArray[float32]) Pointer          `gd:"packed_float32_array_unsafe"`
			Access func(p PackedArray[float32], idx int) float32 `gd:"packed_float32_array_access"`
		}
		Float64s struct {
			Unsafe func(p PackedArray[float64]) Pointer          `gd:"packed_float64_array_unsafe"`
			Access func(p PackedArray[float64], idx int) float64 `gd:"packed_float64_array_access"`
		}
		Int32s struct {
			Unsafe func(p PackedArray[int32]) Pointer        `gd:"packed_int32_array_unsafe"`
			Access func(p PackedArray[int32], idx int) int32 `gd:"packed_int32_array_access"`
		}
		Int64s struct {
			Unsafe func(p PackedArray[int64]) Pointer                             `gd:"packed_int64_array_unsafe"`
			Access func(p PackedArray[int64], idx int, result CallReturns[int64]) `gd:"packed_int64_array_access"`
		}
		Strings struct {
			Unsafe func(p PackedArray[String]) Pointer         `gd:"packed_string_array_unsafe"`
			Access func(p PackedArray[String], idx int) String `gd:"packed_string_array_access"`
		}
		Vector2s struct {
			Unsafe func(p PackedArray[Vector2.XY]) Pointer                                  `gd:"packed_vector2_array_unsafe"`
			Access func(p PackedArray[Vector2.XY], idx int, result CallReturns[Vector2.XY]) `gd:"packed_vector2_array_access"`
		}
		Vector3s struct {
			Unsafe func(p PackedArray[Vector3.XYZ]) Pointer                                   `gd:"packed_vector3_array_unsafe"`
			Access func(p PackedArray[Vector3.XYZ], idx int, result CallReturns[Vector3.XYZ]) `gd:"packed_vector3_array_access"`
		}
		Vector4s struct {
			Unsafe func(p PackedArray[Vector4.XYZW]) Pointer                                    `gd:"packed_vector4_array_unsafe"`
			Access func(p PackedArray[Vector4.XYZW], idx int, result CallReturns[Vector4.XYZW]) `gd:"packed_vector4_array_access"`
		}
		Colors struct {
			Unsafe func(p PackedArray[Color.RGBA]) Pointer                                  `gd:"packed_color_array_unsafe"`
			Access func(p PackedArray[Color.RGBA], idx int, result CallReturns[Color.RGBA]) `gd:"packed_color_array_access"`
		}
	}*/
	/*Array struct {
		Get func(p Array, idx int, result CallReturns[Variant]) `gd:"array_get"`
		Set func(p Array, idx int, value Variant)               `gd:"array_set"`
	}*/
	/*Dictionaries struct {
		Get func(dict Dictionary, index Variant, result CallReturns[Variant]) `gd:"packed_dictionary_access"`
		Set func(dict Dictionary, index, value Variant)                       `gd:"packed_dictionary_modify"`
	}*/
	/*Callables struct {
		Create func(id FunctionID, object ObjectID, result CallReturns[Callable]) `gd:"callable_create"`
		Lookup func(Callable) FunctionID                                          `gd:"callable_lookup"`
	}*/
	/*Objects struct {
		Make func(name StringName) Object                                                                                                               `gd:"object_make"`
		Call func(obj Object, method MethodForClass, result CallReturns[Variant], arg_count int, args CallAccepts[Variant], err CallReturns[CallError]) `gd:"object_call"`
		Name func(obj Object) StringName                                                                                                                `gd:"object_name"`
		Type func(name StringName) ObjectType                                                                                                           `gd:"object_type"`
		Cast func(obj Object, to ObjectType) Object                                                                                                     `gd:"object_cast"`

		ID struct {
			Get           func(obj Object, id CallReturns[ObjectID]) `gd:"object_id"`
			InsideVariant func(v Variant, id CallReturns[ObjectID])  `gd:"object_id_inside_variant"`
		}
		Lookup func(id ObjectID) Object     `gd:"object_lookup"`
		Global func(name StringName) Object `gd:"object_global"`

		Method struct {
			Lookup func(name StringName, method StringName, hash int64) MethodForClass `gd:"object_method_lookup"`
		}
		Unsafe struct {
			Call func(obj Object, fn MethodForClass, result CallReturns[any], shape Shape, args CallAccepts[any]) `gd:"object_unsafe_call"`
			Free func(obj Object)                                                                                 `gd:"object_unsafe_free"`
		}
		Script struct {
			Make func(fn ExtensionInstanceID) ScriptInstance                                                                                          `gd:"object_script_make"`
			Call func(obj Object, name StringName, result CallReturns[Variant], arg_count int, args CallAccepts[Variant], err CallReturns[CallError]) `gd:"object_script_call"`

			Setup func(obj Object, script ScriptInstance)          `gd:"object_script_setup"`
			Fetch func(obj Object, language Object) ScriptInstance `gd:"object_script_fetch"`

			DefinesMethod    func(obj Object, method StringName) bool                         `gd:"object_script_defines_method"`
			AddPropertyState func(fn FunctionID, arg Pointer, name StringName, state Variant) `gd:"object_script_property_state_add"`

			Placeholder struct {
				Create func(language, script, owner Object) ScriptInstance       `gd:"object_script_placeholder_create"`
				Update func(script ScriptInstance, array Array, dict Dictionary) `gd:"object_script_placeholder_update"`
			}
		}
		Extension struct {
			Setup func(obj Object, name StringName, class ExtensionInstanceID) `gd:"object_extension_setup"`
			Fetch func(obj Object) ExtensionInstanceID                         `gd:"object_extension_fetch"`
			Close func(obj Object)                                             `gd:"object_extension_close"`
		}
	}*/
	/*RefCounted struct {
		Get func(ref RefCounted) Object      `gd:"ref_get_object"`
		Set func(ref RefCounted, obj Object) `gd:"ref_set_object"`
	}*/
	/*Editor struct {
		AddDocumentation func(xml string)      `gd:"editor_add_documentation"`
		AddPlugin        func(name StringName) `gd:"editor_add_plugin"`
		EndPlugin        func(name StringName) `gd:"editor_end_plugin"`
	}*/
	/*ClassDB struct {
		Register struct {
			Class            func(class, parent_class StringName, id ExtensionClassID, virtual, abstract, exposed, runtime bool, icon_path String) `gd:"classdb_register"`
			Methods          func(class StringName, methods MethodList)                                                                            `gd:"classdb_register_methods"`
			Constant         func(class, enum, name StringName, value int64, bitfield bool)                                                        `gd:"classdb_register_constant"`
			Property         func(class StringName, property PropertyList, setter, getter StringName)                                              `gd:"classdb_register_property"`
			PropertyIndexed  func(class StringName, property PropertyList, setter, getter StringName, index int)                                   `gd:"classdb_register_property_indexed"`
			PropertyGroup    func(class StringName, group, prefix String)                                                                          `gd:"classdb_register_property_group"`
			PropertySubgroup func(class StringName, subgroup, prefix String)                                                                       `gd:"classdb_register_property_sub_group"`
			Signal           func(class, signal StringName, args PropertyList)                                                                     `gd:"classdb_register_signal"`
			Removal          func(class StringName)                                                                                                `gd:"classdb_register_removal"`
		}
	}*/
}

type InitializationLevel uint32

const (
	InitializationLevelCore    InitializationLevel = 0
	InitializationLevelServers InitializationLevel = 1
	InitializationLevelScene   InitializationLevel = 2
	InitializationLevelEditor  InitializationLevel = 3
)

type String [1]Pointer
type StringName [1]Pointer
type NodePath [1]Pointer

type Array [1]Pointer
type Dictionary [1]Pointer
type Callable [2]uint64
type Signal [2]uint64

type MethodForClass Pointer

type MethodForBuiltinType Pointer

type Object = gdunsafe.Object
type ObjectType = gdunsafe.ObjectType
type ObjectID gdunsafe.ObjectID
type RefCounted Pointer

type Variant = gdunsafe.Variant
type VariantOperator uint32

type Iterator Variant

type ClassLibrary Pointer

type TaskID Pointer
type FunctionID Pointer
type CallableID Pointer

type ExtensionClassID Pointer
type ExtensionInstanceID Pointer
type ExtensionBindingID Pointer

type ScriptInstance Pointer

type CallReturns[T any] unsafe.Pointer
type CallAccepts[T any] unsafe.Pointer
type CallMutates[T any] unsafe.Pointer

type MaybeError struct{}

type PropertyList Pointer
type MethodList Pointer

const (
	Equal VariantOperator = iota
	NotEqual
	Less
	LessEqual
	Greater
	GreaterEqual
	Add
	Subtract
	Multiply
	Divide
	Negate
	Positive
	Module
	Power
	ShiftLeft
	ShiftRight
	BitAnd
	BitOr
	BitXor
	BitNegate
	LogicalAnd
	LogicalOr
	LogicalXor
	LogicalNegate
	In
)

type ArgumentMetadata uint32

const (
	ArgumentMetadataNone ArgumentMetadata = iota
	ArgumentMetadataIntIsInt8
	ArgumentMetadataIntIsInt16
	ArgumentMetadataIntIsInt32
	ArgumentMetadataIntIsInt64
	ArgumentMetadataIntIsUint8
	ArgumentMetadataIntIsUint16
	ArgumentMetadataIntIsUint32
	ArgumentMetadataIntIsUint64
	ArgumentMetadataRealIsFloat32
	ArgumentMetadataRealIsFloat64
	ArgumentMetadataIntIsChar16
	ArgumentMetadataIntIsChar32
)

type MethodFlags uint32

const (
	MethodFlagNormal   MethodFlags = 1
	MethodFlagEditor   MethodFlags = 2
	MethodFlagConst    MethodFlags = 4
	MethodFlagVirtual  MethodFlags = 8
	MethodFlagVararg   MethodFlags = 16
	MethodFlagStatic   MethodFlags = 32
	MethodFlagsDefault MethodFlags = MethodFlagNormal
)

// Shape is used to correctly transfer data for unsafe calls into the engine.
type Shape = gdunsafe.Shape

const (
	ShapeEmpty Shape = iota

	ShapeBytes1
	ShapeBytes2
	ShapeBytes4
	ShapeBytes8
	ShapeBytes4x2
	ShapeBytes4x3
	ShapeBytes8x2
	ShapeBytes4x4
	ShapeBytes8x3
	ShapeBytes4x6
	ShapeBytes4x9
	ShapeBytes4x12
	ShapeBytes4x16
)

func ShapeVariants(count int) Shape {
	if count == 0 {
		return 0
	}
	if count > 16 {
		panic("ShapeVariants: count must be between 0 and 16")
	}
	var shape Shape
	for i := range count {
		shape |= SizeVariant << ((i + 1) * 4)
	}
	return shape
}

type AnyVariant interface {
	Variant | ~byte | ~bool | ~int64 | ~float64 | ~float32 | ~int32 | String | ~Vector2i.XY | Vector2.XY | ~Rect2.PositionSize |
		Rect2i.PositionSize | ~Vector3.XYZ | ~Vector3i.XYZ |
		~Transform2D.OriginXY | ~Vector4.XYZW | ~Vector4i.XYZW | ~Plane.NormalD | ~Quaternion.IJKX | ~AABB.PositionSize |
		~Basis.XYZ | ~Transform3D.BasisOrigin | ~Projection.XYZW | ~Color.RGBA | StringName | NodePath | ~uint64 |
		Object | Callable | Signal | Dictionary | Array | PackedArray[byte] | PackedArray[int32] | PackedArray[int64] |
		PackedArray[float32] | PackedArray[float64] | PackedArray[String] | PackedArray[Vector2.XY] |
		PackedArray[Vector3.XYZ] | PackedArray[Color.RGBA] | PackedArray[Vector4.XYZW] | CallError
}

type AnyPointer interface {
	Variant | String | Callable | Signal | Dictionary | Array | PackedArray[byte] |
		StringName | NodePath | PackedArray[int32] | PackedArray[int64] |
		PackedArray[float32] | PackedArray[float64] | PackedArray[String] | PackedArray[Vector2.XY] |
		PackedArray[Vector3.XYZ] | PackedArray[Color.RGBA] | PackedArray[Vector4.XYZW]
}

const (
	SizeVariant     Shape = ShapeBytes8x3
	SizeBool        Shape = ShapeBytes1
	SizeInt         Shape = ShapeBytes8
	SizeFloat       Shape = ShapeBytes8
	SizeVector2     Shape = ShapeBytes4x2
	SizeVector3     Shape = ShapeBytes4x3
	SizeVector4     Shape = ShapeBytes4x4
	SizeColor       Shape = ShapeBytes4x4
	SizeRect2       Shape = ShapeBytes4x4
	SizeRect2i      Shape = ShapeBytes4x4
	SizeVector2i    Shape = ShapeBytes4x2
	SizeVector3i    Shape = ShapeBytes4x3
	SizeVector4i    Shape = ShapeBytes4x4
	SizeTransform2D Shape = ShapeBytes4x6
	SizeTransform3D Shape = ShapeBytes4x12
	SizePlane       Shape = ShapeBytes4x4
	SizeQuaternion  Shape = ShapeBytes4x4
	SizeAABB        Shape = ShapeBytes4x6
	SizeBasis       Shape = ShapeBytes4x9
	SizeProjection  Shape = ShapeBytes4x16
	SizeRID         Shape = ShapeBytes8
	SizeCallable    Shape = ShapeBytes8x2
	SizeSignal      Shape = ShapeBytes8x2

	SizeCallError Shape = ShapeBytes4x3
)

func init() {
	if SizeVariant.SizeResult() != int(unsafe.Sizeof(Variant{})) || SizeVariant.Alignment() != int(unsafe.Alignof(Variant{})) {
		panic("Size of Variant does not match expected size")
	}
	if SizeBool.SizeResult() != int(unsafe.Sizeof(bool(false))) || SizeBool.Alignment() != int(unsafe.Alignof(bool(false))) {
		panic("Size of Bool does not match expected size")
	}
	if SizeInt.SizeResult() != int(unsafe.Sizeof(int64(0))) || SizeInt.Alignment() != int(unsafe.Alignof(int64(0))) {
		panic("Size of Int does not match expected size")
	}
	if SizeFloat.SizeResult() != int(unsafe.Sizeof(float64(0))) || SizeFloat.Alignment() != int(unsafe.Alignof(float64(0))) {
		panic("Size of Float does not match expected size")
	}
	if SizeString.SizeResult() != int(unsafe.Sizeof(String{})) || SizeString.Alignment() != int(unsafe.Alignof(String{})) {
		panic("Size of String does not match expected size")
	}
	if SizeVector2.SizeResult() != int(unsafe.Sizeof(Vector2.XY{})) || SizeVector2.Alignment() != int(unsafe.Alignof(Vector2.XY{})) {
		panic("Size of Vector2 does not match expected size")
	}
	if SizeVector3.SizeResult() != int(unsafe.Sizeof(Vector3.XYZ{})) || SizeVector3.Alignment() != int(unsafe.Alignof(Vector3.XYZ{})) {
		panic("Size of Vector3 does not match expected size")
	}
	if SizeVector4.SizeResult() != int(unsafe.Sizeof(Vector4.XYZW{})) || SizeVector4.Alignment() != int(unsafe.Alignof(Vector4.XYZW{})) {
		panic("Size of Vector4 does not match expected size")
	}
	if SizeColor.SizeResult() != int(unsafe.Sizeof(Color.RGBA{})) || SizeColor.Alignment() != int(unsafe.Alignof(Color.RGBA{})) {
		panic("Size of Color does not match expected size")
	}
	if SizeRect2.SizeResult() != int(unsafe.Sizeof(Rect2.PositionSize{})) || SizeRect2.Alignment() != int(unsafe.Alignof(Rect2.PositionSize{})) {
		panic("Size of Rect2 does not match expected size")
	}
	if SizeRect2i.SizeResult() != int(unsafe.Sizeof(Rect2i.PositionSize{})) || SizeRect2i.Alignment() != int(unsafe.Alignof(Rect2i.PositionSize{})) {
		panic("Size of Rect2i does not match expected size")
	}
	if SizeVector2i.SizeResult() != int(unsafe.Sizeof(Vector2i.XY{})) || SizeVector2i.Alignment() != int(unsafe.Alignof(Vector2i.XY{})) {
		panic("Size of Vector2i does not match expected size")
	}
	if SizeVector3i.SizeResult() != int(unsafe.Sizeof(Vector3i.XYZ{})) || SizeVector3i.Alignment() != int(unsafe.Alignof(Vector3i.XYZ{})) {
		panic("Size of Vector3i does not match expected size")
	}
	if SizeVector4i.SizeResult() != int(unsafe.Sizeof(Vector4i.XYZW{})) || SizeVector4i.Alignment() != int(unsafe.Alignof(Vector4i.XYZW{})) {
		panic("Size of Vector4i does not match expected size")
	}
	if SizeTransform2D.SizeResult() != int(unsafe.Sizeof(Transform2D.OriginXY{})) || SizeTransform2D.Alignment() != int(unsafe.Alignof(Transform2D.OriginXY{})) {
		panic("Size of Transform2D does not match expected size")
	}
	if SizeTransform3D.SizeResult() != int(unsafe.Sizeof(Transform3D.BasisOrigin{})) || SizeTransform3D.Alignment() != int(unsafe.Alignof(Transform3D.BasisOrigin{})) {
		panic("Size of Transform3D does not match expected size")
	}
	if SizePlane.SizeResult() != int(unsafe.Sizeof(Plane.NormalD{})) || SizePlane.Alignment() != int(unsafe.Alignof(Plane.NormalD{})) {
		panic("Size of Plane does not match expected size")
	}
	if SizeQuaternion.SizeResult() != int(unsafe.Sizeof(Quaternion.IJKX{})) || SizeQuaternion.Alignment() != int(unsafe.Alignof(Quaternion.IJKX{})) {
		panic("Size of Quaternion does not match expected size")
	}
	if SizeAABB.SizeResult() != int(unsafe.Sizeof(AABB.PositionSize{})) || SizeAABB.Alignment() != int(unsafe.Alignof(AABB.PositionSize{})) {
		panic("Size of AABB does not match expected size")
	}
	if SizeBasis.SizeResult() != int(unsafe.Sizeof(Basis.XYZ{})) || SizeBasis.Alignment() != int(unsafe.Alignof(Basis.XYZ{})) {
		panic("Size of Basis does not match expected size")
	}
	if SizeProjection.SizeResult() != int(unsafe.Sizeof(Projection.XYZW{})) || SizeProjection.Alignment() != int(unsafe.Alignof(Projection.XYZW{})) {
		panic("Size of Projection does not match expected size")
	}
	if SizeStringName.SizeResult() != int(unsafe.Sizeof(StringName{})) || SizeStringName.Alignment() != int(unsafe.Alignof(StringName{})) {
		panic("Size of StringName does not match expected size")
	}
	if SizeNodePath.SizeResult() != int(unsafe.Sizeof([1]Pointer{})) || SizeNodePath.Alignment() != int(unsafe.Alignof([1]Pointer{})) {
		panic("Size of NodePath does not match expected size")
	}
	if SizeRID.SizeResult() != int(unsafe.Sizeof([1]uint64{})) || SizeRID.Alignment() != int(unsafe.Alignof([1]uint64{})) {
		panic("Size of RID does not match expected size")
	}
	if SizeObject.SizeResult() != int(unsafe.Sizeof(Object(0))) || SizeObject.Alignment() != int(unsafe.Alignof(Object(0))) {
		panic("Size of Object does not match expected size")
	}
	if SizeCallable.SizeResult() != int(unsafe.Sizeof(Callable{})) || SizeCallable.Alignment() != int(unsafe.Alignof(Callable{})) {
		panic("Size of Callable does not match expected size")
	}
	if SizeSignal.SizeResult() != int(unsafe.Sizeof([2]uint64{})) || SizeSignal.Alignment() != int(unsafe.Alignof([2]uint64{})) {
		panic("Size of Signal does not match expected size")
	}
	if SizeDictionary.SizeResult() != int(unsafe.Sizeof(Dictionary{})) || SizeDictionary.Alignment() != int(unsafe.Alignof(Dictionary{})) {
		panic("Size of Dictionary does not match expected size")
	}
	if SizeArray.SizeResult() != int(unsafe.Sizeof(Array{})) || SizeArray.Alignment() != int(unsafe.Alignof(Array{})) {
		panic("Size of Array does not match expected size")
	}
	if SizePackedArray.SizeResult() != int(unsafe.Sizeof(PackedArray[byte]{})) || SizePackedArray.Alignment() != int(unsafe.Alignof(PackedArray[byte]{})) {
		panic("Size of PackedArray does not match expected size")
	}
}

type CallError = gdunsafe.Error

type Packable interface {
	byte | int32 | int64 | float32 | float64 | Color.RGBA | Vector2.XY | Vector3.XYZ | Vector4.XYZW | String
}
