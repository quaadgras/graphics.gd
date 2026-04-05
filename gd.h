#include <stdint.h>
#include <stdbool.h>

#ifndef __EMSCRIPTEN__
    #define INT64(n) int64_t n
    #define SHAPE uint64_t
    #define ANY void*
    #define UINT uint64_t
    #define OBJECT_ID(n) uint64_t n
    #define UINT64(n) uint64_t n
    #define BUFFER char*
    #define STRING const char*
    #define RESULT_POINTER
    #define RETURNS(t) t
#else
    #define ANY uint32_t
    #define UINT uint32_t
    #define UINT64(n) uint32_t n##_1, uint32_t n##_2
    #define INT64(n) uint32_t n##_1, uint32_t n##_2
    #define SHAPE(n) uint32_t n##_1, uint32_t n##_2
    #define OBJECT_ID(n) uint32_t n##_1, uint32_t n##_2
    #define BUFFER uint32_t
    #define STRING std::string
    #define RESULT_POINTER , uint32_t result
    #define RETURNS(t) void
    extern "C" {
#endif

#define VARIANT_ARG(n) uint64_t n##_1, uint64_t n##_2, uint64_t n##_3
#define PACKED_ARRAY_ARG(n) uintptr_t n##_1, uintptr_t n##_2

typedef struct {
    uint64_t tag;
    uint64_t payload[2];
} Variant;

typedef const Variant* const* VariadicVariants;

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

typedef uint8_t StringEncoding;
static const StringEncoding Latin1  = 0;
static const StringEncoding UTF8    = 1;
static const StringEncoding UTF16LE = 2;
static const StringEncoding UTF16BE = 3;
static const StringEncoding UTF32   = 4;
static const StringEncoding Wide    = 5;

typedef uint32_t InitializationLevel;
typedef uint32_t VariantType;
typedef uint32_t VariantOperator;
typedef uint32_t MethodFlags;
typedef uint32_t ArgumentMetadata;

typedef int64_t Int;
#ifndef __EMSCRIPTEN__
typedef void* UnsafePointer;
#else
typedef uintptr_t UnsafePointer;
#endif

typedef struct PackedArray {
    uintptr_t opaque[2];
} PackedArray;

typedef uintptr_t Object;
typedef uint64_t  ObjectID;
typedef uint64_t  RID;
typedef uintptr_t ObjectType;
typedef uintptr_t RefCounted;
typedef uintptr_t String;
typedef uintptr_t StringName;
typedef uintptr_t Array;
typedef uintptr_t Dictionary;
typedef uintptr_t MethodForClass;
typedef uintptr_t ScriptInstance;
typedef uintptr_t TaskID;
typedef uintptr_t FunctionID;
typedef uintptr_t ExtensionClassID;
typedef uintptr_t ExtensionInstanceID;
typedef uintptr_t ExtensionBindingID;
typedef uintptr_t PropertyList;
typedef uintptr_t MethodList;

// Callable is a generic first-class-function represented by the engine. The underlying function may
// be defined by the engine, a script, or an extension. Extensions use [CallableID] to identify any
// callables they've created. These are extension-specific opaque identifiers (sometimes pointers).
typedef struct {
    uint64_t opaque[2];
} Callable;

#define CALLABLE_ARG(n) uint64_t n##_1, uint64_t n##_2
#define CALLABLE_ARG_GET(n) (Callable){{n##_1, n##_2}}
#define CALLABLE_ARG_PUT(n) n.opaque[0], n.opaque[1]

// CallableID is an opaque pointer-sized extension-specific identifier for a [Callable]. Either
// store a pinned pointer to your callable here, or a handle that you can lookup from a table.
typedef uintptr_t CallableID;

void       gd_callable_create(CallableID id, ObjectID owner, Callable* result);
CallableID gd_callable_lookup(CALLABLE_ARG(c));

// CALLBACK declares a callback function. On Emscripten it expands to a static
// function pointer plus a setter (gd_set_<name>) so that engine.js can register
// gdextension WASM exports by name into Emscripten's indirect function table.
// The engine then calls the extension through call_indirect with no JS intermediary.
// Everywhere else it's a plain extern declaration resolved by the linker.
#ifdef __EMSCRIPTEN__
    #define CALLBACK(ret, name, params)                                          \
        static ret (*name) params;                                               \
        EMSCRIPTEN_KEEPALIVE void gd_set_callback_##name(uintptr_t fn) {         \
            *(uintptr_t*)&name = fn;                                             \
        }
#else
    #define CALLBACK(ret, name, params) extern ret name params
#endif

CALLBACK(void,     gd_on_callable_called, (CallableID c, Variant* ret, Int argc, VariadicVariants args, CallError* err));
CALLBACK(bool,     gd_on_callable_verify, (CallableID c));
CALLBACK(void,     gd_on_callable_delete, (CallableID c));
CALLBACK(uint32_t, gd_on_callable_hashed, (CallableID c));
CALLBACK(Int,      gd_on_callable_sorted, (CallableID a, CallableID b));
CALLBACK(String,   gd_on_callable_string, (CallableID c));
CALLBACK(Int,      gd_on_callable_length, (CallableID c));

extern void gd_on_editor_class_in_use_detection(PACKED_ARRAY_ARG(a), PackedStringArray* result);

extern void gd_on_engine_init(InitializationLevel level);
extern void gd_on_engine_exit(InitializationLevel level);

extern ExtensionBindingID gd_on_extension_binding_created(ExtensionInstanceID inst);
extern void               gd_on_extension_binding_removed(ExtensionInstanceID inst, ExtensionBindingID p1);
extern bool               gd_on_extension_binding_reference(ExtensionInstanceID inst, bool p1);

extern Object     gd_on_extension_class_create(ExtensionClassID class_name, bool notify_postinitialize);
extern FunctionID gd_on_extension_class_method(ExtensionClassID class_name, StringName method, uint32_t hash);
extern FunctionID gd_on_extension_class_caller(ExtensionClassID class_name, StringName method, uint32_t hash);

extern bool   gd_on_extension_instance_set(ExtensionInstanceID inst, StringName property, VARIANT_ARG(val));
extern bool   gd_on_extension_instance_get(ExtensionInstanceID inst, StringName property, Variant* p2);
extern bool   gd_on_extension_instance_property_has_default(ExtensionInstanceID inst, StringName property);
extern bool   gd_on_extension_instance_property_get_default(ExtensionInstanceID inst, StringName property, Variant* result);
extern bool   gd_on_extension_instance_property_validation(ExtensionInstanceID inst, StringName property);
extern void   gd_on_extension_instance_notification(ExtensionInstanceID inst, int32_t what, bool reverse);
extern String gd_on_extension_instance_stringify(ExtensionInstanceID inst);
extern bool   gd_on_extension_instance_reference(ExtensionInstanceID inst, bool increment);
extern RID    gd_on_extension_instance_rid(ExtensionInstanceID inst);
extern void   gd_on_extension_instance_checked_call(ExtensionInstanceID inst, FunctionID fn, UnsafePointer result, UnsafePointer args);
extern void   gd_on_extension_instance_variant_call(ExtensionInstanceID inst, FunctionID fn, Variant* result, VariadicVariants args);
extern void   gd_on_extension_instance_dynamic_call(ExtensionInstanceID inst, FunctionID fn, Variant* result, Int count, VariadicVariants args, CallError* err);
extern void   gd_on_extension_instance_free(ExtensionInstanceID inst);
extern void   gd_on_extension_instance_called(ExtensionInstanceID inst, FunctionID fn, UnsafePointer result, UnsafePointer args);

extern PropertyList gd_on_extension_instance_property_list(ExtensionInstanceID inst);

extern bool       gd_on_extension_script_categorization(ExtensionInstanceID inst, PropertyList p1);
extern uint32_t   gd_on_extension_script_get_property_type(ExtensionInstanceID inst, CallError* err);
extern Object     gd_on_extension_script_get_owner(ExtensionInstanceID inst);
extern void       gd_on_extension_script_get_property_state(ExtensionInstanceID inst, FunctionID add, uintptr_t arg);
extern MethodList gd_on_extension_script_get_methods(ExtensionInstanceID inst);
extern bool       gd_on_extension_script_has_method(ExtensionInstanceID inst, uintptr_t p1);
extern Int        gd_on_extension_script_get_method_argument_count(ExtensionInstanceID inst, StringName property);
extern Object     gd_on_extension_script_get(ExtensionInstanceID inst);
extern bool       gd_on_extension_script_is_placeholder(ExtensionInstanceID inst);
extern Object     gd_on_extension_script_get_language(ExtensionInstanceID inst);

extern void gd_on_first_frame(void);
extern void gd_on_every_frame(void);
extern void gd_on_final_frame(void);
extern void gd_on_worker_thread_pool_task(TaskID task);
extern void gd_on_worker_thread_pool_group_task(TaskID task, uint32_t n);

void gd_array_get(Array a, Int i, Variant* result);
void gd_array_set(Array a, Int i, VARIANT_ARG(v));

FunctionID gd_builtin_name(StringName utility, int64_t hash);
void       gd_builtin_call(FunctionID fn, UnsafePointer result, uint64_t shape, UnsafePointer args);

String gd_variant_type_name(VariantType t);
void   gd_variant_type_make(VariantType t, Variant* result, Int arg_count, Variant args[], CallError* err);
void   gd_variant_type_call(VariantType t, StringName static_method_name, Variant* result, Int arg_count, Variant args[], CallError* err);

bool gd_variant_type_convertable(VariantType t, VariantType to, bool strict);
bool gd_variant_type_has_property(VariantType t, StringName property);

void gd_variant_type_setup_array(Array a, VariantType elem, StringName class_name, VARIANT_ARG(v));
void gd_variant_type_setup_dictionary(Dictionary d,
    VariantType key, StringName key_class_name, VARIANT_ARG(key_script),
    VariantType val, StringName val_class_name, VARIANT_ARG(val_script)
);

// gd_variant_type_fetch_constant reads the value of a named constant for the given variant type
// into the memory pointed to by 'result'.
void gd_variant_type_fetch_constant(VariantType t, StringName constant, UnsafePointer result);

FunctionID gd_variant_type_unsafe_constructor(VariantType t, Int n);
FunctionID gd_variant_type_evaluator(VariantOperator op, VariantType a, VariantType b);
FunctionID gd_variant_type_setter(VariantType t, StringName property);
FunctionID gd_variant_type_getter(VariantType t, StringName property);

// gd_variant_type_builtin_method returns a [FunctionID] for the named method on the given
// variant type, validated against the expected API 'hash'. The returned ID is used with
// [gd_variant_type_unsafe_call].
FunctionID gd_variant_type_builtin_method(VariantType t, StringName method, int64_t hash);

void gd_variant_type_unsafe_call(UnsafePointer self, FunctionID fn, UnsafePointer result, uint64_t shape, UnsafePointer args);
void gd_variant_type_unsafe_make(FunctionID constructor, UnsafePointer result, uint64_t shape, UnsafePointer args);
void gd_variant_type_unsafe_free(VariantType t, uint64_t shape, UnsafePointer args);

void gd_classdb_FileAccess_write(Object file, char* buf, Int len);
Int  gd_classdb_FileAccess_read(Object file, char* buf, Int cap);

UnsafePointer gd_classdb_Image_unsafe(Object img);
uint8_t       gd_classdb_Image_access(Object img, Int offset);

void gd_classdb_WorkerThreadPool_add_task(Object pool, TaskID task, bool priority, String description);
void gd_classdb_WorkerThreadPool_add_group_task(Object pool, TaskID task, int32_t elements, int32_t arg, bool priority, String description);

Int gd_classdb_XMLParser_load(Object parser, char* buf, Int cap);

MethodList gd_method_list_make(Int method_count);
void       gd_method_list_free(MethodList list);

void gd_method_list_push(MethodList list,
    StringName name, FunctionID call, MethodFlags method_flags,
    PropertyList return_value_info, PropertyList arguments_info,
    Int count, ANY default_arguments
);

PropertyList gd_property_list_make(Int property_count);

void gd_property_list_push(PropertyList list,
    VariantType t, StringName name, StringName class_name,
    uint32_t hint, String hint_string, uint32_t usage, ArgumentMetadata meta
);

void        gd_property_list_free(PropertyList list);
VariantType gd_property_info_type(PropertyList info);
StringName  gd_property_info_name(PropertyList info);
uint32_t    gd_property_info_hint(PropertyList info);
String      gd_property_info_hint_string(PropertyList info);
StringName  gd_property_info_class_name(PropertyList info);
uint32_t    gd_property_info_usage(PropertyList info);

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

void gd_packed_dictionary_access(Dictionary d, VARIANT_ARG(key), Variant* result);
void gd_packed_dictionary_modify(Dictionary d, VARIANT_ARG(key), VARIANT_ARG(val));

void gd_editor_add_documentation(const char* xml, uint32_t len);
void gd_editor_add_plugin(StringName class_name);
void gd_editor_end_plugin(StringName class_name);

void gd_iterator_make(VARIANT_ARG(v), UnsafePointer result_iter, UnsafePointer err);
bool gd_iterator_next(VARIANT_ARG(v), UnsafePointer iter, UnsafePointer err);
void gd_iterator_load(VARIANT_ARG(v), VARIANT_ARG(i), UnsafePointer result, UnsafePointer err);

// gd_library_location returns the location of the extension's library.
String gd_library_location();

// LogLevel identifies the severity of a log message.
typedef uint32_t LogLevel;
static const LogLevel LogError   = 0;
static const LogLevel LogWarning = 1;

// gd_log prints a log message at the given severity level. 'text' is the human-readable message,
// 'code' is the error/warning code, 'func'/'file'/'line' identify the source location.
void gd_log(LogLevel level,
    const char* text, uint32_t text_len,
    const char* code, uint32_t code_len,
    const char* func, uint32_t func_len,
    const char* file, uint32_t file_len,
    int32_t line, bool notify_editor
);

UnsafePointer gd_memory_malloc(Int size);                        // allocates memory of the given size.
Int           gd_memory_sizeof(StringName struct_name);          // size of the given engine struct in bytes.
UnsafePointer gd_memory_resize(UnsafePointer addr, Int size);    // resizes the memory, preserving contents.
void          gd_memory_clear(UnsafePointer addr, Int size);     // (sets memory to zero) from addr to addr + size.
void          gd_memory_free(UnsafePointer addr);                // frees the memory block at the given address.

void gd_memory_edit_byte(UnsafePointer addr, uint8_t v);
void gd_memory_edit_u16(UnsafePointer addr, uint16_t v);
void gd_memory_edit_u32(UnsafePointer addr, uint32_t v);
void gd_memory_edit_u64(UnsafePointer addr, uint64_t v);
void gd_memory_edit_128(UnsafePointer addr, uint64_t a, uint64_t b);
void gd_memory_edit_256(UnsafePointer addr, uint64_t a, uint64_t b, uint64_t c, uint64_t d);
void gd_memory_edit_512(UnsafePointer addr, uint64_t a, uint64_t b, uint64_t c, uint64_t d, uint64_t e, uint64_t f, uint64_t g, uint64_t h);

uint8_t  gd_memory_load_byte(UnsafePointer addr);
uint16_t gd_memory_load_u16(UnsafePointer addr);
uint32_t gd_memory_load_u32(UnsafePointer addr);
uint64_t gd_memory_load_u64(UnsafePointer addr);

Object     gd_object_make(StringName name);           // constructs a new instance of the class.
StringName gd_object_name(Object obj);                // class name of the object.
ObjectType gd_object_type(StringName name);           // class tag of the object.
Object     gd_object_cast(Object obj, ObjectType to); // cast an object to a specific type.

Object gd_object_lookup(ObjectID id);     // fetch object with given ID (slow).
Object gd_object_global(StringName name); // return the named singleton.

void gd_object_call(Object obj, MethodForClass method, Variant* result, Int arg_count, Variant args[], CallError* err);
void gd_object_extension_setup(Object obj, StringName name, ExtensionInstanceID inst);

ExtensionInstanceID gd_object_extension_fetch(Object obj);  // lookup the associated extension instance.
void                gd_object_extension_close(Object obj);  // remove the extension binding from the object.

ObjectID       gd_object_id(Object obj);                                                        // get the instance ID of the given object.
ObjectID       gd_object_id_inside_variant(VARIANT_ARG(v));                                     // extract the object ID from a variant.
MethodForClass gd_object_method_lookup(StringName class_name, StringName method, int64_t hash); // lookup method for class.

ScriptInstance gd_object_script_make(ExtensionInstanceID fn);             // create a new script instance.
void           gd_object_script_setup(Object obj, ScriptInstance script); // attach a script instance.
ScriptInstance gd_object_script_fetch(Object obj, Object language);       // fetch associated script instance.

void gd_object_script_call(Object obj, StringName name, Variant* result, Int arg_count, Variant args[], CallError* err);
bool gd_object_script_defines_method(Object obj, StringName method);
void gd_object_script_property_state_add(FunctionID fn, uintptr_t arg, StringName name, VARIANT_ARG(state));

ScriptInstance gd_object_script_placeholder_create(Object language, Object script, Object owner);
void           gd_object_script_placeholder_update(ScriptInstance script, Array array, Dictionary dict);

void gd_object_shaped_call(Object obj, MethodForClass fn, UnsafePointer result, uint64_t shape, UnsafePointer args);
void gd_object_unsafe_free(Object obj);

UnsafePointer gd_packed_array_access(VariantType type, PACKED_ARRAY_ARG(pa), Int idx);
UnsafePointer gd_packed_array_modify(VariantType type, PACKED_ARRAY_ARG(pa), Int idx); // copy on write trigger.

Object gd_ref_get_object(RefCounted ref);
void   gd_ref_set_object(RefCounted ref, Object obj);

int32_t       gd_string_access(String s, Int idx);  // character at the given index.
String        gd_string_resize(String s, Int size); // resizes to the given size.
UnsafePointer gd_string_unsafe(String s);           // unsafe buffer access.

String gd_string_append(String s, String other);     // append string, returns updated 's'.
String gd_string_append_rune(String s, int32_t ch);  // append rune, returns updated 's'.

String     gd_string_decode(StringEncoding enc, const char* s, Int len);
Int        gd_string_encode(StringEncoding enc, String s, char* buf, Int cap);
StringName gd_string_intern(StringEncoding enc, const char* s, Int len);

void        gd_variant_zero(Variant* result); // initialize to nil/zero.
void        gd_variant_copy(VARIANT_ARG(v), Variant* result);
Int         gd_variant_hash(VARIANT_ARG(v)); // sorting hash.
bool        gd_variant_bool(VARIANT_ARG(v)); // truthy value.
String      gd_variant_text(VARIANT_ARG(v)); // stringify
VariantType gd_variant_type(VARIANT_ARG(v)); // typeof

void gd_variant_call(VARIANT_ARG(v), StringName method, Variant* result, Int arg_count, Variant args[], CallError* err);
bool gd_variant_eval(VariantOperator op, VARIANT_ARG(a), VARIANT_ARG(b), Variant* result);

void gd_variant_deep_copy(VARIANT_ARG(v), Variant* result);
Int  gd_variant_deep_hash(VARIANT_ARG(v), Int recursion_count);

bool gd_variant_get_index(VARIANT_ARG(v), VARIANT_ARG(key), Variant* result);
bool gd_variant_get_array(VARIANT_ARG(v), Int idx, Variant* result, CallError* err);
bool gd_variant_get_field(VARIANT_ARG(v), StringName field, Variant* result);
bool gd_variant_has_index(VARIANT_ARG(v), VARIANT_ARG(idx));
bool gd_variant_has_method(VARIANT_ARG(v), StringName method);
bool gd_variant_set_index(VARIANT_ARG(v), VARIANT_ARG(key), VARIANT_ARG(val));
bool gd_variant_set_array(VARIANT_ARG(v), Int idx, VARIANT_ARG(val), UnsafePointer err);
bool gd_variant_set_field(VARIANT_ARG(v), StringName field, VARIANT_ARG(val));


void gd_variant_unsafe_call(FunctionID fn, UnsafePointer result, uint64_t shape, UnsafePointer args);
void gd_variant_unsafe_eval(FunctionID fn, UnsafePointer result, uint64_t shape, UnsafePointer args);
void gd_variant_unsafe_free(VARIANT_ARG(v));

void gd_variant_unsafe_make_native(VariantType vtype, VARIANT_ARG(v), uint64_t shape, UnsafePointer result);
void gd_variant_unsafe_from_native(VariantType vtype, Variant* result, uint64_t shape, UnsafePointer args);

uintptr_t gd_variant_unsafe_internal_pointer(VariantType vtype, VARIANT_ARG(v));

void gd_variant_unsafe_get_field(FunctionID getter, UnsafePointer result, uint64_t shape, UnsafePointer args);
void gd_variant_unsafe_get_array(VariantType vtype, Int idx, UnsafePointer result, uint64_t shape, UnsafePointer args);
void gd_variant_unsafe_get_index(VariantType vtype, UnsafePointer result, uint64_t shape, UnsafePointer args);
void gd_variant_unsafe_set_field(FunctionID setter, uint64_t shape, UnsafePointer args);
void gd_variant_unsafe_set_array(VariantType vtype, Int idx, uint64_t shape, UnsafePointer args);
void gd_variant_unsafe_set_index(VariantType vtype, uint64_t shape, UnsafePointer args);

uint32_t gd_version_major(void);     // major version
uint32_t gd_version_minor(void);     // minor release number
uint32_t gd_version_patch(void);     // patch number
uint32_t gd_version_hex(void);       // hexadecimal representation of the version number.
String   gd_version_status(void);    // ie. "stable", "beta", "alpha"
String   gd_version_build(void);     // build type
String   gd_version_hash(void);      // commit hash
uint64_t gd_version_timestamp(void); // version control and/or release timestamp
String   gd_version_string(void);    // complete version string

#ifdef __EMSCRIPTEN__
}
#endif
