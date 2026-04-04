#include <stdint.h>
#include <stdbool.h>

typedef struct {
    uint64_t tag;
    uint64_t payload[2];
} Variant;

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

typedef uintptr_t Object;
typedef uintptr_t ObjectType;
typedef uint64_t ObjectID;
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

extern void gd_on_callable_call(CallableID c, Variant* result, int64_t arg_count, Variant* args, CallError* err);
extern bool gd_on_callable_validation(CallableID c);
extern void gd_on_callable_free(CallableID c);
extern uint32_t gd_on_callable_hash(CallableID c);
extern bool gd_on_callable_compare(CallableID a, CallableID b);
extern bool gd_on_callable_less_than(CallableID a, CallableID b);
extern String gd_on_callable_stringify(CallableID c, CallError* err);
extern int64_t gd_on_callable_get_argument_count(CallableID c, CallError* err);

extern void gd_on_editor_class_in_use_detection(uint64_t p0, uint64_t p1, PackedStringArray* result);

extern void gd_on_engine_init(InitializationLevel level);
extern void gd_on_engine_exit(InitializationLevel level);

extern ExtensionBindingID gd_on_extension_binding_created(ExtensionInstanceID inst);
extern void gd_on_extension_binding_removed(ExtensionInstanceID inst, ExtensionBindingID p1);
extern bool gd_on_extension_binding_reference(ExtensionInstanceID inst, bool p1);

extern Object gd_on_extension_class_create(ExtensionClassID class, bool notify_postinitialize);
extern FunctionID gd_on_extension_class_method(ExtensionClassID class, StringName method, uint32_t hash);
extern FunctionID gd_on_extension_class_caller(ExtensionClassID class, StringName method, uint32_t hash);

extern bool gd_on_extension_instance_set(ExtensionInstanceID inst, uintptr_t p1, uint64_t p2, uint64_t p3, uint64_t p4);
extern bool gd_on_extension_instance_get(ExtensionInstanceID inst, uintptr_t p1, Variant* p2);
extern PropertyList gd_on_extension_instance_property_list(ExtensionInstanceID inst);
extern bool gd_on_extension_instance_property_has_default(ExtensionInstanceID inst, StringName property);
extern bool gd_on_extension_instance_property_get_default(ExtensionInstanceID inst, StringName property, Variant* result);
extern bool gd_on_extension_instance_property_validation(ExtensionInstanceID inst, StringName property);
extern void gd_on_extension_instance_notification(ExtensionInstanceID inst, int32_t what, bool reverse);
extern String gd_on_extension_instance_stringify(ExtensionInstanceID inst);
extern bool gd_on_extension_instance_reference(ExtensionInstanceID inst, bool increment);
extern void gd_on_extension_instance_rid(ExtensionInstanceID inst, uint64_t* rid);
extern void gd_on_extension_instance_checked_call(ExtensionInstanceID inst, FunctionID fn, void* result, void* args);
extern void gd_on_extension_instance_variant_call(ExtensionInstanceID inst, FunctionID fn, Variant* result, Variant* args);
extern void gd_on_extension_instance_dynamic_call(ExtensionInstanceID inst, FunctionID fn, Variant* result, int64_t count, Variant* args, CallError* err);
extern void gd_on_extension_instance_free(ExtensionInstanceID inst);
extern void gd_on_extension_instance_called(ExtensionInstanceID inst, FunctionID fn, void* result, void* args);

extern bool gd_on_extension_script_categorization(ExtensionInstanceID inst, PropertyList p1);
extern uint32_t gd_on_extension_script_get_property_type(ExtensionInstanceID inst, CallError* err);
extern Object gd_on_extension_script_get_owner(ExtensionInstanceID inst);
extern void gd_on_extension_script_get_property_state(ExtensionInstanceID inst, FunctionID add, uintptr_t arg);
extern MethodList gd_on_extension_script_get_methods(ExtensionInstanceID inst);
extern bool gd_on_extension_script_has_method(ExtensionInstanceID inst, uintptr_t p1);
extern int64_t gd_on_extension_script_get_method_argument_count(ExtensionInstanceID inst, StringName property);
extern Object gd_on_extension_script_get(ExtensionInstanceID inst);
extern bool gd_on_extension_script_is_placeholder(ExtensionInstanceID inst);
extern Object gd_on_extension_script_get_language(ExtensionInstanceID inst);

extern void gd_on_first_frame(void);
extern void gd_on_every_frame(void);
extern void gd_on_final_frame(void);
extern void gd_on_worker_thread_pool_task(TaskID task);
extern void gd_on_worker_thread_pool_group_task(TaskID task, uint32_t n);

void gd_array_get(Array a, int64_t i, void*);
void gd_array_set(Array a, int64_t i, uint64_t v1, uint64_t v2, uint64_t v3);

FunctionID gd_builtin_name(StringName utility, int64_t hash);
void gd_builtin_call(FunctionID utility, void* result, Shape shape, void* args);

String gd_variant_type_name(VariantType t);
void gd_variant_type_make(VariantType t, void* result, int64_t arg_count, void* args, void* err);
void gd_variant_type_call(VariantType t, StringName static_method_name, void* result, int64_t arg_count, void* args, void* err);
bool gd_variant_type_convertable(VariantType t, VariantType to, bool strict);
void gd_variant_type_setup_array(Array a, VariantType elem, StringName class_name, uint64_t v1, uint64_t v2, uint64_t v3);
void gd_variant_type_setup_dictionary(Dictionary d,
    VariantType key, StringName key_class_name, uint64_t k_script_1, uint64_t k_script_2, uint64_t k_script_3,
    VariantType val, StringName val_class_name, uint64_t v_script_1, uint64_t v_script_2, uint64_t v_script_3
);
void gd_variant_type_fetch_constant(VariantType t, StringName constant, void* result);
FunctionID gd_variant_type_unsafe_constructor(VariantType t, int64_t n);
FunctionID gd_variant_type_evaluator(VariantOperator op, VariantType a, VariantType b);
FunctionID gd_variant_type_setter(VariantType t, StringName property);
FunctionID gd_variant_type_getter(VariantType t, StringName property);
bool gd_variant_type_has_property(VariantType t, StringName property);
FunctionID gd_variant_type_builtin_method(VariantType t, StringName method, int64_t hash);
void gd_variant_type_unsafe_call(void* self, FunctionID fn, void* result, Shape shape, void* args);
void gd_variant_type_unsafe_make(FunctionID constructor, void* result, Shape shape, void* args);
void gd_variant_type_unsafe_free(VariantType t, Shape shape, void* args);

void gd_callable_create(CallableID fn, ObjectID object, void* result);
FunctionID gd_callable_lookup(uint64_t c1, uint64_t c2);

void gd_classdb_FileAccess_write(Object file, char* buf, int64_t len);
int64_t gd_classdb_FileAccess_read(Object file, char* buf, int64_t cap);
uintptr_t gd_classdb_Image_unsafe(Object img);
uint8_t gd_classdb_Image_access(Object img, int64_t offset);
void gd_classdb_WorkerThreadPool_add_task(Object pool, TaskID task, bool priority, String description);
void gd_classdb_WorkerThreadPool_add_group_task(Object pool, TaskID task, int32_t elements, int32_t arg, bool priority, String description);
int64_t gd_classdb_XMLParser_load(Object parser, char* buf, int64_t cap);

MethodList gd_method_list_make(int64_t method_count);
void gd_method_list_push(MethodList list,
    StringName name, FunctionID call, MethodFlags method_flags,
    PropertyList return_value_info, PropertyList arguments_info,
    int64_t count, void* default_arguments
);
void gd_method_list_free(MethodList list);

PropertyList gd_property_list_make(int64_t property_count);
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
    StringName class, StringName parent_class,
    ExtensionClassID id, bool is_virtual, bool is_abstract,
    bool is_exposed, bool is_runtime, String icon_path
);
void gd_classdb_register_methods(StringName class, MethodList methods);
void gd_classdb_register_constant(
    StringName class, StringName enum_name, StringName constant_name,
    int64_t value, bool bitfield
);
void gd_classdb_register_property(StringName class, PropertyList property, StringName setter, StringName getter);
void gd_classdb_register_property_indexed(
    StringName class, PropertyList property,
    StringName setter, StringName getter, int64_t index
);
void gd_classdb_register_property_group(StringName class, String group, String prefix);
void gd_classdb_register_property_sub_group(StringName class, String subgroup, String prefix);
void gd_classdb_register_signal(StringName class, StringName signal, PropertyList args);
void gd_classdb_register_removal(StringName class);

void gd_packed_dictionary_access(Dictionary d, uint64_t i1, uint64_t i2, uint64_t i3, void* result);
void gd_packed_dictionary_modify(Dictionary d, uint64_t i1, uint64_t i2, uint64_t i3, uint64_t v1, uint64_t v2, uint64_t v3);

void gd_editor_add_documentation(const char* xml, int64_t len);
void gd_editor_add_plugin(StringName class);
void gd_editor_end_plugin(StringName class);

void gd_iterator_make(uint64_t v1, uint64_t v2, uint64_t v3, void* result_iter, void* err);
bool gd_iterator_next(uint64_t v1, uint64_t v2, uint64_t v3, void* iter, void* err);
void gd_iterator_load(uint64_t v1, uint64_t v2, uint64_t v3, uint64_t i1, uint64_t i2, uint64_t i3, void* result, void* err);

String gd_library_location();

void gd_log_error(
    const char* text, int64_t text_len,
    const char* code, int64_t code_len,
    const char* func, int64_t func_len,
    const char* file, int64_t file_len,
    int32_t line, bool notify_editor
);
void gd_log_warning(
    const char* text, int64_t text_len,
    const char* code, int64_t code_len,
    const char* func, int64_t func_len,
    const char* file, int64_t file_len,
    int32_t line, bool notify_editor
);

uintptr_t gd_memory_malloc(int64_t size);
int64_t gd_memory_sizeof(StringName struct_name);
uintptr_t gd_memory_resize(uintptr_t addr, int64_t size);
void gd_memory_clear(uintptr_t addr, int64_t size);
void gd_memory_free(uintptr_t addr);
void gd_memory_edit_byte(uintptr_t addr, uint8_t v);
void gd_memory_edit_u16(uintptr_t addr, uint16_t v);
void gd_memory_edit_u32(uintptr_t addr, uint32_t v);
void gd_memory_edit_u64(uintptr_t addr, uint64_t v);
void gd_memory_edit_128(uintptr_t addr, uint64_t v1, uint64_t v2);
void gd_memory_edit_256(uintptr_t addr, uint64_t v1, uint64_t v2, uint64_t v3, uint64_t v4);
void gd_memory_edit_512(uintptr_t addr, uint64_t v1, uint64_t v2, uint64_t v3, uint64_t v4, uint64_t v5, uint64_t v6, uint64_t v7, uint64_t v8);
uint8_t gd_memory_load_byte(uintptr_t addr);
uint16_t gd_memory_load_u16(uintptr_t addr);
uint32_t gd_memory_load_u32(uintptr_t addr);

Object gd_object_make(StringName name);
void gd_object_call(Object obj, MethodForClass method, void* result, int64_t arg_count, void* args, void* err);
StringName gd_object_name(Object obj);
ObjectType gd_object_type(StringName name);
Object gd_object_cast(Object obj, ObjectType to);
Object gd_object_lookup(ObjectID id);
Object gd_object_global(StringName name);
void gd_object_extension_setup(Object obj, StringName name, ExtensionInstanceID class);
ExtensionInstanceID gd_object_extension_fetch(Object obj);
void gd_object_extension_close(Object obj);
void gd_object_id(Object obj, void* id);
void gd_object_id_inside_variant(uint64_t v1, uint64_t v2, uint64_t v3, void* id);
MethodForClass gd_object_method_lookup(StringName name, StringName method, int64_t hash);
ScriptInstance gd_object_script_make(ExtensionInstanceID fn);
void gd_object_script_call(Object obj, StringName name, void* result, int64_t arg_count, void* args, void* err);
void gd_object_script_setup(Object obj, ScriptInstance script);
ScriptInstance gd_object_script_fetch(Object obj, Object language);
bool gd_object_script_defines_method(Object obj, StringName method);
void gd_object_script_property_state_add(FunctionID fn, uintptr_t arg, StringName name, uint64_t state_1, uint64_t state_2, uint64_t state_3);
ScriptInstance gd_object_script_placeholder_create(Object language, Object script, Object owner);
void gd_object_script_placeholder_update(ScriptInstance script, Array array, Dictionary dict);
void gd_object_unsafe_call(Object obj, MethodForClass fn, void* result, Shape shape, void* args);
void gd_object_unsafe_free(Object obj);

uintptr_t gd_packed_byte_array_unsafe(uint64_t p1, uint64_t p2);
uint8_t gd_packed_byte_array_access(uint64_t p1, uint64_t p2, int64_t idx);
uintptr_t gd_packed_color_array_unsafe(uint64_t p1, uint64_t p2);
void gd_packed_color_array_access(uint64_t p1, uint64_t p2, int64_t idx, void* result);
uintptr_t gd_packed_float32_array_unsafe(uint64_t p1, uint64_t p2);
float gd_packed_float32_array_access(uint64_t p1, uint64_t p2, int64_t idx);
uintptr_t gd_packed_float64_array_unsafe(uint64_t p1, uint64_t p2);
double gd_packed_float64_array_access(uint64_t p1, uint64_t p2, int64_t idx);
uintptr_t gd_packed_int32_array_unsafe(uint64_t p1, uint64_t p2);
int32_t gd_packed_int32_array_access(uint64_t p1, uint64_t p2, int64_t idx);
uintptr_t gd_packed_int64_array_unsafe(uint64_t p1, uint64_t p2);
void gd_packed_int64_array_access(uint64_t p1, uint64_t p2, int64_t idx, void* result);
uintptr_t gd_packed_string_array_unsafe(uint64_t p1, uint64_t p2);
String gd_packed_string_array_access(uint64_t p1, uint64_t p2, int64_t idx);
uintptr_t gd_packed_vector2_array_unsafe(uint64_t p1, uint64_t p2);
void gd_packed_vector2_array_access(uint64_t p1, uint64_t p2, int64_t idx, void* result);
uintptr_t gd_packed_vector3_array_unsafe(uint64_t p1, uint64_t p2);
void gd_packed_vector3_array_access(uint64_t p1, uint64_t p2, int64_t idx, void* result);
uintptr_t gd_packed_vector4_array_unsafe(uint64_t p1, uint64_t p2);
void gd_packed_vector4_array_access(uint64_t p1, uint64_t p2, int64_t idx, void* result);

Object gd_ref_get_object(RefCounted ref);
void gd_ref_set_object(RefCounted ref, Object obj);

int32_t gd_string_access(String s, int64_t idx);
String gd_string_resize(String s, int64_t size);
uintptr_t gd_string_unsafe(String s);
String gd_string_append(String s, String other);
String gd_string_append_rune(String s, int32_t ch);
String gd_string_decode_latin1(const char* s, int64_t len);
String gd_string_decode_utf8(const char* s, int64_t len);
String gd_string_decode_utf16(const char* s, int64_t len, bool little_endian);
String gd_string_decode_utf32(const char* s, int64_t len);
String gd_string_decode_wide(const char* s, int64_t len);
int64_t gd_string_encode_latin1(String s, char* buf, int64_t cap);
int64_t gd_string_encode_utf8(String s, char* buf, int64_t cap);
int64_t gd_string_encode_utf16(String s, char* buf, int64_t cap);
int64_t gd_string_encode_utf32(String s, char* buf, int64_t cap);
int64_t gd_string_encode_wide(String s, char* buf, int64_t cap);
StringName gd_string_intern_latin1(const char* s, int64_t len);
StringName gd_string_intern_utf8(const char* s, int64_t len);

bool gd_thread_is_main();

void gd_variant_zero(void* result);
void gd_variant_copy(uint64_t v1, uint64_t v2, uint64_t v3, void* result);
void gd_variant_call(uint64_t v1, uint64_t v2, uint64_t v3, StringName method, void* result, int64_t arg_count, void* args, void* err);
bool gd_variant_eval(VariantOperator op, uint64_t a1, uint64_t a2, uint64_t a3, uint64_t b1, uint64_t b2, uint64_t b3, void* result);
void gd_variant_hash(uint64_t v1, uint64_t v2, uint64_t v3, void* hash);
bool gd_variant_bool(uint64_t v1, uint64_t v2, uint64_t v3);
String gd_variant_text(uint64_t v1, uint64_t v2, uint64_t v3);
VariantType gd_variant_type(uint64_t v1, uint64_t v2, uint64_t v3);
void gd_variant_deep_copy(uint64_t v1, uint64_t v2, uint64_t v3, void* result);
void gd_variant_deep_hash(uint64_t v1, uint64_t v2, uint64_t v3, int64_t recursion_count, void* hash);
bool gd_variant_get_index(uint64_t v1, uint64_t v2, uint64_t v3, uint64_t key1, uint64_t key2, uint64_t key3, void* result);
bool gd_variant_get_array(uint64_t v1, uint64_t v2, uint64_t v3, int64_t idx, void* result, void* err);
bool gd_variant_get_field(uint64_t v1, uint64_t v2, uint64_t v3, StringName field, void* result);
bool gd_variant_has_index(uint64_t v1, uint64_t v2, uint64_t v3, uint64_t idx1, uint64_t idx2, uint64_t idx3);
bool gd_variant_has_method(uint64_t v1, uint64_t v2, uint64_t v3, StringName method);
bool gd_variant_set_index(uint64_t v1, uint64_t v2, uint64_t v3, uint64_t key1, uint64_t key2, uint64_t key3, uint64_t val1, uint64_t val2, uint64_t val3);
bool gd_variant_set_array(uint64_t v1, uint64_t v2, uint64_t v3, int64_t idx, uint64_t val1, uint64_t val2, uint64_t val3, void* err);
bool gd_variant_set_field(uint64_t v1, uint64_t v2, uint64_t v3, StringName field, uint64_t val1, uint64_t val2, uint64_t val3);
void gd_variant_unsafe_call(FunctionID fn, void* result, Shape shape, void* args);
void gd_variant_unsafe_eval(FunctionID fn, void* result, Shape shape, void* args);
void gd_variant_unsafe_free(uint64_t v1, uint64_t v2, uint64_t v3);
void gd_variant_unsafe_make_native(VariantType vtype, uint64_t v1, uint64_t v2, uint64_t v3, Shape shape, void* result);
void gd_variant_unsafe_from_native(VariantType vtype, void* result, Shape shape, void* args);
uintptr_t gd_variant_unsafe_internal_pointer(VariantType vtype, uint64_t v1, uint64_t v2, uint64_t v3);
void gd_variant_unsafe_get_field(FunctionID getter, void* result, Shape shape, void* args);
void gd_variant_unsafe_get_array(VariantType vtype, int64_t idx, void* result, Shape shape, void* args);
void gd_variant_unsafe_get_index(VariantType vtype, void* result, Shape shape, void* args);
void gd_variant_unsafe_set_field(FunctionID setter, Shape shape, void* args);
void gd_variant_unsafe_set_array(VariantType vtype, int64_t idx, Shape shape, void* args);
void gd_variant_unsafe_set_index(VariantType vtype, Shape shape, void* args);

uint32_t gd_version_major();
uint32_t gd_version_minor();
uint32_t gd_version_patch();
uint32_t gd_version_hex();
String gd_version_status();
String gd_version_build();
String gd_version_hash();
void gd_version_timestamp(void*);
String gd_version_string();
