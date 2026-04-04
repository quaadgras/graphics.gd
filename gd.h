#include <stdint.h>
#include <stdbool.h>

#ifndef __EMSCRIPTEN__
    #define VARIANT_ARG_OLD(n) uint64_t n##_1, uint64_t n##_2, uint64_t n##_3
    #define INT int64_t
    #define INT64(n) int64_t n
    #define SHAPE uint64_t
    #define ANY void*
    #define UINT uint64_t
    #define OBJECT_ID(n) uint64_t n
    #define UINT64(n) uint64_t n
    #define CALLABLE_ARG(n) uint64_t n##_1, uint64_t n##_2
    #define BUFFER char*
    #define STRING const char*
    #define RESULT_POINTER
    #define RETURNS(t) t
#else
    #define VARIANT_ARG_OLD(n) uint32_t n##_1, uint32_t n##_2, uint32_t n##_3, uint32_t n##_4, uint32_t n##_5, uint32_t n##_6
    #define INT int32_t
    #define ANY uint32_t
    #define UINT uint32_t
    #define UINT64(n) uint32_t n##_1, uint32_t n##_2
    #define INT64(n) uint32_t n##_1, uint32_t n##_2
    #define SHAPE(n) uint32_t n##_1, uint32_t n##_2
    #define OBJECT_ID(n) uint32_t n##_1, uint32_t n##_2
    #define CALLABLE_ARG(n) uint32_t n##_1, uint32_t n##_2, uint32_t n##_3, uint32_t n##_4
    #define BUFFER uint32_t
    #define STRING std::string
    #define RESULT_POINTER , uint32_t result
    #define RETURNS(t) void
    extern "C" {
#endif

#define VARIANT_ARG(n) uint64_t n##_1, uint64_t n##_2, uint64_t n##_3
#define PACKED_ARRAY_ARG(n) UINT n##_1, UINT n##_2

typedef struct {
    uint64_t tag;
    uint64_t payload[2];
} Variant;

typedef const Variant* const* VariantArguments;

typedef struct {
    uint64_t array;
    uint64_t length;
} PackedStringArray;

typedef struct {
	uint32_t error;
	int32_t argument;
	int32_t expected;
} CallError;

typedef uint64_t Shape;

static const Shape ShapeEmpty = 0;
static const Shape ShapeBytes1 = 1;
static const Shape ShapeBytes2 = 2;
static const Shape ShapeBytes4 = 3;
static const Shape ShapeBytes8 = 4;
static const Shape ShapeBytes4x2 = 5;
static const Shape ShapeBytes4x3 = 6;
static const Shape ShapeBytes8x2 = 7;
static const Shape ShapeBytes4x4 = 8;
static const Shape ShapeBytes8x3 = 9;
static const Shape ShapeBytes4x6 = 10;
static const Shape ShapeBytes4x9 = 11;
static const Shape ShapeBytes4x12 = 12;
static const Shape ShapeBytes4x16 = 13;

typedef uint32_t InitializationLevel;
typedef uint32_t VariantType;
typedef uint32_t VariantOperator;
typedef uint32_t MethodFlags;
typedef uint32_t ArgumentMetadata;

typedef int64_t Int;

typedef uintptr_t Object;
typedef uintptr_t ObjectType;

typedef uintptr_t RefCounted;
typedef uintptr_t String;
typedef uintptr_t StringName;
typedef uintptr_t Array;
typedef uintptr_t Dictionary;
typedef uintptr_t MethodForClass;
typedef uintptr_t ScriptInstance;

typedef uintptr_t TaskID;
typedef uintptr_t CallableID;
typedef uintptr_t FunctionID;
typedef uintptr_t ExtensionClassID;
typedef uintptr_t ExtensionInstanceID;
typedef uintptr_t ExtensionBindingID;

typedef uintptr_t PropertyList;
typedef uintptr_t MethodList;

extern void gd_on_callable_call(CallableID c, Variant* result, INT arg_count, VariantArguments args, CallError* err);
extern bool gd_on_callable_validation(CallableID c);
extern void gd_on_callable_free(CallableID c);
extern uint32_t gd_on_callable_hash(CallableID c);
extern bool gd_on_callable_compare(CallableID a, CallableID b);
extern bool gd_on_callable_less_than(CallableID a, CallableID b);
extern String gd_on_callable_stringify(CallableID c, CallError* err);
extern INT gd_on_callable_get_argument_count(CallableID c, CallError* err);

extern void gd_on_editor_class_in_use_detection(PACKED_ARRAY_ARG(a), PackedStringArray* result);

extern void gd_on_engine_init(InitializationLevel level);
extern void gd_on_engine_exit(InitializationLevel level);

extern ExtensionBindingID gd_on_extension_binding_created(ExtensionInstanceID inst);
extern void gd_on_extension_binding_removed(ExtensionInstanceID inst, ExtensionBindingID p1);
extern bool gd_on_extension_binding_reference(ExtensionInstanceID inst, bool p1);

extern Object gd_on_extension_class_create(ExtensionClassID class_name, bool notify_postinitialize);
extern FunctionID gd_on_extension_class_method(ExtensionClassID class_name, StringName method, uint32_t hash);
extern FunctionID gd_on_extension_class_caller(ExtensionClassID class_name, StringName method, uint32_t hash);

extern bool gd_on_extension_instance_set(ExtensionInstanceID inst, StringName property, VARIANT_ARG_OLD(val));
extern bool gd_on_extension_instance_get(ExtensionInstanceID inst, StringName property, Variant* p2);
extern PropertyList gd_on_extension_instance_property_list(ExtensionInstanceID inst);
extern bool gd_on_extension_instance_property_has_default(ExtensionInstanceID inst, StringName property);
extern bool gd_on_extension_instance_property_get_default(ExtensionInstanceID inst, StringName property, Variant* result);
extern bool gd_on_extension_instance_property_validation(ExtensionInstanceID inst, StringName property);
extern void gd_on_extension_instance_notification(ExtensionInstanceID inst, int32_t what, bool reverse);
extern String gd_on_extension_instance_stringify(ExtensionInstanceID inst);
extern bool gd_on_extension_instance_reference(ExtensionInstanceID inst, bool increment);
extern RETURNS(uint64_t) gd_on_extension_instance_rid(ExtensionInstanceID inst RESULT_POINTER);
extern void gd_on_extension_instance_checked_call(ExtensionInstanceID inst, FunctionID fn, void* result, void* args);
extern void gd_on_extension_instance_variant_call(ExtensionInstanceID inst, FunctionID fn, Variant* result, VariantArguments args);
extern void gd_on_extension_instance_dynamic_call(ExtensionInstanceID inst, FunctionID fn, Variant* result, INT count, VariantArguments args, CallError* err);
extern void gd_on_extension_instance_free(ExtensionInstanceID inst);
extern void gd_on_extension_instance_called(ExtensionInstanceID inst, FunctionID fn, void* result, void* args);

extern bool gd_on_extension_script_categorization(ExtensionInstanceID inst, PropertyList p1);
extern uint32_t gd_on_extension_script_get_property_type(ExtensionInstanceID inst, CallError* err);
extern Object gd_on_extension_script_get_owner(ExtensionInstanceID inst);
extern void gd_on_extension_script_get_property_state(ExtensionInstanceID inst, FunctionID add, uintptr_t arg);
extern MethodList gd_on_extension_script_get_methods(ExtensionInstanceID inst);
extern bool gd_on_extension_script_has_method(ExtensionInstanceID inst, uintptr_t p1);
extern INT gd_on_extension_script_get_method_argument_count(ExtensionInstanceID inst, StringName property);
extern Object gd_on_extension_script_get(ExtensionInstanceID inst);
extern bool gd_on_extension_script_is_placeholder(ExtensionInstanceID inst);
extern Object gd_on_extension_script_get_language(ExtensionInstanceID inst);

extern void gd_on_first_frame(void);
extern void gd_on_every_frame(void);
extern void gd_on_final_frame(void);
extern void gd_on_worker_thread_pool_task(TaskID task);
extern void gd_on_worker_thread_pool_group_task(TaskID task, uint32_t n);

RETURNS(Variant) gd_array_get(Array a, Int i RESULT_POINTER);
void gd_array_set(Array a, Int i, VARIANT_ARG(v));

FunctionID gd_builtin_name(StringName utility, INT64(hash));
void gd_builtin_call(FunctionID utility, ANY result, SHAPE(shape), ANY args);

String gd_variant_type_name(VariantType t);
void gd_variant_type_make(VariantType t, ANY result, INT arg_count, ANY args, ANY err);
void gd_variant_type_call(VariantType t, StringName static_method_name, ANY result, INT arg_count, ANY args, ANY err);
bool gd_variant_type_convertable(VariantType t, VariantType to, bool strict);
void gd_variant_type_setup_array(Array a, VariantType elem, StringName class_name, VARIANT_ARG_OLD(v));
void gd_variant_type_setup_dictionary(Dictionary d,
    VariantType key, StringName key_class_name, VARIANT_ARG_OLD(key_script),
    VariantType val, StringName val_class_name, VARIANT_ARG_OLD(val_script)
);
void gd_variant_type_fetch_constant(VariantType t, StringName constant, ANY result);
FunctionID gd_variant_type_unsafe_constructor(VariantType t, INT n);
FunctionID gd_variant_type_evaluator(VariantOperator op, VariantType a, VariantType b);
FunctionID gd_variant_type_setter(VariantType t, StringName property);
FunctionID gd_variant_type_getter(VariantType t, StringName property);
bool gd_variant_type_has_property(VariantType t, StringName property);
FunctionID gd_variant_type_builtin_method(VariantType t, StringName method, INT64(hash));
void gd_variant_type_unsafe_call(ANY self, FunctionID fn, ANY result, SHAPE(shape), ANY args);
void gd_variant_type_unsafe_make(FunctionID constructor, ANY result, SHAPE(shape), ANY args);
void gd_variant_type_unsafe_free(VariantType t, SHAPE(shape), ANY args);

void gd_callable_create(CallableID fn, OBJECT_ID(object), ANY result);
FunctionID gd_callable_lookup(CALLABLE_ARG(c));

void gd_classdb_FileAccess_write(Object file, BUFFER buf, INT len);
INT gd_classdb_FileAccess_read(Object file, BUFFER buf, INT cap);
uintptr_t gd_classdb_Image_unsafe(Object img);
uint8_t gd_classdb_Image_access(Object img, INT offset);
void gd_classdb_WorkerThreadPool_add_task(Object pool, TaskID task, bool priority, String description);
void gd_classdb_WorkerThreadPool_add_group_task(Object pool, TaskID task, int32_t elements, int32_t arg, bool priority, String description);
INT gd_classdb_XMLParser_load(Object parser, BUFFER buf, INT cap);

MethodList gd_method_list_make(INT method_count);
void gd_method_list_push(MethodList list,
    StringName name, FunctionID call, MethodFlags method_flags,
    PropertyList return_value_info, PropertyList arguments_info,
    INT count, ANY default_arguments
);
void gd_method_list_free(MethodList list);

PropertyList gd_property_list_make(INT property_count);
void gd_property_list_push(PropertyList list,
    VariantType t, StringName name, StringName class_name,
    uint32_t hint, String hint_string, uint32_t usage, ArgumentMetadata meta
);

void gd_property_list_free(PropertyList list);
VariantType gd_property_info_type(PropertyList info);
StringName gd_property_info_name(PropertyList info);
StringName gd_property_info_class_name(PropertyList info);
uint32_t gd_property_info_hint(PropertyList info);
String gd_property_info_hint_string(PropertyList info);
uint32_t gd_property_info_usage(PropertyList info);

void gd_classdb_register(
    StringName class_name, StringName parent_class,
    ExtensionClassID id, bool is_virtual, bool is_abstract,
    bool is_exposed, bool is_runtime, String icon_path
);
void gd_classdb_register_methods(StringName class_name, MethodList methods);
void gd_classdb_register_constant(
    StringName class_name, StringName enum_name, StringName constant_name,
    INT64(value), bool bitfield
);
void gd_classdb_register_property(StringName class_name, PropertyList property, StringName setter, StringName getter);
void gd_classdb_register_property_indexed(
    StringName class_name, PropertyList property,
    StringName setter, StringName getter, INT64(index)
);
void gd_classdb_register_property_group(StringName class_name, String group, String prefix);
void gd_classdb_register_property_sub_group(StringName class_name, String subgroup, String prefix);
void gd_classdb_register_signal(StringName class_name, StringName signal, PropertyList args);
void gd_classdb_register_removal(StringName class_name);

void gd_packed_dictionary_access(Dictionary d, VARIANT_ARG_OLD(key), ANY result);
void gd_packed_dictionary_modify(Dictionary d, VARIANT_ARG_OLD(key), VARIANT_ARG_OLD(val));

void gd_editor_add_documentation(STRING xml, INT len);
void gd_editor_add_plugin(StringName class_name);
void gd_editor_end_plugin(StringName class_name);

void gd_iterator_make(VARIANT_ARG_OLD(v), ANY result_iter, ANY err);
bool gd_iterator_next(VARIANT_ARG_OLD(v), ANY iter, ANY err);
void gd_iterator_load(VARIANT_ARG_OLD(v), VARIANT_ARG_OLD(i), ANY result, ANY err);

String gd_library_location();

void gd_log_error(
    STRING text, INT text_len,
    STRING code, INT code_len,
    STRING func, INT func_len,
    STRING file, INT file_len,
    int32_t line, bool notify_editor
);
void gd_log_warning(
    STRING text, INT text_len,
    STRING code, INT code_len,
    STRING func, INT func_len,
    STRING file, INT file_len,
    int32_t line, bool notify_editor
);

uintptr_t gd_memory_malloc(INT size);
INT gd_memory_sizeof(StringName struct_name);
uintptr_t gd_memory_resize(uintptr_t addr, INT size);
void gd_memory_clear(uintptr_t addr, INT size);
void gd_memory_free(uintptr_t addr);
void gd_memory_edit_byte(uintptr_t addr, uint8_t v);
void gd_memory_edit_u16(uintptr_t addr, uint16_t v);
void gd_memory_edit_u32(uintptr_t addr, uint32_t v);
void gd_memory_edit_u64(uintptr_t addr, UINT64(v));
void gd_memory_edit_128(uintptr_t addr, UINT64(v1), UINT64(v2));
void gd_memory_edit_256(uintptr_t addr, UINT64(v1), UINT64(v2), UINT64(v3), UINT64(v4));
void gd_memory_edit_512(uintptr_t addr, UINT64(v1), UINT64(v2), UINT64(v3), UINT64(v4), UINT64(v5), UINT64(v6), UINT64(v7), UINT64(v8));
uint8_t gd_memory_load_byte(uintptr_t addr);
uint16_t gd_memory_load_u16(uintptr_t addr);
uint32_t gd_memory_load_u32(uintptr_t addr);
uint64_t gd_memory_load_u64(uintptr_t addr);

Object gd_object_make(StringName name);
void gd_object_call(Object obj, MethodForClass method, ANY result, INT arg_count, ANY args, ANY err);
StringName gd_object_name(Object obj);
ObjectType gd_object_type(StringName name);
Object gd_object_cast(Object obj, ObjectType to);
Object gd_object_lookup(OBJECT_ID(id));
Object gd_object_global(StringName name);
void gd_object_extension_setup(Object obj, StringName name, ExtensionInstanceID class_name);
ExtensionInstanceID gd_object_extension_fetch(Object obj);
void gd_object_extension_close(Object obj);
void gd_object_id(Object obj, ANY id);
void gd_object_id_inside_variant(VARIANT_ARG_OLD(v), ANY id);
MethodForClass gd_object_method_lookup(StringName name, StringName method, INT64(hash));
ScriptInstance gd_object_script_make(ExtensionInstanceID fn);
void gd_object_script_call(Object obj, StringName name, ANY result, INT arg_count, ANY args, ANY err);
void gd_object_script_setup(Object obj, ScriptInstance script);
ScriptInstance gd_object_script_fetch(Object obj, Object language);
bool gd_object_script_defines_method(Object obj, StringName method);
void gd_object_script_property_state_add(FunctionID fn, uintptr_t arg, StringName name, VARIANT_ARG_OLD(state));
ScriptInstance gd_object_script_placeholder_create(Object language, Object script, Object owner);
void gd_object_script_placeholder_update(ScriptInstance script, Array array, Dictionary dict);
void gd_object_unsafe_call(Object obj, MethodForClass fn, ANY result, SHAPE(shape), ANY args);
void gd_object_unsafe_free(Object obj);

uintptr_t gd_packed_byte_array_unsafe(PACKED_ARRAY_ARG(pa));
uint8_t gd_packed_byte_array_access(PACKED_ARRAY_ARG(pa), INT idx);
uintptr_t gd_packed_color_array_unsafe(PACKED_ARRAY_ARG(pa));
void gd_packed_color_array_access(PACKED_ARRAY_ARG(pa), INT idx, ANY result);
uintptr_t gd_packed_float32_array_unsafe(PACKED_ARRAY_ARG(pa));
float gd_packed_float32_array_access(PACKED_ARRAY_ARG(pa), INT idx);
uintptr_t gd_packed_float64_array_unsafe(PACKED_ARRAY_ARG(pa));
double gd_packed_float64_array_access(PACKED_ARRAY_ARG(pa), INT idx);
uintptr_t gd_packed_int32_array_unsafe(PACKED_ARRAY_ARG(pa));
int32_t gd_packed_int32_array_access(PACKED_ARRAY_ARG(pa), INT idx);
uintptr_t gd_packed_int64_array_unsafe(PACKED_ARRAY_ARG(pa));
void gd_packed_int64_array_access(PACKED_ARRAY_ARG(pa), INT idx, ANY result);
uintptr_t gd_packed_string_array_unsafe(PACKED_ARRAY_ARG(pa));
String gd_packed_string_array_access(PACKED_ARRAY_ARG(pa), INT idx);
uintptr_t gd_packed_vector2_array_unsafe(PACKED_ARRAY_ARG(pa));
void gd_packed_vector2_array_access(PACKED_ARRAY_ARG(pa), INT idx, ANY result);
uintptr_t gd_packed_vector3_array_unsafe(PACKED_ARRAY_ARG(pa));
void gd_packed_vector3_array_access(PACKED_ARRAY_ARG(pa), INT idx, ANY result);
uintptr_t gd_packed_vector4_array_unsafe(PACKED_ARRAY_ARG(pa));
void gd_packed_vector4_array_access(PACKED_ARRAY_ARG(pa), INT idx, ANY result);

Object gd_ref_get_object(RefCounted ref);
void gd_ref_set_object(RefCounted ref, Object obj);

int32_t gd_string_access(String s, INT idx);
String gd_string_resize(String s, INT size);
uintptr_t gd_string_unsafe(String s);
String gd_string_append(String s, String other);
String gd_string_append_rune(String s, int32_t ch);
String gd_string_decode_latin1(STRING s, INT len);
String gd_string_decode_utf8(STRING s, INT len);
String gd_string_decode_utf16(STRING s, INT len, bool little_endian);
String gd_string_decode_utf32(STRING s, INT len);
String gd_string_decode_wide(STRING s, INT len);
INT gd_string_encode_latin1(String s, BUFFER buf, INT cap);
INT gd_string_encode_utf8(String s, BUFFER buf, INT cap);
INT gd_string_encode_utf16(String s, BUFFER buf, INT cap);
INT gd_string_encode_utf32(String s, BUFFER buf, INT cap);
INT gd_string_encode_wide(String s, BUFFER buf, INT cap);
StringName gd_string_intern_latin1(STRING s, INT len);
StringName gd_string_intern_utf8(STRING s, INT len);

bool gd_thread_is_main();

void gd_variant_zero(ANY result);
void gd_variant_copy(VARIANT_ARG_OLD(v), ANY result);
void gd_variant_call(VARIANT_ARG_OLD(v), StringName method, ANY result, INT arg_count, ANY args, ANY err);
bool gd_variant_eval(VariantOperator op, VARIANT_ARG_OLD(a), VARIANT_ARG_OLD(b), ANY result);
void gd_variant_hash(VARIANT_ARG_OLD(v), ANY hash);
bool gd_variant_bool(VARIANT_ARG_OLD(v));
String gd_variant_text(VARIANT_ARG_OLD(v));
VariantType gd_variant_type(VARIANT_ARG_OLD(v));
void gd_variant_deep_copy(VARIANT_ARG_OLD(v), ANY result);
void gd_variant_deep_hash(VARIANT_ARG_OLD(v), INT recursion_count, ANY hash);
bool gd_variant_get_index(VARIANT_ARG_OLD(v), VARIANT_ARG_OLD(key), ANY result);
bool gd_variant_get_array(VARIANT_ARG_OLD(v), INT idx, ANY result, ANY err);
bool gd_variant_get_field(VARIANT_ARG_OLD(v), StringName field, ANY result);
bool gd_variant_has_index(VARIANT_ARG_OLD(v), VARIANT_ARG_OLD(idx));
bool gd_variant_has_method(VARIANT_ARG_OLD(v), StringName method);
bool gd_variant_set_index(VARIANT_ARG_OLD(v), VARIANT_ARG_OLD(key), VARIANT_ARG_OLD(val));
bool gd_variant_set_array(VARIANT_ARG_OLD(v), INT idx, VARIANT_ARG_OLD(val), ANY err);
bool gd_variant_set_field(VARIANT_ARG_OLD(v), StringName field, VARIANT_ARG_OLD(val));
void gd_variant_unsafe_call(FunctionID fn, ANY result, SHAPE(shape), ANY args);
void gd_variant_unsafe_eval(FunctionID fn, ANY result, SHAPE(shape), ANY args);
void gd_variant_unsafe_free(VARIANT_ARG_OLD(v));
void gd_variant_unsafe_make_native(VariantType vtype, VARIANT_ARG_OLD(v), SHAPE(shape), ANY result);
void gd_variant_unsafe_from_native(VariantType vtype, ANY result, SHAPE(shape), ANY args);
uintptr_t gd_variant_unsafe_internal_pointer(VariantType vtype, VARIANT_ARG_OLD(v));
void gd_variant_unsafe_get_field(FunctionID getter, ANY result, SHAPE(shape), ANY args);
void gd_variant_unsafe_get_array(VariantType vtype, INT idx, ANY result, SHAPE(shape), ANY args);
void gd_variant_unsafe_get_index(VariantType vtype, ANY result, SHAPE(shape), ANY args);
void gd_variant_unsafe_set_field(FunctionID setter, SHAPE(shape), ANY args);
void gd_variant_unsafe_set_array(VariantType vtype, INT idx, SHAPE(shape), ANY args);
void gd_variant_unsafe_set_index(VariantType vtype, SHAPE(shape), ANY args);

uint32_t gd_version_major(void);
uint32_t gd_version_minor(void);
uint32_t gd_version_patch(void);
uint32_t gd_version_hex(void);
String gd_version_status(void);
String gd_version_build(void);
String gd_version_hash(void);
uint64_t gd_version_timestamp(void);
String gd_version_string(void);

#ifdef __EMSCRIPTEN__
}
#endif
