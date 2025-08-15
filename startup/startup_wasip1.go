package startup

import (
	"fmt"

	gd "graphics.gd/internal"
)

//go:wasmimport gdextension read_result_buffer
func read_result_buffer(uint32, uint32, uint32) uint64

//go:wasmimport gdextension write_params_buffer
func write_params_buffer(uint32, uint32, uint32, uint64)

//go:wasmimport gdextension write_params_buffer8
func write_params_buffer16(uint32, uint32, uint64, uint64, uint64, uint64, uint64, uint64, uint64, uint64)

//go:wasmimport gdextension get_godot_version_major
func get_godot_version_major() uint32

//go:wasmimport gdextension get_godot_version_minor
func get_godot_version_minor() uint32

//go:wasmimport gdextension get_godot_version_patch
func get_godot_version_patch() uint32

//go:wasmimport gdextension mem_alloc
func mem_alloc(uint64) uint64

//go:wasmimport gdextension mem_realloc
func mem_realloc(ptr uint64, size uint64) uint64

//go:wasmimport gdextension mem_free
func mem_free(ptr uint64)

//go:wasmimport gdextension print_error
func print_error(string, string, string, int32, bool)

//go:wasmimport gdextension print_error_with_message
func print_error_with_message(string, string, string, string, int32, bool)

//go:wasmimport gdextension print_warning
func print_warning(string, string, string, int32, bool)

//go:wasmimport gdextension print_warning_with_message
func print_warning_with_message(string, string, string, string, int32, bool)

//go:wasmimport gdextension print_script_error
func print_script_error(string, string, string, int32, bool)

//go:wasmimport gdextension print_script_error_with_message
func print_script_error_with_message(string, string, string, string, int32, bool)

//go:wasmimport gdextension get_native_struct_size
func get_native_struct_size(uint64) uint64

//go:wasmimport gdextension variant_new_copy
func variant_new_copy(v1, v2, v3 uint64)

//go:wasmimport gdextension variant_destroy
func variant_destroy(v1, v2, v3 uint64)

//go:wasmimport gdextension variant_call
func variant_call(v1, v2, v3 uint64, method uint64, argc uint32) gd.CallErrorType

//go:wasmimport gdextension variant_call_static
func variant_call_static(vtype gd.VariantType, method uint64, argc uint32) gd.CallErrorType

//go:wasmimport gdextension variant_evaluate
func variant_evaluate(op gd.Operator, a1, a2, a3, b1, b2, b3 uint64) bool

//go:wasmimport gdextension variant_set
func variant_set(v1, v2, v3, a1, a2, a3, b1, b2, b3 uint64) bool

//go:wasmimport gdextension variant_set_named
func variant_set_named(v1, v2, v3 uint64, name uint32, b1, b2, b3 uint64) bool

//go:wasmimport gdextension variant_set_keyed
func variant_set_keyed(v1, v2, v3, a1, a2, a3, b1, b2, b3 uint64) bool

//go:wasmimport gdextension variant_set_indexed
func variant_set_indexed(v1, v2, v3 uint64, index int64, b1, b2, b3 uint64) int32

//go:wasmimport gdextension variant_get
func variant_get(v1, v2, v3, k1, k2, k3 uint64) bool

//go:wasmimport gdextension variant_get_named
func variant_get_named(v1, v2, v3, name uint64) bool

//go:wasmimport gdextension variant_get_keyed
func variant_get_keyed(v1, v2, v3, k1, k2, k3 uint64) bool

//go:wasmimport gdextension variant_get_indexed
func variant_get_indexed(v1, v2, v3 uint64, index int64) bool

//go:wasmimport gdextension variant_hash
func variant_hash(v1, v2, v3 uint64) int64

//go:wasmimport gdextension variant_recursive_hash
func variant_recursive_hash(v1, v2, v3 uint64, c int64) int64

//go:wasmimport gdextension variant_booleanize
func variant_booleanize(v1, v2, v3 uint64) bool

//go:wasmimport gdextension variant_duplicate
func variant_duplicate(v1, v2, v3 uint64, deep bool)

//go:wasmimport gdextension variant_stringify
func variant_stringify(v1, v2, v3 uint64) uint64

//go:wasmimport gdextension variant_get_type
func variant_get_type(v1, v2, v3 uint64) gd.VariantType

//go:wasmimport gdextension variant_has_method
func variant_has_method(v1, v2, v3 uint64, method uint64) bool

//go:wasmimport gdextension variant_get_member
func variant_get_member(vtype gd.VariantType, name uint64) bool

//go:wasmimport gdextension variant_has_key
func variant_has_key(v1, v2, v3, k1, k2, k3 uint64) bool

//go:wasmimport gdextension variant_get_object_instance_id
func variant_get_object_instance_id(v1, v2, v3 uint64) int64

//go:wasmimport gdextension variant_get_type_name
func variant_get_type_name(vtype gd.VariantType) uint64

//go:wasmimport gdextension variant_can_convert
func variant_can_convert(v1, v2, v3 uint64, to gd.VariantType) bool

//go:wasmimport gdextension get_variant_to_type_constructor
func get_variant_to_type_constructor(vtype gd.VariantType) uint64

//go:wasmimport gdextension call_variant_to_type_constructor
func call_variant_to_type_constructor(uint64)

// If we are starting up under wasmip1 then that means we are peforming a hot-reload.
// we need to link up with our host and initialise the Global state, we will then
// receive all of the active extension classes so that they can be re-created.
func init() {
	gd.Global.GetGodotVersion = func() gd.Version {
		return gd.Version{
			Major: get_godot_version_major(),
			Minor: get_godot_version_minor(),
			Patch: get_godot_version_patch(),
		}
	}
	fmt.Println("Godot Version: ", gd.Global.GetGodotVersion().Major)
}
