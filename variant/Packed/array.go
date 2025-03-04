// Package Packed contains specific Array types that are more efficient and convenient to work with.
package Packed

import (
	"iter"
	"unsafe"

	GenericArray "graphics.gd/variant/Array"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Vector2"
	"graphics.gd/variant/Vector3"
	"graphics.gd/variant/Vector4"
)

// Type supported by [Array].
type Type interface {
	~byte | ~Color.RGBA | ~float32 | ~float64 | ~int32 | ~int64 | ~Vector2.XY | ~Vector3.XYZ | ~Vector4.XYZW
}

// Array contains comparable elements of type T Similar to [Array.Contains[T]] but for a supported [Type]
// is more efficient and convienient to work with.
type Array[T Type] GenericArray.Contains[T]

// New creates a new array with the given values.
func New[T Type](values ...T) Array[T] { return (Array[T])(GenericArray.New(values...)) }

// Index returns the element at the given index.
func (array Array[T]) Index(idx int) T { return (GenericArray.Contains[T])(array).Index(idx) } //gd:PackedArray.get

// Iter returns an iterator for the array.
func (array Array[T]) Iter() iter.Seq2[int, T] {
	return (GenericArray.Contains[T])(array).Iter()
}

// Values returns an iterator on the values of the array.
func (array Array[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := 0; i < array.Len(); i++ {
			if !yield(array.Index(i)) {
				break
			}
		}
	}
}

// Append an element to the end of the array.
func (array *Array[T]) Append(value T) { (*GenericArray.Contains[T])(array).Append(value) } //gd:PackedArray.push_back PackedArray.append

// AppendArray appends all elements of another array to the end of this array.
func (array *Array[T]) AppendArray(other Array[T]) { //gd:PackedArray.append_array
	(*GenericArray.Contains[T])(array).AppendArray((GenericArray.Contains[T])(other))
}

// Clear clears the array. This is equivalent to using resize with a size of 0.
func (array Array[T]) Clear() { GenericArray.Clear((GenericArray.Contains[T])(array)) } //gd:PackedArray.clear

// Count returns the number of times an element is in the array.
func (array Array[T]) Count(value T) int { //gd:PackedArray.count
	return GenericArray.Count((GenericArray.Contains[T])(array), value)
}

// Duplicate creates a copy of the array, and returns it.
func (array Array[T]) Duplicate() Array[T] { //gd:PackedArray.duplicate
	return (Array[T])(GenericArray.Duplicate((GenericArray.Contains[T])(array)))
}

// Fill assigns the given value to all elements in the array. This can typically be used
// together with resize to create an array with a given size and initialized elements.
func (array Array[T]) Fill(value T) { GenericArray.Fill((GenericArray.Contains[T])(array), value) } //gd:PackedArray.fill

// Find searches the array for a value and returns its index or -1 if not found.
func (array Array[T]) Find(value T) int { //gd:PackedArray.find
	return GenericArray.Find((GenericArray.Contains[T])(array), value)
}

// Has returns true if the array contains the given value.
func (array Array[T]) Has(value T) bool { //gd:PackedArray.has
	return GenericArray.Has(value, (GenericArray.Contains[T])(array))
}

// Insert inserts a new element at a given position in the array. The position must be
// valid, or at the end of the array (idx == size()).
func (array *Array[T]) Insert(idx int, value T) { //gd:PackedArray.insert
	(*GenericArray.Contains[T])(array).Insert(idx, value)
}

// IsEmpty Returns true if the array is empty.
func (array Array[T]) IsEmpty() bool { return GenericArray.IsEmpty((GenericArray.Contains[T])(array)) } //gd:PackedArray.is_empty

// RemoveAt removes an element from the array by index.
func (array Array[T]) RemoveAt(idx int) { //gd:PackedArray.remove_at
	GenericArray.Remove((GenericArray.Contains[T])(array), idx)
}

// Resize sets the size of the array. If the array is grown, reserves elements at the end of the array.
// If the array is shrunk, truncates the array to the new size. Calling resize once and assigning the
// new values is faster than adding new elements one by one.
func (array *Array[T]) Resize(size int) { (*GenericArray.Contains[T])(array).Resize(size) } //gd:PackedArray.resize

// Reverse reverses the order of the elements in the array.
func (array Array[T]) Reverse() { GenericArray.Reverse((GenericArray.Contains[T])(array)) } //gd:PackedArray.reverse

// FindLast searches the array in reverse order for a value and returns its index or -1 if not found.
func (array Array[T]) FindLast(value T) int { //gd:PackedArray.rfind
	return GenericArray.FindLast((GenericArray.Contains[T])(array), value)
}

// SetIndex sets the value of the element at the given index.
func (array *Array[T]) SetIndex(idx int, value T) { //gd:PackedArray.set
	(*GenericArray.Contains[T])(array).SetIndex(idx, value)
}

// Len returns the number of elements in the array.
func (array Array[T]) Len() int { return (GenericArray.Contains[T])(array).Len() } //gd:PackedArray.size

// Slice returns a slice of the array from the given begin index to the given end index.
func (array Array[T]) Slice(begin, end int) Array[T] { //gd:PackedArray.slice
	return (Array[T])(GenericArray.Slice((GenericArray.Contains[T])(array), begin, end))
}

// Bytes returns the underlying data in the array as [Bytes].
func (array Array[T]) Bytes() Bytes { //gd:PackedArray.to_byte_array
	var bytes Bytes
	bytes.Resize(array.Len() * int(unsafe.Sizeof([1]T{})))
	for i := 0; i < array.Len(); i++ {
		elem := array.Index(i)
		for j := 0; j < int(unsafe.Sizeof([1]T{})); j++ {
			bytes.SetIndex(i*int(unsafe.Sizeof([1]T{}))+j, *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&elem)) + uintptr(j))))
		}
	}
	return bytes
}
