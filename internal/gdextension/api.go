// Package gdextension is the graphics.gd authorative Go representation of the Godot C GDExtension API.
package gdextension

import (
	"structs"
	"unsafe"

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
	"runtime.link/api"
	"runtime.link/api/stub"
)

var Host API
var On = api.Import[Callbacks](stub.API, "", nil)

type Callbacks struct {
	Init func(level InitializationLevel) `gd:"on_init"`
	Exit func(level InitializationLevel) `gd:"on_exit"`

	Frames struct {
		First func() `gd:"on_first_frame"`
		Every func() `gd:"on_every_frame"`
		Final func() `gd:"on_final_frame"`
	}
	Extension struct {
		Binding struct {
			Created   func(instance ExtensionInstanceID) ExtensionBindingID          `gd:"on_extension_binding_created"`
			Removed   func(instance ExtensionInstanceID, binding ExtensionBindingID) `gd:"on_extension_binding_removed"`
			Reference func(instance ExtensionInstanceID, increment bool) bool        `gd:"on_extension_binding_reference"`
		}
		Instance struct {
			Set                func(instance ExtensionInstanceID, field StringName, value Variant) bool                                                                 `gd:"on_extension_instance_set"`
			Get                func(instance ExtensionInstanceID, field StringName, result Returns[Variant]) bool                                                       `gd:"on_extension_instance_get"`
			PropertyList       func(instance ExtensionInstanceID) PropertyList                                                                                          `gd:"on_extension_instance_property_list"`
			PropertyHasDefault func(instance ExtensionInstanceID, field StringName) bool                                                                                `gd:"on_extension_instance_property_has_default"`
			PropertyGetDefault func(instance ExtensionInstanceID, field StringName, result Returns[Variant]) bool                                                       `gd:"on_extension_instance_property_get_default"`
			PropertyValidation func(instance ExtensionInstanceID, field PropertyList) bool                                                                              `gd:"on_extension_instance_property_validation"`
			Notification       func(instance ExtensionInstanceID, reverse bool)                                                                                         `gd:"on_extension_instance_notification"`
			Stringify          func(instance ExtensionInstanceID) String                                                                                                `gd:"on_extension_instance_stringify"`
			Reference          func(instance ExtensionInstanceID, increment bool) bool                                                                                  `gd:"on_extension_instance_reference"`
			RID                func(instance ExtensionInstanceID) uint64                                                                                                `gd:"on_extension_instance_rid"`
			Call               func(instance ExtensionInstanceID, fn FunctionID, result Returns[Variant], arg_count int, args Accepts[Variant], err Returns[CallError]) `gd:"on_extension_instance_call"`
			CallChecked        func(instance ExtensionInstanceID, fn FunctionID, result Returns[Variant], args Accepts[Variant])                                        `gd:"on_extension_instance_call_checked"`
			Unsafe             struct {
				Call func(instance ExtensionInstanceID, fn FunctionID, result Returns[any], args Accepts[any]) `gd:"on_extension_instance_unsafe_call"`
			}
			Free func(instance ExtensionInstanceID) `gd:"on_extension_instance_free"`
		}
		Class struct {
			Create func(class ExtensionClassID, notify_postinitialize bool) Object         `gd:"on_extension_class_create"`
			Method func(class ExtensionClassID, method StringName, hash uint32) FunctionID `gd:"on_extension_class_method"`
		}
		Script struct {
			Categorization      func(instance ExtensionInstanceID, into PropertyList) bool      `gd:"on_extension_script_categorization"`
			PropertyType        func(field StringName, err Returns[CallError]) VariantType      `gd:"on_extension_script_get_property_type"`
			Owner               func(instance ExtensionInstanceID) Object                       `gd:"on_extension_script_get_owner"`
			PropertyState       func(instance ExtensionInstanceID, add FunctionID, arg Pointer) `gd:"on_extension_script_get_property_state"`
			Methods             func(instance ExtensionInstanceID) MethodList                   `gd:"on_extension_script_get_methods"`
			HasMethod           func(instance ExtensionInstanceID, method StringName) bool      `gd:"on_extension_script_has_method"`
			MethodArgumentCount func(instance ExtensionInstanceID, method StringName) int       `gd:"on_extension_script_get_method_argument_count"`
			Get                 func(instance ExtensionInstanceID) Object                       `gd:"on_extension_script_get"`
			IsPlaceholder       func(instance ExtensionInstanceID) bool                         `gd:"on_extension_script_is_placeholder"`
			Language            func(instance ExtensionInstanceID) Object                       `gd:"on_extension_script_get_language"`
		}
	}
	Callables struct {
		Call          func(fn FunctionID, result Returns[Variant], arg_count int, args Accepts[Variant], err Returns[CallError]) `gd:"on_callable_call"`
		Validation    func(fn FunctionID) bool                                                                                   `gd:"on_callable_validation"`
		Free          func(fn FunctionID)                                                                                        `gd:"on_callable_free"`
		Hash          func(fn FunctionID) int64                                                                                  `gd:"on_callable_hash"`
		Compare       func(fn FunctionID, other FunctionID) bool                                                                 `gd:"on_callable_compare"`
		LessThan      func(fn FunctionID, other FunctionID) bool                                                                 `gd:"on_callable_less_than"`
		Stringify     func(fn FunctionID, err Returns[CallError]) String                                                         `gd:"on_callable_stringify"`
		ArgumentCount func(fn FunctionID, err Returns[CallError]) int                                                            `gd:"on_callable_get_argument_count"`
	}
	Editor struct {
		ClassInUseDetection func(classes PackedArray, result Returns[PackedArray]) `gd:"on_editor_class_in_use_detection"`
	}
	Tasks struct {
		Run        func(task TaskID)           `gd:"on_worker_thread_pool_task"`
		RunInGroup func(task TaskID, n uint32) `gd:"on_worker_thread_pool_group_task"`
	}
}

type API struct {
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
		Malloc func(size int) Pointer               `gd:"memory_malloc"`
		Sizeof func(name StringName) int            `gd:"memory_sizeof"`
		Resize func(addr Pointer, size int) Pointer `gd:"memory_resize"`

		Load struct {
			Byte   func(addr Pointer) byte   `gd:"memory_load_byte"`
			Uint16 func(addr Pointer) uint16 `gd:"memory_load_u16"`
			Uint32 func(addr Pointer) uint32 `gd:"memory_load_u32"`
			Uint64 func(addr Pointer) uint64 `gd:"memory_load_u64"`
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
	}
	Log struct {
		Error   func(text, code, fn, file string, line int32, notify_editor bool) `gd:"log_error"`
		Warning func(text, code, fn, file string, line int32, notify_editor bool) `gd:"log_warning"`
	}
	Builtins struct {
		Name func(utility StringName, hash int64) FunctionID                                  `gd:"builtin_name"`
		Call func(fn FunctionID, result CallReturns[any], shape Shape, args CallAccepts[any]) `gd:"builtin_call"`
	}
	Variants struct {
		Zero func(result CallReturns[Variant])                                                                                                     `gd:"variant_zero"`
		Copy func(v Variant, result CallReturns[Variant])                                                                                          `gd:"variant_copy"`
		Call func(v Variant, method StringName, result CallReturns[Variant], arg_count int, args CallAccepts[Variant], err CallReturns[CallError]) `gd:"variant_call"`
		Eval func(op VariantOperator, a, b Variant, call CallReturns[Variant]) bool                                                                `gd:"variant_eval"`
		Hash func(v Variant) int64                                                                                                                 `gd:"variant_hash"`
		Bool func(v Variant) bool                                                                                                                  `gd:"variant_bool"`
		Text func(v Variant) String                                                                                                                `gd:"variant_text"`
		Type func(v Variant) VariantType                                                                                                           `gd:"variant_type"`
		Deep struct {
			Copy func(v Variant, result CallReturns[Variant]) `gd:"variant_deep_copy"`
			Hash func(v Variant, recursion_count int64) int64 `gd:"variant_deep_hash"`
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
	VariantTypes struct {
		Name            func(vtype VariantType) String                                                                                                                                  `gd:"variant_type_name"`
		Make            func(vtype VariantType, result CallReturns[Variant], arg_count int, args CallAccepts[Variant], err CallReturns[CallError])                                      `gd:"variant_type_make"`
		Call            func(vtype VariantType, static StringName, result CallReturns[Variant], arg_count int, args CallAccepts[Variant], err CallReturns[CallError])                   `gd:"variant_type_call"`
		Convertable     func(vtype VariantType, to VariantType, strict bool) bool                                                                                                       `gd:"variant_type_convertable"`
		SetupArray      func(array Array, vtype VariantType, class_name StringName, script Variant)                                                                                     `gd:"variant_type_setup_array"`
		SetupDictionary func(dict Dictionary, key_type VariantType, key_class_name StringName, key_script Variant, val_type VariantType, val_class_name StringName, val_script Variant) `gd:"variant_type_setup_dictionary"`
		FetchConstant   func(vtype VariantType, constant StringName, result CallReturns[Variant])                                                                                       `gd:"variant_type_fetch_constant"`
		BuiltinMethod   func(vtype VariantType, builtin StringName, hash int64) FunctionID                                                                                              `gd:"variant_type_builtin_method"`
		Constructor     func(vtype VariantType, n int) FunctionID                                                                                                                       `gd:"variant_type_unsafe_constructor"`
		Evaluator       func(op VariantOperator, a, b VariantType) FunctionID                                                                                                           `gd:"variant_type_evaluator"`
		Setter          func(vtype VariantType, field StringName) FunctionID                                                                                                            `gd:"variant_type_setter"`
		Getter          func(vtype VariantType, field StringName) FunctionID                                                                                                            `gd:"variant_type_getter"`
		HasProperty     func(vtype VariantType, field StringName) bool                                                                                                                  `gd:"variant_type_has_property"`

		Unsafe struct {
			Call func(fn FunctionID, result CallReturns[any], shape Shape, args CallAccepts[any])          `gd:"variant_type_unsafe_call"`
			Make func(constructor FunctionID, result CallReturns[any], shape Shape, args CallAccepts[any]) `gd:"variant_type_unsafe_make"`
			Free func(vtype VariantType, shape Shape, args CallAccepts[any])                               `gd:"variant_type_unsafe_free"`
		}
	}
	Iterators struct {
		Make func(v Variant, result CallReturns[Iterator], err CallReturns[CallError])                     `gd:"iterator_make"`
		Next func(v Variant, result CallReturns[Iterator], iter Iterator, err CallReturns[CallError]) bool `gd:"iterator_next"`
		Load func(v Variant, result CallReturns[Variant], iter Iterator, err CallReturns[CallError])       `gd:"iterator_load"`
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
	}
	Packed struct {
		Bytes struct {
			Unsafe func(p PackedArray) Pointer       `gd:"packed_byte_array_unsafe"`
			Access func(p PackedArray, idx int) byte `gd:"packed_byte_array_access"`
		}
		Float32s struct {
			Unsafe func(p PackedArray) Pointer          `gd:"packed_float32_array_unsafe"`
			Access func(p PackedArray, idx int) float32 `gd:"packed_float32_array_access"`
		}
		Float64s struct {
			Unsafe func(p PackedArray) Pointer          `gd:"packed_float64_array_unsafe"`
			Access func(p PackedArray, idx int) float64 `gd:"packed_float64_array_access"`
		}
		Int32s struct {
			Unsafe func(p PackedArray) Pointer        `gd:"packed_int32_array_unsafe"`
			Access func(p PackedArray, idx int) int32 `gd:"packed_int32_array_access"`
		}
		Int64s struct {
			Unsafe func(p PackedArray) Pointer        `gd:"packed_int64_array_unsafe"`
			Access func(p PackedArray, idx int) int64 `gd:"packed_int64_array_access"`
		}
		Strings struct {
			Unsafe func(p PackedArray) Pointer         `gd:"packed_string_array_unsafe"`
			Access func(p PackedArray, idx int) String `gd:"packed_string_array_access"`
		}
		Vector2s struct {
			Unsafe func(p PackedArray) Pointer                                  `gd:"packed_vector2_array_unsafe"`
			Access func(p PackedArray, idx int, result CallReturns[Vector2.XY]) `gd:"packed_vector2_array_access"`
		}
		Vector3s struct {
			Unsafe func(p PackedArray) Pointer                                   `gd:"packed_vector3_array_unsafe"`
			Access func(p PackedArray, idx int, result CallReturns[Vector3.XYZ]) `gd:"packed_vector3_array_access"`
		}
		Vector4s struct {
			Unsafe func(p PackedArray) Pointer                                    `gd:"packed_vector4_array_unsafe"`
			Access func(p PackedArray, idx int, result CallReturns[Vector4.XYZW]) `gd:"packed_vector4_array_access"`
		}
		Colors struct {
			Unsafe func(p PackedArray) Pointer                                  `gd:"packed_color_array_unsafe"`
			Access func(p PackedArray, idx int, result CallReturns[Color.RGBA]) `gd:"packed_color_array_access"`
		}
	}
	Array struct {
		Get func(p Array, idx int, result CallReturns[Variant]) `gd:"array_get"`
		Set func(p Array, idx int, value Variant)               `gd:"array_set"`
	}
	Dictionaries struct {
		Access func(dict Dictionary, index Variant, result CallReturns[Variant]) `gd:"packed_dictionary_access"`
		Modify func(dict Dictionary, index, value Variant)                       `gd:"packed_dictionary_modify"`
	}
	Callables struct {
		Create func(id CallableID, object ObjectID, result CallReturns[Callable]) `gd:"callable_create"`
		Lookup func(Callable) CallableID                                          `gd:"callable_lookup"`
	}
	Objects struct {
		Make func(name StringName) Object                                                                                                           `gd:"object_make"`
		Call func(obj Object, method FunctionID, result CallReturns[Variant], arg_count int, args CallAccepts[Variant], err CallReturns[CallError]) `gd:"object_call"`
		Name func(obj Object) StringName                                                                                                            `gd:"object_name"`
		Type func(name StringName) ObjectType                                                                                                       `gd:"object_type"`
		Cast func(obj Object, to ObjectType) Object                                                                                                 `gd:"object_cast"`

		ID struct {
			Get           func(obj Object) ObjectID `gd:"object_id"`
			InsideVariant func(v Variant) ObjectID  `gd:"object_id_inside_variant"`
		}
		Lookup func(id ObjectID) Object     `gd:"object_lookup"`
		Global func(name StringName) Object `gd:"object_global"`

		Method struct {
			Lookup func(name StringName, method StringName, hash int64) FunctionID `gd:"object_method_lookup"`
		}
		Unsafe struct {
			Call func(obj Object, fn FunctionID, result CallReturns[any], shape Shape, args CallAccepts[any]) `gd:"object_unsafe_call"`
			Free func(obj Object)                                                                             `gd:"object_unsafe_free"`
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
			Fetch func(obj Object) ExtensionClassID                            `gd:"object_extension_fetch"`
			Close func(obj Object)                                             `gd:"object_extension_close"`
		}
	}
	RefCounted struct {
		Get func(ref RefCounted) Object      `gd:"ref_get_object"`
		Set func(ref RefCounted, obj Object) `gd:"ref_set_object"`
	}
	Editor struct {
		AddDocumentation func(xml string)      `gd:"editor_add_documentation"`
		AddPlugin        func(name StringName) `gd:"editor_add_plugin"`
		EndPlugin        func(name StringName) `gd:"editor_end_plugin"`
	}
	ClassDB struct {
		PropertyList struct {
			Make func(length int) PropertyList `gd:"property_list_make"`
			Push func(info PropertyList,
				vtype VariantType,
				name StringName,
				class_name StringName,
				hint uint32,
				hint_string String,
				usage uint32,
				meta ArgumentMetadata,
			) `gd:"property_list_push"`
			Info struct {
				Type      func(info PropertyList) VariantType `gd:"property_info_type"`
				Name      func(info PropertyList) StringName  `gd:"property_info_name"`
				ClassName func(info PropertyList) StringName  `gd:"property_info_class_name"`
				Hint      func(info PropertyList) uint32      `gd:"property_info_hint"`
				HinString func(info PropertyList) String      `gd:"property_info_hint_string"`
				Usage     func(info PropertyList) uint32      `gd:"property_info_usage"`
			}
			Free func(info PropertyList) `gd:"property_list_free"`
		}
		MethodList struct {
			Make func(length int) MethodList `gd:"method_list_make"`
			Push func(info MethodList,
				name StringName,
				call FunctionID,
				method_flags MethodFlags,
				has_return_value bool,
				return_value_info PropertyList,
				argument_count uint32,
				arguments_info PropertyList,
				default_argument_count int,
				default_arguments CallAccepts[Variant],
			) `gd:"method_list_push"`
			Free func(info MethodList) `gd:"method_list_free"`
		}
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
		Image struct {
			Unsafe func(img Object) Pointer       `gd:"classdb_Image_unsafe"`
			Access func(img Object, idx int) byte `gd:"classdb_Image_access"`
		}
		XMLParser struct {
			Load func(parser Object, buf []byte) int `gd:"classdb_XMLParser_load"`
		}
		FileAccess struct {
			Write func(file Object, buf []byte)     `gd:"classdb_FileAccess_write"`
			Read  func(file Object, buf []byte) int `gd:"classdb_FileAccess_read"`
		}
		WorkerThreadPool struct {
			AddTask      func(pool Object, fn TaskID, priority bool, description String)                             `gd:"classdb_WorkerThreadPool_add_task"`
			AddGroupTask func(pool Object, fn TaskID, elements int32, task int32, priority bool, description String) `gd:"classdb_WorkerThreadPool_add_group_task"`
		}
	}
}

type InitializationLevel uint32

const (
	InitCore    InitializationLevel = 0
	InitServers InitializationLevel = 1
	InitScene   InitializationLevel = 2
	InitEditor  InitializationLevel = 3
)

type String Pointer
type StringName Pointer
type NodePath Pointer

type Array Pointer
type Dictionary Pointer
type Callable [2]uint64

type Object Pointer
type ObjectType Pointer
type ObjectID uint64
type RefCounted Pointer

type Variant [3]uint64
type VariantOperator uint32

type Iterator [3]uint64

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

type CallErrorType uint32

const (
	CallOK               CallErrorType = iota
	CallInvalidMethod    CallErrorType = 1
	CallInvalidArguments CallErrorType = 2
	CallTooManyArguments CallErrorType = 3
	CallTooFewArguments  CallErrorType = 4
	CallInstanceIsNull   CallErrorType = 5
	CallMethodNotConst   CallErrorType = 6
)

// Shape is used to correctly transfer data for unsafe calls into the engine.
type Shape uint64

func (shape Shape) SizeResult() (size int) {
	switch shape & 0xF {
	case SizeBytes0:
		return 0
	case SizeBytes1:
		return 1
	case SizeBytes2:
		return 2
	case SizeBytes4:
		return 4
	case SizeBytes8:
		return 8
	case SizeBytes12:
		return 12
	case SizeBytes16:
		return 16
	case SizeBytes24:
		return 24
	case SizeBytes32:
		return 32
	case SizeBytes36:
		return 36
	case SizeBytes40:
		return 40
	case SizeBytes48:
		return 48
	case SizeBytes64:
		return 64
	case SizeBytes72:
		return 72
	case SizeBytes96:
		return 96
	default:
		return 128
	}
}

func (shape Shape) SizeArguments() (size int) {
	for i := 1; i < 16; i++ {
		switch (shape >> (i * 4)) & 0xF {
		case SizeBytes0:
			continue
		case SizeBytes1:
			size += 1
		case SizeBytes2:
			size += 2
		case SizeBytes4:
			size += 4
		case SizeBytes8:
			size += 8
		case SizeBytes12:
			size += 12
		case SizeBytes16:
			size += 16
		case SizeBytes24:
			size += 24
		case SizeBytes32:
			size += 32
		case SizeBytes36:
			size += 36
		case SizeBytes40:
			size += 40
		case SizeBytes48:
			size += 48
		case SizeBytes64:
			size += 64
		case SizeBytes72:
			size += 72
		case SizeBytes96:
			size += 96
		case SizeBytes128:
			size += 128
		}
	}
	return
}

const (
	SizeBytes0 Shape = iota
	SizeBytes1
	SizeBytes2
	SizeBytes4
	SizeBytes8
	SizeBytes12
	SizeBytes16
	SizeBytes24
	SizeBytes32
	SizeBytes36
	SizeBytes40
	SizeBytes48
	SizeBytes64
	SizeBytes72
	SizeBytes96
	SizeBytes128
)

type VariantType uint32

func (vtype VariantType) String() string {
	name := Host.VariantTypes.Name(vtype)
	var buf = make([]byte, 32)
	n := Host.Strings.Encode.UTF8(name, buf)
	return unsafe.String(&buf[0], min(n, len(buf)))
}

const (
	TypeNil                VariantType = 0
	TypeBool               VariantType = 1
	TypeInt                VariantType = 2
	TypeFloat              VariantType = 3
	TypeString             VariantType = 4
	TypeVector2            VariantType = 5
	TypeVector2i           VariantType = 6
	TypeRect2              VariantType = 7
	TypeRect2i             VariantType = 8
	TypeVector3            VariantType = 9
	TypeVector3i           VariantType = 10
	TypeTransform2D        VariantType = 11
	TypeVector4            VariantType = 12
	TypeVector4i           VariantType = 13
	TypePlane              VariantType = 14
	TypeQuaternion         VariantType = 15
	TypeAABB               VariantType = 16
	TypeBasis              VariantType = 17
	TypeTransform3D        VariantType = 18
	TypeProjection         VariantType = 19
	TypeColor              VariantType = 20
	TypeStringName         VariantType = 21
	TypeNodePath           VariantType = 22
	TypeRID                VariantType = 23
	TypeObject             VariantType = 24
	TypeCallable           VariantType = 25
	TypeSignal             VariantType = 26
	TypeDictionary         VariantType = 27
	TypeArray              VariantType = 28
	TypePackedByteArray    VariantType = 29
	TypePackedInt32Array   VariantType = 30
	TypePackedInt64Array   VariantType = 31
	TypePackedFloat32Array VariantType = 32
	TypePackedFloat64Array VariantType = 33
	TypePackedStringArray  VariantType = 34
	TypePackedVector2Array VariantType = 35
	TypePackedVector3Array VariantType = 36
	TypePackedColorArray   VariantType = 37
	TypePackedVector4Array VariantType = 38
	TypeMax                VariantType = 39
)

const (
	SizeVariant     Shape = SizeBytes24
	SizeBool        Shape = SizeBytes1
	SizeInt         Shape = SizeBytes8
	SizeFloat       Shape = SizeBytes8
	SizeVector2     Shape = SizeBytes8
	SizeVector3     Shape = SizeBytes12
	SizeVector4     Shape = SizeBytes16
	SizeColor       Shape = SizeBytes16
	SizeRect2       Shape = SizeBytes16
	SizeRect2i      Shape = SizeBytes16
	SizeVector2i    Shape = SizeBytes8
	SizeVector3i    Shape = SizeBytes12
	SizeVector4i    Shape = SizeBytes16
	SizeTransform2D Shape = SizeBytes24
	SizeTransform3D Shape = SizeBytes48
	SizePlane       Shape = SizeBytes16
	SizeQuaternion  Shape = SizeBytes16
	SizeAABB        Shape = SizeBytes24
	SizeBasis       Shape = SizeBytes36
	SizeProjection  Shape = SizeBytes64
	SizeRID         Shape = SizeBytes8
	SizeCallable    Shape = SizeBytes16
	SizeSignal      Shape = SizeBytes16
)

func init() {
	if SizeVariant.SizeResult() != int(unsafe.Sizeof(Variant{})) {
		panic("Size of Variant does not match expected size")
	}
	if SizeBool.SizeResult() != int(unsafe.Sizeof(bool(false))) {
		panic("Size of Bool does not match expected size")
	}
	if SizeInt.SizeResult() != int(unsafe.Sizeof(int64(0))) {
		panic("Size of Int does not match expected size")
	}
	if SizeFloat.SizeResult() != int(unsafe.Sizeof(float64(0))) {
		panic("Size of Float does not match expected size")
	}
	if SizeString.SizeResult() != int(unsafe.Sizeof(String(0))) {
		panic("Size of String does not match expected size")
	}
	if SizeVector2.SizeResult() != int(unsafe.Sizeof(Vector2.XY{})) {
		panic("Size of Vector2 does not match expected size")
	}
	if SizeVector3.SizeResult() != int(unsafe.Sizeof(Vector3.XYZ{})) {
		panic("Size of Vector3 does not match expected size")
	}
	if SizeVector4.SizeResult() != int(unsafe.Sizeof(Vector4.XYZW{})) {
		panic("Size of Vector4 does not match expected size")
	}
	if SizeColor.SizeResult() != int(unsafe.Sizeof(Color.RGBA{})) {
		panic("Size of Color does not match expected size")
	}
	if SizeRect2.SizeResult() != int(unsafe.Sizeof(Rect2.PositionSize{})) {
		panic("Size of Rect2 does not match expected size")
	}
	if SizeRect2i.SizeResult() != int(unsafe.Sizeof(Rect2i.PositionSize{})) {
		panic("Size of Rect2i does not match expected size")
	}
	if SizeVector2i.SizeResult() != int(unsafe.Sizeof(Vector2i.XY{})) {
		panic("Size of Vector2i does not match expected size")
	}
	if SizeVector3i.SizeResult() != int(unsafe.Sizeof(Vector3i.XYZ{})) {
		panic("Size of Vector3i does not match expected size")
	}
	if SizeVector4i.SizeResult() != int(unsafe.Sizeof(Vector4i.XYZW{})) {
		panic("Size of Vector4i does not match expected size")
	}
	if SizeTransform2D.SizeResult() != int(unsafe.Sizeof(Transform2D.OriginXY{})) {
		panic("Size of Transform2D does not match expected size")
	}
	if SizeTransform3D.SizeResult() != int(unsafe.Sizeof(Transform3D.BasisOrigin{})) {
		panic("Size of Transform3D does not match expected size")
	}
	if SizePlane.SizeResult() != int(unsafe.Sizeof(Plane.NormalD{})) {
		panic("Size of Plane does not match expected size")
	}
	if SizeQuaternion.SizeResult() != int(unsafe.Sizeof(Quaternion.IJKX{})) {
		panic("Size of Quaternion does not match expected size")
	}
	if SizeAABB.SizeResult() != int(unsafe.Sizeof(AABB.PositionSize{})) {
		panic("Size of AABB does not match expected size")
	}
	if SizeBasis.SizeResult() != int(unsafe.Sizeof(Basis.XYZ{})) {
		panic("Size of Basis does not match expected size")
	}
	if SizeProjection.SizeResult() != int(unsafe.Sizeof(Projection.XYZW{})) {
		panic("Size of Projection does not match expected size")
	}
	if SizeStringName.SizeResult() != int(unsafe.Sizeof(StringName(0))) {
		panic("Size of StringName does not match expected size")
	}
	if SizeNodePath.SizeResult() != int(unsafe.Sizeof([1]Pointer{})) {
		panic("Size of NodePath does not match expected size")
	}
	if SizeRID.SizeResult() != int(unsafe.Sizeof([1]uint64{})) {
		panic("Size of RID does not match expected size")
	}
	if SizeObject.SizeResult() != int(unsafe.Sizeof(Object(0))) {
		panic("Size of Object does not match expected size")
	}
	if SizeCallable.SizeResult() != int(unsafe.Sizeof(Callable{})) {
		panic("Size of Callable does not match expected size")
	}
	if SizeSignal.SizeResult() != int(unsafe.Sizeof([2]uint64{})) {
		panic("Size of Signal does not match expected size")
	}
	if SizeDictionary.SizeResult() != int(unsafe.Sizeof(Dictionary(0))) {
		panic("Size of Dictionary does not match expected size")
	}
	if SizeArray.SizeResult() != int(unsafe.Sizeof(Array(0))) {
		panic("Size of Array does not match expected size")
	}
	if SizePackedArray.SizeResult() != int(unsafe.Sizeof(PackedArray{})) {
		panic("Size of PackedArray does not match expected size")
	}
}

type CallError struct {
	_ structs.HostLayout

	Type     CallErrorType
	Argument int32
	Expected int32
}

func (err CallError) Err() error {
	if err.Type == CallOK {
		return nil
	}
	return err
}

func (err CallError) Error() string {
	switch err.Type {
	case CallInvalidMethod:
		return "Call Invalid Method"
	case CallInvalidArguments:
		return "Call Invalid Arguments"
	case CallTooManyArguments:
		return "Call Too Many Arguments"
	case CallTooFewArguments:
		return "Call Too Few Arguments"
	case CallInstanceIsNull:
		return "Call Instance Is Null"
	case CallMethodNotConst:
		return "Call Method Not Const"
	default:
		return "Unknown Call Error"
	}
}
