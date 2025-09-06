package gd

import (
	"strings"
	"unsafe"

	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdmemory"
	"graphics.gd/internal/pointers"
	StringType "graphics.gd/variant/String"
)

func (p *PackedFloat32Array) Pointer() *PackedFloat32Array { return p }
func (p *PackedFloat64Array) Pointer() *PackedFloat64Array { return p }
func (p *PackedInt32Array) Pointer() *PackedInt32Array     { return p }
func (p *PackedInt64Array) Pointer() *PackedInt64Array     { return p }
func (p *PackedColorArray) Pointer() *PackedColorArray     { return p }
func (p *PackedStringArray) Pointer() *PackedStringArray   { return p }
func (p *PackedVector2Array) Pointer() *PackedVector2Array { return p }
func (p *PackedVector3Array) Pointer() *PackedVector3Array { return p }
func (p *PackedVector4Array) Pointer() *PackedVector4Array { return p }

func (p PackedInt32Array) AsSlice() []int32 {
	return gdmemory.IntoSlice[int32](gdextension.Host.Packed.Int32s.Unsafe(pointers.Get(p)), int(p.Size()))
}
func (p PackedInt64Array) AsSlice() []int64 {
	return gdmemory.IntoSlice[int64](gdextension.Host.Packed.Int64s.Unsafe(pointers.Get(p)), int(p.Size()))
}
func (p PackedFloat32Array) AsSlice() []float32 {
	return gdmemory.IntoSlice[float32](gdextension.Host.Packed.Float32s.Unsafe(pointers.Get(p)), int(p.Size()))
}
func (p PackedFloat64Array) AsSlice() []float64 {
	return gdmemory.IntoSlice[float64](gdextension.Host.Packed.Float64s.Unsafe(pointers.Get(p)), int(p.Size()))
}
func (p PackedVector2Array) AsSlice() []Vector2 {
	return gdmemory.IntoSlice[Vector2](gdextension.Host.Packed.Vector2s.Unsafe(pointers.Get(p)), int(p.Size()))
}
func (p PackedVector3Array) AsSlice() []Vector3 {
	return gdmemory.IntoSlice[Vector3](gdextension.Host.Packed.Vector3s.Unsafe(pointers.Get(p)), int(p.Size()))
}
func (p PackedVector4Array) AsSlice() []Vector4 {
	return gdmemory.IntoSlice[Vector4](gdextension.Host.Packed.Vector4s.Unsafe(pointers.Get(p)), int(p.Size()))
}
func (p PackedColorArray) AsSlice() []Color {
	return gdmemory.IntoSlice[Color](gdextension.Host.Packed.Colors.Unsafe(pointers.Get(p)), int(p.Size()))
}
func (p PackedStringArray) Strings() []string {
	var s = make([]string, p.Size())
	for i := Int(0); i < p.Size(); i++ {
		s[i] = p.Index(i).String()
	}
	return s
}

func (p PackedByteArray) Index(idx Int) byte {
	return gdextension.Host.Packed.Bytes.Access(pointers.Get(p), int(idx))
}
func (p PackedByteArray) ToByteArray() PackedByteArray { return p.Duplicate() }
func (p PackedByteArray) SetIndex(idx Int, value byte) {
	gdmemory.Set[byte](gdextension.Host.Packed.Bytes.Unsafe(pointers.Get(p))+gdextension.Pointer(idx)*gdextension.Pointer(unsafe.Sizeof(value)), value)
}

// Bytes returns a copy of the byte array as a byte slice.
func (p PackedByteArray) Bytes() []byte {
	return gdmemory.IntoSlice[byte](gdextension.Host.Packed.Bytes.Unsafe(pointers.Get(p)), int(p.Size()))
}

func (p PackedByteArray) Len() int { return int(p.Size()) }
func (p PackedByteArray) Cap() int { return int(p.Size()) }

func (p PackedByteArray) Free() {
	if ptr, ok := pointers.End(p); ok {
		gdextension.Free(gdextension.TypePackedByteArray, &ptr)
	}
}

func (p PackedInt32Array) Index(idx Int) int32 {
	return gdextension.Host.Packed.Int32s.Access(pointers.Get(p), int(idx))
}
func (p PackedInt32Array) SetIndex(idx Int, value int32) {
	gdmemory.Set[int32](gdextension.Host.Packed.Int32s.Unsafe(pointers.Get(p))+gdextension.Pointer(idx)*gdextension.Pointer(unsafe.Sizeof(value)), value)
}

func (p PackedInt32Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		gdextension.Free(gdextension.TypePackedInt32Array, &ptr)
	}
}

func (p PackedInt32Array) Len() int { return int(p.Size()) }
func (p PackedInt32Array) Cap() int { return int(p.Size()) }

func (p PackedInt64Array) Index(idx Int) int64 {
	var i64 int64
	gdextension.Host.Packed.Int64s.Access(pointers.Get(p), int(idx), gdextension.CallReturns[int64](unsafe.Pointer(&i64)))
	return i64
}

func (p PackedInt64Array) SetIndex(idx Int, value int64) {
	gdmemory.Set[int64](gdextension.Host.Packed.Int64s.Unsafe(pointers.Get(p))+gdextension.Pointer(idx)*gdextension.Pointer(unsafe.Sizeof(value)), value)
}

func (p PackedInt64Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		gdextension.Free(gdextension.TypePackedInt64Array, &ptr)
	}
}

func (p PackedInt64Array) Len() int { return int(p.Size()) }
func (p PackedInt64Array) Cap() int { return int(p.Size()) }

func (p PackedFloat32Array) Index(idx Int) float32 {
	return gdextension.Host.Packed.Float32s.Access(pointers.Get(p), int(idx))
}

func (p PackedFloat32Array) SetIndex(idx Int, value float32) {
	gdmemory.Set[float32](gdextension.Host.Packed.Float32s.Unsafe(pointers.Get(p))+gdextension.Pointer(idx)*gdextension.Pointer(unsafe.Sizeof(value)), value)
}

func (p PackedFloat32Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		gdextension.Free(gdextension.TypePackedFloat32Array, &ptr)
	}
}

func (p PackedFloat32Array) Len() int { return int(p.Size()) }
func (p PackedFloat32Array) Cap() int { return int(p.Size()) }

func (p PackedFloat64Array) Index(idx Int) float64 {
	return gdextension.Host.Packed.Float64s.Access(pointers.Get(p), int(idx))
}

func (p PackedFloat64Array) SetIndex(idx Int, value float64) {
	gdmemory.Set[float64](gdextension.Host.Packed.Float64s.Unsafe(pointers.Get(p))+gdextension.Pointer(idx)*gdextension.Pointer(unsafe.Sizeof(value)), value)
}

func (p PackedFloat64Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		gdextension.Free(gdextension.TypePackedFloat64Array, &ptr)
	}
}

func (p PackedFloat64Array) Len() int { return int(p.Size()) }
func (p PackedFloat64Array) Cap() int { return int(p.Size()) }

func (p PackedStringArray) String() string {
	var builder strings.Builder
	builder.WriteString("[")
	size := int(p.Size())
	for i := 0; i < size; i++ {
		builder.WriteString(p.Index(Int(i)).String())
		if i < size-1 {
			builder.WriteString(" ")
		}
	}
	builder.WriteString("]")
	return builder.String()
}

func (p PackedStringArray) Index(idx Int) String {
	return pointers.Raw[String](gdextension.Host.Packed.Strings.Access(pointers.Get(p), int(idx))).Copy()
}

func (p PackedStringArray) SetIndex(idx Int, value String) {
	raw, _ := pointers.End(value.Copy())
	gdmemory.Set[gdextension.String](gdextension.Host.Packed.Strings.Unsafe(pointers.Get(p))+gdextension.Pointer(idx)*gdextension.Pointer(unsafe.Sizeof(gdextension.String{})), raw)
}

func (p PackedStringArray) AsSlice() []String {
	var ptr = gdextension.Host.Packed.Strings.Unsafe(pointers.Get(p))
	var slice = make([]String, p.Size())
	for i := range slice {
		slice[i] = pointers.Raw[String](gdmemory.Get[gdextension.String](ptr + gdextension.Pointer(i)*gdextension.Pointer(unsafe.Sizeof([1]gdextension.String{}[0])))).Copy()
	}
	return slice
}

func (p PackedStringArray) Free() {
	if ptr, ok := pointers.End(p); ok {
		gdextension.Free(gdextension.TypePackedStringArray, &ptr)
	}
}

func (p PackedVector2Array) Index(idx Int) Vector2 {
	return gdextension.IndexPacked(gdextension.Host.Packed.Vector2s.Access, pointers.Get(p), int(idx))
}

func (p PackedVector2Array) SetIndex(idx Int, value Vector2) {
	gdmemory.Set[Vector2](gdextension.Host.Packed.Vector2s.Unsafe(pointers.Get(p))+gdextension.Pointer(idx)*gdextension.Pointer(unsafe.Sizeof(value)), value)
}

func (p PackedVector2Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		gdextension.Free(gdextension.TypePackedVector2Array, &ptr)
	}
}

func (p PackedVector2Array) Len() int { return int(p.Size()) }
func (p PackedVector2Array) Cap() int { return int(p.Size()) }

func (p PackedVector3Array) Index(idx Int) Vector3 {
	return gdextension.IndexPacked(gdextension.Host.Packed.Vector3s.Access, pointers.Get(p), int(idx))
}

func (p PackedVector3Array) SetIndex(idx Int, value Vector3) {
	gdmemory.Set[Vector3](gdextension.Host.Packed.Vector3s.Unsafe(pointers.Get(p))+gdextension.Pointer(idx)*gdextension.Pointer(unsafe.Sizeof(value)), value)
}

func (p PackedVector3Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		gdextension.Free(gdextension.TypePackedVector3Array, &ptr)
	}
}

func (p PackedVector3Array) Len() int { return int(p.Size()) }
func (p PackedVector3Array) Cap() int { return int(p.Size()) }

func (p PackedVector4Array) Index(idx Int) Vector4 {
	return gdextension.IndexPacked(gdextension.Host.Packed.Vector4s.Access, pointers.Get(p), int(idx))
}

func (p PackedVector4Array) SetIndex(idx Int, value Vector4) {
	gdmemory.Set[Vector4](gdextension.Host.Packed.Vector4s.Unsafe(pointers.Get(p))+gdextension.Pointer(idx)*gdextension.Pointer(unsafe.Sizeof(value)), value)
}

func (p PackedVector4Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		gdextension.Free(gdextension.TypePackedVector4Array, &ptr)
	}
}

func (p PackedVector4Array) Len() int { return int(p.Size()) }
func (p PackedVector4Array) Cap() int { return int(p.Size()) }

func (p PackedColorArray) Index(idx Int) Color {
	return gdextension.IndexPacked(gdextension.Host.Packed.Colors.Access, pointers.Get(p), int(idx))
}

func (p PackedColorArray) SetIndex(idx Int, value Color) {
	gdmemory.Set[Color](gdextension.Host.Packed.Colors.Unsafe(pointers.Get(p))+gdextension.Pointer(idx)*gdextension.Pointer(unsafe.Sizeof(value)), value)
}

func (p PackedColorArray) Free() {
	if ptr, ok := pointers.End(p); ok {
		gdextension.Free(gdextension.TypePackedColorArray, &ptr)
	}
}

func (p PackedColorArray) Len() int { return int(p.Size()) }
func (p PackedColorArray) Cap() int { return int(p.Size()) }

func NewPackedByteArray() PackedByteArray {
	return pointers.New[PackedByteArray](gdextension.Make[gdextension.PackedArray[byte]](builtin.creation.PackedByteArray[0], 0, nil))
}

// PackedByteSlice returns a [PackedByteArray] from a byte slice.
func NewPackedByteSlice(data []byte) PackedByteArray {
	var array = NewPackedByteArray()
	array.Resize(Int(len(data)))
	gdmemory.LoadSlice[byte](gdextension.Host.Packed.Bytes.Unsafe(pointers.Get(array)), data)
	return array
}

func NewPackedColorArray() PackedColorArray {
	return pointers.New[PackedColorArray](gdextension.Make[gdextension.PackedArray[Color]](builtin.creation.PackedColorArray[0], 0, nil))
}

func NewPackedColorSlice(data []Color) PackedColorArray {
	var array = NewPackedColorArray()
	array.Resize(Int(len(data)))
	gdmemory.LoadSlice[Color](gdextension.Host.Packed.Colors.Unsafe(pointers.Get(array)), data)
	return array
}

func NewPackedFloat32Array() PackedFloat32Array {
	return pointers.New[PackedFloat32Array](gdextension.Make[gdextension.PackedArray[float32]](builtin.creation.PackedFloat32Array[0], 0, nil))
}

func NewPackedFloat32Slice(data []float32) PackedFloat32Array {
	var array = NewPackedFloat32Array()
	array.Resize(Int(len(data)))
	gdmemory.LoadSlice[float32](gdextension.Host.Packed.Float32s.Unsafe(pointers.Get(array)), data)
	return array
}

func NewPackedInt32Array() PackedInt32Array {
	return pointers.New[PackedInt32Array](gdextension.Make[gdextension.PackedArray[int32]](builtin.creation.PackedInt32Array[0], 0, nil))
}

func NewPackedInt32Slice(data []int32) PackedInt32Array {
	var array = NewPackedInt32Array()
	array.Resize(Int(len(data)))
	gdmemory.LoadSlice[int32](gdextension.Host.Packed.Int32s.Unsafe(pointers.Get(array)), data)
	return array
}

func NewPackedStringArray() PackedStringArray {
	return pointers.New[PackedStringArray](gdextension.Make[gdextension.PackedArray[gdextension.String]](builtin.creation.PackedStringArray[0], 0, nil))
}

func NewPackedStringSlice(data []string) PackedStringArray {
	var array = NewPackedStringArray()
	array.Resize(Int(len(data)))
	for i, str := range data {
		array.SetIndex(Int(i), NewString(str))
	}
	return array
}

func NewPackedReadableStringSlice(data []StringType.Readable) PackedStringArray {
	var array = NewPackedStringArray()
	array.Resize(Int(len(data)))
	for i, str := range data {
		_, raw := StringType.Proxy(str, StringCacheCheck, NewStringProxy)
		array.SetIndex(Int(i), pointers.Load[String](raw))
	}
	return array
}

func NewPackedVector2Array() PackedVector2Array {
	return pointers.New[PackedVector2Array](gdextension.Make[gdextension.PackedArray[Vector2]](builtin.creation.PackedVector2Array[0], 0, nil))
}

func NewPackedVector2Slice(data []Vector2) PackedVector2Array {
	var array = NewPackedVector2Array()
	array.Resize(Int(len(data)))
	gdmemory.LoadSlice[Vector2](gdextension.Host.Packed.Vector2s.Unsafe(pointers.Get(array)), data)
	return array
}

func NewPackedVector3Array() PackedVector3Array {
	return pointers.New[PackedVector3Array](gdextension.Make[gdextension.PackedArray[Vector3]](builtin.creation.PackedVector3Array[0], 0, nil))
}

func NewPackedVector3Slice(data []Vector3) PackedVector3Array {
	var array = NewPackedVector3Array()
	array.Resize(Int(len(data)))
	gdmemory.LoadSlice[Vector3](gdextension.Host.Packed.Vector3s.Unsafe(pointers.Get(array)), data)
	return array
}

func NewPackedVector4Array() PackedVector4Array {
	return pointers.New[PackedVector4Array](gdextension.Make[gdextension.PackedArray[Vector4]](builtin.creation.PackedVector4Array[0], 0, nil))
}

func NewPackedVector4Slice(data []Vector4) PackedVector4Array {
	var array = NewPackedVector4Array()
	array.Resize(Int(len(data)))
	gdmemory.LoadSlice[Vector4](gdextension.Host.Packed.Vector4s.Unsafe(pointers.Get(array)), data)
	return array
}

func NewPackedInt64Array() PackedInt64Array {
	return pointers.New[PackedInt64Array](gdextension.Make[gdextension.PackedArray[int64]](builtin.creation.PackedInt64Array[0], 0, nil))
}

func NewPackedInt64Slice(data []int64) PackedInt64Array {
	var array = NewPackedInt64Array()
	array.Resize(Int(len(data)))
	gdmemory.LoadSlice[int64](gdextension.Host.Packed.Int64s.Unsafe(pointers.Get(array)), data)
	return array
}

func NewPackedFloat64Array() PackedFloat64Array {
	return pointers.New[PackedFloat64Array](gdextension.Make[gdextension.PackedArray[float64]](builtin.creation.PackedFloat64Array[0], 0, nil))
}

func NewPackedFloat64Slice(data []float64) PackedFloat64Array {
	var array = NewPackedFloat64Array()
	array.Resize(Int(len(data)))
	gdmemory.LoadSlice[float64](gdextension.Host.Packed.Float64s.Unsafe(pointers.Get(array)), data)
	return array
}
