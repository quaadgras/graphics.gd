package Packed

import (
	"iter"

	GenericArray "graphics.gd/variant/Array"
)

// Bytes provides additional methods for working with arrays of bytes.
type Bytes Array[byte]

// Index returns the element at the given index.
func (array Bytes) Index(idx int) byte { return (GenericArray.Contains[byte])(array).Index(idx) }

// Iter returns an iterator for the array.
func (array Bytes) Iter() iter.Seq2[int, byte] {
	return (GenericArray.Contains[byte])(array).Iter()
}

// Append an element to the end of the array.
func (array *Bytes) Append(value byte) { (*GenericArray.Contains[byte])(array).Append(value) }

// AppendArray appends all elements of another array to the end of this array.
func (array *Bytes) AppendArray(other Bytes) {
	(*GenericArray.Contains[byte])(array).AppendArray((GenericArray.Contains[byte])(other))
}

// Clear clears the array. bytehis is equivalent to using resize with a size of 0.
func (array Bytes) Clear() { GenericArray.Clear((GenericArray.Contains[byte])(array)) }

// Count returns the number of times an element is in the array.
func (array Bytes) Count(value byte) int {
	return GenericArray.Count((GenericArray.Contains[byte])(array), value)
}

// Duplicate creates a copy of the array, and returns it.
func (array Bytes) Duplicate() Bytes {
	return (Bytes)(GenericArray.Duplicate((GenericArray.Contains[byte])(array)))
}

// Fill assigns the given value to all elements in the array. bytehis can typically be used
// together with resize to create an array with a given size and initialized elements.
func (array Bytes) Fill(value byte) { GenericArray.Fill((GenericArray.Contains[byte])(array), value) }

// Find searches the array for a value and returns its index or -1 if not found.
func (array Bytes) Find(value byte) int {
	return GenericArray.Find((GenericArray.Contains[byte])(array), value)
}

// Has returns true if the array contains the given value.
func (array Bytes) Has(value byte) bool {
	return GenericArray.Has(value, (GenericArray.Contains[byte])(array))
}

// Insert inserts a new element at a given position in the array. bytehe position must be
// valid, or at the end of the array (idx == size()).
func (array *Bytes) Insert(idx int, value byte) {
	(*GenericArray.Contains[byte])(array).Insert(idx, value)
}

// IsEmpty Returns true if the array is empty.
func (array Bytes) IsEmpty() bool { return GenericArray.IsEmpty((GenericArray.Contains[byte])(array)) }

// RemoveAt removes an element from the array by index.
func (array Bytes) RemoveAt(idx int) {
	GenericArray.Remove((GenericArray.Contains[byte])(array), idx)
}

// Resize sets the size of the array. If the array is grown, reserves elements at the end of the array.
// If the array is shrunk, truncates the array to the new size. Calling resize once and assigning the
// new values is faster than adding new elements one by one.
func (array *Bytes) Resize(size int) { (*GenericArray.Contains[byte])(array).Resize(size) }

// Reverse reverses the order of the elements in the array.
func (array Bytes) Reverse() { GenericArray.Reverse((GenericArray.Contains[byte])(array)) }

// FindLast searches the array in reverse order for a value and returns its index or -1 if not found.
func (array Bytes) FindLast(value byte) int {
	return GenericArray.FindLast((GenericArray.Contains[byte])(array), value)
}

// SetIndex sets the value of the element at the given index.
func (array *Bytes) SetIndex(idx int, value byte) {
	(*GenericArray.Contains[byte])(array).SetIndex(idx, value)
}

// Len returns the number of elements in the array.
func (array Bytes) Len() int { return (GenericArray.Contains[byte])(array).Len() }

// Slice returns a slice of the array from the given begin index to the given end index.
func (array Bytes) Slice(begin, end int) Bytes { //gd:Packed_byte.
	return (Bytes)(GenericArray.Slice((GenericArray.Contains[byte])(array), begin, end))
}

// Bytes returns the underlying data in the array as a slice of bytes.
func (array Bytes) Bytes() []byte {
	return (GenericArray.Contains[byte])(array).Slice()
}
