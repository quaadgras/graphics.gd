package Packed

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"encoding/binary"
	"encoding/hex"
	"iter"
	"math"
	"unsafe"

	"graphics.gd/variant"
	GenericArray "graphics.gd/variant/Array"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector4"
)

type CompressionMode int

const (
	CompressionFastLZ CompressionMode = iota
	CompressionDeflate
	CompressionZstandard // not implemented
	CompressionGzip
	CompressionBrotli // not implemented
)

// Bytes provides additional methods for working with arrays of bytes.
type Bytes struct {
	Array[byte]
	order binary.ByteOrder
}

func BytesFrom(data ...byte) Bytes {
	return Bytes{Array: Array[byte](GenericArray.New(data...))}
}

// Index returns the element at the given index.
func (array Bytes) Index(idx int) byte { return (GenericArray.Contains[byte])(array.Array).Index(idx) } //gd:PackedByteArray.get

// Iter returns an iterator for the array.
func (array Bytes) Iter() iter.Seq2[int, byte] {
	return (GenericArray.Contains[byte])(array.Array).Iter()
}

// Append an element to the end of the array.
func (array *Bytes) Append(value byte) { (*GenericArray.Contains[byte])(&array.Array).Append(value) } //gd:PackedByteArray.push_back PackedByteArray.append

// AppendArray appends all elements of another array to the end of this array.
func (array *Bytes) AppendArray(other Bytes) { //gd:PackedByteArray.append_array
	(*GenericArray.Contains[byte])(&array.Array).AppendArray((GenericArray.Contains[byte])(other.Array))
}

// Clear clears the array. bytehis is equivalent to using resize with a size of 0.
func (array Bytes) Clear() { GenericArray.Clear((GenericArray.Contains[byte])(array.Array)) } //gd:PackedByteArray.clear

// Count returns the number of times an element is in the array.
func (array Bytes) Count(value byte) int { //gd:PackedByteArray.count
	return GenericArray.Count((GenericArray.Contains[byte])(array.Array), value)
}

// Duplicate creates a copy of the array, and returns it.
func (array Bytes) Duplicate() Bytes { //gd:PackedByteArray.duplicate
	return Bytes{Array: Array[byte](GenericArray.Duplicate((GenericArray.Contains[byte])(array.Array))), order: array.order}
}

// Fill assigns the given value to all elements in the array. bytehis can typically be used
// together with resize to create an array with a given size and initialized elements.
func (array Bytes) Fill(value byte) { //gd:PackedByteArray.fill
	GenericArray.Fill((GenericArray.Contains[byte])(array.Array), value)
}

// Find searches the array for a value and returns its index or -1 if not found.
func (array Bytes) Find(value byte) int { //gd:PackedByteArray.find
	return GenericArray.Find((GenericArray.Contains[byte])(array.Array), value)
}

// Has returns true if the array contains the given value.
func (array Bytes) Has(value byte) bool { //gd:PackedByteArray.has
	return GenericArray.Has(value, (GenericArray.Contains[byte])(array.Array))
}

// Insert inserts a new element at a given position in the array. bytehe position must be
// valid, or at the end of the array (idx == size()).
func (array *Bytes) Insert(idx int, value byte) { //gd:PackedByteArray.insert
	(*GenericArray.Contains[byte])(&array.Array).Insert(idx, value)
}

// IsEmpty Returns true if the array is empty.
func (array Bytes) IsEmpty() bool { //gd:PackedByteArray.is_empty
	return GenericArray.IsEmpty((GenericArray.Contains[byte])(array.Array))
}

// RemoveAt removes an element from the array by index.
func (array Bytes) RemoveAt(idx int) { //gd:PackedByteArray.remove_at
	GenericArray.Remove((GenericArray.Contains[byte])(array.Array), idx)
}

// Resize sets the size of the array. If the array is grown, reserves elements at the end of the array.
// If the array is shrunk, truncates the array to the new size. Calling resize once and assigning the
// new values is faster than adding new elements one by one.
func (array *Bytes) Resize(size int) { (*GenericArray.Contains[byte])(&array.Array).Resize(size) } //gd:PackedByteArray.resize

// Reverse reverses the order of the elements in the array.
func (array Bytes) Reverse() { GenericArray.Reverse((GenericArray.Contains[byte])(array.Array)) } //gd:PackedByteArray.reverse

// FindLast searches the array in reverse order for a value and returns its index or -1 if not found.
func (array Bytes) FindLast(value byte) int { //gd:PackedByteArray.rfind
	return GenericArray.FindLast((GenericArray.Contains[byte])(array.Array), value)
}

// SetIndex sets the value of the element at the given index.
func (array *Bytes) SetIndex(idx int, value byte) { //gd:PackedByteArray.set
	(*GenericArray.Contains[byte])(&array.Array).SetIndex(idx, value)
}

// Len returns the number of elements in the array.
func (array Bytes) Len() int { return (GenericArray.Contains[byte])(array.Array).Len() } //gd:PackedByteArray.size

// Slice returns a slice of the array from the given begin index to the given end index.
func (array Bytes) Slice(begin, end int) Bytes { //gd:PackedByteArray.slice
	return Bytes{Array: Array[byte](GenericArray.Slice((GenericArray.Contains[byte])(array.Array), begin, end)), order: array.order}
}

// Sort sorts the array in ascending order. The final order is dependent on the
// "less than" (<) comparison between elements.
func (array Bytes) Sort() { GenericArray.Sort((GenericArray.Contains[byte])(array.Array)) } //gd:PackedByteArray.sort

// BinarySearch returns the index of value in the sorted array. If it cannot be found, returns where value should
// be inserted to keep the array sorted. The algorithm used is binary search. The returned index comes before all
// existing elements equal to value in the array.
//
// Note: Calling BinarySearch on an unsorted array will result in unexpected behavior. Use [Sort] before calling
// this method.
func (array Bytes) BinarySearch(value byte, before bool) int { //gd:PackedByteArray.bsearch
	return GenericArray.BinarySearch((GenericArray.Contains[byte])(array.Array), value, before)
}

// Erase Removes the first occurrence of a value from the array and returns true. If the value does not exist in the array, nothing
// happens and false is returned. To remove an element by index, use [Array.RemoveAt] instead.
func (array Bytes) Erase(value byte) bool { //gd:PackedByteArray.erase
	for i := 0; i < array.Len(); i++ {
		if array.Index(i) == value {
			array.RemoveAt(i)
			return true
		}
	}
	return false
}

// Bytes returns the underlying data in the array as a slice of bytes.
func (array Bytes) Bytes() []byte {
	return (GenericArray.Contains[byte])(array.Array).Slice()
}

// ToHex returns a hexadecimal representation of this array as a String.
func (array Bytes) ToHex() string { //gd:PackedByteArray.hex_encode
	return hex.EncodeToString(array.Bytes())
}

// Compress returns a new PackedByteArray with the data compressed. Set the compression mode using one of
// [CompressionMode]'s constants.
func (array Bytes) Compress(mode CompressionMode) Bytes { //gd:PackedByteArray.compress
	switch mode {
	case CompressionFastLZ:
		out := make([]byte, int(float64(array.Len())*1.05))
		out = out[0:fastlz_compress(array.Bytes(), out)]
		return Bytes{Array: Array[byte](GenericArray.New(out...)), order: array.order}
	case CompressionDeflate:
		var out bytes.Buffer
		w, _ := flate.NewWriter(&out, 1)
		w.Write(array.Bytes())
		w.Close()
		return Bytes{Array: Array[byte](GenericArray.New(out.Bytes()...)), order: array.order}
	case CompressionZstandard:
		panic("not implemented")
	case CompressionGzip:
		var out bytes.Buffer
		w := gzip.NewWriter(&out)
		w.Write(array.Bytes())
		w.Close()
		return Bytes{Array: Array[byte](GenericArray.New(out.Bytes()...)), order: array.order}
	case CompressionBrotli:
		panic("not implemented")
	default:
		return array
	}
}

// Decompress returns a new PackedByteArray with the data decompressed. Set buffer_size to the size of the
// uncompressed data. Set the compression mode using one of CompressionMode's constants.
func (array Bytes) DecompressSize(buffer_size int, mode CompressionMode) Bytes { //gd:PackedByteArray.decompress
	out := make([]byte, buffer_size)
	switch mode {
	case CompressionFastLZ:
		out = out[0:fastlz_decompress(array.Bytes(), out)]
	case CompressionDeflate:
		r := flate.NewReader(bytes.NewReader(array.Bytes()))
		n, _ := r.Read(out)
		out = out[0:n]
		r.Close()
	case CompressionZstandard:
		panic("not implemented")
	case CompressionGzip:
		r, _ := gzip.NewReader(bytes.NewReader(array.Bytes()))
		n, _ := r.Read(out)
		out = out[0:n]
		r.Close()
	case CompressionBrotli:
		panic("not implemented")
	default:
		panic("not implemented")
	}
	return Bytes{Array: Array[byte](GenericArray.New(out...)), order: array.order}
}

// DecompressUpto returns a new PackedByteArray with the data decompressed. Set the compression mode using
// one of CompressionMode's constants. This method only accepts brotli, gzip, and deflate compression modes.
//
// This method is potentially slower than decompress, as it may have to re-allocate its output buffer multiple
// times while decompressing, whereas decompress knows it's output buffer size from the beginning.
func (array Bytes) DecompressUpto(max_output_size int, mode CompressionMode) Bytes { //gd:PackedByteArray.decompress_dynamic
	return array.DecompressSize(max_output_size, mode)
}

// DecodeFloat64 decodes a 64-bit floating-point number from the bytes starting at offset. Fails if the
// byte count is insufficient. Returns 0.0 if a valid number can't be decoded.
func (array Bytes) DecodeFloat64(offset uintptr) float64 { //gd:PackedByteArray.decode_double
	if offset+8 > uintptr(array.Len()) {
		return 0.0
	}
	var buf [8]byte
	for i := range 8 {
		buf[i] = array.Index(int(offset) + i)
	}
	return math.Float64frombits(array.ByteOrder().Uint64(buf[:]))
}

// DecodeFloat32 decodes a 32-bit floating-point number from the bytes starting at offset. Fails if the byte
// count is insufficient. Returns 0.0 if a valid number can't be decoded.
func (array Bytes) DecodeFloat32(offset uintptr) float32 { //gd:PackedByteArray.decode_float PackedByteArray.decode_half
	if offset+4 > uintptr(array.Len()) {
		return 0.0
	}
	var buf [4]byte
	for i := range 4 {
		buf[i] = array.Index(int(offset) + i)
	}
	return math.Float32frombits(array.ByteOrder().Uint32(buf[:]))
}

// DecodeInt8 decodes a 8-bit signed integer number from the bytes starting at offset. Fails if the byte count is
// insufficient. Returns 0 if a valid number can't be decoded.
func (array Bytes) DecodeInt8(offset uintptr) int8 { //gd:PackedByteArray.decode_s8
	if offset+1 > uintptr(array.Len()) {
		return 0
	}
	b := array.Index(int(offset))
	return *(*int8)(unsafe.Pointer(&b))
}

// DecodeInt16 decodes a 16-bit signed integer number from the bytes starting at offset. Fails if the byte count
// is insufficient. Returns 0 if a valid number can't be decoded.
func (array Bytes) DecodeInt16(offset uintptr) int16 { //gd:PackedByteArray.decode_s16
	if offset+2 > uintptr(array.Len()) {
		return 0
	}
	var buf [2]byte
	for i := range 2 {
		buf[i] = array.Index(int(offset) + i)
	}
	return int16(array.ByteOrder().Uint16(buf[:]))
}

// DecodeInt32 decodes a 32-bit signed integer number from the bytes starting at offset. Fails if the byte count
// is insufficient. Returns 0 if a valid number can't be decoded.
func (array Bytes) DecodeInt32(offset uintptr) int32 { //gd:PackedByteArray.decode_s32
	if offset+4 > uintptr(array.Len()) {
		return 0
	}
	var buf [4]byte
	for i := range 4 {
		buf[i] = array.Index(int(offset) + i)
	}
	return int32(array.ByteOrder().Uint32(buf[:]))
}

// DecodeInt64 decodes a 64-bit signed integer number from the bytes starting at offset. Fails if the byte count
// is insufficient. Returns 0 if a valid number can't be decoded.
func (array Bytes) DecodeInt64(offset uintptr) int64 { //gd:PackedByteArray.decode_s64
	if offset+8 > uintptr(array.Len()) {
		return 0
	}
	var buf [8]byte
	for i := range 8 {
		buf[i] = array.Index(int(offset) + i)
	}
	return int64(array.ByteOrder().Uint64(buf[:]))
}

// DecodeUint8 decodes a 8-bit unsigned integer number from the bytes starting at offset. Fails if the byte count
// is insufficient. Returns 0 if a valid number can't be decoded.
func (array Bytes) DecodeUint8(offset uintptr) uint8 { //gd:PackedByteArray.decode_u8
	if offset+1 > uintptr(array.Len()) {
		return 0
	}
	return array.Index(int(offset))
}

// DecodeUint16 decodes a 16-bit unsigned integer number from the bytes starting at offset. Fails if the byte count
// is insufficient. Returns 0 if a valid number can't be decoded.
func (array Bytes) DecodeUint16(offset uintptr) uint16 { //gd:PackedByteArray.decode_u16
	if offset+2 > uintptr(array.Len()) {
		return 0
	}
	var buf [2]byte
	for i := range 2 {
		buf[i] = array.Index(int(offset) + i)
	}
	return array.ByteOrder().Uint16(buf[:])
}

// DecodeUint32 decodes a 32-bit unsigned integer number from the bytes starting at offset. Fails if the byte count
// is insufficient. Returns 0 if a valid number can't be decoded.
func (array Bytes) DecodeUint32(offset uintptr) uint32 { //gd:PackedByteArray.decode_u32
	if offset+4 > uintptr(array.Len()) {
		return 0
	}
	var buf [4]byte
	for i := range 4 {
		buf[i] = array.Index(int(offset) + i)
	}
	return array.ByteOrder().Uint32(buf[:])
}

// DecodeUint64 decodes a 64-bit unsigned integer number from the bytes starting at offset. Fails if the byte count
// is insufficient. Returns 0 if a valid number can't be decoded.
func (array Bytes) DecodeUint64(offset uintptr) uint64 { //gd:PackedByteArray.decode_u64
	if offset+8 > uintptr(array.Len()) {
		return 0
	}
	var buf [8]byte
	for i := range 8 {
		buf[i] = array.Index(int(offset) + i)
	}
	return array.ByteOrder().Uint64(buf[:])
}

// Decode a variant-encoded value from the bytes starting at offset. Returns null if a valid variant can't be
// decoded.
func (array Bytes) Decode(offset uintptr) any { //gd:PackedByteArray.decode_var
	val, _ := variant.UnmarshalAny(array.Bytes()[offset:])
	return val
}

// Decodes a size of a Variant from the bytes starting at offset. Requires at least 4 bytes of data
// starting at the offset, otherwise fails.
func (array Bytes) DecodeSize(offset uintptr) uintptr { //gd:PackedByteArray.decode_var_size
	val, _ := variant.UnmarshalSize(array.Bytes()[offset:])
	return val
}

// EncodeFloat64 encodes a 64-bit floating-point number as bytes at the index of offset bytes. The
// array must have at least 8 bytes of allocated space, starting at the offset.
func (array Bytes) EncodeFloat64(offset uintptr, value float64) { //gd:PackedByteArray.encode_double
	if offset+8 > uintptr(array.Len()) {
		return
	}
	var buf [8]byte
	array.ByteOrder().PutUint64(buf[:], math.Float64bits(value))
	for i := range buf {
		array.SetIndex(int(offset)+i, buf[i])
	}
}

// EncodeFloat32 encodes a 32-bit floating-point number as bytes at the index of offset bytes. The
// array must have at least 4 bytes of allocated space, starting at the offset.
func (array Bytes) EncodeFloat32(offset uintptr, value float32) { //gd:PackedByteArray.encode_float PackedByteArray.encode_half
	if offset+4 > uintptr(array.Len()) {
		return
	}
	var buf [4]byte
	array.ByteOrder().PutUint32(buf[:], math.Float32bits(value))
	for i := range buf {
		array.SetIndex(int(offset)+i, buf[i])
	}
}

// EncodeInt8 encodes a 8-bit signed integer number as bytes at the index of offset bytes. The
// array must have at least 1 byte of allocated space, starting at the offset.
func (array Bytes) EncodeInt8(offset uintptr, value int8) { //gd:PackedByteArray.encode_s8
	if offset+1 > uintptr(array.Len()) {
		return
	}
	array.SetIndex(int(offset), *(*byte)(unsafe.Pointer(&value)))
}

// EncodeInt16 encodes a 16-bit signed integer number as bytes at the index of offset bytes. The
// array must have at least 2 bytes of allocated space, starting at the offset.
func (array Bytes) EncodeInt16(offset uintptr, value int16) { //gd:PackedByteArray.encode_s16
	if offset+2 > uintptr(array.Len()) {
		return
	}
	var buf [2]byte
	array.ByteOrder().PutUint16(buf[:], *(*uint16)(unsafe.Pointer(&value)))
	for i := range buf {
		array.SetIndex(int(offset)+i, buf[i])
	}
}

// EncodeInt32 encodes a 32-bit signed integer number as bytes at the index of offset bytes. The
// array must have at least 4 bytes of allocated space, starting at the offset.
func (array Bytes) EncodeInt32(offset uintptr, value int32) { //gd:PackedByteArray.encode_s32
	if offset+4 > uintptr(array.Len()) {
		return
	}
	var buf [4]byte
	array.ByteOrder().PutUint32(buf[:], *(*uint32)(unsafe.Pointer(&value)))
	for i := range buf {
		array.SetIndex(int(offset)+i, buf[i])
	}
}

// EncodeInt64 encodes a 64-bit signed integer number as bytes at the index of offset bytes. The
// array must have at least 8 bytes of allocated space, starting at the offset.
func (array Bytes) EncodeInt64(offset uintptr, value int64) { //gd:PackedByteArray.encode_s64
	if offset+8 > uintptr(array.Len()) {
		return
	}
	var buf [8]byte
	array.ByteOrder().PutUint64(buf[:], *(*uint64)(unsafe.Pointer(&value)))
	for i := range buf {
		array.SetIndex(int(offset)+i, buf[i])
	}
}

// EncodeUint8 encodes a 8-bit unsigned integer number as bytes at the index of offset bytes. The
// array must have at least 1 byte of allocated space, starting at the offset.
func (array Bytes) EncodeUint8(offset uintptr, value uint8) { //gd:PackedByteArray.encode_u8
	if offset+1 > uintptr(array.Len()) {
		return
	}
	array.SetIndex(int(offset), value)
}

// EncodeUint16 encodes a 16-bit unsigned integer number as bytes at the index of offset bytes. The
// array must have at least 2 bytes of allocated space, starting at the offset.
func (array Bytes) EncodeUint16(offset uintptr, value uint16) { //gd:PackedByteArray.encode_u16
	if offset+2 > uintptr(array.Len()) {
		return
	}
	var buf [2]byte
	array.ByteOrder().PutUint16(buf[:], value)
	for i := range buf {
		array.SetIndex(int(offset)+i, buf[i])
	}
}

// EncodeUint32 encodes a 32-bit unsigned integer number as bytes at the index of offset bytes. The
// array must have at least 4 bytes of allocated space, starting at the offset.
func (array Bytes) EncodeUint32(offset uintptr, value uint32) { //gd:PackedByteArray.encode_u32
	if offset+4 > uintptr(array.Len()) {
		return
	}
	var buf [4]byte
	array.ByteOrder().PutUint32(buf[:], value)
	for i := range buf {
		array.SetIndex(int(offset)+i, buf[i])
	}
}

// EncodeUint64 encodes a 64-bit unsigned integer number as bytes at the index of offset bytes. The
// array must have at least 8 bytes of allocated space, starting at the offset.
func (array Bytes) EncodeUint64(offset uintptr, value uint64) { //gd:PackedByteArray.encode_u64
	if offset+8 > uintptr(array.Len()) {
		return
	}
	var buf [8]byte
	array.ByteOrder().PutUint64(buf[:], value)
	for i := range buf {
		array.SetIndex(int(offset)+i, buf[i])
	}
}

// Encode a variant-encoded value at the index of offset bytes. Sufficient space must be allocated,
// depending on the encoded variant's size
func (array Bytes) Encode(offset uintptr, value any) { //gd:PackedByteArray.encode_var
	buf, _ := variant.Marshal(value)
	if offset+uintptr(len(buf)) > uintptr(array.Len()) {
		return
	}
	for i, b := range buf {
		array.SetIndex(int(offset)+i, b)
	}
}

// HasVariantAt returns true if a valid Variant value can be decoded at the byte_offset.
// Returns false otherwise
func (array Bytes) HasVariantAt(offset int) bool { //gd:PackedByteArray.has_encoded_var
	_, err := variant.UnmarshalAny(array.Bytes()[offset:])
	return err == nil
}

// ToFloat32s returns a copy of the data converted to a Array[float32], where each block of 4
// bytes has been converted to a 32-bit floating point number.
// The size of the input array must be a multiple of 4 (size of 32-bit float). The size of the new
// array will be byte_array.size() / 4.
// If the original data can't be converted to 32-bit floats, the resulting data is undefined.
func (array Bytes) ToFloat32s() Array[float32] { //gd:PackedByteArray.to_float32_array
	if array.Len()%4 != 0 {
		return Array[float32]{}
	}
	var converted Array[float32]
	converted.Resize(array.Len() / 4)
	for i := 0; i < array.Len(); i += 4 {
		var buf [4]byte
		for j := range 4 {
			buf[j] = array.Index(i + j)
		}
		converted.SetIndex(i/4, math.Float32frombits(array.ByteOrder().Uint32(buf[:])))
	}
	return converted
}

// ToFloat64s returns a copy of the data converted to a Array[float64], where each block of 8
// bytes has been converted to a 64-bit floating point number.
// The size of the input array must be a multiple of 8 (size of 64-bit float). The size of the new
// array will be byte_array.size() / 8.
// If the original data can't be converted to 64-bit floats, the resulting data is undefined.
func (array Bytes) ToFloat64s() Array[float64] { //gd:PackedByteArray.to_float64_array
	if array.Len()%8 != 0 {
		return Array[float64]{}
	}
	var converted Array[float64]
	converted.Resize(array.Len() / 8)
	for i := 0; i < array.Len(); i += 8 {
		var buf [8]byte
		for j := range 8 {
			buf[j] = array.Index(i + j)
		}
		converted.SetIndex(i/8, math.Float64frombits(array.ByteOrder().Uint64(buf[:])))
	}
	return converted
}

// ToInt32s returns a copy of the data converted to a Array[int32], where each block of 4
// bytes has been converted to a signed 32-bit integer.
//
// The size of the input array must be a multiple of 4 (size of 32-bit integer). The size of the
// new array will be byte_array.size() / 4.
//
// If the original data can't be converted to signed 32-bit integers, the resulting data is undefined.
func (array Bytes) ToInt32s() Array[int32] { //gd:PackedByteArray.to_int32_array
	if array.Len()%4 != 0 {
		return Array[int32]{}
	}
	var converted Array[int32]
	converted.Resize(array.Len() / 4)
	for i := 0; i < array.Len(); i += 4 {
		var buf [4]byte
		for j := range 4 {
			buf[j] = array.Index(i + j)
		}
		var u32 uint32
		converted.SetIndex(i/4, *(*int32)(unsafe.Pointer(&u32)))
	}
	return converted
}

// ToInt64s returns a copy of the data converted to a Array[int64], where each block of 8
// bytes has been converted to a signed 64-bit integer.
// The size of the input array must be a multiple of 8 (size of 64-bit integer). The size of the
// new array will be byte_array.size() / 8.
// If the original data can't be converted to signed 64-bit integers, the resulting data is undefined.
func (array Bytes) ToInt64s() Array[int64] { //gd:PackedByteArray.to_int64_array
	if array.Len()%8 != 0 {
		return Array[int64]{}
	}
	var converted Array[int64]
	converted.Resize(array.Len() / 8)
	for i := 0; i < array.Len(); i += 8 {
		var buf [8]byte
		for j := range 8 {
			buf[j] = array.Index(i + j)
		}
		var u64 uint64 = array.ByteOrder().Uint64(buf[:])
		converted.SetIndex(i/8, *(*int64)(unsafe.Pointer(&u64)))
	}
	return converted
}

// ToVector2s returns a copy of the data converted to a [Array[Vector2.XY]], where each block of 8 bytes
// or 16 bytes (32-bit or 64-bit)  has been converted to a [Vector2.XY]
// The size of the input array must be a multiple of 8 (size of 64-bit integer). The size of the
// new array will be byte_array.size() / (8 or 16)
// If the original data can't be converted to Vector2s, the resulting data is undefined.
func (array Bytes) ToVector2s() Array[Vector2.XY] { //gd:PackedByteArray.to_vector2_array
	if array.Len()%8 != 0 {
		return Array[Vector2.XY]{}
	}
	var converted Array[Vector2.XY]
	converted.Resize(array.Len() / 8)
	for i := 0; i < array.Len(); i += 8 {
		var buf [8]byte
		for j := range 8 {
			buf[j] = array.Index(i + j)
		}
		x := math.Float32frombits(array.ByteOrder().Uint32(buf[0:4]))
		y := math.Float32frombits(array.ByteOrder().Uint32(buf[4:8]))
		converted.SetIndex(i/8, Vector2.XY{X: x, Y: y})
	}
	return converted
}

// ToVector3s returns a copy of the data converted to a [Array[Vector3.XYZ]], where each block of 12 bytes
// or 24 bytes (32-bit or 64-bit)  has been converted to a [Vector3.XYZ]
// The size of the input array must be a multiple of 12 (size of 3 32-bit floats). The size of the
// new array will be byte_array.size() / (12 or 24)
// If the original data can't be converted to Vector3s, the resulting data is undefined.
func (array Bytes) ToVector3s() Array[Vector3.XYZ] { //gd:PackedByteArray.to_vector3_array
	if array.Len()%12 != 0 {
		return Array[Vector3.XYZ]{}
	}
	var converted Array[Vector3.XYZ]
	converted.Resize(array.Len() / 12)
	for i := 0; i < array.Len(); i += 12 {
		var buf [12]byte
		for j := range 12 {
			buf[j] = array.Index(i + j)
		}
		x := math.Float32frombits(array.ByteOrder().Uint32(buf[0:4]))
		y := math.Float32frombits(array.ByteOrder().Uint32(buf[4:8]))
		z := math.Float32frombits(array.ByteOrder().Uint32(buf[8:12]))
		converted.SetIndex(i/12, Vector3.XYZ{X: x, Y: y, Z: z}) // should be Vector3 but not implemented yet
	}
	return converted
}

// ToVector4s returns a copy of the data converted to a [Array[Vector4.XYZW]], where each block of 16 bytes
// or 32 bytes (32-bit or 64-bit)  has been converted to a [Vector4.XYZW]
// The size of the input array must be a multiple of 16 (size of 4 32-bit floats). The size of the
// new array will be byte_array.size() / (16 or 32)
// If the original data can't be converted to Vector4s, the resulting data is undefined.
func (array Bytes) ToVector4s() Array[Vector4.XYZW] { //gd:PackedByteArray.to_vector4_array
	if array.Len()%16 != 0 {
		return Array[Vector4.XYZW]{}
	}
	var converted Array[Vector4.XYZW]
	converted.Resize(array.Len() / 16)
	for i := 0; i < array.Len(); i += 16 {
		var buf [16]byte
		for j := range 16 {
			buf[j] = array.Index(i + j)
		}
		x := math.Float32frombits(array.ByteOrder().Uint32(buf[0:4]))
		y := math.Float32frombits(array.ByteOrder().Uint32(buf[4:8]))
		z := math.Float32frombits(array.ByteOrder().Uint32(buf[8:12]))
		w := math.Float32frombits(array.ByteOrder().Uint32(buf[12:16]))
		converted.SetIndex(i/16, Vector4.XYZW{X: x, Y: y, Z: z, W: w})
	}
	return converted
}

// ToColors returns a copy of the data converted to a Array[Color.RGBA], where each block of 16
// bytes has been converted to a Color.RGBA
//
// The size of the input array must be a multiple of 16 (size of Color.RGBA). The size of the new
// array will be byte_array.size() / 16.
// If the original data can't be converted to Colors, the resulting data is undefined.
func (array Bytes) ToColors() Array[Color.RGBA] { //gd:PackedByteArray.to_color_array
	if array.Len()%16 != 0 {
		return Array[Color.RGBA]{}
	}
	var converted Array[Color.RGBA]
	converted.Resize(array.Len() / 16)
	for i := 0; i < array.Len(); i += 16 {
		var buf [16]byte
		for j := range 16 {
			buf[j] = array.Index(i + j)
		}
		r := math.Float32frombits(array.ByteOrder().Uint32(buf[0:4]))
		g := math.Float32frombits(array.ByteOrder().Uint32(buf[4:8]))
		b := math.Float32frombits(array.ByteOrder().Uint32(buf[8:12]))
		a := math.Float32frombits(array.ByteOrder().Uint32(buf[12:16]))
		converted.SetIndex(i/16, Color.RGBA{R: r, G: g, B: b, A: a})
	}
	return converted
}

// ByteOrder returns the current byte order used for encoding and decoding multi-byte values.
func (array Bytes) ByteOrder() binary.ByteOrder {
	if array.order == nil {
		return binary.LittleEndian
	}
	return array.order
}

// SetByteOrder sets the byte order used for encoding and decoding multi-byte values (big endian by default).
func (array Bytes) SetByteOrder(order binary.ByteOrder) { //gd:PackedByteArray.bswap16 PackedByteArray.bswap32 PackedByteArray.bswap64
	array.order = order
}
