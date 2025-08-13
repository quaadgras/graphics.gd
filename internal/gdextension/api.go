// Package gdextension is the graphics.gd authorative Go representation of the Godot C GDExtension API.
package gdextension

import (
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector4"
)

type String uintptr
type Object uintptr
type StringName uintptr
type Pointer uintptr
type Function uintptr
type Variant [3]uint64
type VariantType uint32
type VariantOperator uint32
type Iterator [3]uint64
type CallbackID uintptr
type PackedArray uintptr
type Array uintptr
type Dictionary uintptr
type ClassLibrary uintptr
type ExtensionID uintptr
type ObjectType uintptr
type ObjectID uint64
type RefCounted uintptr
type ScriptInstance uintptr
type Callable uintptr
type PropertyList uintptr

type CallReturns[T any] struct{}
type CallAccepts[T any] Call
type MaybeError struct{}

type Call uintptr

type API struct {
	Call struct {
		New func() Call        `gd:"call_new"`
		Add func(Call) Pointer `gd:"call_add"`
		Get func(Call) Pointer `gd:"call_get"`
		Err func(Call) int64   `gd:"call_err"`
		End func(Call) bool    `gd:"call_end"`
	}
	Version struct {
		Major     func() uint32 `gd:"version_major"`
		Minor     func() uint32 `gd:"version_minor"`
		Patch     func() uint32 `gd:"version_patch"`
		Hex       func() uint32 `gd:"version_hex"`
		Status    func() String `gd:"version_status"`
		Build     func() String `gd:"version_build"`
		Hash      func() String `gd:"version_hash"`
		Timestamp func() uint64 `gd:"version_timestamp"`
		String    func() String `gd:"version_string"`
	}
	Memory struct {
		Allocate func(uintptr) Pointer          `gd:"malloc"`
		Size     func(StringName) uint64        `gd:"sizeof"`
		Resize   func(Pointer, uintptr) Pointer `gd:"realloc"`
		Load     struct {
			Byte   func(Pointer, uintptr) byte   `gd:"load_byte"`
			Uint16 func(Pointer, uintptr) uint16 `gd:"load_u16"`
			Uint32 func(Pointer, uintptr) uint32 `gd:"load_u32"`
			Uint64 func(Pointer, uintptr) uint64 `gd:"load_u64"`
		}
		Send struct {
			Byte   func(Pointer, uintptr, byte)   `gd:"send_byte"`
			Uint16 func(Pointer, uintptr, uint16) `gd:"send_u16"`
			Uint32 func(Pointer, uintptr, uint32) `gd:"send_u32"`
			Uint64 func(Pointer, uintptr, uint64) `gd:"send_u64"`

			Bits256 func(Pointer, uintptr, uint64, uint64, uint64, uint64)                                 `gd:"send_256"`
			Bits512 func(Pointer, uintptr, uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64) `gd:"send_512"`
		}
		Free func(Pointer) `gd:"free"`
	}
	Log struct {
		Error   func(string, string, string, string, int32, bool) `gd:"log_error"`
		Warning func(string, string, string, string, int32, bool) `gd:"log_warning"`
	}
	Builtins struct {
		Find func(StringName, int64) Function                                           `gd:"builtin_lookup"`
		Call func(Function, int32, CallAccepts[any]) (CallReturns[Variant], MaybeError) `gd:"builtin_call"`
	}
	Variants struct {
		Zero func(Call) CallReturns[Variant]                                                           `gd:"variant_zero"`
		Copy func(Call, Variant) CallReturns[Variant]                                                  `gd:"variant_copy"`
		Call func(Variant, StringName, CallAccepts[Variant], int64) (CallReturns[Variant], MaybeError) `gd:"variant_call"`
		Eval func(Call, VariantOperator, Variant, Variant) (CallReturns[Variant], bool)                `gd:"variant_eval"`
		Hash func(Variant) int64                                                                       `gd:"variant_hash"`
		Bool func(Variant) bool                                                                        `gd:"variant_bool"`
		Text func(Variant) String                                                                      `gd:"variant_text"`
		Type func(Variant) VariantType                                                                 `gd:"variant_type"`
		Safe struct {
			Object func(Variant) ObjectID `gd:"variant_safe_object"`
		}
		Deep struct {
			Copy func(Call, Variant) CallReturns[Variant] `gd:"variant_deep_copy"`
			Hash func(Variant, int64) int64               `gd:"variant_deep_hash"`
		}
		Set struct {
			Index func(Call, Variant, Variant, Variant) (CallReturns[Variant], bool, MaybeError) `gd:"variant_set_index"`
			Array func(Call, Variant, int64, Variant) (CallReturns[Variant], bool)               `gd:"variant_set_array"`
			Field func(Call, Variant, StringName, Variant) (CallReturns[Variant], bool)          `gd:"variant_set_field"`
		}
		Get struct {
			Index func(Call, Variant, Variant) (CallReturns[Variant], bool)           `gd:"variant_get_index"`
			Array func(Call, Variant, int64) (CallReturns[Variant], bool, MaybeError) `gd:"variant_get_array"`
			Field func(Call, Variant, StringName) (CallReturns[Variant], bool)        `gd:"variant_get_field"`
		}
		Has struct {
			Field  func(Call, Variant, StringName) bool            `gd:"variant_has_field"`
			Index  func(Call, Variant, Variant) (bool, MaybeError) `gd:"variant_has_index"`
			Method func(Call, Variant, StringName) bool            `gd:"variant_has_method"`
		}
		Unsafe struct {
			From func(VariantType, CallAccepts[any]) CallReturns[Variant] `gd:"variant_unsafe_from"`
			Eval func(VariantOperator, CallAccepts[any]) CallReturns[any] `gd:"variant_unsafe_eval"`
			Call func(Function, int32, CallAccepts[any]) CallReturns[any] `gd:"variant_unsafe_call"`
			Free func(Variant)                                            `gd:"variant_unsafe_free"`

			Set struct {
				Field func(CallAccepts[any], VariantType)        `gd:"variant_unsafe_set_field"`
				Array func(CallAccepts[any], VariantType, int64) `gd:"variant_unsafe_set_array"`
				Index func(CallAccepts[any], VariantType)        `gd:"variant_unsafe_set_index"`
			}
			Get struct {
				Field func(CallAccepts[any], VariantType) CallReturns[any]        `gd:"variant_unsafe_get_field"`
				Array func(CallAccepts[any], VariantType, int64) CallReturns[any] `gd:"variant_unsafe_get_array"`
				Index func(CallAccepts[any], VariantType) CallReturns[any]        `gd:"variant_unsafe_get_index"`
			}
			Pointer func(Variant) Pointer `gd:"variant_unsafe_pointer"`
		}
	}
	VariantTypes struct {
		Name           func(VariantType) String                                                                      `gd:"variant_type_name"`
		From           func(Call, VariantType, Variant) CallReturns[any]                                             `gd:"variant_type_from"`
		Make           func(VariantType, int32, CallAccepts[Variant]) (CallReturns[Variant], MaybeError)             `gd:"variant_type_make"`
		Call           func(VariantType, StringName, int64, CallAccepts[Variant]) (CallReturns[Variant], MaybeError) `gd:"variant_type_call"`
		Convertable    func(VariantType, VariantType, bool) bool                                                     `gd:"variant_type_convertable"`
		LockArray      func(Array, VariantType, StringName, Variant) bool                                            `gd:"variant_type_lock_array"`
		LockDictionary func(Dictionary, VariantType, StringName, Variant, VariantType, StringName, Variant) bool     `gd:"variant_type_lock_dictionary"`
		Constant       func(Call, VariantType, StringName) CallReturns[Variant]                                      `gd:"variant_type_constant"`
		BuiltinMethod  func(VariantType, StringName, int64) Function                                                 `gd:"variant_type_builtin_method"`

		Unsafe struct {
			Constructor func(CallAccepts[any], VariantType) CallReturns[any] `gd:"variant_type_unsafe_constructor"`
			Free        func(CallAccepts[any], VariantType) CallReturns[any] `gd:"variant_type_unsafe_free"`
		}
	}
	Iterators struct {
		Make func(Call, Variant) (CallReturns[Iterator], bool, MaybeError)           `gd:"iterator_make"`
		Next func(Call, Variant, Iterator) (CallReturns[Iterator], bool, MaybeError) `gd:"iterator_next"`
		Load func(Call, Variant, Iterator) (CallReturns[Variant], bool, MaybeError)  `gd:"iterator_load"`
	}
	Strings struct {
		Access func(String, int64) rune   `gd:"string_access"`
		Resize func(String, int64) String `gd:"string_resize"`
		Unsafe func(String) Pointer       `gd:"string_unsafe"`

		Decode struct {
			Latin1 func(string) String       `gd:"string_decode_latin1"`
			UTF8   func(string) String       `gd:"string_decode_utf8"`
			UTF16  func(string, bool) String `gd:"string_decode_utf16"`
			UTF32  func(string) String       `gd:"string_decode_utf32"`
			Wide   func(string) String       `gd:"string_decode_wide"`
		}
		Encode struct {
			Latin1 func(String, Pointer, int64) int64 `gd:"string_encode_latin1"`
			UTF8   func(String, Pointer, int64) int64 `gd:"string_encode_utf8"`
			UTF16  func(String, Pointer, int64) int64 `gd:"string_encode_utf16"`
			UTF32  func(String, Pointer, int64) int64 `gd:"string_encode_utf32"`
			Wide   func(String, Pointer, int64) int64 `gd:"string_encode_wide"`
		}
		Append struct {
			String func(String, String) String `gd:"string_append"`
			Rune   func(String, rune) String   `gd:"string_append_rune"`
		}
		Intern struct {
			Latin1 func(string, bool) StringName `gd:"string_intern_latin1"`
			UTF8   func(string) StringName       `gd:"string_intern_utf8"`
		}
	}
	Packed struct {
		Bytes struct {
			Unsafe func(PackedArray) Pointer     `gd:"packed_byte_array_unsafe"`
			Access func(PackedArray, int64) byte `gd:"packed_byte_array_access"`
		}
		Float32s struct {
			Unsafe func(PackedArray) Pointer        `gd:"packed_float32_array_unsafe"`
			Access func(PackedArray, int64) float32 `gd:"packed_float32_array_access"`
		}
		Float64s struct {
			Unsafe func(PackedArray) Pointer        `gd:"packed_float64_array_unsafe"`
			Access func(PackedArray, int64) float64 `gd:"packed_float64_array_access"`
		}
		Int32s struct {
			Unsafe func(PackedArray) Pointer      `gd:"packed_int32_array_unsafe"`
			Access func(PackedArray, int64) int32 `gd:"packed_int32_array_access"`
		}
		Int64s struct {
			Unsafe func(PackedArray) Pointer      `gd:"packed_int64_array_unsafe"`
			Access func(PackedArray, int64) int64 `gd:"packed_int64_array_access"`
		}
		Strings struct {
			Unsafe func(PackedArray) Pointer       `gd:"packed_string_array_unsafe"`
			Access func(PackedArray, int64) String `gd:"packed_string_array_access"`
		}
		Vector2s struct {
			Unsafe func(PackedArray) Pointer                              `gd:"packed_vector2_array_unsafe"`
			Access func(Call, PackedArray, int64) CallReturns[Vector2.XY] `gd:"packed_vector2_array_access"`
		}
		Vector3s struct {
			Unsafe func(PackedArray) Pointer                               `gd:"packed_vector3_array_unsafe"`
			Access func(Call, PackedArray, int64) CallReturns[Vector3.XYZ] `gd:"packed_vector3_array_access"`
		}
		Vector4s struct {
			Unsafe func(PackedArray) Pointer                                `gd:"packed_vector4_array_unsafe"`
			Access func(Call, PackedArray, int64) CallReturns[Vector4.XYZW] `gd:"packed_vector4_array_access"`
		}
		Colors struct {
			Unsafe func(PackedArray) Pointer                              `gd:"packed_color_array_unsafe"`
			Access func(Call, PackedArray, int64) CallReturns[Color.RGBA] `gd:"packed_color_array_access"`
		}
		Variants struct {
			Unsafe func(PackedArray) Pointer                           `gd:"packed_variant_array_unsafe"`
			Access func(Call, PackedArray, int64) CallReturns[Variant] `gd:"packed_variant_array_access"`
		}
	}
	Dictionaries struct {
		Access func(Call, Dictionary, Variant) CallReturns[Variant] `gd:"packed_dictionary_access"`
		Modify func(PackedArray, Variant, Variant)                  `gd:"packed_dictionary_modify"`
	}
	Callables struct {
		Create func(Call, CallbackID, ClassLibrary) CallReturns[Callable] `gd:"callable_create"`
		Lookup func(Callable, ClassLibrary) CallbackID                    `gd:"callable_lookup"`
	}

	Objects struct {
		Make func(StringName) Object                                                                        `gd:"object_make"`
		Call func(Call, Object, StringName, int64, CallAccepts[Variant]) (CallReturns[Variant], MaybeError) `gd:"object_call"`
		Free func(Object)                                                                                   `gd:"object_free"`
		Name func(Object) StringName                                                                        `gd:"object_name"`
		Type func(Object) ObjectType                                                                        `gd:"object_type"`
		Cast func(Object, ObjectType) Object                                                                `gd:"object_cast"`

		ID     func(Object) ObjectID   `gd:"object_id"`
		Lookup func(ObjectID) Object   `gd:"object_lookup"`
		Global func(StringName) Object `gd:"object_global"`

		Method struct {
			Lookup func(Object, StringName, int64) Function                        `gd:"object_method_lookup"`
			Unsafe func(Call, Function, Object, CallAccepts[any]) CallReturns[any] `gd:"object_method_unsafe"`
		}
		Script struct {
			Make func(CallbackID) ScriptInstance                                                                `gd:"object_script_make"`
			Call func(Call, Object, StringName, int64, CallAccepts[Variant]) (CallReturns[Variant], MaybeError) `gd:"object_script_call"`
			Find func(Object, Object) ScriptInstance                                                            `gd:"object_script_find"`

			DefinesMethod func(Object, StringName) bool `gd:"object_script_defines_method"`

			Placeholder struct {
				Create func(Object, Object, Object) ScriptInstance `gd:"object_script_placeholder_create"`
				Update func(ScriptInstance, Array, Dictionary)     `gd:"object_script_placeholder_update"`
			}
		}
		Extension struct {
			Setup func(Object, ClassLibrary, ExtensionID) `gd:"object_extension_setup"`
			Fetch func(Object, ClassLibrary) ExtensionID  `gd:"object_extension_fetch"`
			Close func(Object, ClassLibrary)              `gd:"object_extension_close"`
		}
	}
	RefCounted struct {
		Get func(RefCounted) Object  `gd:"ref_get_object"`
		Set func(RefCounted, Object) `gd:"ref_set_object"`
	}
	Editor struct {
		AddDocumentation func(string)     `gd:"editor_add_documentation"`
		AddPlugin        func(StringName) `gd:"editor_add_plugin"`
		EndPlugin        func(StringName) `gd:"editor_end_plugin"`
	}
	ClassDB struct {
		PropertyList struct {
			New func() PropertyList
			Add func(info PropertyList,
				vtype VariantType,
				name StringName,
				class_name StringName,
				hint uint32,
				hint_string String,
				usage uint32,
			)
			End func(info PropertyList)
		}
		Register struct {
			Class  func(ClassLibrary, StringName, StringName, bool, bool, bool, bool, String) `gd:"classdb_register"`
			Method func(lib ClassLibrary,
				name StringName,
				call CallbackID,
				method_flags uint32,
				has_return_value bool,
				return_value_info PropertyList,
				return_value_metadata uint32,
				argument_count uint32,
				arguments_info PropertyList,
				default_argument_count uint32,
				default_arguments CallAccepts[Variant],
			) `gd:"classdb_register_method"`
			Constant         func(ClassLibrary, StringName, StringName, StringName, int64, bool)         `gd:"classdb_register_constant"`
			Property         func(ClassLibrary, StringName, PropertyList, StringName, StringName)        `gd:"classdb_register_property"`
			PropertyIndexed  func(ClassLibrary, StringName, PropertyList, StringName, StringName, int64) `gd:"classdb_register_property_indexed"`
			PropertyGroup    func(ClassLibrary, StringName, String, String)                              `gd:"classdb_register_property_group"`
			PropertySubgroup func(ClassLibrary, StringName, String, String)                              `gd:"classdb_register_property_sub_group"`
			Signal           func(ClassLibrary, StringName, StringName, PropertyList, int64)             `gd:"classdb_register_signal"`
			Removal          func(ClassLibrary, StringName)                                              `gd:"classdb_register_removal"`
		}
		Image struct {
			Unsafe func(Object) Pointer     `gd:"classdb_Image_unsafe"`
			Access func(Object, int64) byte `gd:"classdb_Image_access"`
		}
		XMLParser struct {
			Load func(Object, []byte) int64 `gd:"classdb_XMLParser_load"`
		}
		FileAccess struct {
			Write func(Object, []byte)     `gd:"classdb_FileAccess_write"`
			Read  func(Object, []byte) int `gd:"classdb_FileAccess_read"`
		}
		WorkerThreadPool struct {
			AddTask      func(Object, CallbackID, bool, String)               `gd:"classdb_WorkerThreadPool_add_task"`
			AddGroupTask func(Object, CallbackID, int32, int32, bool, String) `gd:"classdb_WorkerThreadPool_add_group_task"`
		}
	}
}
