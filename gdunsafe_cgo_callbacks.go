//go:build cgo

package gdunsafe

// #include "gd.h"
import "C"
import (
	"unsafe"

	"graphics.gd/variant"
)

// Lifecycle callbacks

//export go_on_engine_init
func go_on_engine_init(level C.gd_initialization_level) {
	l := InitializationLevel(level)
	for _, fn := range onEngineInit {
		fn(l)
	}
}

//export go_on_engine_exit
func go_on_engine_exit(level C.gd_initialization_level) {
	l := InitializationLevel(level)
	for _, fn := range onEngineExit {
		fn(l)
	}
}

//export go_on_yield
func go_on_yield() {}

//export go_on_first_frame
func go_on_first_frame() {
	for _, fn := range onFirstFrame {
		fn()
	}
}

//export go_on_every_frame
func go_on_every_frame() {
	for _, fn := range onEveryFrame {
		fn()
	}
}

//export go_on_final_frame
func go_on_final_frame() {
	for _, fn := range onFinalFrame {
		fn()
	}
}

// Extension class callbacks

//export go_on_extension_class_create
func go_on_extension_class_create(id C.gd_extension_class_id, notify C.bool) C.struct_Object {
	obj := classes.Get(uintptr(id)).Create(bool(notify))
	return *(*C.struct_Object)(unsafe.Pointer(&obj.raw))
}

//export go_on_extension_class_method
func go_on_extension_class_method(id C.gd_extension_class_id, method C.struct_StringName, hash C.uint32_t) C.gd_extension_method_id {
	name := *(*StringName)(unsafe.Pointer(&method))
	fn := classes.Get(uintptr(id)).Method(name, uint32(hash))
	if fn == nil {
		return 0
	}
	return C.gd_extension_method_id(functions.New(fn))
}

//export go_on_extension_class_caller
func go_on_extension_class_caller(id C.gd_extension_class_id, method C.struct_StringName, hash C.uint32_t) C.gd_extension_method_id {
	name := *(*StringName)(unsafe.Pointer(&method))
	fn := classes.Get(uintptr(id)).Method(name, uint32(hash))
	if fn == nil {
		return 0
	}
	return C.gd_extension_method_id(functions.New(fn))
}

// Extension instance callbacks

//export go_on_extension_instance_set
func go_on_extension_instance_set(inst C.gd_extension_object_id, property C.struct_StringName, v_1, v_2, v_3 C.uint64_t) C.bool {
	name := *(*StringName)(unsafe.Pointer(&property))
	return C.bool(instances.Get(uintptr(inst)).Set(name, Variant{uint64(v_1), uint64(v_2), uint64(v_3)}))
}

//export go_on_extension_instance_get
func go_on_extension_instance_get(inst C.gd_extension_object_id, property C.struct_StringName, result *C.struct_Variant) C.bool {
	name := *(*StringName)(unsafe.Pointer(&property))
	v, ok := instances.Get(uintptr(inst)).Get(name)
	if ok {
		*(*Variant)(unsafe.Pointer(result)) = v
	}
	return C.bool(ok)
}

//export go_on_extension_instance_property_has_default
func go_on_extension_instance_property_has_default(inst C.gd_extension_object_id, property C.struct_StringName) C.bool {
	name := *(*StringName)(unsafe.Pointer(&property))
	return C.bool(instances.Get(uintptr(inst)).HasDefault(name))
}

//export go_on_extension_instance_property_get_default
func go_on_extension_instance_property_get_default(inst C.gd_extension_object_id, property C.struct_StringName, result *C.struct_Variant) C.bool {
	name := *(*StringName)(unsafe.Pointer(&property))
	v, ok := instances.Get(uintptr(inst)).GetDefault(name)
	if ok {
		*(*Variant)(unsafe.Pointer(result)) = v
	}
	return C.bool(ok)
}

//export go_on_extension_instance_property_list
func go_on_extension_instance_property_list(inst C.gd_extension_object_id) C.gd_property_list_t {
	return C.gd_property_list_t(instances.Get(uintptr(inst)).PropertyList())
}

//export go_on_extension_instance_property_validation
func go_on_extension_instance_property_validation(inst C.gd_extension_object_id, vtype C.VariantType, name, class_name C.struct_StringName, hint C.uint32_t, hint_string C.struct_String, usage C.uint32_t) C.gd_property_list_t {
	n := *(*StringName)(unsafe.Pointer(&name))
	cn := *(*StringName)(unsafe.Pointer(&class_name))
	hs := *(*String)(unsafe.Pointer(&hint_string))
	ok := instances.Get(uintptr(inst)).ValidateProperty(Property{
		Type:       variant.Type(vtype),
		Name:       n,
		ClassName:  cn,
		Hint:       uint32(hint),
		HintString: hs,
		Usage:      uint32(usage),
	})
	if ok {
		return 0 // null = valid, no changes
	}
	// return empty property list to signal invalid
	pl := MakePropertyList(0)
	return C.gd_property_list_t(pl)
}

//export go_on_extension_instance_notification
func go_on_extension_instance_notification(inst C.gd_extension_object_id, what C.int32_t, reverse C.bool) {
	instances.Get(uintptr(inst)).Notification(int32(what), bool(reverse))
}

//export go_on_extension_instance_stringify
func go_on_extension_instance_stringify(inst C.gd_extension_object_id, ok *C.bool, out *C.gd_addr) {
	s := instances.Get(uintptr(inst)).UnsafeString()
	if s.raw != (gdString{}) {
		*(*gdString)(unsafe.Pointer(out)) = s.raw
		*ok = true
	} else {
		*ok = false
	}
}

//export go_on_extension_instance_reference
func go_on_extension_instance_reference(inst C.gd_extension_object_id) {
	instances.Get(uintptr(inst)).Reference(true)
}

//export go_on_extension_instance_unreference
func go_on_extension_instance_unreference(inst C.gd_extension_object_id) C.bool {
	return C.bool(instances.Get(uintptr(inst)).Reference(false))
}

//export go_on_extension_instance_rid
func go_on_extension_instance_rid(inst C.gd_extension_object_id) C.RID {
	return C.RID(instances.Get(uintptr(inst)).RID())
}

//export go_on_extension_instance_checked_call
func go_on_extension_instance_checked_call(fn C.gd_extension_method_id, inst C.gd_extension_object_id, args C.gd_addr, result C.gd_addr) {
	var i ExtensionInstance
	if inst != 0 {
		i = instances.Get(uintptr(inst))
	}
	functions.Get(uintptr(fn)).PointerCall(i, MutablePointer(result), PointerTo[Pointer](args))
}

//export go_on_extension_instance_dynamic_call
func go_on_extension_instance_dynamic_call(fn C.gd_extension_method_id, inst C.gd_extension_object_id, args C.gd_unsafe_variants, count C.int64_t, result *C.struct_Variant, err *C.gd_error) {
	var i ExtensionInstance
	if inst != 0 {
		i = instances.Get(uintptr(inst))
	}
	v, e := functions.Get(uintptr(fn)).DynamicCall(i, Variants{
		first: PointerTo[PointerTo[Variant]](uintptr(unsafe.Pointer(args))),
		count: int(count),
	})
	*(*Variant)(unsafe.Pointer(result)) = v
	*(*Error)(unsafe.Pointer(err)) = e
}

//export go_on_extension_instance_free
func go_on_extension_instance_free(inst C.gd_extension_object_id) {
	i := instances.Get(uintptr(inst))
	if f, ok := i.(interface{ Free() }); ok {
		f.Free()
	}
	instances.Del(uintptr(inst))
}

//export go_on_extension_instance_called
func go_on_extension_instance_called(inst C.gd_extension_object_id, name C.struct_StringName, fn C.gd_extension_method_id, args C.gd_addr, result C.gd_addr) {
	i := instances.Get(uintptr(inst))
	functions.Get(uintptr(fn)).PointerCall(i, MutablePointer(result), PointerTo[Pointer](args))
}

// Callable callbacks

//export go_on_callable_called
func go_on_callable_called(c C.gd_extension_callable_t, args C.gd_unsafe_variants, argc C.int64_t, ret *C.struct_Variant, err *C.gd_error) {
	v, e := callables.Get(uintptr(c)).Call(Variants{
		first: PointerTo[PointerTo[Variant]](uintptr(unsafe.Pointer(args))),
		count: int(argc),
	})
	*(*Variant)(unsafe.Pointer(ret)) = v
	*(*Error)(unsafe.Pointer(err)) = e
}

//export go_on_callable_verify
func go_on_callable_verify(c C.gd_extension_callable_t) C.bool {
	return C.bool(callables.Get(uintptr(c)).IsValid())
}

//export go_on_callable_delete
func go_on_callable_delete(c C.gd_extension_callable_t) {
	callables.Del(uintptr(c))
}

//export go_on_callable_hashed
func go_on_callable_hashed(c C.gd_extension_callable_t) C.uint32_t {
	return C.uint32_t(callables.Get(uintptr(c)).Hash())
}

//export go_on_callable_equal
func go_on_callable_equal(a, b C.gd_extension_callable_t) C.bool {
	return C.bool(callables.Get(uintptr(a)).Compare(callables.Get(uintptr(b))) == 0)
}

//export go_on_callable_less_than
func go_on_callable_less_than(a, b C.gd_extension_callable_t) C.bool {
	return C.bool(callables.Get(uintptr(a)).Compare(callables.Get(uintptr(b))) < 0)
}

//export go_on_callable_string
func go_on_callable_string(c C.gd_extension_callable_t, ok *C.bool, out C.gd_addr) {
	s := callables.Get(uintptr(c)).UnsafeString()
	*ok = C.bool(s.raw != (gdString{}))
	*(*gdString)(unsafe.Pointer(out)) = s.raw
}

//export go_on_callable_length
func go_on_callable_length(c C.gd_extension_callable_t, ok *C.bool) C.int64_t {
	*ok = true
	return C.int64_t(callables.Get(uintptr(c)).ArgumentCount())
}

// Extension binding callbacks

//export go_on_extension_binding_created
func go_on_extension_binding_created(token C.uintptr_t, inst C.gd_extension_object_id) C.gd_extension_binding_id {
	return 0
}

//export go_on_extension_binding_removed
func go_on_extension_binding_removed(token C.uintptr_t, inst C.gd_extension_object_id, binding C.gd_extension_binding_id) {
}

//export go_on_extension_binding_reference
func go_on_extension_binding_reference(token C.uintptr_t, inst C.gd_extension_object_id, reference C.bool) C.bool {
	return false
}

// Worker thread pool callbacks

//export go_on_worker_thread_pool_task
func go_on_worker_thread_pool_task(task C.gd_extension_task_id) {
}

//export go_on_worker_thread_pool_group_task
func go_on_worker_thread_pool_group_task(task C.gd_extension_task_id, n C.uint32_t) {
}

// Editor callback

//export go_on_editor_class_in_use_detection
func go_on_editor_class_in_use_detection(packed C.gd_addr) {
}

// Extension script callbacks

//export go_on_extension_script_property_iter
func go_on_extension_script_property_iter(inst C.gd_extension_object_id, addFunc C.gd_property_iterator_t, arg C.uintptr_t) {
	script, ok := instances.Get(uintptr(inst)).(ExtensionScript)
	if !ok {
		return
	}
	script.ExportedProperties(func(name StringName, value Variant) bool {
		gd_script_yield_property(gd_property_iterator_t(addFunc), uintptr(arg), name.raw, value[0], value[1], value[2])
		return true
	})
}

//export go_on_extension_script_categorization
func go_on_extension_script_categorization(inst C.gd_extension_object_id, propertyInfo C.gd_addr) C.bool {
	script, ok := instances.Get(uintptr(inst)).(ExtensionScript)
	if !ok {
		return false
	}
	return C.bool(script.PropertyCategory() != StringName{})
}

//export go_on_extension_script_get_property_type
func go_on_extension_script_get_property_type(inst C.gd_extension_object_id, property C.struct_StringName, ok *C.bool) C.VariantType {
	script, s_ok := instances.Get(uintptr(inst)).(ExtensionScript)
	if !s_ok {
		*ok = false
		return 0
	}
	name := *(*StringName)(unsafe.Pointer(&property))
	*ok = true
	return C.VariantType(script.PropertyType(name))
}

//export go_on_extension_script_get_owner
func go_on_extension_script_get_owner(inst C.gd_extension_object_id) C.struct_Object {
	script, ok := instances.Get(uintptr(inst)).(ExtensionScript)
	if !ok {
		return C.struct_Object{}
	}
	obj := script.Owner()
	return *(*C.struct_Object)(unsafe.Pointer(&obj.raw))
}

//export go_on_extension_script_get_methods
func go_on_extension_script_get_methods(inst C.gd_extension_object_id, count *C.uint32_t) C.gd_addr {
	script, ok := instances.Get(uintptr(inst)).(ExtensionScript)
	if !ok {
		*count = 0
		return nil
	}
	ml := script.MethodList()
	// TODO: return the method info array and count
	_ = ml
	*count = 0
	return nil
}

//export go_on_extension_script_has_method
func go_on_extension_script_has_method(inst C.gd_extension_object_id, method C.struct_StringName) C.bool {
	script, ok := instances.Get(uintptr(inst)).(ExtensionScript)
	if !ok {
		return false
	}
	name := *(*StringName)(unsafe.Pointer(&method))
	return C.bool(script.HasMethod(name))
}

//export go_on_extension_script_get_method_argument_count
func go_on_extension_script_get_method_argument_count(inst C.gd_extension_object_id, method C.struct_StringName, ok *C.bool) C.int64_t {
	script, s_ok := instances.Get(uintptr(inst)).(ExtensionScript)
	if !s_ok {
		*ok = false
		return 0
	}
	name := *(*StringName)(unsafe.Pointer(&method))
	*ok = true
	return C.int64_t(script.MethodArgumentCount(name))
}

//export go_on_extension_script_get
func go_on_extension_script_get(inst C.gd_extension_object_id) C.struct_Object {
	script, ok := instances.Get(uintptr(inst)).(ExtensionScript)
	if !ok {
		return C.struct_Object{}
	}
	obj := script.Script()
	return *(*C.struct_Object)(unsafe.Pointer(&obj.raw))
}

//export go_on_extension_script_is_placeholder
func go_on_extension_script_is_placeholder(inst C.gd_extension_object_id) C.bool {
	script, ok := instances.Get(uintptr(inst)).(ExtensionScript)
	if !ok {
		return false
	}
	return C.bool(script.IsPlaceholder())
}

//export go_on_extension_script_get_language
func go_on_extension_script_get_language(inst C.gd_extension_object_id) C.struct_Object {
	script, ok := instances.Get(uintptr(inst)).(ExtensionScript)
	if !ok {
		return C.struct_Object{}
	}
	obj := script.ScriptLanguage()
	return *(*C.struct_Object)(unsafe.Pointer(&obj.raw))
}

// ensure variant import is used
var _ variant.Type
