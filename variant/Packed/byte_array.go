// Package Packed provides functions for working with packed arrays.
package Packed

import (
	"encoding/binary"
	"encoding/hex"
	"math"
	"unsafe"

	gd "graphics.gd/internal"
	"graphics.gd/variant"
)

// ByteArray is an array specifically designed to hold bytes. Packs data tightly, so it saves memory for
// large array sizes.
//
// ByteArray also provides methods to encode/decode various types to/from bytes. The way values are
// encoded is an implementation detail and shouldn't be relied upon when interacting with external apps.
type ByteArray struct {
	Array[ByteArray, byte, gd.Int, gd.PackedByteArray, *gd.PackedByteArray]
}

func (array ByteArray) conv(c byte) gd.Int { return gd.Int(c) }
func (array ByteArray) wrap(c gd.Int) byte { return byte(c) }
func (array ByteArray) alloc() gd.PackedByteArray {
	return gd.NewPackedByteArray()
}

func (array ByteArray) make(local []byte, proxy gd.PackedByteArray) ByteArray {
	array.local = local
	array.proxy = proxy
	return array
}

func (array ByteArray) less(a, b byte) bool {
	return a < b
}

type CompressionMode int

const (
	CompressionFastLZ CompressionMode = iota
	CompressionDeflate
	CompressionZstandard
	CompressionGzip
	CompressionBrotli
)

// Compress returns a new PackedByteArray with the data compressed. Set the compression mode using one of
// [CompressionMode]'s constants.
func (array ByteArray) Compress(mode CompressionMode) ByteArray { //gd:PackedByteArray.compress
	array.proxy = array.Proxy()
	array.proxy = array.proxy.Compress(gd.Int(mode))
	return array
}

// Decompress returns a new PackedByteArray with the data decompressed. Set buffer_size to the size of the
// uncompressed data. Set the compression mode using one of CompressionMode's constants.
func (array ByteArray) DecompressSize(buffer_size int, mode CompressionMode) ByteArray { //gd:PackedByteArray.decompress
	array.proxy = array.Proxy()
	array.proxy = array.proxy.Decompress(gd.Int(buffer_size), gd.Int(mode))
	return array
}

// DecompressUpto returns a new PackedByteArray with the data decompressed. Set the compression mode using
// one of CompressionMode's constants. This method only accepts brotli, gzip, and deflate compression modes.
//
// This method is potentially slower than decompress, as it may have to re-allocate its output buffer multiple
// times while decompressing, whereas decompress knows it's output buffer size from the beginning.
func (array ByteArray) DecompressUpto(max_output_size int, mode CompressionMode) ByteArray { //gd:PackedByteArray.decompress_dynamic
	array.proxy = array.Proxy()
	array.proxy = array.proxy.DecompressDynamic(gd.Int(max_output_size), gd.Int(mode))
	return array
}

// DecodeFloat64 decodes a 64-bit floating-point number from the bytes starting at offset. Fails if the
// byte count is insufficient. Returns 0.0 if a valid number can't be decoded.
func (array ByteArray) DecodeFloat64(offset uintptr) float64 { //gd:PackedByteArray.decode_double
	if array.proxy != (gd.PackedByteArray{}) {
		return float64(array.proxy.DecodeDouble(gd.Int(offset)))
	}
	if offset+8 > uintptr(len(array.local)) {
		return 0.0
	}
	return *(*float64)(unsafe.Pointer(&array.local[offset]))
}

// DecodeFloat32 decodes a 32-bit floating-point number from the bytes starting at offset. Fails if the byte
// count is insufficient. Returns 0.0 if a valid number can't be decoded.
func (array ByteArray) DecodeFloat32(offset uintptr) float32 { //gd:PackedByteArray.decode_float PackedByteArray.decode_half
	if array.proxy != (gd.PackedByteArray{}) {
		return float32(array.proxy.DecodeFloat(gd.Int(offset)))
	}
	if offset+4 > uintptr(len(array.local)) {
		return 0.0
	}
	return *(*float32)(unsafe.Pointer(&array.local[offset]))
}

// DecodeInt8 decodes a 8-bit signed integer number from the bytes starting at offset. Fails if the byte count is
// insufficient. Returns 0 if a valid number can't be decoded.
func (array ByteArray) DecodeInt8(offset uintptr) int8 { //gd:PackedByteArray.decode_s8
	if array.proxy != (gd.PackedByteArray{}) {
		return int8(array.proxy.DecodeS8(gd.Int(offset)))
	}
	if offset+1 > uintptr(len(array.local)) {
		return 0
	}
	return int8(array.local[offset])
}

// DecodeInt16 decodes a 16-bit signed integer number from the bytes starting at offset. Fails if the byte count
// is insufficient. Returns 0 if a valid number can't be decoded.
func (array ByteArray) DecodeInt16(offset uintptr) int16 { //gd:PackedByteArray.decode_s16
	if array.proxy != (gd.PackedByteArray{}) {
		return int16(array.proxy.DecodeS16(gd.Int(offset)))
	}
	if offset+2 > uintptr(len(array.local)) {
		return 0
	}
	return *(*int16)(unsafe.Pointer(&array.local[offset]))
}

// DecodeInt32 decodes a 32-bit signed integer number from the bytes starting at offset. Fails if the byte count
// is insufficient. Returns 0 if a valid number can't be decoded.
func (array ByteArray) DecodeInt32(offset uintptr) int32 { //gd:PackedByteArray.decode_s32
	if array.proxy != (gd.PackedByteArray{}) {
		return int32(array.proxy.DecodeS32(gd.Int(offset)))
	}
	if offset+4 > uintptr(len(array.local)) {
		return 0
	}
	return *(*int32)(unsafe.Pointer(&array.local[offset]))
}

// DecodeInt64 decodes a 64-bit signed integer number from the bytes starting at offset. Fails if the byte count
// is insufficient. Returns 0 if a valid number can't be decoded.
func (array ByteArray) DecodeInt64(offset uintptr) int64 { //gd:PackedByteArray.decode_s64
	if array.proxy != (gd.PackedByteArray{}) {
		return int64(array.proxy.DecodeS64(gd.Int(offset)))
	}
	if offset+8 > uintptr(len(array.local)) {
		return 0
	}
	return *(*int64)(unsafe.Pointer(&array.local[offset]))
}

// DecodeUint8 decodes a 8-bit unsigned integer number from the bytes starting at offset. Fails if the byte count
// is insufficient. Returns 0 if a valid number can't be decoded.
func (array ByteArray) DecodeUint8(offset uintptr) uint8 { //gd:PackedByteArray.decode_u8
	if array.proxy != (gd.PackedByteArray{}) {
		return uint8(array.proxy.DecodeU8(gd.Int(offset)))
	}
	if offset+1 > uintptr(len(array.local)) {
		return 0
	}
	return array.local[offset]
}

// DecodeUint16 decodes a 16-bit unsigned integer number from the bytes starting at offset. Fails if the byte count
// is insufficient. Returns 0 if a valid number can't be decoded.
func (array ByteArray) DecodeUint16(offset uintptr) uint16 { //gd:PackedByteArray.decode_u16
	if array.proxy != (gd.PackedByteArray{}) {
		return uint16(array.proxy.DecodeU16(gd.Int(offset)))
	}
	if offset+2 > uintptr(len(array.local)) {
		return 0
	}
	return *(*uint16)(unsafe.Pointer(&array.local[offset]))
}

// DecodeUint32 decodes a 32-bit unsigned integer number from the bytes starting at offset. Fails if the byte count
// is insufficient. Returns 0 if a valid number can't be decoded.
func (array ByteArray) DecodeUint32(offset uintptr) uint32 { //gd:PackedByteArray.decode_u32
	if array.proxy != (gd.PackedByteArray{}) {
		return uint32(array.proxy.DecodeU32(gd.Int(offset)))
	}
	if offset+4 > uintptr(len(array.local)) {
		return 0
	}
	return *(*uint32)(unsafe.Pointer(&array.local[offset]))
}

// DecodeUint64 decodes a 64-bit unsigned integer number from the bytes starting at offset. Fails if the byte count
// is insufficient. Returns 0 if a valid number can't be decoded.
func (array ByteArray) DecodeUint64(offset uintptr) uint64 { //gd:PackedByteArray.decode_u64
	if array.proxy != (gd.PackedByteArray{}) {
		return uint64(array.proxy.DecodeU64(gd.Int(offset)))
	}
	if offset+8 > uintptr(len(array.local)) {
		return 0
	}
	return *(*uint64)(unsafe.Pointer(&array.local[offset]))
}

// DecodeString decodes a length-prefixed string from the bytes starting at offset. Fails if the byte count is
// insufficient. Returns an empty string if a valid string can't be decoded.
func (array ByteArray) decodeString(offset uintptr) string {
	length := uintptr(array.DecodeUint32(offset))
	offset += 4
	if offset+length > uintptr(len(array.local)) {
		return ""
	}
	return string(array.local[offset : offset+length])
}

// Decode a variant-encoded value from the bytes starting at offset. Returns null if a valid variant can't be
// decoded.
func (array ByteArray) Decode(offset uintptr) any { //gd:PackedByteArray.decode_var
	if array.proxy != (gd.PackedByteArray{}) {
		return array.proxy.DecodeVar(gd.Int(offset), false).Interface()
	}
	val, _ := variant.UnmarshalAny(array.local[offset:])
	return val
}

// Decodes a size of a Variant from the bytes starting at offset. Requires at least 4 bytes of data
// starting at the offset, otherwise fails.
func (array ByteArray) DecodeSize(offset uintptr) uintptr { //gd:PackedByteArray.decode_var_size
	if array.proxy != (gd.PackedByteArray{}) {
		return uintptr(array.proxy.DecodeVarSize(gd.Int(offset), false))
	}
	val, _ := variant.UnmarshalSize(array.local[offset:])
	return val
}

// EncodeFloat64 encodes a 64-bit floating-point number as bytes at the index of offset bytes. The
// array must have at least 8 bytes of allocated space, starting at the offset.
func (array ByteArray) EncodeFloat64(offset uintptr, value float64) { //gd:PackedByteArray.encode_double
	if array.proxy != (gd.PackedByteArray{}) {
		array.proxy.EncodeDouble(gd.Int(offset), gd.Float(value))
		return
	}
	if offset+8 > uintptr(len(array.local)) {
		return
	}
	*(*float64)(unsafe.Pointer(&array.local[offset])) = value
}

// EncodeFloat32 encodes a 32-bit floating-point number as bytes at the index of offset bytes. The
// array must have at least 4 bytes of allocated space, starting at the offset.
func (array ByteArray) EncodeFloat32(offset uintptr, value float32) { //gd:PackedByteArray.encode_float PackedByteArray.encode_half
	if array.proxy != (gd.PackedByteArray{}) {
		array.proxy.EncodeFloat(gd.Int(offset), gd.Float(value))
		return
	}
	if offset+4 > uintptr(len(array.local)) {
		return
	}
	*(*float32)(unsafe.Pointer(&array.local[offset])) = value
}

// EncodeInt8 encodes a 8-bit signed integer number as bytes at the index of offset bytes. The
// array must have at least 1 byte of allocated space, starting at the offset.
func (array ByteArray) EncodeInt8(offset uintptr, value int8) { //gd:PackedByteArray.encode_s8
	if array.proxy != (gd.PackedByteArray{}) {
		array.proxy.EncodeS8(gd.Int(offset), gd.Int(value))
		return
	}
	if offset+1 > uintptr(len(array.local)) {
		return
	}
	array.local[offset] = byte(value)
}

// EncodeInt16 encodes a 16-bit signed integer number as bytes at the index of offset bytes. The
// array must have at least 2 bytes of allocated space, starting at the offset.
func (array ByteArray) EncodeInt16(offset uintptr, value int16) { //gd:PackedByteArray.encode_s16
	if array.proxy != (gd.PackedByteArray{}) {
		array.proxy.EncodeS16(gd.Int(offset), gd.Int(value))
		return
	}
	if offset+2 > uintptr(len(array.local)) {
		return
	}
	*(*int16)(unsafe.Pointer(&array.local[offset])) = value
}

// EncodeInt32 encodes a 32-bit signed integer number as bytes at the index of offset bytes. The
// array must have at least 4 bytes of allocated space, starting at the offset.
func (array ByteArray) EncodeInt32(offset uintptr, value int32) { //gd:PackedByteArray.encode_s32
	if array.proxy != (gd.PackedByteArray{}) {
		array.proxy.EncodeS32(gd.Int(offset), gd.Int(value))
		return
	}
	if offset+4 > uintptr(len(array.local)) {
		return
	}
	*(*int32)(unsafe.Pointer(&array.local[offset])) = value
}

// EncodeInt64 encodes a 64-bit signed integer number as bytes at the index of offset bytes. The
// array must have at least 8 bytes of allocated space, starting at the offset.
func (array ByteArray) EncodeInt64(offset uintptr, value int64) { //gd:PackedByteArray.encode_s64
	if array.proxy != (gd.PackedByteArray{}) {
		array.proxy.EncodeS64(gd.Int(offset), gd.Int(value))
		return
	}
	if offset+8 > uintptr(len(array.local)) {
		return
	}
	*(*int64)(unsafe.Pointer(&array.local[offset])) = value
}

// EncodeUint8 encodes a 8-bit unsigned integer number as bytes at the index of offset bytes. The
// array must have at least 1 byte of allocated space, starting at the offset.
func (array ByteArray) EncodeUint8(offset uintptr, value uint8) { //gd:PackedByteArray.encode_u8
	if array.proxy != (gd.PackedByteArray{}) {
		array.proxy.EncodeU8(gd.Int(offset), gd.Int(value))
		return
	}
	if offset+1 > uintptr(len(array.local)) {
		return
	}
	array.local[offset] = value
}

// EncodeUint16 encodes a 16-bit unsigned integer number as bytes at the index of offset bytes. The
// array must have at least 2 bytes of allocated space, starting at the offset.
func (array ByteArray) EncodeUint16(offset uintptr, value uint16) { //gd:PackedByteArray.encode_u16
	if array.proxy != (gd.PackedByteArray{}) {
		array.proxy.EncodeU16(gd.Int(offset), gd.Int(value))
		return
	}
	if offset+2 > uintptr(len(array.local)) {
		return
	}
	*(*uint16)(unsafe.Pointer(&array.local[offset])) = value
}

// EncodeUint32 encodes a 32-bit unsigned integer number as bytes at the index of offset bytes. The
// array must have at least 4 bytes of allocated space, starting at the offset.
func (array ByteArray) EncodeUint32(offset uintptr, value uint32) { //gd:PackedByteArray.encode_u32
	if array.proxy != (gd.PackedByteArray{}) {
		array.proxy.EncodeU32(gd.Int(offset), gd.Int(value))
		return
	}
	if offset+4 > uintptr(len(array.local)) {
		return
	}
	*(*uint32)(unsafe.Pointer(&array.local[offset])) = value
}

// EncodeUint64 encodes a 64-bit unsigned integer number as bytes at the index of offset bytes. The
// array must have at least 8 bytes of allocated space, starting at the offset.
func (array ByteArray) EncodeUint64(offset uintptr, value uint64) { //gd:PackedByteArray.encode_u64
	if array.proxy != (gd.PackedByteArray{}) {
		array.proxy.EncodeU64(gd.Int(offset), gd.Int(value))
		return
	}
	if offset+8 > uintptr(len(array.local)) {
		return
	}
	*(*uint64)(unsafe.Pointer(&array.local[offset])) = value
}

// Encode a variant-encoded value at the index of offset bytes. Sufficient space must be allocated,
// depending on the encoded variant's size
func (array ByteArray) Encode(offset uintptr, value any) { //gd:PackedByteArray.encode_var
	if array.proxy != (gd.PackedByteArray{}) {
		array.proxy.EncodeVar(gd.Int(offset), variant.New(value), false)
		return
	}
	buf, _ := variant.Marshal(value)
	if offset+uintptr(len(buf)) > uintptr(len(array.local)) {
		return
	}
	copy(array.local[offset:], buf)
}

// ToHex returns a hexadecimal representation of this array as a String.
func (array ByteArray) ToHex() string { //gd:PackedByteArray.hex_encode
	if array.proxy != (gd.PackedByteArray{}) {
		return array.proxy.HexEncode().String()
	}
	return hex.EncodeToString(array.local)
}

// HasVariantAt returns true if a valid Variant value can be decoded at the byte_offset.
// Returns false otherwise
func (array ByteArray) HasVariantAt(offset int) bool { //gd:PackedByteArray.has_encoded_var
	if array.proxy != (gd.PackedByteArray{}) {
		return bool(array.proxy.HasEncodedVar(gd.Int(offset), false))
	}
	_, err := variant.UnmarshalAny(array.local[offset:])
	return err == nil
}

// ToFloat32Array returns a copy of the data converted to a PackedFloat32Array, where each block of 4
// bytes has been converted to a 32-bit floating point number.
// The size of the input array must be a multiple of 4 (size of 32-bit float). The size of the new
// array will be byte_array.size() / 4.
// If the original data can't be converted to 32-bit floats, the resulting data is undefined.
func (array ByteArray) ToFloat32Array() Float32Array { //gd:PackedByteArray.to_float32_array
	var converted Float32Array
	if array.proxy != (gd.PackedByteArray{}) {
		converted.proxy = array.proxy.ToFloat32Array()
		return converted
	}
	if len(array.local)%4 != 0 {
		return Float32Array{}
	}
	converted.local = make([]float32, len(array.local)/4)
	for i := 0; i < len(array.local); i += 4 {
		converted.local[i/4] = math.Float32frombits(binary.LittleEndian.Uint32(array.local[i : i+4]))
	}
	return converted
}

// ToFloat64Array returns a copy of the data converted to a PackedFloat64Array, where each block of 8
// bytes has been converted to a 64-bit floating point number.
// The size of the input array must be a multiple of 8 (size of 64-bit float). The size of the new
// array will be byte_array.size() / 8.
// If the original data can't be converted to 64-bit floats, the resulting data is undefined.
func (array ByteArray) ToFloat64Array() Float64Array { //gd:PackedByteArray.to_float64_array
	var converted Float64Array
	if array.proxy != (gd.PackedByteArray{}) {
		converted.proxy = array.proxy.ToFloat64Array()
		return converted
	}
	if len(array.local)%8 != 0 {
		return Float64Array{}
	}
	converted.local = make([]float64, len(array.local)/8)
	for i := 0; i < len(array.local); i += 8 {
		converted.local[i/8] = math.Float64frombits(binary.LittleEndian.Uint64(array.local[i : i+8]))
	}
	return converted
}

// ToInt32Array returns a copy of the data converted to a PackedInt32Array, where each block of 4
// bytes has been converted to a signed 32-bit integer.
//
// The size of the input array must be a multiple of 4 (size of 32-bit integer). The size of the
// new array will be byte_array.size() / 4.
//
// If the original data can't be converted to signed 32-bit integers, the resulting data is undefined.
func (array ByteArray) ToInt32Array() Int32Array { //gd:PackedByteArray.to_int32_array
	var converted Int32Array
	if array.proxy != (gd.PackedByteArray{}) {
		converted.proxy = array.proxy.ToInt32Array()
		return converted
	}
	if len(array.local)%4 != 0 {
		return Int32Array{}
	}
	converted.local = make([]int32, len(array.local)/4)
	for i := 0; i < len(converted.local); i++ {
		converted.local[i] = *(*int32)(unsafe.Pointer(&array.local[i*4]))
	}
	return converted
}

// ToInt64Array returns a copy of the data converted to a PackedInt64Array, where each block of 8
// bytes has been converted to a signed 64-bit integer.
// The size of the input array must be a multiple of 8 (size of 64-bit integer). The size of the
// new array will be byte_array.size() / 8.
// If the original data can't be converted to signed 64-bit integers, the resulting data is undefined.
func (array ByteArray) ToInt64Array() Int64Array { //gd:PackedByteArray.to_int64_array
	var converted Int64Array
	if array.proxy != (gd.PackedByteArray{}) {
		converted.proxy = array.proxy.ToInt64Array()
		return converted
	}
	if len(array.local)%8 != 0 {
		return Int64Array{}
	}
	converted.local = make([]int64, len(array.local)/8)
	for i := 0; i < len(converted.local); i++ {
		converted.local[i] = *(*int64)(unsafe.Pointer(&array.local[i*8]))
	}
	return converted
}
