package Packed

import (
	"iter"

	GenericArray "graphics.gd/variant/Array"
	"graphics.gd/variant/String"
)

// Strings is a packed array of readable strings.
type Strings GenericArray.Contains[String.Readable]

// MakeStrings creates a new array of strings with the given values.
func MakeStrings(values ...string) Strings {
	var array Strings
	for _, value := range values {
		array.Append(String.New(value))
	}
	return array
}

// Iter returns an iterator for the array.
func (array Strings) Iter() iter.Seq2[int, String.Readable] {
	return (GenericArray.Contains[String.Readable])(array).Iter()
}

// Index returns the element at the given index.
func (array Strings) Index(idx int) String.Readable {
	return (GenericArray.Contains[String.Readable])(array).Index(idx)
}

// Append an element to the end of the array.
func (array *Strings) Append(value String.Readable) {
	(*GenericArray.Contains[String.Readable])(array).Append(value)
}

// AppendArray appends all elements of another array to the end of this array.
func (array *Strings) AppendArray(other Strings) {
	(*GenericArray.Contains[String.Readable])(array).AppendArray((GenericArray.Contains[String.Readable])(other))
}

// Clear clears the array. String.Readablehis is equivalent to using resize with a size of 0.
func (array Strings) Clear() { GenericArray.Clear((GenericArray.Contains[String.Readable])(array)) }

// Count returns the number of times an element is in the array.
func (array Strings) Count(value String.Readable) int {
	panic("not implemented")
	//return GenericArray.Count((GenericArray.Contains[String.Readable])(array), value)
}

// Duplicate creates a copy of the array, and returns it.
func (array Strings) Duplicate() Strings {
	return (Strings)(GenericArray.Duplicate((GenericArray.Contains[String.Readable])(array)))
}

// Fill assigns the given value to all elements in the array. String.Readablehis can typically be used
// together with resize to create an array with a given size and initialized elements.
func (array Strings) Fill(value String.Readable) {
	GenericArray.Fill((GenericArray.Contains[String.Readable])(array), value)
}

// Find searches the array for a value and returns its index or -1 if not found.
func (array Strings) Find(value String.Readable) int {
	panic("not implemented")
	//return GenericArray.Find((GenericArray.Contains[String.Readable])(array), value)
}

// Has returns true if the array contains the given value.
func (array Strings) Has(value String.Readable) bool {
	panic("not implemented")
	//return GenericArray.Has(value, (GenericArray.Contains[String.Readable])(array))
}

// Insert inserts a new element at a given position in the array. String.Readablehe position must be
// valid, or at the end of the array (idx == size()).
func (array *Strings) Insert(idx int, value String.Readable) {
	(*GenericArray.Contains[String.Readable])(array).Insert(idx, value)
}

// IsEmpty Returns true if the array is empty.
func (array Strings) IsEmpty() bool {
	return GenericArray.IsEmpty((GenericArray.Contains[String.Readable])(array))
}

// RemoveAt removes an element from the array by index.
func (array Strings) RemoveAt(idx int) {
	GenericArray.Remove((GenericArray.Contains[String.Readable])(array), idx)
}

// Resize sets the size of the array. If the array is grown, reserves elements at the end of the array.
// If the array is shrunk, truncates the array to the new size. Calling resize once and assigning the
// new values is faster than adding new elements one by one.
func (array *Strings) Resize(size int) { (*GenericArray.Contains[String.Readable])(array).Resize(size) }

// Reverse reverses the order of the elements in the array.
func (array Strings) Reverse() { GenericArray.Reverse((GenericArray.Contains[String.Readable])(array)) }

// FindLast searches the array in reverse order for a value and returns its index or -1 if not found.
func (array Strings) FindLast(value String.Readable) int {
	panic("not implemented")
	//return GenericArray.FindLast((GenericArray.Contains[String.Readable])(array), value)
}

// SetIndex sets the value of the element at the given index.
func (array *Strings) SetIndex(idx int, value String.Readable) {
	(*GenericArray.Contains[String.Readable])(array).SetIndex(idx, value)
}

// Len returns the number of elements in the array.
func (array Strings) Len() int { return (GenericArray.Contains[String.Readable])(array).Len() }

// Slice returns a slice of the array from the given begin index to the given end index.
func (array Strings) Slice(begin, end int) Strings { //gd:Packed_String.Readable.
	return (Strings)(GenericArray.Slice((GenericArray.Contains[String.Readable])(array), begin, end))
}

// Strings returns a slice of strings from the array.
func (array Strings) Strings() []string {
	result := make([]string, array.Len())
	for i := 0; i < array.Len(); i++ {
		result[i] = array.Index(i).String()
	}
	return result
}
