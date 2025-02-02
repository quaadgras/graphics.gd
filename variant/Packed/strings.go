package Packed

import (
	"iter"
	"sort"

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
func (array *Strings) Append(value String.Readable) { //gd:PackedStringArray.push_back PackedStringArray.append
	(*GenericArray.Contains[String.Readable])(array).Append(value)
}

// AppendArray appends all elements of another array to the end of this array.
func (array *Strings) AppendArray(other Strings) { //gd:PackedStringArray.append_array
	(*GenericArray.Contains[String.Readable])(array).AppendArray((GenericArray.Contains[String.Readable])(other))
}

// Clear clears the array. String.Readablehis is equivalent to using resize with a size of 0.
func (array Strings) Clear() { GenericArray.Clear((GenericArray.Contains[String.Readable])(array)) } //gd:PackedStringArray.clear

// Count returns the number of times an element is in the array.
func (array Strings) Count(value String.Readable) int { //gd:PackedStringArray.count
	var count int
	for i := range array.Len() {
		if String.Comparison(array.Index(i), value) == 0 {
			count++
		}
	}
	return count
}

// Duplicate creates a copy of the array, and returns it.
func (array Strings) Duplicate() Strings { //gd:PackedStringArray.duplicate
	return (Strings)(GenericArray.Duplicate((GenericArray.Contains[String.Readable])(array)))
}

// Fill assigns the given value to all elements in the array. String.Readablehis can typically be used
// together with resize to create an array with a given size and initialized elements.
func (array Strings) Fill(value String.Readable) { //gd:PackedStringArray.fill
	GenericArray.Fill((GenericArray.Contains[String.Readable])(array), value)
}

// Find searches the array for a value and returns its index or -1 if not found.
func (array Strings) Find(value String.Readable) int { //gd:PackedStringArray.find
	for i := range array.Len() {
		if String.Comparison(array.Index(i), value) == 0 {
			return i
		}
	}
	return -1
}

// Has returns true if the array contains the given value.
func (array Strings) Has(value String.Readable) bool { //gd:PackedStringArray.has
	return array.Find(value) != -1
}

// Insert inserts a new element at a given position in the array. String.Readablehe position must be
// valid, or at the end of the array (idx == size()).
func (array *Strings) Insert(idx int, value String.Readable) { //gd:PackedStringArray.insert
	(*GenericArray.Contains[String.Readable])(array).Insert(idx, value)
}

// IsEmpty Returns true if the array is empty.
func (array Strings) IsEmpty() bool { //gd:PackedStringArray.is_empty
	return GenericArray.IsEmpty((GenericArray.Contains[String.Readable])(array))
}

// RemoveAt removes an element from the array by index.
func (array Strings) RemoveAt(idx int) { //gd:PackedStringArray.remove_at
	GenericArray.Remove((GenericArray.Contains[String.Readable])(array), idx)
}

// Resize sets the size of the array. If the array is grown, reserves elements at the end of the array.
// If the array is shrunk, truncates the array to the new size. Calling resize once and assigning the
// new values is faster than adding new elements one by one.
func (array *Strings) Resize(size int) { (*GenericArray.Contains[String.Readable])(array).Resize(size) } //gd:PackedStringArray.resize

// Reverse reverses the order of the elements in the array.
func (array Strings) Reverse() { GenericArray.Reverse((GenericArray.Contains[String.Readable])(array)) } //gd:PackedStringArray.reverse

// FindLast searches the array in reverse order for a value and returns its index or -1 if not found.
func (array Strings) FindLast(value String.Readable) int { //gd:PackedStringArray.rfind
	for i := array.Len() - 1; i >= 0; i-- {
		if String.Comparison(array.Index(i), value) == 0 {
			return i
		}
	}
	return -1
}

// SetIndex sets the value of the element at the given index.
func (array *Strings) SetIndex(idx int, value String.Readable) { //gd:PackedStringArray.set
	(*GenericArray.Contains[String.Readable])(array).SetIndex(idx, value)
}

// Len returns the number of elements in the array.
func (array Strings) Len() int { return (GenericArray.Contains[String.Readable])(array).Len() } //gd:PackedStringArray.size

// Slice returns a slice of the array from the given begin index to the given end index.
func (array Strings) Slice(begin, end int) Strings { //gd:PackedStringArray.slice
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

type stringsSorter struct {
	array Strings
}

func (s stringsSorter) Swap(i, j int) {
	v, w := s.array.Index(i), s.array.Index(j)
	s.array.SetIndex(i, w)
	s.array.SetIndex(j, v)
}
func (s stringsSorter) Less(i, j int) bool {
	return String.Comparison(s.array.Index(i), s.array.Index(j)) < 0
}
func (s stringsSorter) Len() int { return s.array.Len() }

// Sort sorts the array in ascending order. The final order is dependent on the
// "less than" (<) comparison between elements.
func (array Strings) Sort() { //gd:PackedStringArray.sort
	sort.Sort(stringsSorter{array: array})
}

// BinarySearch returns the index of value in the sorted array. If it cannot be found, returns where value should
// be inserted to keep the array sorted. The algorithm used is binary search. The returned index comes before all
// existing elements equal to value in the array.
//
// Note: Calling BinarySearch on an unsorted array will result in unexpected behavior. Use [Sort] before calling
// this method.
func (array Strings) BinarySearch(value String.Readable, before bool) int { //gd:PackedStringArray.bsearch
	less := func(a, b String.Readable) bool {
		return String.Comparison(a, b) < 0
	}
	return GenericArray.Contains[String.Readable](array).BinarySearchFunc(less, value, before)
}
