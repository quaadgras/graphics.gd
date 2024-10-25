package Array

import (
	"fmt"
	"iter"
	"math/rand"
	"reflect"
	"sort"

	"grow.graphics/gd/gdmaths/Int"
)

// Of is an array data structure that can contain a sequence of elements of T. Elements
// are accessed by a numerical index starting at 0. Negative indices are used to count
// from the back (-1 is the last element, -2 is the second to last, etc.).
type Of[T any] interface {
	// Lookup returns the element at index. If index is negative, it counts from the back
	// (-1 is the last element, -2 is the second to last, etc.). If index is out of
	// bounds, it wraps around.
	Lookup(int) T
	// Mutate sets the element at index to element. If index is negative, it counts from
	// the back (-1 is the last element, -2 is the second to last, etc.). If index is
	// out of bounds, it wraps around. If the array is read-only, this method panics.
	Mutate(int, T)
	// Length returns the number of elements in the array.
	Length() int
	// Resize sets the array's number of elements to size. If size is smaller than the
	// array's current size, the elements at the end are removed. If size is greater,
	// new default elements (usually null) are added, depending on the array's type.
	Resize(int)
	// Freeze makes the array read-only. The array's elements cannot be overridden with
	// different values, and their order cannot change. Does not apply to nested
	// elements, such as dictionaries.
	Freeze()
	// Static returns true if the array is read-only. See [Of.Freeze].
	Static() bool
}

type ordered interface {
	~int8 | ~int16 | ~int32 | ~int64 | ~int | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uint | ~uintptr | ~float32 | ~float64 | ~string
}

type local[T any] struct {
	memory []T
	static bool
}

func (s *local[T]) Lookup(index int) T {
	index = Int.Posmod(index, s.Length())
	return s.memory[index]
}
func (s *local[T]) Mutate(index int, element T) {
	index = Int.Posmod(index, s.Length())
	s.memory[index] = element
}
func (s *local[T]) Length() int { return len(s.memory) }
func (s *local[T]) Resize(size int) {
	if size < 0 {
		panic(fmt.Errorf("invalid size %d", size))
	}
	if len(s.memory) < size {
		s.memory = append(s.memory, make([]T, size-len(s.memory))...)
	} else {
		s.memory = s.memory[:size]
	}
}
func (s *local[T]) Freeze()      { s.static = true }
func (s *local[T]) Static() bool { return s.static }

// New creates a new array with the given elements.
func New[T any](elements ...T) Of[T] {
	return &local[T]{
		memory: []T(elements),
	}
}

// All calls the given function on each element in the array and returns true if the fn returns
// true for all elements in the array. If the fn returns false for one array element or more,
// this method returns false. See also [Any], [Filter], [Map], and [Reduce].
func All[T any](fn func(T) bool, array Of[T]) bool {
	if array == nil {
		return true
	}
	for i := range array.Length() {
		if !fn(array.Lookup(i)) {
			return false
		}
	}
	return true
}

// Any calls the given function on each element in the array and returns true if the function
// returns true for one or more elements in the array. If the function returns false for all
// elements in the array, this method returns false. See also [Any], [Filter], [Map], and [Reduce].
func Any[T any](fn func(T) bool, array Of[T]) bool {
	if array == nil {
		return false
	}
	for i := range array.Length() {
		if fn(array.Lookup(i)) {
			return true
		}
	}
	return false
}

// Append appends value at the end of the array (alias of PushBack).
func Append[T any](array Of[T], value T) {
	if array == nil {
		return
	}
	array.Resize(array.Length() + 1)
	array.Mutate(array.Length()-1, value)
}

// AppendTo appends another array at the end of this array.
func AppendTo[T any](array, other Of[T]) {
	if array == nil || other == nil {
		return
	}
	for i := range other.Length() {
		Append(array, other.Lookup(i))
	}
}

// Assign assigns elements of another array into the array. Resizes the
// array to match array. Performs type conversions if the array is typed.
func Assign[T any](array, other Of[T]) {
	if array == nil {
		return
	}
	if other == nil {
		array.Resize(0)
		return
	}
	for i := range other.Length() {
		if i < array.Length() {
			array.Mutate(i, other.Lookup(i))
		} else {
			Append(array, other.Lookup(i))
		}
	}
	array.Resize(other.Length())
}

// Last returns the last element of the array. If the array is empty, fails and returns null. See also [First].
func Last[T any](array Of[T]) T {
	if array == nil || array.Length() == 0 {
		var zero T
		return zero
	}
	return array.Lookup(-1)
}

// BinarySearch returns the index of value in the sorted array. If it cannot be found, returns where value should
// be inserted to keep the array sorted. The algorithm used is binary search. The returned index comes before all
// existing elements equal to value in the array.
//
// Note: Calling BinarySearch on an unsorted array will result in unexpected behavior. Use [Sort] before calling
// this method.
func BinarySearch[T ordered](array Of[T], value T, before bool) int {
	return BinarySearchFunc(func(a, b T) bool { return a < b }, array, value, before)
}

// BinarySearchFunc returns the index of value in the sorted array. If it cannot be found, returns
// where value should be inserted to keep the array sorted (using func for the comparisons). The
// algorithm used is binary search.
//
// Similar to [SortFunc], func is called as many times as necessary, receiving one array element
// and value as arguments. The function should return true if the array element should be behind
// value, otherwise it should return false.
//
// If before is true (as by default), the returned index comes before all existing elements equal
// to value in the array.
func BinarySearchFunc[T comparable](fn func(T, T) bool, array Of[T], value T, before bool) int {
	if array == nil {
		return -1
	}
	var lo = 0
	var hi = array.Length()
	if before {
		for lo < hi {
			var mid = (lo + hi) / 2
			if fn(array.Lookup(mid), value) {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
	} else {
		for lo < hi {
			var mid = (lo + hi) / 2
			if fn(value, array.Lookup(mid)) {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
	}
	return lo
}

// Clear removes all elements from the array. This is equivalent to using resize with a size of 0.
func Clear[T any](array Of[T]) {
	if array == nil {
		return
	}
	array.Resize(0)
}

// Count returns the number of times an element is in the array.
func Count[T comparable](array Of[T], value T) int {
	if array == nil {
		return 0
	}
	var count int
	for i := range array.Length() {
		if array.Lookup(i) == value {
			count++
		}
	}
	return count
}

// Duplicate returns a new copy of the array.
//
// A shallow copy is returned: all nested Array and Dictionary elements are
// shared with the original array. Modifying them in one array will also affect
// them in the other.
func Duplicate[T any](array Of[T]) Of[T] {
	if array == nil {
		return nil
	}
	var copy = New[T]()
	Assign(copy, array)
	return copy
}

// Erase finds and removes the first occurrence of value from the array. If value does not
// exist in the array, nothing happens. To remove an element by index, use [Remove] instead.
//
// Note: This method shifts every element's index after the removed value back, which may
// have a noticeable performance cost, especially on larger arrays.
//
// Note: Erasing elements while iterating over arrays is not supported and will result in
// unpredictable behavior.
func Erase[T comparable](array Of[T], value T) {
	if array == nil {
		return
	}
	var index = Find(array, value)
	if index != -1 {
		Remove(array, index)
	}
}

// Fill assigns the given value to all elements in the array.
func Fill[T any](array Of[T], value T) {
	if array == nil {
		return
	}
	for i := range array.Length() {
		array.Mutate(i, value)
	}
}

// Insert inserts a new element (value) at a given index (position) in the array.
// position should be between 0 and the array's size.
//
// Note: Every element's index after position needs to be shifted forward, which
// may have a noticeable performance cost, especially on larger arrays.
func Insert[T any](array Of[T], position int, value T) {
	if array == nil {
		return
	}
	index := Int.Posmod(position, array.Length())
	array.Resize(array.Length() + 1)
	for i := array.Length() - 1; i > index; i-- {
		array.Mutate(i, array.Lookup(i-1))
	}
	array.Mutate(index, value)
}

// Filter calls the given function on each element in the array and returns a new, filtered Array.
// See also [Any], [All], [Map] and [Reduce].
func Filter[T any](fn func(T) bool, array Of[T]) Of[T] {
	if array == nil {
		return nil
	}
	var result = New[T]()
	for i := range array.Length() {
		if fn(array.Lookup(i)) {
			Append(result, array.Lookup(i))
		}
	}
	return result
}

// Find returns the index of the first occurrence of what in this array, or -1 if there are none.
//
// Note: If you just want to know whether the array contains what, use [Has].
//
// Note: For performance reasons, the search is affected by what's Variant.Type. For example, 7
// (int) and 7.0 (float) are not considered equal for this method.
func Find[T comparable](array Of[T], what T) int {
	if array == nil {
		return -1
	}
	for i := range array.Length() {
		if array.Lookup(i) == what {
			return i
		}
	}
	return -1
}

// First returns the first element of the array. See also [Last].
func First[T any](array Of[T]) T {
	if array == nil || array.Length() == 0 {
		return [1]T{}[0]
	}
	return array.Lookup(0)
}

// Type returns the type of the elements in the array.
func Type[T any](array Of[T]) reflect.Type { return reflect.TypeOf([0]T{}).Elem() }

// Has returns true if the array contains the given value.
func Has[T comparable](array Of[T], value T) bool { return Find(array, value) != -1 }

// Hash returns a hashed 32-bit integer value representing the array and its contents.
//
// Note: Arrays with equal hash values are not guaranteed to be the same, as a result
// of hash collisions. On the countrary, arrays with different hash values are
// guaranteed to be different.
func Hash[T any](array Of[T]) uint32 {
	panic("not implemented") // TODO/FIXME
}

// IsEmpty returns true if the array is empty ([]). See also [Size].
func IsEmpty[T any](array Of[T]) bool { return array == nil || array.Length() == 0 }

// IsReadOnly returns true if the array is read-only.
func IsReadOnly[T any](array Of[T]) bool { return array == nil || array.Static() }

// IsTyped returns true if the array is typed. Typed arrays can only
// contain elements of a specific type, as defined by the typed array constructor.
func IsTyped[T any](array Of[T]) bool { return Type(array).Kind() != reflect.Interface }

// MakeReadOnly makes the array read-only. The array's elements cannot be overridden with
// different values, and their order cannot change. Does not apply to nested elements,
// such as dictionaries.
func MakeReadOnly[T any](array Of[T]) {
	if array == nil {
		return
	}
	array.Freeze()
}

// Map calls the given function for each element in the array and returns a new array
// filled with values returned by the method.
func Map[T, U any](fn func(T) U, array Of[T]) Of[U] {
	if array == nil {
		return nil
	}
	var result = New[U]()
	result.Resize(array.Length())
	for i := range array.Length() {
		result.Mutate(i, fn(array.Lookup(i)))
	}
	return result
}

// Max returns the maximum value contained in the array, if all elements can be compared.
// Otherwise, returns the zero value for T. See also [Min].
//
// To find the maximum value using a custom comparator, you can use [Reduce].
func Max[T ordered](array Of[T]) T {
	if array == nil || array.Length() == 0 {
		return [1]T{}[0]
	}
	var max = array.Lookup(0)
	for i := 1; i < array.Length(); i++ {
		if array.Lookup(i) > max {
			max = array.Lookup(i)
		}
	}
	return max
}

// Min returns the minimum value contained in the array, if all elements can be compared.
// Otherwise, returns the zero value for T. See also [Max].
//
// To find the minimum value using a custom comparator, you can use [Reduce].
func Min[T ordered](array Of[T]) T {
	if array == nil || array.Length() == 0 {
		return [1]T{}[0]
	}
	var min = array.Lookup(0)
	for i := 1; i < array.Length(); i++ {
		if array.Lookup(i) < min {
			min = array.Lookup(i)
		}
	}
	return min
}

// PickRandom returns a random element from the array. Generates an error and returns
// the zero value for T if the array is empty.
func PickRandom[T any](array Of[T]) T {
	if array == nil || array.Length() == 0 {
		return [1]T{}[0]
	}
	return array.Lookup(rand.Intn(array.Length()))
}

// PopAt removes and returns the element of the array at index position. If negative,
// position is considered relative to the end of the array. Returns the zero value for
// T if the array is empty. If position is out of bounds, an error message is also generated.
//
// Note: This method shifts every element's index after position back, which may have
// a noticeable performance cost, especially on larger arrays.
func PopAt[T any](array Of[T], position int) T {
	if array == nil || array.Length() == 0 {
		return [1]T{}[0]
	}
	index := Int.Posmod(position, array.Length())
	value := array.Lookup(index)
	for i := index; i < array.Length()-1; i++ {
		array.Mutate(i, array.Lookup(i+1))
	}
	array.Resize(array.Length() - 1)
	return value
}

// PopBack removes and returns the last element of the array. Returns the zero value for
// T if the array is empty, without generating an error. See also [PopFront].
func PopBack[T any](array Of[T]) T { return PopAt(array, -1) }

// PopFront removes and returns the first element of the array. Returns the zero value for
// T if the array is empty, without generating an error. See also [PopBack].
func PopFront[T any](array Of[T]) T { return PopAt(array, 0) }

// PushBack appends an element at the end of the array. See also [PushFront].
func PushBack[T any](array Of[T], value T) {
	if array == nil {
		return
	}
	array.Resize(array.Length() + 1)
	array.Mutate(array.Length()-1, value)
}

// PushFront prepends an element at the beginning of the array. See also [PushBack].
//
// Note: This method shifts every other element's index forward, which may have a
// noticeable performance cost, especially on larger arrays.
func PushFront[T any](array Of[T], value T) {
	if array == nil {
		return
	}
	array.Resize(array.Length() + 1)
	for i := array.Length() - 1; i > 0; i-- {
		array.Mutate(i, array.Lookup(i-1))
	}
	array.Mutate(0, value)
}

// Reduce Calls the given function for each element in array, accumulates the result in accum,
// then returns it.
//
// The method takes two arguments: the current value of accum and the current array element.
// If accum is null (as by default), the iteration will start from the second element, with
// the first one used as initial value of accum.
func Reduce[T, U any](array Of[T], fn func(U, T) U, accum U) U {
	if array == nil {
		return accum
	}
	for i := range array.Length() {
		accum = fn(accum, array.Lookup(i))
	}
	return accum
}

// Remove removes the element from the array at the given index (position).
//
// If you need to return the removed element, use [PopAt]. To remove an element
// by value, use [Erase] instead.
//
// Note: This method shifts every element's index after position back, which may
// have a noticeable performance cost, especially on larger arrays.
func Remove[T any](array Of[T], position int) {
	if array == nil {
		return
	}
	index := Int.Posmod(position, array.Length())
	for i := index; i < array.Length()-1; i++ {
		array.Mutate(i, array.Lookup(i+1))
	}
	array.Resize(array.Length() - 1)
}

// Resize changes the size of the array. If the new size is smaller than the current size,
// the array is truncated. If the new size is larger, the array is padded with the zero
// value for T.
func Resize[T any](array Of[T], size int) {
	if array == nil {
		return
	}
	array.Resize(size)
}

// Reverse reverses the order of elements in the array.
func Reverse[T any](array Of[T]) {
	if array == nil {
		return
	}
	for i, j := 0, array.Length()-1; i < j; i, j = i+1, j-1 {
		v, w := array.Lookup(i), array.Lookup(j)
		array.Mutate(i, w)
		array.Mutate(j, v)
	}
}

// FindLast returns the index of the last occurrence of what in this array,
// or -1 if there are none.
func FindLast[T comparable](array Of[T], what T) int {
	if array == nil {
		return -1
	}
	for i := array.Length() - 1; i >= 0; i-- {
		if array.Lookup(i) == what {
			return i
		}
	}
	return -1
}

// Shuffle shuffles all elements of the array in a random order.
func Shuffle[T any](array Of[T]) {
	if array == nil {
		return
	}
	for i := array.Length() - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		v, w := array.Lookup(i), array.Lookup(j)
		array.Mutate(i, w)
		array.Mutate(j, v)
	}
}

// Size returns the number of elements in the array. Empty arrays ([]) always return 0. See also [IsEmpty].
func Size[T any](array Of[T]) int {
	if array == nil {
		return 0
	}
	return array.Length()
}

// Slice returns a new Array containing this array's elements, from index begin (inclusive) to end (exclusive).
// If either begin or end are negative, their value is relative to the end of the array.
func Slice[T any](array Of[T], from, upto int) Of[T] {
	if array == nil {
		return nil
	}
	begin := from
	if begin < 0 {
		begin += array.Length()
	}
	end := upto
	if end < 0 {
		end += array.Length()
	}
	begin = min(begin, end)
	end = max(begin, end)
	begin = Int.Clamp(begin, 0, array.Length())
	end = Int.Clamp(end, 0, array.Length())
	if begin >= end {
		return nil
	}
	result := New[T]()
	result.Resize(end - begin)
	for i := begin; i < end; i++ {
		result.Mutate(i-begin, array.Lookup(i))
	}
	return result
}

type sorter[T any] struct {
	array Of[T]
	less  func(T, T) bool
}

func (s sorter[T]) Swap(i, j int) {
	v, w := s.array.Lookup(i), s.array.Lookup(j)
	s.array.Mutate(i, w)
	s.array.Mutate(j, v)
}
func (s sorter[T]) Less(i, j int) bool {
	return s.less(s.array.Lookup(i), s.array.Lookup(j))
}
func (s sorter[T]) Len() int { return s.array.Length() }

// Sort sorts the array in ascending order. The final order is dependent on the
// "less than" (<) comparison between elements.
func Sort[T ordered](array Of[T]) {
	if array == nil {
		return
	}
	sort.Sort(sorter[T]{array: array, less: func(a, b T) bool { return a < b }})
}

// SortFunc sorts the array using a custom function.
//
// fn is called as many times as necessary, receiving two array elements as arguments.
// The function should return true if the first element should be moved before the
// second one, otherwise it should return false.
func SortFunc[T any](less func(a, b T) bool, array Of[T]) {
	if array == nil {
		return
	}
	sort.Sort(sorter[T]{array: array, less: less})
}

// Iter returns a new iterator for this array.
func Iter[T any](array Of[T]) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i := range array.Length() {
			if !yield(i, array.Lookup(i)) {
				return
			}
		}
	}
}
