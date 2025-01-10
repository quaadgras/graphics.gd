// Package gdextension is the graphics.gd authorative Go representation of the Godot C GDExtension API.
package gdextension

import (
	"runtime/cgo"
	"structs"
	"unsafe"
)

type API struct {
	get_godot_version                                  func(Uninitialized[*GodotVersion])
	mem_alloc                                          func(uintptr) unsafe.Pointer
	mem_realloc                                        func(unsafe.Pointer, uintptr) unsafe.Pointer
	mem_free                                           func(unsafe.Pointer)
	print_error                                        func(string, string, string, int32, Bool)
	print_error_with_message                           func(string, string, string, string, int32, Bool)
	print_warning                                      func(string, string, string, int32, Bool)
	print_warning_with_message                         func(string, string, string, string, int32, Bool)
	print_script_error                                 func(string, string, string, int32, Bool)
	print_script_error_with_message                    func(string, string, string, string, int32, Bool)
	get_native_struct_size                             func(Const[StringNamePtr]) uint64
	variant_new_copy                                   func(Uninitialized[VariantPtr], Const[VariantPtr])
	variant_new_nil                                    func(Uninitialized[VariantPtr])
	variant_destroy                                    func(VariantPtr)
	variant_call                                       func(VariantPtr, Const[StringNamePtr], UnsafeArray[Const[VariantPtr]], Int, Uninitialized[VariantPtr], *CallError)
	variant_call_static                                func(VariantType, Const[StringNamePtr], UnsafeArray[Const[VariantPtr]], Int, Uninitialized[VariantPtr], *CallError)
	variant_evaluate                                   func(VariantOperator, Const[VariantPtr], Const[VariantPtr], Uninitialized[VariantPtr], *Bool)
	variant_set                                        func(VariantPtr, Const[VariantPtr], Const[VariantPtr], *Bool)
	variant_set_named                                  func(VariantPtr, Const[StringNamePtr], Const[VariantPtr], *Bool)
	variant_set_keyed                                  func(VariantPtr, Const[VariantPtr], Const[VariantPtr], *Bool)
	variant_set_indexed                                func(VariantPtr, Int, Const[VariantPtr], *Bool, *Bool)
	variant_get                                        func(VariantPtr, VariantPtr, Uninitialized[VariantPtr], *Bool)
	variant_get_named                                  func(Const[VariantPtr], Const[StringNamePtr], Uninitialized[VariantPtr], *Bool)
	variant_get_keyed                                  func(Const[VariantPtr], Const[VariantPtr], Uninitialized[VariantPtr], *Bool)
	variant_get_indexed                                func(Const[VariantPtr], Int, VariantPtr, *Bool)
	variant_iter_init                                  func(Const[VariantPtr], Uninitialized[VariantPtr], *Bool) Bool
	variant_iter_next                                  func(Const[VariantPtr], VariantPtr, *Bool) Bool
	variant_iter_get                                   func(Const[VariantPtr], VariantPtr, Uninitialized[VariantPtr], *Bool)
	variant_hash                                       func(Const[VariantPtr]) Int
	variant_recursive_hash                             func(Const[VariantPtr], Int) Int
	variant_hash_compare                               func(Const[VariantPtr], Const[VariantPtr]) Bool
	variant_booleanize                                 func(Const[VariantPtr]) Bool
	variant_duplicate                                  func(Const[VariantPtr], VariantPtr, Bool)
	variant_stringify                                  func(Const[VariantPtr], StringPtr)
	variant_get_type                                   func(Const[VariantPtr]) VariantType
	variant_has_method                                 func(Const[VariantPtr], Const[StringNamePtr]) Bool
	variant_has_member                                 func(VariantType, Const[StringNamePtr]) Bool
	variant_has_key                                    func(Const[VariantPtr], Const[VariantPtr], *Bool) Bool
	variant_get_object_instance_id                     func(Const[VariantPtr]) ObjectInstanceID
	variant_get_type_name                              func(VariantType, Uninitialized[StringPtr])
	variant_can_convert                                func(VariantType, VariantType) Bool
	variant_can_convert_strict                         func(VariantType, VariantType) Bool
	get_variant_from_type_constructor                  func(VariantType) VariantFromTypeConstructorFunc
	get_variant_to_type_constructor                    func(VariantType) TypeFromVariantConstructorFunc
	variant_get_ptr_internal_getter                    func(VariantType) VariantGetInternalPtrFunc
	variant_get_ptr_operator_evaluator                 func(VariantOperator, VariantType, VariantType) PtrOperatorEvaluator
	variant_get_ptr_builtin_method                     func(VariantType, Const[StringNamePtr], Int) PtrBuiltInMethod
	variant_get_ptr_constructor                        func(VariantType, int32) PtrConstructor
	variant_get_ptr_destructor                         func(VariantType) PtrDestructor
	variant_construct                                  func(VariantType, Uninitialized[VariantPtr], UnsafeArray[Const[VariantPtr]], int32, *CallError)
	variant_get_ptr_setter                             func(VariantType) PtrSetter
	variant_get_ptr_getter                             func(VariantType, Const[StringNamePtr]) PtrGetter
	variant_get_ptr_indexed_setter                     func(VariantType) PtrIndexedSetter
	variant_get_ptr_indexed_getter                     func(VariantType) PtrIndexedGetter
	variant_get_ptr_keyed_setter                       func(VariantType) PtrKeyedSetter
	variant_get_ptr_keyed_getter                       func(VariantType) PtrKeyedGetter
	variant_get_ptr_keyed_checker                      func(VariantType) PtrKeyedChecker
	variant_get_constant_value                         func(VariantType, Const[StringNamePtr], Uninitialized[VariantPtr])
	variant_get_ptr_utility_function                   func(StringNamePtr, Int) PtrUtilityFunction
	string_new_with_latin1_chars                       func(Uninitialized[StringPtr], string)
	string_new_with_utf8_chars                         func(Uninitialized[StringPtr], string)
	string_new_with_utf16_chars                        func(Uninitialized[StringPtr], string)
	string_new_with_utf32_chars                        func(Uninitialized[StringPtr], string)
	string_new_with_wide_chars                         func(Uninitialized[StringPtr], string)
	string_new_with_latin1_chars_and_len               func(Uninitialized[StringPtr], string, Int)
	string_new_with_utf8_chars_and_len2                func(Uninitialized[StringPtr], string, Int) Int
	string_new_with_utf16_chars_and_len2               func(Uninitialized[StringPtr], string, Int, Bool) Int
	string_new_with_utf32_chars_and_len                func(Uninitialized[StringPtr], string, Int)
	string_new_with_wide_chars_and_len                 func(Uninitialized[StringPtr], string, Int)
	string_to_latin1_chars                             func(Const[StringPtr], UnsafeArray[*byte], Int) Int
	string_to_utf8_chars                               func(Const[StringPtr], UnsafeArray[*byte], Int) Int
	string_to_utf16_chars                              func(Const[StringPtr], UnsafeArray[*byte], Int) Int
	string_to_utf32_chars                              func(Const[StringPtr], UnsafeArray[*byte], Int) Int
	string_to_wide_chars                               func(Const[StringPtr], UnsafeArray[*byte], Int) Int
	string_operator_index_const                        func(StringPtr, Int) rune
	string_operator_plus_eq_string                     func(StringPtr, Const[StringPtr])
	string_operator_plus_eq_char                       func(StringPtr, rune)
	string_operator_plus_eq_cstr                       func(StringPtr, string)
	string_operator_plus_eq_wcstr                      func(StringPtr, string)
	string_operator_plus_eq_c32str                     func(StringPtr, string)
	string_resize                                      func(StringPtr, Int) Int
	string_name_new_with_latin1_chars                  func(Uninitialized[StringNamePtr], string, Bool)
	string_name_new_with_utf8_chars                    func(Uninitialized[StringNamePtr], string)
	string_name_new_with_utf8_chars_and_len            func(Uninitialized[StringNamePtr], string, Int)
	xml_parser_open_buffer                             func(Const[ArrayPtr], UnsafeArray[*byte], uintptr) Int
	file_access_store_buffer                           func(Const[ObjectPtr], UnsafeArray[*byte], uint64)
	file_access_get_buffer                             func(Const[ObjectPtr], UnsafeArray[*byte], uint64) uint64
	image_ptrw                                         func(ObjectPtr) *byte
	image_ptr                                          func(ObjectPtr) *byte
	worker_thread_pool_add_native_group_task           func(ObjectPtr, Func[func(cgo.Handle, uint32)], cgo.Handle, int, int, Bool, Const[StringPtr])
	worker_thread_pool_add_native_task                 func(ObjectPtr, Func[func(cgo.Handle)], cgo.Handle, Bool, Const[StringPtr])
	packed_byte_array_operator_index                   func(PackedArrayPtr, Int) *byte
	packed_byte_array_operator_index_const             func(PackedArrayPtr, Int) *byte
	packed_float32_array_operator_index                func(PackedArrayPtr, Int) *float32
	packed_float32_array_operator_index_const          func(PackedArrayPtr, Int) *float32
	packed_float64_array_operator_index                func(PackedArrayPtr, Int) *float64
	packed_float64_array_operator_index_const          func(PackedArrayPtr, Int) *float64
	packed_int32_array_operator_index                  func(PackedArrayPtr, Int) *int32
	packed_int32_array_operator_index_const            func(PackedArrayPtr, Int) *int32
	packed_int64_array_operator_index                  func(PackedArrayPtr, Int) *int64
	packed_int64_array_operator_index_const            func(PackedArrayPtr, Int) *int64
	packed_string_array_operator_index                 func(PackedArrayPtr, Int) StringPtr
	packed_string_array_operator_index_const           func(PackedArrayPtr, Int) StringPtr
	packed_vector2_array_operator_index                func(PackedArrayPtr, Int) TypePtr
	packed_vector2_array_operator_index_const          func(PackedArrayPtr, Int) TypePtr
	packed_vector3_array_operator_index                func(PackedArrayPtr, Int) TypePtr
	packed_vector3_array_operator_index_const          func(PackedArrayPtr, Int) TypePtr
	packed_vector4_array_operator_index                func(PackedArrayPtr, Int) TypePtr
	packed_vector4_array_operator_index_const          func(PackedArrayPtr, Int) TypePtr
	packed_color_array_operator_index                  func(PackedArrayPtr, Int) TypePtr
	packed_color_array_operator_index_const            func(PackedArrayPtr, Int) TypePtr
	array_operator_index                               func(ArrayPtr, Int) VariantPtr
	array_operator_index_const                         func(ArrayPtr, Int) VariantPtr
	array_ref                                          func(ArrayPtr, Const[ArrayPtr]) ArrayPtr
	array_set_typed                                    func(ArrayPtr, VariantType, Const[StringNamePtr], Const[VariantPtr])
	dictionary_operator_index                          func(DictionaryPtr, Const[VariantPtr]) VariantPtr
	dictionary_operator_index_const                    func(DictionaryPtr, Const[VariantPtr]) VariantPtr
	dictionary_set_typed                               func(DictionaryPtr, VariantType, Const[StringNamePtr], Const[VariantPtr], VariantType, Const[StringNamePtr], Const[VariantPtr])
	object_method_bind_call                            func(MethodBindPtr, ObjectPtr, UnsafeArray[Const[VariantPtr]], Int, Uninitialized[VariantPtr], *CallError)
	object_method_bind_ptrcall                         func(MethodBindPtr, ClassInstancePtr, UnsafeArray[Const[TypePtr]], TypePtr)
	object_destroy                                     func(ObjectPtr)
	global_get_singleton                               func(Const[StringNamePtr]) ObjectPtr
	object_get_instance_binding                        func(ObjectPtr, ClassLibraryPtr, *InstanceBindingCallbacks) cgo.Handle
	object_set_instance_binding                        func(ObjectPtr, ClassInstancePtr, cgo.Handle, *InstanceBindingCallbacks)
	object_free_instance_binding                       func(ObjectPtr, ClassLibraryPtr)
	object_set_instance                                func(ObjectPtr, Const[StringNamePtr], ClassInstancePtr)
	object_get_class_name                              func(Const[ObjectPtr], ClassLibraryPtr, Uninitialized[StringNamePtr])
	object_cast_to                                     func(Const[ObjectPtr], ClassTag) ObjectPtr
	object_get_instance_from_id                        func(ObjectInstanceID) ObjectPtr
	object_get_instance_id                             func(Const[ObjectPtr]) ObjectInstanceID
	object_has_script_method                           func(Const[ObjectPtr], Const[StringNamePtr]) Bool
	object_call_script_method                          func(Const[ObjectPtr], Const[StringNamePtr], UnsafeArray[Const[VariantPtr]], Int, Uninitialized[VariantPtr], *CallError)
	ref_get_object                                     func(Const[RefPtr]) ObjectPtr
	ref_set_object                                     func(RefPtr, ObjectPtr)
	script_instance_create3                            func(ScriptInstanceInfo3, ScriptInstanceDataPtr) ScriptInstancePtr
	placeholder_script_instance_create                 func(ObjectPtr, ObjectPtr, ObjectPtr) ScriptInstancePtr
	placeholder_script_instance_update                 func(ScriptInstancePtr, Const[ArrayPtr], Const[DictionaryPtr])
	object_get_script_instance                         func(Const[ObjectPtr], ObjectPtr) ScriptInstanceDataPtr
	callable_custom_create2                            func(Uninitialized[CallablePtr], CallableCustomInfo2)
	callable_custom_get_userdata                       func(CallablePtr, ClassLibraryPtr) cgo.Handle
	classdb_construct_object2                          func(StringNamePtr) ObjectPtr
	classdb_get_method_bind                            func(StringNamePtr, StringNamePtr, Int) MethodBindPtr
	classdb_get_class_tag                              func(StringNamePtr) ClassTag
	classdb_register_extension_class4                  func(ClassLibraryPtr, Const[StringNamePtr], Const[StringNamePtr], *ClassCreationInfo4)
	classdb_register_extension_class_method            func(ClassLibraryPtr, StringNamePtr, *ClassMethodInfo)
	classdb_register_extension_class_virtual_method    func(ClassLibraryPtr, StringNamePtr, *ClassVirtualMethodInfo)
	classdb_register_extension_class_integer_constant  func(ClassLibraryPtr, StringNamePtr, StringNamePtr, StringNamePtr, Int, Bool)
	classdb_register_extension_class_property          func(ClassLibraryPtr, StringNamePtr, *PropertyInfo, StringNamePtr, StringNamePtr)
	classdb_register_extension_class_property_indexed  func(ClassLibraryPtr, StringNamePtr, *PropertyInfo, StringNamePtr, StringNamePtr, Int)
	classdb_register_extension_class_property_group    func(ClassLibraryPtr, StringNamePtr, StringPtr, StringPtr)
	classdb_register_extension_class_property_subgroup func(ClassLibraryPtr, StringNamePtr, StringPtr, StringPtr)
	classdb_register_extension_class_signal            func(ClassLibraryPtr, StringNamePtr, StringNamePtr, []PropertyInfo, Int)
	classdb_unregister_extension_class                 func(ClassLibraryPtr, StringNamePtr)
	get_library_path                                   func(ClassLibraryPtr, Uninitialized[StringPtr])
	editor_add_plugin                                  func(StringNamePtr)
	editor_remove_plugin                               func(StringNamePtr)
	editor_help_load_xml_from_utf8_chars               func(string)
	editor_help_load_xml_from_utf8_chars_and_len       func(string, Int)
}

type (
	Bool                  uint8
	Int                   int64
	String                string   // in-memory representation of Godot's String, solo pointer-sized.
	Object                uintptr  // in-memory representation of Godot's Object, solo pointer-sized.
	Array                 uintptr  // in-memory representation of Godot's Array, solo pointer-sized.
	Dictionary            uintptr  // in-memory representation of Godot's Dictionary, solo pointer-sized.
	Variant               [24]byte // in-memory representation of Godot's Variant, 24 bytes.
	StringName            uintptr  // in-memory representation of Godot's StringName, solo pointer-sized.
	ClassLibraryPtr       uintptr  // provided by the entry point.
	VariantType           uint32
	MethodBindPtr         uintptr
	CallablePtr           [2]uintptr
	ScriptInstanceDataPtr uintptr
	ClassTag              uintptr
	PackedArray           [2]uintptr // in-memory representation of Godot's PackedArray, two pointer-sized.

	VariantOperator uint32

	ObjectInstanceID uint64
)

type (
	StringNamePtr     *StringName
	StringPtr         *String
	PackedArrayPtr    *PackedArray
	ClassInstancePtr  cgo.Handle
	VariantPtr        *Variant
	TypePtr           unsafe.Pointer
	ObjectPtr         *Object
	ScriptInstancePtr *Object
	ScriptLanguagePtr *Object
	RefPtr            *Object
	DictionaryPtr     *Dictionary
	ArrayPtr          *Array
)

type PropertyInfo struct {
	_ structs.HostLayout

	_type       VariantType
	name        StringNamePtr
	class_name  StringNamePtr
	hint        uint32
	hint_string StringPtr
	usage       uint32
}

type ClassVirtualMethodInfo struct {
	_ structs.HostLayout

	name                  StringNamePtr
	method_flags          uint32
	return_value          PropertyInfo
	return_value_metadata MethodArgumentMetadata
	argument_count        uint32
	arguments             UnsafeArray[PropertyInfo]
	argument_metadata     UnsafeArray[MethodArgumentMetadata]
}

type CallError struct {
	_ structs.HostLayout

	error    CallErrorType
	argument int32
	expected int32
}

type CallErrorType uint32

type ClassMethodInfo struct {
	_ structs.HostLayout

	name                   StringNamePtr
	method_userdata        cgo.Handle
	call_func              Func[func(cgo.Handle, ClassInstancePtr, Const[VariantPtr], Int, VariantPtr, *CallError)]
	ptrcall_func           Func[func(cgo.Handle, ClassInstancePtr, Const[TypePtr], TypePtr)]
	method_flags           uint32
	has_return_value       Bool
	return_value_info      *PropertyInfo
	return_value_metadata  MethodArgumentMetadata
	argument_count         uint32
	arguments_info         UnsafeArray[PropertyInfo]
	arguments_metadata     UnsafeArray[MethodArgumentMetadata]
	default_argument_count uint32
	default_arguments      UnsafeArray[Variant]
}

type MethodArgumentMetadata uint32

type UnsafeArray[T any] *T
type Func[T any] unsafe.Pointer

type ClassCreationInfo4 struct {
	_ structs.HostLayout

	is_virtual                  Bool
	is_abstract                 Bool
	is_exposed                  Bool
	is_runtime                  Bool
	icon_path                   Const[StringNamePtr]
	set_func                    ClassSet
	get_func                    ClassGet
	get_property_list_func      ClassGetPropertyList
	free_property_list_func     ClassFreePropertyList2
	property_can_revert_func    ClassPropertyCanRevert
	property_get_revert_func    ClassPropertyGetRevert
	validate_property_func      ClassValidateProperty
	notification_func           ClassNotification2
	to_string_func              ClassToString
	reference                   ClassReference
	unreference                 ClassUnreference
	create_instance_func        ClassCreateInstance2
	free_instance               ClassFreeInstance
	recreate_instance_func      ClassRecreateInstance
	get_virtual_func            ClassGetVirtual
	get_virtual_call_data_func  ClassGetVirtualCallData
	call_virtual_with_data_func ClassCallVirtualWithData
}

type (
	ClassSet                 Func[func(ClassInstancePtr, StringNamePtr, Const[VariantPtr])]
	ClassGet                 Func[func(ClassInstancePtr, StringNamePtr, VariantPtr)]
	ClassGetPropertyList     Func[func(ClassInstancePtr, *uint32) UnsafeArray[PropertyInfo]]
	ClassFreePropertyList2   Func[func(ClassInstancePtr, UnsafeArray[PropertyInfo], uint32)]
	ClassPropertyCanRevert   Func[func(ClassInstancePtr, Const[StringNamePtr]) Bool]
	ClassPropertyGetRevert   Func[func(ClassInstancePtr, Const[StringNamePtr], VariantPtr) Bool]
	ClassValidateProperty    Func[func(ClassInstancePtr, Const[StringNamePtr], *PropertyInfo) Bool]
	ClassNotification2       Func[func(ClassInstancePtr, int32, Bool)]
	ClassToString            Func[func(ClassInstancePtr, *Bool, StringPtr)]
	ClassReference           Func[func(ClassInstancePtr)]
	ClassUnreference         Func[func(ClassInstancePtr)]
	ClassCreateInstance2     Func[func(cgo.Handle, Bool) ObjectPtr]
	ClassFreeInstance        Func[func(cgo.Handle, ClassInstancePtr)]
	ClassRecreateInstance    Func[func(cgo.Handle, ObjectPtr) ClassInstancePtr]
	ClassGetVirtual          Func[func(cgo.Handle, Const[StringNamePtr]) Func[func(ClassInstancePtr, UnsafeArray[Const[TypePtr]], TypePtr)]]
	ClassGetVirtualCallData  Func[func(cgo.Handle, Const[StringNamePtr]) cgo.Handle]
	ClassCallVirtualWithData Func[func(ClassInstancePtr, Const[StringNamePtr], cgo.Handle, UnsafeArray[Const[TypePtr]], TypePtr)]
)

type CallableCustomInfo2 struct {
	_ structs.HostLayout

	callable_userdata       cgo.Handle
	token                   ClassLibraryPtr
	object_id               ObjectInstanceID
	call_func               CallableCustomCall
	is_valid_func           CallableCustomIsValid
	free_func               CallableCustomFree
	hash_func               CallableCustomHash
	equal_func              CallableCustomEqual
	less_than_func          CallableCustomLessThan
	to_string_func          CallableCustomToString
	get_argument_count_func CallableCustomGetArgumentCount
}

type (
	CallableCustomCall             Func[func(cgo.Handle, UnsafeArray[Const[VariantPtr]], Int, VariantPtr, *CallError)]
	CallableCustomIsValid          Func[func(cgo.Handle) Bool]
	CallableCustomFree             Func[func(cgo.Handle)]
	CallableCustomHash             Func[func(cgo.Handle) uint32]
	CallableCustomEqual            Func[func(cgo.Handle, cgo.Handle) Bool]
	CallableCustomLessThan         Func[func(cgo.Handle, cgo.Handle) Bool]
	CallableCustomToString         Func[func(cgo.Handle, *Bool, StringPtr)]
	CallableCustomGetArgumentCount Func[func(cgo.Handle, *Bool) Int]
)

type ScriptInstanceInfo3 struct {
	_ structs.HostLayout

	set_func                       ScriptInstanceSet
	get_func                       ScriptInstanceGet
	get_property_list_func         ScriptInstanceGetPropertyList
	free_property_list_func        ScriptInstanceFreePropertyList2
	get_class_category_func        ScriptInstanceGetClassCategory
	property_can_revert_func       ScriptInstancePropertyCanRevert
	property_get_revert_func       ScriptInstancePropertyGetRevert
	get_owner_func                 ScriptInstanceGetOwner
	get_property_state_func        ScriptInstanceGetPropertyState
	get_method_list_func           ScriptInstanceGetMethodList
	free_method_list_func          ScriptInstanceFreeMethodList2
	get_property_type_func         ScriptInstanceGetPropertyType
	validate_property_func         ScriptInstanceValidateProperty
	has_method_func                ScriptInstanceHasMethod
	get_method_argument_count_func ScriptInstanceGetMethodArgumentCount
	call_func                      ScriptInstanceCall
	notification_func              ScriptInstanceNotification2
	to_string_func                 ScriptInstanceToString
	refcount_incremented_func      ScriptInstanceRefCountIncremented
	refcount_decremented_func      ScriptInstanceRefCountDecremented
	get_script_func                ScriptInstanceGetScript
	is_placeholder_func            ScriptInstanceIsPlaceholder
	set_fallback_func              ScriptInstanceSet
	get_fallback_func              ScriptInstanceGet
	get_language_func              ScriptInstanceGetLanguage
	free_func                      ScriptInstanceFree
}

type (
	ScriptInstanceSet                    Func[func(ScriptInstanceDataPtr, Const[StringNamePtr], Const[VariantPtr]) Bool]
	ScriptInstanceGet                    Func[func(ScriptInstanceDataPtr, Const[StringNamePtr], VariantPtr) Bool]
	ScriptInstanceGetPropertyList        Func[func(ScriptInstanceDataPtr, *uint32) UnsafeArray[PropertyInfo]]
	ScriptInstanceFreePropertyList2      Func[func(ScriptInstanceDataPtr, UnsafeArray[PropertyInfo], uint32)]
	ScriptInstanceGetClassCategory       Func[func(ScriptInstanceDataPtr, *PropertyInfo) Bool]
	ScriptInstancePropertyCanRevert      Func[func(ScriptInstanceDataPtr, Const[StringNamePtr]) Bool]
	ScriptInstancePropertyGetRevert      Func[func(ScriptInstanceDataPtr, Const[StringNamePtr], VariantPtr) Bool]
	ScriptInstanceGetOwner               Func[func(ScriptInstanceDataPtr) ObjectPtr]
	ScriptInstancePropertyStateAdd       Func[func(StringNamePtr, Const[VariantPtr], cgo.Handle)]
	ScriptInstanceGetPropertyState       Func[func(ScriptInstanceDataPtr, ScriptInstancePropertyStateAdd, cgo.Handle)]
	ScriptInstanceGetMethodList          Func[func(ScriptInstanceDataPtr, *uint32) UnsafeArray[MethodInfo]]
	ScriptInstanceFreeMethodList2        Func[func(ScriptInstanceDataPtr, UnsafeArray[MethodInfo], uint32)]
	ScriptInstanceGetPropertyType        Func[func(ScriptInstanceDataPtr, Const[StringNamePtr], *Bool) VariantType]
	ScriptInstanceValidateProperty       Func[func(ScriptInstanceDataPtr, *PropertyInfo) Bool]
	ScriptInstanceHasMethod              Func[func(ScriptInstanceDataPtr, Const[StringNamePtr]) Bool]
	ScriptInstanceGetMethodArgumentCount Func[func(ScriptInstanceDataPtr, Const[StringNamePtr], *Bool) Int]
	ScriptInstanceCall                   Func[func(ScriptInstanceDataPtr, Const[StringNamePtr], UnsafeArray[Const[VariantPtr]], Int, VariantPtr, *CallError)]
	ScriptInstanceNotification2          Func[func(ScriptInstanceDataPtr, Int, Bool)]
	ScriptInstanceToString               Func[func(ScriptInstanceDataPtr, *Bool, StringPtr)]
	ScriptInstanceRefCountIncremented    Func[func(ScriptInstanceDataPtr)]
	ScriptInstanceRefCountDecremented    Func[func(ScriptInstanceDataPtr) Bool]
	ScriptInstanceGetScript              Func[func(ScriptInstanceDataPtr) ObjectPtr]
	ScriptInstanceIsPlaceholder          Func[func(ScriptInstanceDataPtr) Bool]
	ScriptInstanceGetLanguage            Func[func(ScriptInstanceDataPtr) ScriptLanguagePtr]
	ScriptInstanceFree                   Func[func(ScriptInstanceDataPtr)]
)

type MethodInfo struct {
	_ structs.HostLayout

	name                   StringNamePtr
	return_value           PropertyInfo
	flags                  uint32
	id                     int32
	argument_count         uint32
	arguments              UnsafeArray[PropertyInfo]
	default_argument_count uint32
	default_arguments      UnsafeArray[Variant]
}

type InstanceBindingCallbacks struct {
	_ structs.HostLayout

	create_callback    BindingCreateCallback
	free_callback      BindingFreeCallback
	reference_callback BindingReferenceCallback
}

type (
	BindingCreateCallback    Func[func(ClassLibraryPtr, cgo.Handle) cgo.Handle]
	BindingFreeCallback      Func[func(ClassLibraryPtr, cgo.Handle, cgo.Handle)]
	BindingReferenceCallback Func[func(ClassLibraryPtr, cgo.Handle, Bool) Bool]
)

type (
	PtrUtilityFunction   Func[func(TypePtr, UnsafeArray[Const[TypePtr]], Int)]
	PtrKeyedChecker      Func[func(Const[VariantPtr], Const[VariantPtr]) uint32]
	PtrKeyedGetter       Func[func(Const[TypePtr], Const[TypePtr], TypePtr)]
	PtrKeyedSetter       Func[func(Const[TypePtr], Const[TypePtr], Const[TypePtr])]
	PtrIndexedGetter     Func[func(Const[TypePtr], Int, TypePtr)]
	PtrIndexedSetter     Func[func(Const[TypePtr], Int, Const[TypePtr])]
	PtrGetter            Func[func(Const[TypePtr], TypePtr)]
	PtrSetter            Func[func(Const[TypePtr], Const[TypePtr])]
	PtrDestructor        Func[func(TypePtr)]
	PtrConstructor       Func[func(TypePtr, UnsafeArray[Const[TypePtr]])]
	PtrBuiltInMethod     Func[func(TypePtr, UnsafeArray[Const[TypePtr]], TypePtr, Int)]
	PtrOperatorEvaluator Func[func(Const[TypePtr], Const[TypePtr], TypePtr)]

	VariantGetInternalPtrFunc Func[func(Const[VariantPtr]) TypePtr]

	TypeFromVariantConstructorFunc Func[func(TypePtr, VariantPtr)]
	VariantFromTypeConstructorFunc Func[func(VariantPtr, TypePtr)]
)

type GodotVersion struct {
	_ structs.HostLayout

	major  uint32
	minor  uint32
	patch  uint32
	string UnsafeString
}

type Const[T any] struct {
	RO T
}

type Uninitialized[T any] struct {
	WO T
}

type UnsafeString *byte
