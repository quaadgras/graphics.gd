#pragma once
#ifndef GD
#define GD

#include <stdint.h>
#include <stdbool.h>

#if defined(__cplusplus)
    extern "C" {
#endif

/// gd_shape describes the size and memory layout of up to 16 components. Used for passing memory from one address space to another.
typedef uint64_t gd_shape;

#define GD_BYTES_1x1  1  /// 1 byte, alignment 1
#define GD_BYTES_2x1  2  /// 2 bytes, alignment 2
#define GD_BYTES_4x1  3  /// 4 bytes, alignment 4
#define GD_BYTES_8x1  4  /// 8 bytes, alignment 8
#define GD_BYTES_4x2  5  /// 8 bytes, alignment 4
#define GD_BYTES_4x3  6  /// 12 bytes, alignment 4
#define GD_BYTES_8x2  7  /// 16 bytes, alignment 8
#define GD_BYTES_4x4  8  /// 16 bytes, alignment 4
#define GD_BYTES_8x3  9  /// 24 bytes, alignment 8

#if defined(PRECISION_DOUBLE)
    typedef double gd_float;

    #define GD_BYTES_8x4  10 /// 32 bytes, alignment 8
    #define GD_BYTES_8x5  11 /// 40 bytes, alignment 8
    #define GD_BYTES_8x6  12 /// 48 bytes, alignment 8
    #define GD_BYTES_8x9  13 /// 72 bytes, alignment 8
    #define GD_BYTES_8x12 14 /// 96 bytes, alignment 8
    #define GD_BYTES_8x16 15 /// 128 bytes, alignment 8
#else
    typedef float gd_float;

    #define GD_BYTES_4x6  10 /// 24 bytes, alignment 4
    #define GD_BYTES_4x9  11 /// 36 bytes, alignment 4
    #define GD_BYTES_4x12 12 /// 48 bytes, alignment 4
    #define GD_BYTES_4x16 13 /// 60 bytes, alignment 4
#endif

/// gd_extension_initialization_level represents the initialization level of an extension.
typedef uint32_t gd_initialization_level; /// gd_extension_initialization_level_core means the engine core is being started up.
static const gd_initialization_level GD_INIT_CORE = 0; /// GD_INIT_SERVERS means servers are being started up.
static const gd_initialization_level GD_INIT_SERVERS = 1; /// GD_INIT_SCENE means scenes are being started up.
static const gd_initialization_level GD_INIT_SCENE = 2; /// GD_INIT_EDITOR means the editor is being started up.
static const gd_initialization_level GD_INIT_EDITOR = 3;

/// gd_error represents an error that occurred during a gd call.
typedef struct { uint32_t error; int32_t argument; int32_t expected; } gd_error;

/// Variant can hold any engine value.
struct Variant     { uint64_t tag; uint64_t payload[2]; }; /// String of UTF-32 runes.
struct String      { uintptr_t ptr; }; /// Vector2 represents a 2D floating-pointer vector.
struct Vector2     { gd_float x; gd_float y; }; /// Vector2i represents a 2D integer vector.
struct Vector2i    { int32_t x; int32_t y; }; /// Rect2 represents a 2D floating-pointer rectangle.
struct Rect2       { struct Vector2 position; struct Vector2 size; }; /// Rect2i represents a 2D integer rectangle.
struct Rect2i      { struct Vector2i position; struct Vector2i size; }; /// Vector3 represents a 3D floating-pointer vector.
struct Vector3     { gd_float x; gd_float y; gd_float z; }; /// Vector3i represents a 3D integer vector.
struct Vector3i    { int32_t x; int32_t y; int32_t z; }; /// Transform2D represents a 2D transformation.
struct Transform2D { struct Vector2 x; struct Vector2 y; struct Vector2 origin; }; /// Vector4 represents a 4D floating-pointer vector.
struct Vector4     { gd_float x; gd_float y; gd_float z; gd_float w; }; /// Vector4i represents a 4D integer vector.
struct Vector4i    { int32_t x; int32_t y; int32_t z; int32_t w; }; /// Plane represents a 3D plane.
struct Plane       { struct Vector3 normal; gd_float d; }; /// Quaternion represents a 3D rotation.
struct Quaternion  { int32_t i; int32_t j; int32_t k; int32_t x; }; /// AABB represents a 3D axis-aligned bounding box.
struct AABB        { struct Vector3 position; struct Vector3 size; }; /// Basis represents a 3D basis (rotation matrix).
struct Basis       { struct Vector3 rows[3]; }; /// Transform3D represents a 3D transformation.
struct Transform3D { struct Basis basis; struct Vector3 origin; }; /// Projection represents a 3D projection matrix.
struct Projection  { struct Vector3 rows[4]; }; /// Color represents a 4D floating-point color.
struct Color       { gd_float r; gd_float g; gd_float b; gd_float a; }; /// Callable represents a callable object.
struct Callable    { uint64_t opaque[2]; }; /// StringName represents an identifier.
struct StringName  { uintptr_t opaque; }; /// NodePath represents a path to a node in the scene tree.
struct NodePath    { uintptr_t opaque; };

/// RID represents a resource identifier.
typedef uint64_t RID;

/// Object represents an instantiated class.
struct Object             { uintptr_t opaque; }; /// Signal represents a signal that can be emitted by an object.
struct Signal             { uint64_t opaque[2]; }; /// Dictionary represents a key-value dictionary.
struct Dictionary         { uintptr_t opaque; }; /// Array represents a dynamic array of variants.
struct Array              { uintptr_t opaque; }; /// PackedArray represents a generic packed array..
struct PackedArray        { uintptr_t opaque[2]; }; /// PackedByteArray represents a packed array of bytes.
struct PackedByteArray    { uintptr_t opaque[2]; }; /// PackedInt32Array represents a packed array of 32-bit integers.
struct PackedInt32Array   { uintptr_t opaque[2]; }; /// PackedInt64Array represents a packed array of 64-bit integers.
struct PackedInt64Array   { uintptr_t opaque[2]; }; /// PackedFloat32Array represents a packed array of 32-bit floating-point numbers.
struct PackedFloat32Array { uintptr_t opaque[2]; }; /// PackedFloat64Array represents a packed array of 64-bit floating-point numbers.
struct PackedFloat64Array { uintptr_t opaque[2]; }; /// PackedStringArray represents a packed array of strings.
struct PackedStringArray  { uintptr_t opaque[2]; }; /// PackedVector2Array represents a packed array of 2D vectors.
struct PackedVector2Array { uintptr_t opaque[2]; }; /// PackedVector3Array represents a packed array of 3D vectors.
struct PackedVector3Array { uintptr_t opaque[2]; }; /// PackedColorArray represents a packed array of colors.
struct PackedColorArray   { uintptr_t opaque[2]; }; /// PackedVector4Array represents a packed array of 4D vectors.
struct PackedVector4Array { uintptr_t opaque[2]; };

/// RefCounted instance.
struct RefCounted { uintptr_t opaque; }; /// Script instance.
struct Script     { uintptr_t opaque; };

//
// Version Information
//

/// gd_version returns the complete version string.
struct String gd_version(void); /// gd_version_major returns the major version number.
int64_t       gd_version_major(void); /// gd_version_minor returns the minor release number.
int64_t       gd_version_minor(void); /// gd_version_patch returns the patch number.
int64_t       gd_version_patch(void); /// gd_version_hexed returns the hexadecimal representation of the version number.
int64_t       gd_version_hexed(void); /// gd_version_state returns the state (e.g. "stable", "beta", "alpha").
struct String gd_version_state(void); /// gd_version_build returns the build type.
struct String gd_version_build(void); /// gd_version_stamp returns the commit hash
struct String gd_version_stamp(void); /// gd_version_nanos returns the timestamp in unix nanoseconds.
int64_t       gd_version_nanos(void);

//
// Unsafe Engine Memory
//

#if defined(__wasm__)
    /// gd_addr is a pointer/handle/reference to engine memory.
    typedef uintptr_t gd_addr;
#else
    /// gd_addr is a pointer/handle/reference to engine memory.
    typedef void* gd_addr;
#endif

/// gd_sizeof returns the size of the given engine type in bytes.
int64_t gd_sizeof(struct StringName type); /// gd_malloc allocates memory of the given size.
gd_addr gd_malloc(int64_t bytes, bool pad8); /// gd_resize resizes the memory, preserving contents.
gd_addr gd_resize(gd_addr addr, int64_t size, bool pad8); /// gd_memset memset equivalent.
void    gd_memset(gd_addr addr, uint8_t value, int64_t size);

/// gd_memory_bytes1 dereferences a 1-byte value from engine memory.
uint8_t  gd_memory_bytes1(const gd_addr addr); /// gd_memory_bytes2 dereferences a 2-byte value from engine memory.
uint16_t gd_memory_bytes2(const gd_addr addr); /// gd_memory_bytes4 dereferences a 4-byte value from engine memory.
uint32_t gd_memory_bytes4(const gd_addr addr); /// gd_memory_bytes8 dereferences an 8-byte value from engine memory.
uint64_t gd_memory_bytes8(const gd_addr addr);

/// gd_store_bytes1 stores a 1-byte value into engine memory.
void gd_store_bytes1(gd_addr addr, uint8_t v); /// gd_store_bytes2 stores a 2-byte value into engine memory.
void gd_store_bytes2(gd_addr addr, uint16_t v); /// gd_store_bytes4 stores a 4-byte value into engine memory.
void gd_store_bytes4(gd_addr addr, uint32_t v); /// gd_store_bytes8 stores an 8-byte value into engine memory.
void gd_store_bytes8(gd_addr addr, uint64_t v); /// gd_store_pair64 stores a pair of 64-bit values into engine memory.
void gd_store_pair64(gd_addr addr, uint64_t a, uint64_t b); /// gd_store_quad64 stores four 64-bit values into engine memory.
void gd_store_quad64(gd_addr addr, uint64_t a, uint64_t b, uint64_t c, uint64_t d); /// gd_store_octo64 stores eight 64-bit values into engine memory.
void gd_store_octo64(gd_addr addr, uint64_t a, uint64_t b, uint64_t c, uint64_t d, uint64_t e, uint64_t f, uint64_t g, uint64_t h);

/// gd_free frees the memory block at the given address.
void gd_free(gd_addr addr);

//
// Variants
//

/// VariantType, values defined in <variant.h>
typedef uint32_t VariantType; /// VariantOperator, values defined in <variant.h>
typedef uint32_t VariantOperator;

typedef const struct Variant* const* gd_unsafe_variants;

#define VARIANT_ARG(n) uint64_t n##_1, uint64_t n##_2, uint64_t n##_3

/// gd_variant_zero initializes the variant to nil/zero.
void          gd_variant_zero(struct Variant* result); /// gd_variant_copy copies the variant to the result, if deep is true, the variant is copied recursively.
void          gd_variant_copy(VARIANT_ARG(v), struct Variant* result, bool deep); /// gd_variant_hash returns the variant's hash, used for sorting, with the given recursion limit.
int64_t       gd_variant_hash(VARIANT_ARG(v), int64_t recursion_count); /// gd_variant_bool returns the variant's truthy value.
bool          gd_variant_bool(VARIANT_ARG(v)); /// gd_variant_text returns the variant's string representation.
struct String gd_variant_text(VARIANT_ARG(v)); /// gd_variant_type returns the variant's type.
VariantType   gd_variant_type(VARIANT_ARG(v)); /// gd_variant_free frees any underlying memory used by the variant.
void          gd_variant_free(VARIANT_ARG(v)); /// gd_variant_from initializes the variant from the given type, shape, and arguments.
void          gd_variant_from(VariantType vtype, struct Variant* result, const gd_addr args); /// gd_variant_make initializes the variant from the given type and arguments.
void          gd_variant_make(VariantType t, struct Variant* result, int64_t arg_count, struct Variant args[], gd_error* err); /// gd_variant_data returns a pointer to the variant's underlying data.
gd_addr       gd_variant_data(VariantType vtype, VARIANT_ARG(v)); /// gd_variant_call calls the variant's method with the given arguments.
void          gd_variant_call(VARIANT_ARG(v), struct StringName method, struct Variant* result, int64_t arg_count, struct Variant args[], gd_error* err); /// gd_variant_eval evaluates the variant with the given operator and arguments.
bool          gd_variant_eval(VariantOperator op, VARIANT_ARG(a), VARIANT_ARG(b), struct Variant* result); /// gd_variant_get_keyed gets the value of the variant at the given key.
bool          gd_variant_get_keyed(VARIANT_ARG(v), VARIANT_ARG(key), struct Variant* result); /// gd_variant_get_index gets the value of the variant at the given index.
bool          gd_variant_get_index(VARIANT_ARG(v), int64_t idx, struct Variant* result, gd_error* err); /// gd_variant_get_field gets the value of the variant at the given field.
bool          gd_variant_get_field(VARIANT_ARG(v), struct StringName field, struct Variant* result); /// gd_variant_has_key checks if the variant has a key at the given index.
bool          gd_variant_has_key(VARIANT_ARG(v), VARIANT_ARG(idx)); /// gd_variant_has_method checks if the variant has a method with the given name.
bool          gd_variant_has_method(VARIANT_ARG(v), struct StringName method); /// gd_variant_set_keyed sets the value of the variant at the given key.
bool          gd_variant_set_keyed(VARIANT_ARG(v), VARIANT_ARG(key), VARIANT_ARG(val)); /// gd_variant_set_index sets the value of the variant at the given index.
bool          gd_variant_set_index(VARIANT_ARG(v), int64_t idx, VARIANT_ARG(val), gd_error* err); /// gd_variant_set_field sets the value of the variant at the given field.
bool          gd_variant_set_field(VARIANT_ARG(v), struct StringName field, VARIANT_ARG(val));

//
// Builtin Types
//

typedef uintptr_t gd_constructor_id;
typedef uintptr_t gd_evaluator_id;
typedef uintptr_t gd_method_id;
typedef uintptr_t gd_getter_id;
typedef uintptr_t gd_setter_id;
typedef uintptr_t gd_caller_id;

/// gd_evaluator returns an evaluator for the given operator and variant types.
gd_evaluator_id      gd_evaluator(VariantOperator op, VariantType a, VariantType b); /// gd_setter returns a setter for the given variant type and property.
gd_setter_id         gd_setter(VariantType t, struct StringName property); /// gd_getter returns a getter for the given variant type and property.
gd_getter_id         gd_getter(VariantType t, struct StringName property); /// gd_constructor returns a constructor for the given variant type and number of arguments.
gd_constructor_id    gd_constructor(VariantType t, int64_t n); /// gd_caller returns a caller for the given variant type, method name, and hash.
gd_caller_id         gd_builtin_method(VariantType t, struct StringName method, int64_t hash);


/// gd_builtin_call calls the given function with the given arguments.
void gd_builtin_call(gd_addr self, gd_caller_id fn, gd_addr result, gd_shape shape, const gd_addr args); /// gd_builtin_make creates a new instance of the given constructor with the given arguments.
void gd_builtin_make(gd_constructor_id constructor, gd_addr result, gd_shape shape, const gd_addr args); /// gd_builtin_free frees the given value.
void gd_builtin_free(VariantType t, const gd_addr value); /// gd_builtin_from converts the given value to the given variant type.
void gd_builtin_from(VariantType vtype, VARIANT_ARG(v), gd_addr result); /// gd_builtin_eval evaluates the given evaluator with the given arguments.
void gd_builtin_eval(gd_evaluator_id op, gd_addr result, gd_shape shape, const gd_addr args); /// gd_builtin_get_field gets the value of the given field from the given object.
void gd_builtin_get_field(gd_getter_id getter, gd_addr result, gd_shape shape, const gd_addr args); /// gd_builtin_get_array gets the value of the given array element.
void gd_builtin_get_array(VariantType vtype, int64_t idx, gd_addr result, gd_shape shape, const gd_addr args); /// gd_builtin_get_keyed gets the value of the given key from the given object.
void gd_builtin_get_keyed(VariantType vtype, gd_addr result, gd_shape shape, const gd_addr args); /// gd_builtin_set_field sets the value of the given field on the given object.
void gd_builtin_set_field(gd_setter_id setter, gd_shape shape, const gd_addr args); /// gd_builtin_set_array sets the value of the given array element.
void gd_builtin_set_array(VariantType vtype, int64_t idx, gd_shape shape, const gd_addr args); /// gd_builtin_set_keyed sets the value of the given key on the given object.
void gd_builtin_set_keyed(VariantType vtype, gd_shape shape, const gd_addr args);

/// gd_variant_type_name returns the name of the given variant type.
struct String gd_variant_type_name(VariantType t); /// gd_variant_type_call calls the given static method on the given variant type.
void          gd_variant_type_call(VariantType t, struct StringName static_method_name, struct Variant* result, int64_t arg_count, struct Variant args[], gd_error* err); /// gd_variant_type_convertable returns true if the given variant type can be converted to the given variant type.
bool          gd_variant_type_convertable(VariantType t, VariantType to, bool strict); /// gd_variant_type_has_property returns true if the given variant type has the given property.
bool          gd_variant_type_has_property(VariantType t, struct StringName property); /// gd_variant_type_setup_array sets up the given array with the given element type and class name.
void          gd_variant_type_setup_array(struct Array a, VariantType elem, struct StringName class_name, VARIANT_ARG(v)); /// gd_variant_type_setup_dictionary sets up the given dictionary with the given key and value types and class names.
void          gd_variant_type_setup_dictionary(struct Dictionary d,
    VariantType key, struct StringName key_class_name, VARIANT_ARG(key_script),
    VariantType val, struct StringName val_class_name, VARIANT_ARG(val_script)
); /// gd_variant_type_constant returns the value of the given constant for the given variant type.
void          gd_variant_type_constant(VariantType t, struct StringName constant, struct Variant* result);

//
// Packed Arrays
//

#define PACKED_ARRAY_ARG(n) uintptr_t n##_1, uintptr_t n##_2

// gd_packed_array_access returns a pointer to the element at the given index in the packed array.
gd_addr gd_packed_array_access(VariantType type, PACKED_ARRAY_ARG(pa), int64_t idx); // gd_packed_array_modify returns a pointer to the element at the given index in the packed array, triggering a copy-on-write if necessary.
gd_addr gd_packed_array_modify(VariantType type, PACKED_ARRAY_ARG(pa), int64_t idx);

//
// Extension Callbacks
//

/// gd_extension_callable_t is an opaque pointer-sized identifier for a [Callable] (that is extension-specific). Either
/// store a pinned pointer to your callable here, or a handle that you can lookup from a table.
typedef uintptr_t gd_extension_callable_t;
typedef uintptr_t gd_extension_task_id;
typedef uintptr_t gd_extension_class_id;
typedef uintptr_t gd_extension_object_id;
typedef uintptr_t gd_extension_method_id;
typedef uintptr_t gd_extension_script_id;
typedef uintptr_t gd_extension_binding_id;

typedef uintptr_t gd_property_list_t;
typedef uintptr_t gd_method_list_t;
typedef uintptr_t gd_property_iterator_t;

/// gd_extension is the complete set of callbacks that a language extension provides.
/// Function pointer signatures are ABI-compatible with GDExtension callback types,
/// so gd.c can cast them directly without wrapper indirection.
typedef struct {
    void              (*on_engine_init)(void* userdata, uint32_t level);
    void              (*on_engine_exit)(void* userdata, uint32_t level);
    void              (*on_yield)(void);
    void              (*on_first_frame)(void);
    void              (*on_every_frame)(void);
    void              (*on_final_frame)(void);
    void*             (*on_extension_class_create)(void* class_userdata, uint8_t notify);
    void*             (*on_extension_class_method)(void* class_userdata, const void* method, uint32_t hash);
    void*             (*on_extension_class_caller)(void* class_userdata, const void* method, uint32_t hash);
    void              (*on_worker_thread_pool_task)(void* userdata);
    void              (*on_worker_thread_pool_group_task)(void* userdata, uint32_t n);
    void              (*on_editor_class_in_use_detection)(void* packed_string_array);
    void              (*on_extension_script_property_iter)(void* inst, void* add_func, void* userdata);
    uint8_t           (*on_extension_script_categorization)(void* inst, void* property_info);
    uint32_t          (*on_extension_script_get_property_type)(void* inst, const void* prop, uint8_t* ok);
    void*             (*on_extension_script_get_owner)(void* inst);
    const void*       (*on_extension_script_get_methods)(void* inst, uint32_t* r_count);
    uint8_t           (*on_extension_script_has_method)(void* inst, const void* method);
    int64_t           (*on_extension_script_get_method_argument_count)(void* inst, const void* m, uint8_t* ok);
    void*             (*on_extension_script_get)(void* inst);
    uint8_t           (*on_extension_script_is_placeholder)(void* inst);
    void*             (*on_extension_script_get_language)(void* inst);
    void              (*on_callable_called)(void* c, const void* const* args, int64_t argc, void* ret, void* err);
    uint8_t           (*on_callable_verify)(void* c);
    void              (*on_callable_delete)(void* c);
    uint32_t          (*on_callable_hashed)(void* c);
    uint8_t           (*on_callable_equal)(void* a, void* b);
    uint8_t           (*on_callable_less_than)(void* a, void* b);
    void              (*on_callable_string)(void* c, uint8_t* r_is_valid, void* r_out);
    int64_t           (*on_callable_length)(void* c, uint8_t* r_is_valid);
    uint8_t           (*on_extension_instance_set)(void* inst, const void* property, const void* value);
    uint8_t           (*on_extension_instance_get)(void* inst, const void* property, void* result);
    uint8_t           (*on_extension_instance_property_has_default)(void* inst, const void* property);
    uint8_t           (*on_extension_instance_property_get_default)(void* inst, const void* prop, void* r);
    uint8_t           (*on_extension_instance_property_validation)(void* inst, void* property_info);
    void              (*on_extension_instance_notification)(void* inst, int32_t what, uint8_t reverse);
    void              (*on_extension_instance_stringify)(void* inst, uint8_t* r_is_valid, void* r_out);
    void              (*on_extension_instance_reference)(void* inst);
    void              (*on_extension_instance_unreference)(void* inst);
    uint64_t          (*on_extension_instance_rid)(void* inst);
    void              (*on_extension_instance_checked_call)(void* fn, void* inst, const void* const* args, void* ret);
    void              (*on_extension_instance_dynamic_call)(void* fn, void* inst, const void* const* args, int64_t argc, void* ret, void* err);
    void              (*on_extension_instance_free)(void* class_userdata, void* inst);      // GDExtensionClassFreeInstance
    void              (*on_extension_script_free)(void* inst);                                // GDExtensionScriptInstanceFree
    void              (*on_extension_instance_called)(void* inst, const void* name, void* fn, const void* const* args, void* ret);
    const void*       (*on_extension_instance_property_list)(void* inst, uint32_t* r_count);           // GDExtensionClassGetPropertyList / GDExtensionScriptInstanceGetPropertyList
    void              (*on_extension_script_call)(void* inst, const void* method, const void* const* args, int64_t argc, void* ret, void* err); // GDExtensionScriptInstanceCall
    uint8_t           (*on_extension_script_unreference)(void* inst);                                  // GDExtensionScriptInstanceRefCountDecremented
    // Instance bindings
    void*             (*on_extension_binding_created)(void* token, void* inst);
    void              (*on_extension_binding_removed)(void* token, void* inst, void* binding);
    uint8_t           (*on_extension_binding_reference)(void* token, void* binding, uint8_t reference);
} gd_extension;


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

#ifdef __EMSCRIPTEN__
/// gd_on_yield is called occasionally.
CALLBACK(void, gd_on_yield, (void));
#else
/// gd_on_yield is called occasionally.
static inline void gd_on_yield(void) {}
#endif

///
/// Arrays
///

/// gd_array_get_index returns the value at the given index in the array.
void gd_array_get_index(struct Array a, int64_t i, struct Variant* result); /// gd_array_set_index sets the value at the given index in the array.
void gd_array_set_index(struct Array a, int64_t i, VARIANT_ARG(v));

///
/// GlobalScope Functions
///

typedef uintptr_t gd_function_t;

/// gd_function returns a function pointer for the given utility function.
gd_function_t gd_function(struct StringName utility, int64_t hash); /// gd_call calls the given function with the given arguments.
void          gd_call(gd_function_t fn, gd_addr result, gd_shape shape, gd_addr args);

//
// ClassDB
//

/// MethodFlags, values defined in <classdb.h>
typedef uint32_t MethodFlags; /// ArgumentMetadata, values defined in <classdb.h>
typedef uint32_t ArgumentMetadata;

/// gd_method_list_make returns a new method list with the given number of methods.
gd_method_list_t gd_method_list_make(gd_extension* extension, int64_t method_count); /// gd_method_list_free frees the given method list.
void             gd_method_list_free(gd_method_list_t list); /// gd_method_list_push pushes a method onto the given method list.
void             gd_method_list_push(gd_method_list_t list,
    struct StringName name, gd_extension_method_id call, MethodFlags method_flags,
    gd_property_list_t return_value_info, gd_property_list_t arguments_info,
    int64_t count, gd_addr default_arguments
);

/// gd_property_list_make returns a new property list with the given number of properties.
gd_property_list_t gd_property_list_make(int64_t property_count); /// gd_property_list_push pushes a property onto the given property list.
void               gd_property_list_push(gd_property_list_t list,
    VariantType t, struct StringName name, struct StringName class_name,
    uint32_t hint, struct String hint_string, uint32_t usage, ArgumentMetadata meta
); /// gd_property_list_free frees the given property list.
void               gd_property_list_free(gd_property_list_t list);


/// gd_classdb_register registers a new class with the given name, parent class, and ID.
void gd_classdb_register(gd_extension* extension,
    struct StringName class_name, struct StringName parent_class,
    gd_extension_class_id id, bool is_virtual, bool is_abstract,
    bool is_exposed, bool is_runtime, struct String icon_path
); /// gd_classdb_register_methods registers methods for the given class.
void gd_classdb_register_methods(struct StringName class_name, gd_method_list_t methods); /// gd_classdb_register_constant registers a constant for the given class.
void gd_classdb_register_constant(
    struct StringName class_name, struct StringName enum_name, struct StringName constant_name,
    int64_t value, bool bitfield
); /// gd_classdb_register_property registers a property for the given class.
void gd_classdb_register_property(struct StringName class_name, gd_property_list_t property, struct StringName setter, struct StringName getter); /// gd_classdb_register_property_indexed registers an indexed property for the given class.
void gd_classdb_register_property_indexed(
    struct StringName class_name, gd_property_list_t property,
    struct StringName setter, struct StringName getter, int64_t index
); /// gd_classdb_register_property_group registers a property group for the given class.
void gd_classdb_register_property_group(struct StringName class_name, struct String group, struct String prefix); /// gd_classdb_register_property_sub_group registers a property sub-group for the given class.
void gd_classdb_register_property_sub_group(struct StringName class_name, struct String subgroup, struct String prefix); /// gd_classdb_register_signal registers a signal for the given class.
void gd_classdb_register_signal(struct StringName class_name, struct StringName signal, gd_property_list_t args); /// gd_classdb_register_removal unregisters a class from the class database.
void gd_classdb_register_removal(struct StringName class_name);

/// gd_classdb_FileAccess_write writes data to a file.
void     gd_classdb_FileAccess_write(struct Object file, char* buf, int64_t len); /// gd_classdb_FileAccess_read reads data from a file.
int64_t  gd_classdb_FileAccess_read(struct Object file, char* buf, int64_t cap);

/// gd_classdb_Image_memory returns a pointer to the image's memory.
gd_addr gd_classdb_Image_memory(struct Object img); /// gd_classdb_Image_access returns the value at the given offset in the image's memory.
uint8_t gd_classdb_Image_access(struct Object img, int64_t offset);

/// gd_classdb_WorkerThreadPool_add_task adds a task to the worker thread pool.
void gd_classdb_WorkerThreadPool_add_task(gd_extension* extension, struct Object pool, gd_extension_task_id task, bool priority, struct String description); /// gd_classdb_WorkerThreadPool_add_group_task adds a group task to the worker thread pool.
void gd_classdb_WorkerThreadPool_add_group_task(gd_extension* extension, struct Object pool, gd_extension_task_id task, int32_t elements, int32_t arg, bool priority, struct String description);

/// gd_classdb_XMLParser_load loads an XML document into the parser.
int64_t gd_classdb_XMLParser_load(struct Object parser, char* buf, int64_t cap);

//
// Dictionaries
//

/// gd_packed_dictionary_access reads the value of a key from a packed dictionary.
void gd_packed_dictionary_access(struct Dictionary d, VARIANT_ARG(key), struct Variant* result); /// gd_packed_dictionary_modify modifies the value of a key in a packed dictionary.
void gd_packed_dictionary_modify(struct Dictionary d, VARIANT_ARG(key), VARIANT_ARG(val));

//
// Editor
//

/// gd_editor_add_documentation adds documentation XML to the editor.
void gd_editor_add_documentation(const char* xml, uint32_t len); /// gd_editor_add_plugin adds a plugin to the editor.
void gd_editor_add_plugin(struct StringName class_name); /// gd_editor_end_plugin removes a plugin from the editor.
void gd_editor_end_plugin(struct StringName class_name);

///
/// Iterators
///

/// gd_iterator_make creates an iterator from a variant.
void gd_iterator_make(VARIANT_ARG(v), struct Variant* result_iter, gd_error* err); /// gd_iterator_next advances the iterator to the next item.
bool gd_iterator_next(VARIANT_ARG(v), struct Variant* iter, gd_error* err); /// gd_iterator_load loads the value of the current item into 'result'.
void gd_iterator_load(VARIANT_ARG(v), VARIANT_ARG(i), struct Variant* result, gd_error* err);

// LogLevel identifies the severity of a log message.
typedef uint32_t gd_log_level;
static const gd_log_level GD_ERROR   = 0;
static const gd_log_level GD_WARNING = 1;

// gd_log prints a log message at the given severity level. 'text' is the human-readable message,
// 'code' is the error/warning code, 'func'/'file'/'line' identify the source location.
void gd_log(gd_log_level level,
    const char* text, uint32_t text_len,
    const char* code, uint32_t code_len,
    const char* func, uint32_t func_len,
    const char* file, uint32_t file_len,
    int32_t line, bool notify_editor
);

//
// Objects
//

// ObjectID uniquely identifies an object instance.
typedef uint64_t  ObjectID;

/// ClassTag uniquely identifies a class.
struct ClassTag { uintptr_t opaque; };

/// gd_object_make constructs a new instance of the class with the given name.
struct Object     gd_object_make(struct StringName name); /// gd_object_name returns the class name of the object.
struct StringName gd_object_name(struct Object obj); /// gd_object_type returns the class tag of the object.
struct ClassTag   gd_object_type(struct StringName name); /// gd_object_cast casts an object to a specific type.
struct Object     gd_object_cast(struct Object obj, struct ClassTag to); /// gd_object_lookup fetches an object with the given ID (slow).
struct Object     gd_object_lookup(ObjectID id); /// gd_object_global returns the named singleton.
struct Object     gd_object_global(struct StringName name); /// gd_object_call calls a method on the given object.
void              gd_object_call(struct Object obj, gd_method_id method, struct Variant* result, int64_t arg_count, struct Variant args[], gd_error* err); /// gd_object_id returns the instance ID of the given object.
ObjectID          gd_object_id(struct Object obj); /// gd_object_id_inside_variant extracts the object ID from a variant.
ObjectID          gd_object_id_inside_variant(VARIANT_ARG(v)); /// gd_object_free frees the given object.
void              gd_object_free(struct Object obj);

/// gd_method looks up a method for the given class and method name.
gd_method_id gd_method(struct StringName class_name, struct StringName method, int64_t hash); /// gd_method_call calls the given method on the given object.
void        gd_method_call(struct Object obj, gd_method_id fn, gd_addr result, gd_shape shape, const gd_addr args);

//
// Extension Scripts
//

/// gd_script returns the script instance associated with the given object and language.
gd_extension_script_id gd_script(struct Object obj, struct Object language);  /// gd_script_call calls the given method on the given object.
struct Script          gd_script_make(gd_extension* extension, gd_extension_script_id script);
void                   gd_script_call(struct Object obj, struct StringName name, struct Variant* result, int64_t arg_count, struct Variant args[], gd_error* err); /// gd_script_setup attaches a script instance to an object.
void                   gd_script_setup(struct Object obj, gd_extension_script_id script); /// gd_script_defines_method returns true if the given object defines the given method.
bool                   gd_script_defines_method(struct Object obj, struct StringName method); /// gd_object_script_placeholder_create creates a script placeholder object.
gd_extension_script_id gd_object_script_placeholder_create(struct Object language, struct Object script, struct Object owner); /// gd_object_script_placeholder_update updates a script placeholder object.
void                   gd_object_script_placeholder_update(gd_extension_script_id script, struct Array array, struct Dictionary dict);

/// gd_script_yield_property is used to yield a property state to the iterator.
void                  gd_script_yield_property(gd_property_iterator_t fn, uintptr_t arg, struct StringName name, VARIANT_ARG(state));

//
// Callables
//

#define CALLABLE_ARG(n) uint64_t n##_1, uint64_t n##_2
#define CALLABLE_ARG_GET(n) (struct Callable){{n##_1, n##_2}}
#define CALLABLE_ARG_PUT(n) n.opaque[0], n.opaque[1]


/// gd_callable_create creates a new [Callable] with the given id and owner, and stores it in [result].
void                    gd_callable_create(gd_extension* extension, gd_extension_callable_t id, ObjectID owner, struct Callable* result); /// gd_callable_lookup looks up the [Callable] with the given id and returns its opaque pointer.
gd_extension_callable_t gd_callable_lookup(CALLABLE_ARG(c));

///
/// Extension Lifecycle
///

// gd_extension_library_location returns the location of the extension's library.
struct String gd_extension_library_location();

//
// Extension Instances
//

/// gd_extension_object_setup sets the extension instance associated with the given object.
void gd_extension_object_setup(struct Object obj, struct StringName name, gd_extension_object_id inst);

//
// Extension Bindings
//

/// gd_object_lookup_extension_binding looks up the extension binding for the given object.
gd_extension_binding_id gd_object_lookup_extension_binding(gd_extension* extension, struct Object obj);
/// gd_object_attach_extension_binding attaches an extension binding to the given object.
void gd_object_attach_extension_binding(gd_extension* extension, struct Object obj, gd_extension_binding_id binding); /// gd_object_detach_extension_binding detaches the extension binding from the given object.
void gd_object_detach_extension_binding(struct Object obj);

//
// RefCounted
//

/// gd_ref_get_object returns the object associated with the given RefCounted reference.
struct Object gd_ref_get_object(struct RefCounted ref); /// gd_ref_set_object sets the object associated with the given RefCounted reference.
void          gd_ref_set_object(struct RefCounted ref, struct Object obj);

//
// Strings
//

/// gd_encoding describes the character encoding of a string.
typedef uint8_t gd_encoding; /// GD_LATIN1 represents the Latin-1 encoding
static const gd_encoding GD_LATIN1  = 0; /// GD_UTF8 represents UTF-8 encoding
static const gd_encoding GD_UTF8    = 1; /// GD_UTF16_LE represents UTF-16 little-endian encoding
static const gd_encoding GD_UTF16_LE = 2; /// GD_UTF16_BE represents UTF-16 big-endian encoding
static const gd_encoding GD_UTF16_BE = 3; /// GD_UTF32 represents UTF-32 encoding
static const gd_encoding GD_UTF32   = 4; /// GD_WIDE represents wide character encoding
static const gd_encoding GD_WIDE    = 5;

/// gd_string_access returns the character at the given index in the string.
int32_t           gd_string_access(struct String s, int64_t idx); /// gd_string_memory returns a pointer to the underlying buffer of the string.
gd_addr           gd_string_memory(struct String s); /// gd_string_decode decodes a string from the given encoding.
struct String     gd_string_decode(gd_encoding enc, const char* s, int64_t len); /// gd_string_encode encodes a string to the given encoding.
int64_t           gd_string_encode(gd_encoding enc, struct String s, char* buf, int64_t cap); /// gd_string_intern returns a StringName for the given string.
struct StringName gd_string_intern(gd_encoding enc, const char* s, int64_t len); /// gd_string_resize resizes the string to the given size.
struct String     gd_string_resize(struct String s, int64_t size); /// gd_string_append appends a string to the given string, returns updated 's'.
struct String     gd_string_append(struct String s, struct String other); /// gd_string_append_rune appends a rune to the given string, returns updated 's'.
struct String     gd_string_append_rune(struct String s, int32_t ch);

/// gd_extension_init initializes gd.c with the given proc address, library token, and callback table.
bool gd_extension_init(void* p_get_proc_address, void* p_library, void* r_initialization, const gd_extension* extension);

/// GD_EXTENSION defines the extension entry point and callback table for a language.
/// The language implements {Lang}_on_* with a friendly by-value API. The macro generates
/// static wrappers that bridge to the ABI-compatible gd_extension struct, allowing gd.c
/// to cast function pointers directly to GDExtension types with no runtime indirection.
/// On WASM, the wrappers pass variants/StringNames by value to avoid cross-address-space pointer chasing.
#define GD_EXTENSION(Lang)                                                                                                              \
    /* Language-facing callbacks (nice by-value API) */                                                                                 \
    CALLBACK(void,              Lang##_on_engine_init,              (gd_initialization_level level));                                   \
    CALLBACK(void,              Lang##_on_engine_exit,              (gd_initialization_level level));                                   \
    CALLBACK(void,              Lang##_on_yield,                    (void));                                                            \
    CALLBACK(void,              Lang##_on_first_frame,              (void));                                                            \
    CALLBACK(void,              Lang##_on_every_frame,              (void));                                                            \
    CALLBACK(void,              Lang##_on_final_frame,              (void));                                                            \
    CALLBACK(struct Object,     Lang##_on_extension_class_create,   (gd_extension_class_id id, bool notify_postinitialize));             \
    CALLBACK(gd_extension_method_id, Lang##_on_extension_class_method, (gd_extension_class_id id, struct StringName method, uint32_t hash)); \
    CALLBACK(gd_extension_method_id, Lang##_on_extension_class_caller, (gd_extension_class_id id, struct StringName method, uint32_t hash)); \
    CALLBACK(void,              Lang##_on_worker_thread_pool_task,  (gd_extension_task_id task));                                       \
    CALLBACK(void,              Lang##_on_worker_thread_pool_group_task, (gd_extension_task_id task, uint32_t n));                      \
    CALLBACK(void,              Lang##_on_editor_class_in_use_detection, (gd_addr packed_string_array));                                \
    CALLBACK(void,              Lang##_on_extension_script_property_iter, (gd_extension_object_id inst, gd_property_iterator_t op, uintptr_t arg)); \
    CALLBACK(bool,              Lang##_on_extension_script_categorization, (gd_extension_object_id inst, gd_addr property_info));       \
    CALLBACK(VariantType,       Lang##_on_extension_script_get_property_type, (gd_extension_object_id inst, struct StringName property, bool* r_is_valid)); \
    CALLBACK(struct Object,     Lang##_on_extension_script_get_owner, (gd_extension_object_id inst));                                   \
    CALLBACK(gd_addr,           Lang##_on_extension_script_get_methods, (gd_extension_object_id inst, uint32_t* r_count));              \
    CALLBACK(bool,              Lang##_on_extension_script_has_method, (gd_extension_object_id inst, struct StringName method));         \
    CALLBACK(int64_t,           Lang##_on_extension_script_get_method_argument_count, (gd_extension_object_id inst, struct StringName method, bool* r_is_valid)); \
    CALLBACK(struct Object,     Lang##_on_extension_script_get,     (gd_extension_object_id inst));                                     \
    CALLBACK(bool,              Lang##_on_extension_script_is_placeholder, (gd_extension_object_id inst));                              \
    CALLBACK(struct Object,     Lang##_on_extension_script_get_language, (gd_extension_object_id inst));                                \
    CALLBACK(void,              Lang##_on_callable_called,          (gd_extension_callable_t c, gd_unsafe_variants args, int64_t argc, struct Variant* ret, gd_error* err)); \
    CALLBACK(bool,              Lang##_on_callable_verify,          (gd_extension_callable_t c));                                       \
    CALLBACK(void,              Lang##_on_callable_delete,          (gd_extension_callable_t c));                                       \
    CALLBACK(uint32_t,          Lang##_on_callable_hashed,          (gd_extension_callable_t c));                                       \
    CALLBACK(bool,              Lang##_on_callable_equal,           (gd_extension_callable_t a, gd_extension_callable_t b));            \
    CALLBACK(bool,              Lang##_on_callable_less_than,       (gd_extension_callable_t a, gd_extension_callable_t b));            \
    CALLBACK(void,              Lang##_on_callable_string,          (gd_extension_callable_t c, bool* r_is_valid, gd_addr r_out));      \
    CALLBACK(int64_t,           Lang##_on_callable_length,          (gd_extension_callable_t c, bool* r_is_valid));                     \
    CALLBACK(bool,              Lang##_on_extension_instance_set,   (gd_extension_object_id inst, struct StringName property, VARIANT_ARG(val))); \
    CALLBACK(bool,              Lang##_on_extension_instance_get,   (gd_extension_object_id inst, struct StringName property, struct Variant* result)); \
    CALLBACK(bool,              Lang##_on_extension_instance_property_has_default, (gd_extension_object_id inst, struct StringName property)); \
    CALLBACK(bool,              Lang##_on_extension_instance_property_get_default, (gd_extension_object_id inst, struct StringName property, struct Variant* result)); \
    CALLBACK(gd_property_list_t, Lang##_on_extension_instance_property_validation, (gd_extension_object_id inst, VariantType type, struct StringName name, struct StringName class_name, uint32_t hint, struct String hint_string, uint32_t usage)); \
    CALLBACK(void,              Lang##_on_extension_instance_notification, (gd_extension_object_id inst, int32_t what, bool reverse));  \
    CALLBACK(void,              Lang##_on_extension_instance_stringify, (gd_extension_object_id inst, bool* r_is_valid, gd_addr r_out)); \
    CALLBACK(void,              Lang##_on_extension_instance_reference, (gd_extension_object_id inst));                                 \
    CALLBACK(bool,              Lang##_on_extension_instance_unreference, (gd_extension_object_id inst));                               \
    CALLBACK(RID,               Lang##_on_extension_instance_rid,   (gd_extension_object_id inst));                                    \
    CALLBACK(void,              Lang##_on_extension_instance_checked_call, (gd_extension_method_id fn, gd_extension_object_id inst, const gd_addr args, gd_addr result)); \
    CALLBACK(void,              Lang##_on_extension_instance_dynamic_call, (gd_extension_method_id fn, gd_extension_object_id inst, gd_unsafe_variants args, int64_t count, struct Variant* result, gd_error* err)); \
    CALLBACK(void,              Lang##_on_extension_instance_free,  (gd_extension_object_id inst));                                     \
    CALLBACK(void,              Lang##_on_extension_instance_called, (gd_extension_object_id inst, struct StringName name, gd_extension_method_id fn, const gd_addr args, gd_addr result)); \
    CALLBACK(gd_property_list_t, Lang##_on_extension_instance_property_list, (gd_extension_object_id inst));                           \
    CALLBACK(gd_extension_binding_id, Lang##_on_extension_binding_created, (uintptr_t token, gd_extension_object_id inst));            \
    CALLBACK(void,              Lang##_on_extension_binding_removed, (uintptr_t token, gd_extension_object_id inst, gd_extension_binding_id p1)); \
    CALLBACK(bool,              Lang##_on_extension_binding_reference, (uintptr_t token, gd_extension_object_id inst, bool p1));        \
                                                                                                                                       \
    /* Static wrappers: bridge by-value language API → ABI-compatible struct signatures */                                             \
    /* Compiler can inline Lang##_on_* into these when defined in the same TU */                                                       \
    static void _gd_wrap_##Lang##_engine_init(void* ud, gd_initialization_level level) {                                              \
        Lang##_on_engine_init(level);                                                                                                  \
        Lang##_on_yield();                                                                                                             \
    }                                                                                                                                  \
    static void _gd_wrap_##Lang##_engine_exit(void* ud, gd_initialization_level level) {                                              \
        Lang##_on_engine_exit(level);                                                                                                  \
        Lang##_on_yield();                                                                                                             \
    }                                                                                                                                  \
    static void _gd_wrap_##Lang##_first_frame(void) { Lang##_on_first_frame(); Lang##_on_yield(); }                                   \
    static void _gd_wrap_##Lang##_every_frame(void) { Lang##_on_every_frame(); Lang##_on_yield(); }                                   \
    static void _gd_wrap_##Lang##_final_frame(void) { Lang##_on_final_frame(); Lang##_on_yield(); }                                   \
    static void _gd_wrap_##Lang##_instance_free(void* class_userdata, void* inst) {                                                  \
        Lang##_on_extension_instance_free((gd_extension_object_id)inst);                                                               \
    }                                                                                                                                  \
    static void _gd_wrap_##Lang##_script_instance_free(void* inst) {                                                                  \
        Lang##_on_extension_instance_free((gd_extension_object_id)inst);                                                               \
    }                                                                                                                                  \
    static void* _gd_wrap_##Lang##_class_create(void* class_userdata, uint8_t notify) {                                                 \
        return (void*)(Lang##_on_extension_class_create((gd_extension_class_id)class_userdata, (bool)notify).opaque);                   \
    }                                                                                                                                  \
    static void* _gd_wrap_##Lang##_class_method(void* class_userdata, const void* m, uint32_t h) {                                     \
        return (void*)Lang##_on_extension_class_method((gd_extension_class_id)class_userdata, *(const struct StringName*)m, h);         \
    }                                                                                                                                  \
    static void* _gd_wrap_##Lang##_class_caller(void* class_userdata, const void* m, uint32_t h) {                                     \
        return (void*)Lang##_on_extension_class_caller((gd_extension_class_id)class_userdata, *(const struct StringName*)m, h);         \
    }                                                                                                                                  \
    static uint32_t _gd_wrap_##Lang##_script_get_property_type(void* inst, const void* p, uint8_t* ok) {                               \
        return Lang##_on_extension_script_get_property_type((gd_extension_object_id)inst, *(const struct StringName*)p, (bool*)ok);     \
    }                                                                                                                                  \
    static void* _gd_wrap_##Lang##_script_get_owner(void* inst) {                                                                      \
        return (void*)(Lang##_on_extension_script_get_owner((gd_extension_object_id)inst).opaque);                                     \
    }                                                                                                                                  \
    static uint8_t _gd_wrap_##Lang##_script_has_method(void* inst, const void* m) {                                                    \
        return Lang##_on_extension_script_has_method((gd_extension_object_id)inst, *(const struct StringName*)m);                       \
    }                                                                                                                                  \
    static int64_t _gd_wrap_##Lang##_script_get_method_argument_count(void* inst, const void* m, uint8_t* ok) {                        \
        return Lang##_on_extension_script_get_method_argument_count((gd_extension_object_id)inst, *(const struct StringName*)m, (bool*)ok); \
    }                                                                                                                                  \
    static void* _gd_wrap_##Lang##_script_get(void* inst) {                                                                            \
        return (void*)(Lang##_on_extension_script_get((gd_extension_object_id)inst).opaque);                                           \
    }                                                                                                                                  \
    static void* _gd_wrap_##Lang##_script_get_language(void* inst) {                                                                   \
        return (void*)(Lang##_on_extension_script_get_language((gd_extension_object_id)inst).opaque);                                   \
    }                                                                                                                                  \
    static uint8_t _gd_wrap_##Lang##_instance_set(void* inst, const void* p, const void* v) {                                          \
        const struct Variant* var = (const struct Variant*)v;                                                                           \
        return Lang##_on_extension_instance_set((gd_extension_object_id)inst, *(const struct StringName*)p, var->tag, var->payload[0], var->payload[1]); \
    }                                                                                                                                  \
    static uint8_t _gd_wrap_##Lang##_instance_get(void* inst, const void* p, void* r) {                                                \
        return Lang##_on_extension_instance_get((gd_extension_object_id)inst, *(const struct StringName*)p, (struct Variant*)r);        \
    }                                                                                                                                  \
    static uint8_t _gd_wrap_##Lang##_instance_property_has_default(void* inst, const void* p) {                                        \
        return Lang##_on_extension_instance_property_has_default((gd_extension_object_id)inst, *(const struct StringName*)p);           \
    }                                                                                                                                  \
    static uint8_t _gd_wrap_##Lang##_instance_property_get_default(void* inst, const void* p, void* r) {                               \
        return Lang##_on_extension_instance_property_get_default((gd_extension_object_id)inst, *(const struct StringName*)p, (struct Variant*)r); \
    }                                                                                                                                  \
    static uint8_t _gd_wrap_##Lang##_instance_property_validation(void* inst, void* prop) {                                             \
        /* GDExtensionPropertyInfo: {uint32_t type, void* name, void* class_name, uint32_t hint, void* hint_string, uint32_t usage} */ \
        uint32_t* p = (uint32_t*)prop;                                                                                                 \
        gd_property_list_t result = Lang##_on_extension_instance_property_validation((gd_extension_object_id)inst,                      \
            (VariantType)p[0],                                                                                                         \
            *(struct StringName*)((char*)prop + sizeof(uint32_t)),                                                                     \
            *(struct StringName*)((char*)prop + sizeof(uint32_t) + sizeof(void*)),                                                     \
            *(uint32_t*)((char*)prop + sizeof(uint32_t) + 2*sizeof(void*)),                                                            \
            *(struct String*)((char*)prop + sizeof(uint32_t) + 2*sizeof(void*) + sizeof(uint32_t)),                                    \
            *(uint32_t*)((char*)prop + sizeof(uint32_t) + 2*sizeof(void*) + sizeof(uint32_t) + sizeof(void*))                          \
        );                                                                                                                             \
        if (!result) return true; /* null = valid, no changes */                                                                       \
        uint32_t count = *(uint32_t*)result;                                                                                           \
        if (count == 0) { gd_free((gd_addr)result); return false; } /* empty = invalid */                                              \
        /* count == 1: copy updated property info back */                                                                              \
        void* info = *(void**)((char*)result + 2*sizeof(int32_t));                                                                     \
        if (info) __builtin_memcpy(prop, info, sizeof(uint32_t) + 2*sizeof(void*) + sizeof(uint32_t) + sizeof(void*) + sizeof(uint32_t)); \
        gd_property_list_free(result);                                                                                                 \
        return true;                                                                                                                   \
    }                                                                                                                                  \
    static void _gd_wrap_##Lang##_instance_called(void* inst, const void* name, void* fn, const void* const* args, void* ret) {          \
        Lang##_on_extension_instance_called((gd_extension_object_id)inst, *(const struct StringName*)name, (gd_extension_method_id)fn, (const gd_addr)args, (gd_addr)ret); \
    }                                                                                                                                  \
    static const void* _gd_wrap_##Lang##_script_property_list(void* inst, uint32_t* count) {                                          \
        gd_property_list_t handle = Lang##_on_extension_instance_property_list((gd_extension_object_id)inst);                          \
        if (!handle) { *count = 0; return (void*)0; }                                                                                  \
        /* property_list layout: {int32_t push, int32_t size, void* info, void* meta} */                                              \
        *count = *(uint32_t*)handle;                                                                                                   \
        const void* info = *(const void**)((char*)handle + 2*sizeof(int32_t));                                                        \
        void* meta = *(void**)((char*)handle + 2*sizeof(int32_t) + sizeof(void*));                                                    \
        if (meta) gd_free((gd_addr)meta);                                                                                             \
        gd_free((gd_addr)handle);                                                                                                      \
        return info;                                                                                                                   \
    }                                                                                                                                  \
    extern void Lang##_on_extension_instance_dynamic_call(gd_extension_method_id fn, gd_extension_object_id inst,                      \
        gd_unsafe_variants args, int64_t count, struct Variant* result, gd_error* err);                                                \
    static void _gd_wrap_##Lang##_script_call(void* inst, const void* method, const void* const* args, int64_t argc, void* ret, void* err) { \
        Lang##_on_extension_instance_dynamic_call(0, (gd_extension_object_id)inst, (gd_unsafe_variants)args, argc, (struct Variant*)ret, (gd_error*)err); \
    }                                                                                                                                  \
    static uint8_t _gd_wrap_##Lang##_script_unreference(void* inst) {                                                                 \
        return Lang##_on_extension_instance_unreference((gd_extension_object_id)inst);                                                  \
    }                                                                                                                                  \
    static void _gd_wrap_##Lang##_worker_thread_pool_task(void* userdata) {                                                            \
        Lang##_on_worker_thread_pool_task((gd_extension_task_id)userdata);                                                              \
    }                                                                                                                                  \
    static void _gd_wrap_##Lang##_worker_thread_pool_group_task(void* userdata, uint32_t n) {                                          \
        Lang##_on_worker_thread_pool_group_task((gd_extension_task_id)userdata, n);                                                     \
    }                                                                                                                                  \
    static void _gd_wrap_##Lang##_editor_class_in_use_detection(void* packed_string_array) {                                           \
        Lang##_on_editor_class_in_use_detection((gd_addr)packed_string_array);                                                          \
    }                                                                                                                                  \
    static void _gd_wrap_##Lang##_script_property_iter(void* inst, void* add_func, void* userdata) {                                   \
        Lang##_on_extension_script_property_iter((gd_extension_object_id)inst, (gd_property_iterator_t)add_func, (uintptr_t)userdata);  \
    }                                                                                                                                  \
    static uint8_t _gd_wrap_##Lang##_script_categorization(void* inst, void* property_info) {                                          \
        return Lang##_on_extension_script_categorization((gd_extension_object_id)inst, (gd_addr)property_info);                         \
    }                                                                                                                                  \
    static const void* _gd_wrap_##Lang##_script_get_methods(void* inst, uint32_t* r_count) {                                           \
        return (const void*)Lang##_on_extension_script_get_methods((gd_extension_object_id)inst, r_count);                              \
    }                                                                                                                                  \
    static uint8_t _gd_wrap_##Lang##_script_is_placeholder(void* inst) {                                                               \
        return Lang##_on_extension_script_is_placeholder((gd_extension_object_id)inst);                                                 \
    }                                                                                                                                  \
    static void _gd_wrap_##Lang##_callable_called(void* c, const void* const* args, int64_t argc, void* ret, void* err) {              \
        Lang##_on_callable_called((gd_extension_callable_t)c, (gd_unsafe_variants)args, argc, (struct Variant*)ret, (gd_error*)err);    \
    }                                                                                                                                  \
    static uint8_t _gd_wrap_##Lang##_callable_verify(void* c) {                                                                        \
        return Lang##_on_callable_verify((gd_extension_callable_t)c);                                                                   \
    }                                                                                                                                  \
    static void _gd_wrap_##Lang##_callable_delete(void* c) {                                                                           \
        Lang##_on_callable_delete((gd_extension_callable_t)c);                                                                          \
    }                                                                                                                                  \
    static uint32_t _gd_wrap_##Lang##_callable_hashed(void* c) {                                                                       \
        return Lang##_on_callable_hashed((gd_extension_callable_t)c);                                                                   \
    }                                                                                                                                  \
    static uint8_t _gd_wrap_##Lang##_callable_equal(void* a, void* b) {                                                                \
        return Lang##_on_callable_equal((gd_extension_callable_t)a, (gd_extension_callable_t)b);                                        \
    }                                                                                                                                  \
    static uint8_t _gd_wrap_##Lang##_callable_less_than(void* a, void* b) {                                                            \
        return Lang##_on_callable_less_than((gd_extension_callable_t)a, (gd_extension_callable_t)b);                                    \
    }                                                                                                                                  \
    static void _gd_wrap_##Lang##_callable_string(void* c, uint8_t* r_is_valid, void* r_out) {                                         \
        Lang##_on_callable_string((gd_extension_callable_t)c, (bool*)r_is_valid, (gd_addr)r_out);                                       \
    }                                                                                                                                  \
    static int64_t _gd_wrap_##Lang##_callable_length(void* c, uint8_t* r_is_valid) {                                                   \
        return Lang##_on_callable_length((gd_extension_callable_t)c, (bool*)r_is_valid);                                                \
    }                                                                                                                                  \
    static void _gd_wrap_##Lang##_instance_notification(void* inst, int32_t what, uint8_t reverse) {                                   \
        Lang##_on_extension_instance_notification((gd_extension_object_id)inst, what, (bool)reverse);                                   \
    }                                                                                                                                  \
    static void _gd_wrap_##Lang##_instance_stringify(void* inst, uint8_t* r_is_valid, void* r_out) {                                   \
        Lang##_on_extension_instance_stringify((gd_extension_object_id)inst, (bool*)r_is_valid, (gd_addr)r_out);                        \
    }                                                                                                                                  \
    static void _gd_wrap_##Lang##_instance_reference(void* inst) {                                                                     \
        Lang##_on_extension_instance_reference((gd_extension_object_id)inst);                                                           \
    }                                                                                                                                  \
    static void _gd_wrap_##Lang##_instance_unreference(void* inst) {                                                                   \
        Lang##_on_extension_instance_unreference((gd_extension_object_id)inst);                                                         \
    }                                                                                                                                  \
    static uint64_t _gd_wrap_##Lang##_instance_rid(void* inst) {                                                                       \
        return Lang##_on_extension_instance_rid((gd_extension_object_id)inst);                                                          \
    }                                                                                                                                  \
    static void _gd_wrap_##Lang##_instance_checked_call(void* fn, void* inst, const void* const* args, void* ret) {                    \
        Lang##_on_extension_instance_checked_call((gd_extension_method_id)fn, (gd_extension_object_id)inst, (const gd_addr)args, (gd_addr)ret); \
    }                                                                                                                                  \
    static void _gd_wrap_##Lang##_instance_dynamic_call(void* fn, void* inst, const void* const* args, int64_t argc, void* ret, void* err) { \
        Lang##_on_extension_instance_dynamic_call((gd_extension_method_id)fn, (gd_extension_object_id)inst, (gd_unsafe_variants)args, argc, (struct Variant*)ret, (gd_error*)err); \
    }                                                                                                                                  \
    static void* _gd_wrap_##Lang##_binding_created(void* token, void* inst) {                                                          \
        return (void*)Lang##_on_extension_binding_created((uintptr_t)token, (gd_extension_object_id)inst);                              \
    }                                                                                                                                  \
    static void _gd_wrap_##Lang##_binding_removed(void* token, void* inst, void* binding) {                                            \
        Lang##_on_extension_binding_removed((uintptr_t)token, (gd_extension_object_id)inst, (gd_extension_binding_id)binding);          \
    }                                                                                                                                  \
    static uint8_t _gd_wrap_##Lang##_binding_reference(void* token, void* binding, uint8_t reference) {                                \
        return Lang##_on_extension_binding_reference((uintptr_t)token, (gd_extension_object_id)binding, (bool)reference);                \
    }                                                                                                                                  \
                                                                                                                                       \
    static const gd_extension Lang##_gd_extension = {                                                                                  \
        .on_engine_init                                = _gd_wrap_##Lang##_engine_init,                                               \
        .on_engine_exit                                = _gd_wrap_##Lang##_engine_exit,                                               \
        .on_yield                                      = Lang##_on_yield,                                                              \
        .on_first_frame                                = _gd_wrap_##Lang##_first_frame,                                               \
        .on_every_frame                                = _gd_wrap_##Lang##_every_frame,                                               \
        .on_final_frame                                = _gd_wrap_##Lang##_final_frame,                                               \
        .on_extension_class_create                     = _gd_wrap_##Lang##_class_create,                                              \
        .on_extension_class_method                     = _gd_wrap_##Lang##_class_method,                                              \
        .on_extension_class_caller                     = _gd_wrap_##Lang##_class_caller,                                              \
        .on_worker_thread_pool_task                    = _gd_wrap_##Lang##_worker_thread_pool_task,                                   \
        .on_worker_thread_pool_group_task              = _gd_wrap_##Lang##_worker_thread_pool_group_task,                              \
        .on_editor_class_in_use_detection              = _gd_wrap_##Lang##_editor_class_in_use_detection,                              \
        .on_extension_script_property_iter             = _gd_wrap_##Lang##_script_property_iter,                                       \
        .on_extension_script_categorization            = _gd_wrap_##Lang##_script_categorization,                                      \
        .on_extension_script_get_property_type         = _gd_wrap_##Lang##_script_get_property_type,                                  \
        .on_extension_script_get_owner                 = _gd_wrap_##Lang##_script_get_owner,                                          \
        .on_extension_script_get_methods               = _gd_wrap_##Lang##_script_get_methods,                                        \
        .on_extension_script_has_method                = _gd_wrap_##Lang##_script_has_method,                                         \
        .on_extension_script_get_method_argument_count = _gd_wrap_##Lang##_script_get_method_argument_count,                          \
        .on_extension_script_get                       = _gd_wrap_##Lang##_script_get,                                                \
        .on_extension_script_is_placeholder            = _gd_wrap_##Lang##_script_is_placeholder,                                     \
        .on_extension_script_get_language              = _gd_wrap_##Lang##_script_get_language,                                       \
        .on_callable_called                            = _gd_wrap_##Lang##_callable_called,                                             \
        .on_callable_verify                            = _gd_wrap_##Lang##_callable_verify,                                            \
        .on_callable_delete                            = _gd_wrap_##Lang##_callable_delete,                                            \
        .on_callable_hashed                            = _gd_wrap_##Lang##_callable_hashed,                                            \
        .on_callable_equal                             = _gd_wrap_##Lang##_callable_equal,                                             \
        .on_callable_less_than                         = _gd_wrap_##Lang##_callable_less_than,                                         \
        .on_callable_string                            = _gd_wrap_##Lang##_callable_string,                                            \
        .on_callable_length                            = _gd_wrap_##Lang##_callable_length,                                            \
        .on_extension_instance_set                     = _gd_wrap_##Lang##_instance_set,                                              \
        .on_extension_instance_get                     = _gd_wrap_##Lang##_instance_get,                                              \
        .on_extension_instance_property_has_default    = _gd_wrap_##Lang##_instance_property_has_default,                             \
        .on_extension_instance_property_get_default    = _gd_wrap_##Lang##_instance_property_get_default,                             \
        .on_extension_instance_property_validation     = _gd_wrap_##Lang##_instance_property_validation,                               \
        .on_extension_instance_notification            = _gd_wrap_##Lang##_instance_notification,                                      \
        .on_extension_instance_stringify               = _gd_wrap_##Lang##_instance_stringify,                                         \
        .on_extension_instance_reference               = _gd_wrap_##Lang##_instance_reference,                                        \
        .on_extension_instance_unreference             = _gd_wrap_##Lang##_instance_unreference,                                      \
        .on_extension_instance_rid                     = _gd_wrap_##Lang##_instance_rid,                                               \
        .on_extension_instance_checked_call            = _gd_wrap_##Lang##_instance_checked_call,                                      \
        .on_extension_instance_dynamic_call            = _gd_wrap_##Lang##_instance_dynamic_call,                                      \
        .on_extension_instance_free                    = _gd_wrap_##Lang##_instance_free,                                             \
        .on_extension_script_free                      = _gd_wrap_##Lang##_script_instance_free,                                       \
        .on_extension_instance_called                  = _gd_wrap_##Lang##_instance_called,                                           \
        .on_extension_instance_property_list           = _gd_wrap_##Lang##_script_property_list,                                       \
        .on_extension_script_call                               = _gd_wrap_##Lang##_script_call,                                                \
        .on_extension_script_unreference                        = _gd_wrap_##Lang##_script_unreference,                                         \
        .on_extension_binding_created                  = _gd_wrap_##Lang##_binding_created,                                             \
        .on_extension_binding_removed                  = _gd_wrap_##Lang##_binding_removed,                                            \
        .on_extension_binding_reference                = _gd_wrap_##Lang##_binding_reference,                                          \
    };                                                                                                                                 \
    bool Lang##_gd_extension_init(void* p_get_proc_address, void* p_library, void* r_initialization) {                                \
        return gd_extension_init(p_get_proc_address, p_library, r_initialization, &Lang##_gd_extension);                               \
    }

#ifdef __EMSCRIPTEN__
}
#endif
#endif
