// Package gdunsafe provides unsafe access to the gdextension API (no protections for double-free, pointer-aliasing or use-after-free).
package gdunsafe

import "unsafe"

// Version of the engine.
func Version() String { return String{raw: gd_version()} } //gdextension:get_godot_version2

// Major version of the engine.
func VersionMajor() uint32 { return gd_version_major() } //gdextension:get_godot_version2

// Minor version of the engine.
func VersionMinor() uint32 { return gd_version_minor() } //gdextension:get_godot_version2

// Patch version of the engine.
func VersionPatch() uint32 { return gd_version_patch() } //gdextension:get_godot_version2

// Hexed version of the engine.
func VersionHexed() uint32 { return gd_version_hexed() } //gdextension:get_godot_version2

// State of the engine (e.g. "stable", "beta", "alpha").
func VersionState() String { return String{raw: gd_version_state()} } //gdextension:get_godot_version2

// Build type.
func VersionBuild() String { return String{raw: gd_version_build()} } //gdextension:get_godot_version2

// Commit hash of the build.
func VersionStamp() String { return String{raw: gd_version_stamp()} } //gdextension:get_godot_version2

// Timestamp of the engine build.
func VersionNanos() int64 { return gd_version_nanos() } //gdextension:get_godot_version2

type MutableMemory struct {
	addr gd_addr
}

// Malloc allocates optionally 8-byte aligned memory.
func Malloc(size uintptr, pad8 bool) MutableMemory {
	return MutableMemory{addr: gd_malloc(int64(size), pad8)}
}

// Resize resizes optionally 8-byte aligned memory.
func Resize(ptr MutableMemory, size uintptr, pad8 bool) MutableMemory {
	return MutableMemory{addr: gd_resize(ptr.addr, int64(size), pad8)}
}

// Memset sets a block of memory to a given value.
func Memset(ptr MutableMemory, size uintptr, value byte) {
	gd_memset(ptr.addr, value, int64(size))
}

// LibraryPath returns a string representing the location of the current extension.
func LibraryPath() String { return String{raw: gd_extension_library_location()} } //gdextension:get_library_path

// Tag returns the [ClassTag] of the class.
func (class Class) Tag() ClassTag { return ClassTag{raw: gd_object_type(class.raw)} }

// Free releases any resources associated with the given value of type T.
func Free[T Any](val T) {
	gd_builtin_free(uint32(variantTypeOf[T]()), gd_addrOf(shapeFor[T](), unsafe.Pointer(&val)))
}
