package Packed

import (
	"iter"
	"sort"

	gd "grow.graphics/gd/internal"
	"grow.graphics/gd/variant"
	"grow.graphics/gd/variant/Color"
	"grow.graphics/gd/variant/Vector2"
	"grow.graphics/gd/variant/Vector3"
	"grow.graphics/gd/variant/Vector4"
)

type packable interface {
	~byte | ~Color.RGBA | ~float32 | ~float64 | ~int32 | ~int64 | ~string | ~Vector2.XY | ~Vector3.XYZ | ~Vector4.XYZW | gd.String
}

type proxiable interface {
	gd.PackedByteArray | gd.PackedColorArray | gd.PackedFloat32Array | gd.PackedFloat64Array | gd.PackedInt32Array |
		gd.PackedInt64Array | gd.PackedStringArray | gd.PackedVector2Array | gd.PackedVector3Array | gd.PackedVector4Array
}

type methods[A proxiable, T packable] interface {
	Index(gd.Int) T
	Append(value T) gd.Bool
	AppendArray(A)
	Bsearch(T, gd.Bool) gd.Int
	Clear()
	Count(T) gd.Int
	Duplicate() A
	Fill(T)
	Find(T, gd.Int) gd.Int
	Has(T) gd.Bool
	Insert(gd.Int, T) gd.Int
	IsEmpty() gd.Bool
	PushBack(T) gd.Bool
	RemoveAt(gd.Int)
	Resize(gd.Int) gd.Int
	Reverse()
	Rfind(T, gd.Int) gd.Int
	Set(gd.Int, T)
	Size() gd.Int
	Slice(gd.Int, gd.Int) A
	Sort()
	ToByteArray() gd.PackedByteArray
}

type container[Self any, Wrap, Elem packable, Proxy proxiable] interface {
	make([]Wrap, Proxy) Self
	load() ([]Wrap, Proxy)
	less(Wrap, Wrap) bool

	alloc() Proxy

	wrap(Elem) Wrap
	conv(Wrap) Elem
}

// Array is the base type for all packed arrays.
type Array[T container[T, Wrap, Elem, Proxy], Wrap, Elem packable, Proxy proxiable,
	Methods interface {
		*Proxy
		methods[Proxy, Elem]
	},
] struct {
	local []Wrap
	proxy Proxy
}

func (p Array[S, W, T, P, M]) load() ([]W, P) {
	return p.local, p.proxy
}

func (p *Array[S, W, T, P, M]) Iter() iter.Seq2[int, W] {
	var zero S
	return func(yield func(int, W) bool) {
		if p.proxy != ([1]P{}[0]) {
			for i := range M(&p.proxy).Size() {
				if !yield(int(i), zero.wrap(M(&p.proxy).Index(gd.Int(i)))) {
					break
				}
			}
			return
		}
		for i, v := range p.local {
			if !yield(i, v) {
				break
			}
		}
	}
}

// Proxy converts the array into a proxied array and returns it.
func (p *Array[S, W, T, P, M]) Proxy() P {
	var zero S
	if p.proxy == ([1]P{}[0]) {
		p.proxy = zero.alloc()
		M(&p.proxy).Resize(gd.Int(len(p.local)))
		for i, v := range p.local {
			M(&p.proxy).Set(gd.Int(i), zero.conv(v))
		}
		p.local = nil
	}
	return p.proxy
}

// Index returns the element at the specified index.
func (p *Array[S, W, T, P, M]) Index(i int) W { //gd:PackedArray[]
	var zero S
	if p.proxy != ([1]P{}[0]) {
		return zero.wrap(M(&p.proxy).Index(gd.Int(i)))
	}
	return p.local[i]
}

// Append appends an element at the end of the array (alias of [PushBack]).
func (p *Array[S, W, T, P, M]) Append(value W) bool { //gd:PackedArray.append
	var zero S
	if p.proxy != ([1]P{}[0]) {
		return M(&p.proxy).Append(zero.conv(value))
	}
	p.local = append(p.local, value)
	return true
}

// Append appends an element at the end of the array (alias of [PushBack]).
func (p *Array[S, W, T, P, M]) AppendArray(value S) { //gd:PackedArray.append_array
	local, proxy := value.load()
	if p.proxy != ([1]P{}[0]) && proxy != ([1]P{}[0]) {
		M(&p.proxy).AppendArray(proxy)
		return
	}
	wrap := (Array[S, W, T, P, M]{local: local, proxy: proxy})
	for _, v := range wrap.Iter() {
		p.Append(v)
	}
}

// BinarySearch finds the index of an existing value (or the insertion index that maintains
// sorting order, if the value is not yet present in the array) using binary search.
//
// Note: Calling bsearch on an unsorted array results in unexpected behavior.
func (p *Array[S, W, T, P, M]) BinarySearch(value W) int { //gd:PackedArray.bsearch
	var zero S
	if p.proxy != ([1]P{}[0]) {
		return int(M(&p.proxy).Bsearch(zero.conv(value), true))
	}
	const before = true
	var lo = 0
	var hi = p.Size()
	if before {
		for lo < hi {
			var mid = (lo + hi) / 2
			if zero.less(p.Index(mid), value) {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
	} else {
		for lo < hi {
			var mid = (lo + hi) / 2
			if zero.less(value, p.Index(mid)) {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
	}
	return lo
}

// Clear clears the array. This is equivalent to using resize with a size of 0.
func (p *Array[S, W, T, P, M]) Clear() { //gd:PackedArray.clear
	if p.proxy != ([1]P{}[0]) {
		M(&p.proxy).Clear()
		return
	}
	p.local = p.local[:0]
}

// Count returns the number of times an element is in the array.
func (p *Array[S, W, T, P, M]) Count(value W) int { //gd:PackedArray.count
	var zero S
	if p.proxy != ([1]P{}[0]) {
		return int(M(&p.proxy).Count(zero.conv(value)))
	}
	var count int
	for _, v := range p.local {
		if v == value {
			count++
		}
	}
	return count
}

// Duplicate creates a copy of the array, and returns it.
func (p *Array[S, W, T, P, M]) Duplicate() S { //gd:PackedArray.duplicate
	var zero S
	if p.proxy != ([1]P{}[0]) {
		return zero.make(nil, M(&p.proxy).Duplicate())
	}
	return zero.make(append([]W(nil), p.local...), ([1]P{}[0]))
}

// Fill assigns the given value to all elements in the array. This can typically be used
// together with resize to create an array with a given size and initialized elements.
func (p *Array[S, W, T, P, M]) Fill(value W) { //gd:PackedArray.fill
	var zero S
	if p.proxy != ([1]P{}[0]) {
		M(&p.proxy).Fill(zero.conv(value))
		return
	}
	for i := range p.local {
		p.local[i] = value
	}
}

// Find searches the array for a value and returns its index or -1 if not found. Optionally,
// the initial search index can be passed.
//
// Note: Vectors with [math.NaN] elements don't behave the same as other vectors. Therefore,
// the results from this method may not be accurate if NaNs are included.
func (p *Array[S, W, T, P, M]) Find(value W, from int) int { //gd:PackedArray.find
	var zero S
	if p.proxy != ([1]P{}[0]) {
		return int(M(&p.proxy).Find(zero.conv(value), gd.Int(from)))
	}
	for i := from; i < len(p.local); i++ {
		if p.local[i] == value {
			return i
		}
	}
	return -1
}

// Has returns true if the array contains value.
//
// Note: Vectors with NaN elements don't behave the same as other vectors. Therefore, the
// results from this method may not be accurate if NaNs are included.
func (p *Array[S, W, T, P, M]) Has(value W) bool { //gd:PackedArray.has
	var zero S
	if p.proxy != ([1]P{}[0]) {
		return M(&p.proxy).Has(zero.conv(value))
	}
	for _, v := range p.local {
		if v == value {
			return true
		}
	}
	return false
}

// Insert inserts a new element at a given position in the array. The position must be valid,
// or at the end of the array (idx == size()).
func (p *Array[S, W, T, P, M]) Insert(idx int, value W) int { //gd:PackedArray.insert
	var zero S
	if p.proxy != ([1]P{}[0]) {
		return int(M(&p.proxy).Insert(gd.Int(idx), zero.conv(value)))
	}
	if idx == len(p.local) {
		p.local = append(p.local, value)
		return idx
	}
	if idx < 0 || idx > len(p.local) {
		return -1
	}
	p.local = append(p.local, p.local[len(p.local)-1])
	copy(p.local[idx+1:], p.local[idx:])
	p.local[idx] = value
	return idx
}

// IsEmpty returns true if the array is empty.
func (p *Array[S, W, T, P, M]) IsEmpty() bool { //gd:PackedArray.is_empty
	if p.proxy != ([1]P{}[0]) {
		return M(&p.proxy).IsEmpty()
	}
	return len(p.local) == 0
}

// PushBack adds an element at the end of the array.
func (p *Array[S, W, T, P, M]) PushBack(value W) bool { //gd:PackedArray.push_back
	var zero S
	if p.proxy != ([1]P{}[0]) {
		return M(&p.proxy).PushBack(zero.conv(value))
	}
	p.local = append(p.local, value)
	return true
}

// RemoveAt removes an element from the array by index.
func (p *Array[S, W, T, P, M]) RemoveAt(idx int) { //gd:PackedArray.remove_at
	if p.proxy != ([1]P{}[0]) {
		M(&p.proxy).RemoveAt(gd.Int(idx))
		return
	}
	p.local = append(p.local[:idx], p.local[idx+1:]...)
}

// Resize sets the size of the array. If the array is grown, reserves elements at
// the end of the array. If the array is shrunk, truncates the array to the new size.
// Calling resize once and assigning the new values is faster than adding new elements one by one.
func (p *Array[S, W, T, P, M]) Resize(size int) { //gd:PackedArray.resize
	if p.proxy != ([1]P{}[0]) {
		M(&p.proxy).Resize(gd.Int(size))
		return
	}
	if size < len(p.local) {
		p.local = p.local[:size]
	} else if size > len(p.local) {
		p.local = append(p.local, make([]W, size-len(p.local))...)
	}
}

// Reverse reverses the order of the elements in the array.
func (p *Array[S, W, T, P, M]) Reverse() { //gd:PackedArray.reverse
	if p.proxy != ([1]P{}[0]) {
		M(&p.proxy).Reverse()
		return
	}
	for i, j := 0, len(p.local)-1; i < j; i, j = i+1, j-1 {
		p.local[i], p.local[j] = p.local[j], p.local[i]
	}
}

// ReverseFind searches the array in reverse order.
//
// Note: Vectors with NaN elements don't behave the same as other vectors. Therefore, the
// results from this method may not be accurate if NaNs are included.
func (p *Array[S, W, T, P, M]) ReverseFind(value W) int { //gd:PackedArray.rfind
	var zero S
	if p.proxy != ([1]P{}[0]) {
		return int(M(&p.proxy).Rfind(zero.conv(value), -1))
	}
	for i := len(p.local) - 1; i >= 0; i-- {
		if p.local[i] == value {
			return i
		}
	}
	return -1
}

// Set assigns a value to an element in the array. The index must be valid.
func (p *Array[S, W, T, P, M]) Set(idx int, value W) { //gd:PackedArray.set
	var zero S
	if p.proxy != ([1]P{}[0]) {
		M(&p.proxy).Set(gd.Int(idx), zero.conv(value))
		return
	}
	p.local[idx] = value
}

// Size returns the number of elements in the array.
func (p *Array[S, W, T, P, M]) Size() int { //gd:PackedArray.size
	if p.proxy != ([1]P{}[0]) {
		return int(M(&p.proxy).Size())
	}
	return len(p.local)
}

// Slice returns the slice of the the array, from begin (inclusive) to end (exclusive), as a new array.
//
// The absolute value of begin and end will be clamped to the array size, so the default value for end
// makes it slice to the size of the array by default (i.e. arr.slice(1) is a shorthand for arr.slice(1, arr.size())).
//
// If either begin or end are negative, they will be relative to the end of the array (i.e. arr.slice(0, -2)
// is a shorthand for arr.slice(0, arr.size() - 2)).
func (p *Array[S, W, T, P, M]) Slice(begin, end int) S { //gd:PackedArray.slice
	var zero S
	if p.proxy != ([1]P{}[0]) {
		return zero.make(nil, M(&p.proxy).Slice(gd.Int(begin), gd.Int(end)))
	}
	if begin < 0 {
		begin = len(p.local) + begin
	}
	if end < 0 {
		end = len(p.local) + end
	}
	if begin < 0 {
		begin = 0
	}
	if end > len(p.local) {
		end = len(p.local)
	}
	return zero.make(p.local[begin:end], ([1]P{}[0]))
}

// Sort sorts the elements of the array in ascending order.
//
// Note: Vectors with NaN elements don't behave the same as other vectors. Therefore, the results from
// this method may not be accurate if NaNs are included.
func (p *Array[S, W, T, P, M]) Sort() { //gd:PackedArray.sort
	var zero S
	if p.proxy != ([1]P{}[0]) {
		M(&p.proxy).Sort()
		return
	}
	sort.Slice(p.local, func(i, j int) bool {
		return zero.less(p.local[i], p.local[j])
	})
}

// ToByteArray returns a byte array containing the elements of the array encoded as bytes.
func (p *Array[S, W, T, P, M]) ToByteArray() ByteArray { //gd:PackedArray.to_byte_array
	var zero ByteArray
	if p.proxy != ([1]P{}[0]) {
		return zero.make(nil, M(&p.proxy).ToByteArray())
	}
	zero.local, _ = variant.Marshal(p.local)
	return zero
}
