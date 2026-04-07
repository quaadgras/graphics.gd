package gd

import (
	"strings"

	gdunsafe "graphics.gd"
	"graphics.gd/internal/gdextension"
	"graphics.gd/internal/gdmemory"
	"graphics.gd/internal/noescape"
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
	return gdmemory.IntoSlice[int32](gdunsafe.PackedArray[int32](pointers.Get(p)).Access(0), int(p.Size()))
}
func (p PackedInt64Array) AsSlice() []int64 {
	return gdmemory.IntoSlice[int64](gdunsafe.PackedArray[int64](pointers.Get(p)).Access(0), int(p.Size()))
}
func (p PackedFloat32Array) AsSlice() []float32 {
	return gdmemory.IntoSlice[float32](gdunsafe.PackedArray[float32](pointers.Get(p)).Access(0), int(p.Size()))
}
func (p PackedFloat64Array) AsSlice() []float64 {
	return gdmemory.IntoSlice[float64](gdunsafe.PackedArray[float64](pointers.Get(p)).Access(0), int(p.Size()))
}
func (p PackedVector2Array) AsSlice() []Vector2 {
	return gdmemory.IntoSlice[Vector2](gdunsafe.PackedArray[Vector2](pointers.Get(p)).Access(0), int(p.Size()))
}
func (p PackedVector3Array) AsSlice() []Vector3 {
	return gdmemory.IntoSlice[Vector3](gdunsafe.PackedArray[Vector3](pointers.Get(p)).Access(0), int(p.Size()))
}
func (p PackedVector4Array) AsSlice() []Vector4 {
	return gdmemory.IntoSlice[Vector4](gdunsafe.PackedArray[Vector4](pointers.Get(p)).Access(0), int(p.Size()))
}
func (p PackedColorArray) AsSlice() []Color {
	return gdmemory.IntoSlice[Color](gdunsafe.PackedArray[Color](pointers.Get(p)).Access(0), int(p.Size()))
}
func (p PackedStringArray) Strings() []string {
	var s = make([]string, p.Size())
	for i := Int(0); i < p.Size(); i++ {
		s[i] = p.Index(i).String()
	}
	return s
}

func (p PackedByteArray) Index(idx Int) byte {
	return gdunsafe.PackedArray[byte](pointers.Get(p)).Access(idx).Get()
}
func (p PackedByteArray) ToByteArray() PackedByteArray { return p.Duplicate() }
func (p PackedByteArray) SetIndex(idx Int, value byte) {
	gdunsafe.PackedArray[byte](pointers.Get(p)).Modify(idx).Set(value)
}

// Bytes returns a copy of the byte array as a byte slice.
func (p PackedByteArray) Bytes() []byte {
	return gdmemory.IntoSlice[byte](gdunsafe.PackedArray[byte](pointers.Get(p)).Access(0), int(p.Size()))
}

func (p PackedByteArray) Len() int { return int(p.Size()) }
func (p PackedByteArray) Cap() int { return int(p.Size()) }

func (p PackedByteArray) Free() {
	if ptr, ok := pointers.End(p); ok {
		noescape.Free(gdextension.TypePackedByteArray, &ptr)
	}
}

func (p PackedInt32Array) Index(idx Int) int32 {
	return gdunsafe.PackedArray[int32](pointers.Get(p)).Access(idx).Get()
}
func (p PackedInt32Array) SetIndex(idx Int, value int32) {
	gdunsafe.PackedArray[int32](pointers.Get(p)).Modify(idx).Set(value)
}

func (p PackedInt32Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		gdunsafe.Free(gdunsafe.PackedArray[int32](ptr))
	}
}

func (p PackedInt32Array) Len() int { return int(p.Size()) }
func (p PackedInt32Array) Cap() int { return int(p.Size()) }

func (p PackedInt64Array) Index(idx Int) int64 {
	return gdunsafe.PackedArray[int64](pointers.Get(p)).Access(idx).Get()
}

func (p PackedInt64Array) SetIndex(idx Int, value int64) {
	gdunsafe.PackedArray[int64](pointers.Get(p)).Modify(idx).Set(value)
}

func (p PackedInt64Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		gdunsafe.Free(gdunsafe.PackedArray[int64](ptr))
	}
}

func (p PackedInt64Array) Len() int { return int(p.Size()) }
func (p PackedInt64Array) Cap() int { return int(p.Size()) }

func (p PackedFloat32Array) Index(idx Int) float32 {
	return gdunsafe.PackedArray[float32](pointers.Get(p)).Access(idx).Get()
}

func (p PackedFloat32Array) SetIndex(idx Int, value float32) {
	gdunsafe.PackedArray[float32](pointers.Get(p)).Modify(idx).Set(value)
}

func (p PackedFloat32Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		gdunsafe.Free(gdunsafe.PackedArray[float32](ptr))
	}
}

func (p PackedFloat32Array) Len() int { return int(p.Size()) }
func (p PackedFloat32Array) Cap() int { return int(p.Size()) }

func (p PackedFloat64Array) Index(idx Int) float64 {
	return gdunsafe.PackedArray[float64](pointers.Get(p)).Access(idx).Get()
}

func (p PackedFloat64Array) SetIndex(idx Int, value float64) {
	gdunsafe.PackedArray[float64](pointers.Get(p)).Modify(idx).Set(value)
}

func (p PackedFloat64Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		gdunsafe.Free(gdunsafe.PackedArray[float64](ptr))
	}
}

func (p PackedFloat64Array) Len() int { return int(p.Size()) }
func (p PackedFloat64Array) Cap() int { return int(p.Size()) }

func (p PackedStringArray) String() string {
	var builder strings.Builder
	builder.WriteString("[")
	size := int(p.Size())
	for i := range size {
		builder.WriteString(p.Index(Int(i)).String())
		if i < size-1 {
			builder.WriteString(" ")
		}
	}
	builder.WriteString("]")
	return builder.String()
}

func (p PackedStringArray) Index(idx Int) String {
	return pointers.Raw[String](gdextension.String{gdextension.Pointer(gdunsafe.PackedArray[gdunsafe.String](pointers.Get(p)).Access(idx).Get())}).Copy()
}

func (p PackedStringArray) SetIndex(idx Int, value String) {
	raw, _ := pointers.End(value.Copy())
	gdunsafe.PackedArray[gdunsafe.String](pointers.Get(p)).Modify(idx).Set(gdunsafe.String(raw[0]))
}

func (p PackedStringArray) AsSlice() []String {
	var ptr = gdunsafe.PackedArray[gdunsafe.String](pointers.Get(p))
	var slice = make([]String, p.Size())
	for i := range slice {
		slice[i] = pointers.Raw[String](gdextension.String{gdextension.Pointer(ptr.Access(Int(i)).Get())}).Copy()
	}
	return slice
}

func (p PackedStringArray) Free() {
	if ptr, ok := pointers.End(p); ok {
		gdunsafe.Free(gdunsafe.PackedArray[gdunsafe.String](ptr))
	}
}

func (p PackedVector2Array) Index(idx Int) Vector2 {
	return gdunsafe.PackedArray[Vector2](pointers.Get(p)).Access(idx).Get()
}

func (p PackedVector2Array) SetIndex(idx Int, value Vector2) {
	gdunsafe.PackedArray[Vector2](pointers.Get(p)).Modify(idx).Set(value)
}

func (p PackedVector2Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		gdunsafe.Free(gdunsafe.PackedArray[Vector2](ptr))
	}
}

func (p PackedVector2Array) Len() int { return int(p.Size()) }
func (p PackedVector2Array) Cap() int { return int(p.Size()) }

func (p PackedVector3Array) Index(idx Int) Vector3 {
	return gdunsafe.PackedArray[Vector3](pointers.Get(p)).Access(idx).Get()
}

func (p PackedVector3Array) SetIndex(idx Int, value Vector3) {
	gdunsafe.PackedArray[Vector3](pointers.Get(p)).Modify(idx).Set(value)
}

func (p PackedVector3Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		gdunsafe.Free(gdunsafe.PackedArray[Vector3](ptr))
	}
}

func (p PackedVector3Array) Len() int { return int(p.Size()) }
func (p PackedVector3Array) Cap() int { return int(p.Size()) }

func (p PackedVector4Array) Index(idx Int) Vector4 {
	return gdunsafe.PackedArray[Vector4](pointers.Get(p)).Access(idx).Get()
}

func (p PackedVector4Array) SetIndex(idx Int, value Vector4) {
	gdunsafe.PackedArray[Vector4](pointers.Get(p)).Modify(idx).Set(value)
}

func (p PackedVector4Array) Free() {
	if ptr, ok := pointers.End(p); ok {
		gdunsafe.Free(gdunsafe.PackedArray[Vector4](ptr))
	}
}

func (p PackedVector4Array) Len() int { return int(p.Size()) }
func (p PackedVector4Array) Cap() int { return int(p.Size()) }

func (p PackedColorArray) Index(idx Int) Color {
	return gdunsafe.PackedArray[Color](pointers.Get(p)).Access(idx).Get()
}

func (p PackedColorArray) SetIndex(idx Int, value Color) {
	gdunsafe.PackedArray[Color](pointers.Get(p)).Modify(idx).Set(value)
}

func (p PackedColorArray) Free() {
	if ptr, ok := pointers.End(p); ok {
		gdunsafe.Free(gdunsafe.PackedArray[Color](ptr))
	}
}

func (p PackedColorArray) Len() int { return int(p.Size()) }
func (p PackedColorArray) Cap() int { return int(p.Size()) }

func NewPackedByteArray() PackedByteArray {
	return pointers.New[PackedByteArray](noescape.Make[gdextension.PackedArray[byte]](builtin.creation.PackedByteArray[0], 0, nil))
}

// PackedByteSlice returns a [PackedByteArray] from a byte slice.
func NewPackedByteSlice(data []byte) PackedByteArray {
	var array = NewPackedByteArray()
	array.Resize(Int(len(data)))
	gdmemory.LoadSlice[byte](gdunsafe.PackedArray[byte](pointers.Get(array)).Modify(0), data)
	return array
}

func NewPackedColorArray() PackedColorArray {
	return pointers.New[PackedColorArray](noescape.Make[gdextension.PackedArray[Color]](builtin.creation.PackedColorArray[0], 0, nil))
}

func NewPackedColorSlice(data []Color) PackedColorArray {
	var array = NewPackedColorArray()
	array.Resize(Int(len(data)))
	gdmemory.LoadSlice[Color](gdunsafe.PackedArray[Color](pointers.Get(array)).Modify(0), data)
	return array
}

func NewPackedFloat32Array() PackedFloat32Array {
	return pointers.New[PackedFloat32Array](noescape.Make[gdextension.PackedArray[float32]](builtin.creation.PackedFloat32Array[0], 0, nil))
}

func NewPackedFloat32Slice(data []float32) PackedFloat32Array {
	var array = NewPackedFloat32Array()
	array.Resize(Int(len(data)))
	gdmemory.LoadSlice[float32](gdunsafe.PackedArray[float32](pointers.Get(array)).Modify(0), data)
	return array
}

func NewPackedInt32Array() PackedInt32Array {
	return pointers.New[PackedInt32Array](noescape.Make[gdextension.PackedArray[int32]](builtin.creation.PackedInt32Array[0], 0, nil))
}

func NewPackedInt32Slice(data []int32) PackedInt32Array {
	var array = NewPackedInt32Array()
	array.Resize(Int(len(data)))
	gdmemory.LoadSlice[int32](gdunsafe.PackedArray[int32](pointers.Get(array)).Modify(0), data)
	return array
}

func NewPackedStringArray() PackedStringArray {
	return pointers.New[PackedStringArray](noescape.Make[gdextension.PackedArray[gdextension.String]](builtin.creation.PackedStringArray[0], 0, nil))
}

func NewPackedStringSlice(data []string) PackedStringArray {
	var array = NewPackedStringArray()
	array.Resize(Int(len(data)))
	for i, str := range data {
		array.SetIndex(Int(i), NewString(str))
	}
	return array
}

func NewPackedReadableStringSlice(data []StringType.Unicode) PackedStringArray {
	var array = NewPackedStringArray()
	array.Resize(Int(len(data)))
	for i, str := range data {
		_, raw := StringType.Proxy(str, StringCacheCheck, NewStringProxy)
		array.SetIndex(Int(i), pointers.Load[String](raw))
	}
	return array
}

func NewPackedVector2Array() PackedVector2Array {
	return pointers.New[PackedVector2Array](noescape.Make[gdextension.PackedArray[Vector2]](builtin.creation.PackedVector2Array[0], 0, nil))
}

func NewPackedVector2Slice(data []Vector2) PackedVector2Array {
	var array = NewPackedVector2Array()
	array.Resize(Int(len(data)))
	gdmemory.LoadSlice[Vector2](gdunsafe.PackedArray[Vector2](pointers.Get(array)).Modify(0), data)
	return array
}

func NewPackedVector3Array() PackedVector3Array {
	return pointers.New[PackedVector3Array](noescape.Make[gdextension.PackedArray[Vector3]](builtin.creation.PackedVector3Array[0], 0, nil))
}

func NewPackedVector3Slice(data []Vector3) PackedVector3Array {
	var array = NewPackedVector3Array()
	array.Resize(Int(len(data)))
	gdmemory.LoadSlice[Vector3](gdunsafe.PackedArray[Vector3](pointers.Get(array)).Modify(0), data)
	return array
}

func NewPackedVector4Array() PackedVector4Array {
	return pointers.New[PackedVector4Array](noescape.Make[gdextension.PackedArray[Vector4]](builtin.creation.PackedVector4Array[0], 0, nil))
}

func NewPackedVector4Slice(data []Vector4) PackedVector4Array {
	var array = NewPackedVector4Array()
	array.Resize(Int(len(data)))
	gdmemory.LoadSlice[Vector4](gdunsafe.PackedArray[Vector4](pointers.Get(array)).Modify(0), data)
	return array
}

func NewPackedInt64Array() PackedInt64Array {
	return pointers.New[PackedInt64Array](noescape.Make[gdextension.PackedArray[int64]](builtin.creation.PackedInt64Array[0], 0, nil))
}

func NewPackedInt64Slice(data []int64) PackedInt64Array {
	var array = NewPackedInt64Array()
	array.Resize(Int(len(data)))
	gdmemory.LoadSlice[int64](gdunsafe.PackedArray[int64](pointers.Get(array)).Modify(0), data)
	return array
}

func NewPackedFloat64Array() PackedFloat64Array {
	return pointers.New[PackedFloat64Array](noescape.Make[gdextension.PackedArray[float64]](builtin.creation.PackedFloat64Array[0], 0, nil))
}

func NewPackedFloat64Slice(data []float64) PackedFloat64Array {
	var array = NewPackedFloat64Array()
	array.Resize(Int(len(data)))
	gdmemory.LoadSlice[float64](gdunsafe.PackedArray[float64](pointers.Get(array)).Modify(0), data)
	return array
}
