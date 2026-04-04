//go:build cgo

package gdunsafe

// #include "gd.h"
import "C"

type String uintptr

type Array uintptr

func (array Array) Set(index Int, value Variant) {
	C.gd_array_set(C.Array(array), C.int64_t(index), C.uint64_t(value[0]), C.uint64_t(value[1]), C.uint64_t(value[2]))
}

func (array Array) Get(index Int) Variant {
	r := C.gd_array_get(C.Array(array), C.int64_t(index))
	return Variant{uint64(r.tag), uint64(r.payload[0]), uint64(r.payload[1])}
}

func VersionMajor() uint32     { return uint32(C.gd_version_major()) }
func VersionMinor() uint32     { return uint32(C.gd_version_minor()) }
func VersionPatch() uint32     { return uint32(C.gd_version_patch()) }
func VersionHex() uint32       { return uint32(C.gd_version_hex()) }
func VersionStatus() String    { return String(C.gd_version_status()) }
func VersionBuild() String     { return String(C.gd_version_build()) }
func VersionHash() String      { return String(C.gd_version_hash()) }
func VersionTimestamp() uint64 { return uint64(C.gd_version_timestamp()) }
func VersionString() String    { return String(C.gd_version_string()) }
