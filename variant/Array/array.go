// Package Array provides a data structure that holds a sequence of elements.
package Array

import (
	"cmp"
	"iter"
	"math/rand"
	"reflect"
	"sort"

	"graphics.gd/variant"
	"graphics.gd/variant/Int"
)

// Contains is an array data structure that can contain a sequence of elements of T. Elements
// are accessed by a numerical index starting at 0. Negative indices are used to count
// from the back (-1 is the last element, -2 is the second to last, etc.).
//
// Note: Arrays are always passed by reference. To get a copy of an array that can be modified
// independently of the original array, use duplicate.
type Contains[T any] struct {
	state State
	proxy Proxy[T]
}

// Any is an array that can contain any type of element, equivalent to [[]any].
type Any = Contains[variant.Any]

// Nil reference array.
var Nil Any

// New creates a new array with the given elements.
func New[T any](elements ...T) Contains[T] {
	return Contains[T]{
		proxy: &localFirst[T]{
			slice: elements,
		},
	}
}

// Any returns an [Any] array with a shared view on the elements in the array.
func (a *Contains[T]) Any() Any {
	if a.proxy == nil {
		a.proxy = new(localFirst[T])
	}
	return a.proxy.Any(a.state)
}

// Len returns the number of elements in the array.
func (a Contains[T]) Len() int { //gd:Array.size
	if a.proxy == nil {
		return 0
	}
	return a.proxy.Len(a.state)
}

// Index returns the value at the given index. If the index is negative,
// it counts from the end of the array.
func (a Contains[T]) Index(i int) T { //gd:Array[]
	if a.proxy == nil {
		panic("index out of range")
	}
	return a.proxy.Index(a.state, i)
}

// SetIndex sets the value at the given index. If the index is negative,
// it counts from the end of the array.
func (a *Contains[T]) SetIndex(i int, value T) { //gd:Array[]=
	if a.proxy == nil {
		panic("index out of range")
	}
	a.proxy.SetIndex(a.state, i, value)
}

// All calls the given function on each element in the array and returns true if the fn returns
// true for all elements in the array. If the fn returns false for one array element or more,
// this method returns false. See also [Any], [Filter], [Map], and [Reduce].
func All[T any](fn func(T) bool, array Contains[T]) bool { //gd:Array.all
	for _, v := range array.Iter() {
		if !fn(v) {
			return false
		}
	}
	return true
}

// IfAny calls the given function on each element in the array and returns true if the function
// returns true for one or more elements in the array. If the function returns false for all
// elements in the array, this method returns false. See also [Any], [Filter], [Map], and [Reduce].
func IfAny[T any](fn func(T) bool, array Contains[T]) bool { //gd:Array.any
	for _, v := range array.Iter() {
		if fn(v) {
			return true
		}
	}
	return false
}

// Append appends value at the end of the array (alias of PushBack).
func (a *Contains[T]) Append(value T) { //gd:Array.append
	if a.proxy == nil {
		a.proxy = new(localFirst[T])
	}
	l := a.Len()
	a.Resize(l + 1)
	a.SetIndex(l, value)
}

// AppendTo appends another array at the end of this array.
func (a *Contains[T]) AppendArray(other Contains[T]) { //gd:Array.append_array
	for _, v := range other.Iter() {
		a.Append(v)
	}
}

// Assign assigns elements of another array into the array. Resizes the
// array to match array. Performs type conversions if the array is typed.
func (a *Contains[T]) Assign(other Contains[T]) { //gd:Array.assign
	a.Resize(other.Len())
	for i, v := range other.Iter() {
		a.SetIndex(i, v)
	}
}

// Back returns the last element of the array. If the array is empty, fails and returns the zero value for T. See also [First].
func Last[T any](a Contains[T]) T { return a.Index(-1) } //gd:Array.back

// BinarySearch returns the index of value in the sorted array. If it cannot be found, returns where value should
// be inserted to keep the array sorted. The algorithm used is binary search. The returned index comes before all
// existing elements equal to value in the array.
//
// Note: Calling BinarySearch on an unsorted array will result in unexpected behavior. Use [Sort] before calling
// this method.
func BinarySearch[T cmp.Ordered](array Contains[T], value T, before bool) int { //gd:Array.bsearch
	return array.BinarySearchFunc(func(a, b T) bool { return a < b }, value, before)
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
func (a Contains[T]) BinarySearchFunc(fn func(T, T) bool, value T, before bool) int { //gd:Array.bsearch_custom
	var lo = 0
	var hi = a.Len()
	if before {
		for lo < hi {
			var mid = (lo + hi) / 2
			if fn(a.Index(mid), value) {
				lo = mid + 1
			} else {
				hi = mid
			}
		}
	} else {
		for lo < hi {
			var mid = (lo + hi) / 2
			if fn(value, a.Index(mid)) {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
	}
	return lo
}

// Clear removes all elements from the array. This is equivalent to using resize with a size of 0.
func Clear[T any](array Contains[T]) { array.Resize(0) } //gd:Array.clear

// Count returns the number of times an element is in the array.
func Count[T comparable](array Contains[T], value T) int { //gd:Array.count
	var count int
	for i := range array.Len() {
		if array.Index(i) == value {
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
func Duplicate[T any](array Contains[T]) Contains[T] { //gd:Array.duplicate
	var copy = New[T]()
	copy.Assign(array)
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
func Erase[T comparable](array Contains[T], value T) { //gd:Array.erase
	var index = Find(array, value)
	if index != -1 {
		Remove(array, index)
	}
}

// Fill assigns the given value to all elements in the array.
func Fill[T any](array Contains[T], value T) { //gd:Array.fill
	for i := range array.Len() {
		array.SetIndex(i, value)
	}
}

// Insert inserts a new element (value) at a given index (position) in the array.
// position should be between 0 and the array's size.
//
// Note: Every element's index after position needs to be shifted forward, which
// may have a noticeable performance cost, especially on larger arrays.
func (array *Contains[T]) Insert(position int, value T) { //gd:Array.insert
	index := Int.Posmod(position, array.Len())
	array.Resize(array.Len() + 1)
	for i := array.Len() - 1; i > index; i-- {
		array.SetIndex(i, array.Index(i-1))
	}
	array.SetIndex(index, value)
}

// Filter calls the given function on each element in the array and returns a new, filtered Array.
// See also [Any], [All], [Map] and [Reduce].
func Filter[T any](fn func(T) bool, array Contains[T]) Contains[T] { //gd:Array.filter
	var result = New[T]()
	for i := range array.Len() {
		if fn(array.Index(i)) {
			array.Append(array.Index(i))
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
func Find[T comparable](array Contains[T], what T) int { //gd:Array.find
	for i := range array.Len() {
		if array.Index(i) == what {
			return i
		}
	}
	return -1
}

// First returns the first element of the array. See also [Last].
func First[T any](array Contains[T]) T { //gd:Array.front
	return array.Index(0)
}

// Type returns the type of the elements in the array.
func Type[T any](array Contains[T]) reflect.Type { return reflect.TypeOf([0]T{}).Elem() } //gd:Array.get_typed_builtin Array.is_same_typed Array.get_typed_class_name Array.get_typed_script

// Has returns true if the array contains the given value.
func Has[T comparable](value T, array Contains[T]) bool { return Find(array, value) != -1 } //gd:Array.has

// Hash returns a hashed 32-bit integer value representing the array and its contents.
//
// Note: Arrays with equal hash values are not guaranteed to be the same, as a result
// of hash collisions. On the countrary, arrays with different hash values are
// guaranteed to be different.
func Hash[T any](array Contains[T]) uint32 { //gd:Array.hash
	panic("not implemented") // TODO/FIXME
}

// IsEmpty returns true if the array is empty ([]). See also [Size].
func IsEmpty[T any](array Contains[T]) bool { return array.Len() == 0 } //gd:Array.is_empty

// IsReadOnly returns true if the array is read-only.
func IsReadOnly[T any](array Contains[T]) bool { //gd:Array.is_read_only
	if array.proxy == nil {
		return false
	}
	return array.proxy.IsReadOnly(array.state)
}

// IsTyped returns true if the array is typed. Typed arrays can only
// contain elements of a specific type, as defined by the typed array constructor.
func IsTyped[T any](array Contains[T]) bool { return Type(array).Kind() != reflect.Interface } //gd:Array.is_typed

// MakeReadOnly makes the array read-only. The array's elements cannot be overridden with
// different values, and their order cannot change. Does not apply to nested elements,
// such as dictionaries.
func (a *Contains[T]) MakeReadOnly() { //gd:Array.make_read_only
	if a.proxy == nil {
		a.proxy = new(localFirst[T])
	}
	a.proxy.MakeReadOnly(a.state)
}

// Map calls the given function for each element in the array and returns a new array
// filled with values returned by the method.
func Map[T, U any](fn func(T) U, array Contains[T]) Contains[U] { //gd:Array.map
	var result = New[U]()
	result.Resize(array.Len())
	for i, v := range array.Iter() {
		result.SetIndex(i, fn(v))
	}
	return result
}

// Max returns the maximum value contained in the array, if all elements can be compared.
// Otherwise, returns the zero value for T. See also [Min].
//
// To find the maximum value using a custom comparator, you can use [Reduce].
func Max[T cmp.Ordered](array Contains[T]) T { //gd:Array.max
	if array.Len() == 0 {
		return [1]T{}[0]
	}
	var max = array.Index(0)
	for _, v := range array.Iter() {
		if v > max {
			max = v
		}
	}
	return max
}

// Min returns the minimum value contained in the array, if all elements can be compared.
// Otherwise, returns the zero value for T. See also [Max].
//
// To find the minimum value using a custom comparator, you can use [Reduce].
func Min[T cmp.Ordered](array Contains[T]) T { //gd:Array.min
	if array.Len() == 0 {
		return [1]T{}[0]
	}
	var min = array.Index(0)
	for _, v := range array.Iter() {
		if v < min {
			min = v
		}
	}
	return min
}

// PickRandom returns a random element from the array. Generates an error and returns
// the zero value for T if the array is empty.
func PickRandom[T any](array Contains[T]) T { //gd:Array.pick_random
	if array.Len() == 0 {
		return [1]T{}[0]
	}
	return array.Index(rand.Intn(array.Len()))
}

// PopAt removes and returns the element of the array at index position. If negative,
// position is considered relative to the end of the array. Returns the zero value for
// T if the array is empty. If position is out of bounds, an error message is also generated.
//
// Note: This method shifts every element's index after position back, which may have
// a noticeable performance cost, especially on larger arrays.
func PopAt[T any](array Contains[T], position int) T { //gd:Array.pop_at
	if array.Len() == 0 {
		return [1]T{}[0]
	}
	index := Int.Posmod(position, array.Len())
	value := array.Index(index)
	for i := index; i < array.Len()-1; i++ {
		array.SetIndex(i, array.Index(i+1))
	}
	array.Resize(array.Len() - 1)
	return value
}

// PopBack removes and returns the last element of the array. Returns the zero value for
// T if the array is empty, without generating an error. See also [PopFront].
func PopBack[T any](array Contains[T]) T { return PopAt(array, -1) } //gd:Array.pop_back

// PopFront removes and returns the first element of the array. Returns the zero value for
// T if the array is empty, without generating an error. See also [PopBack].
func PopFront[T any](array Contains[T]) T { return PopAt(array, 0) } //gd:Array.pop_front

// PushBack appends an element at the end of the array. See also [PushFront].
func (array *Contains[T]) PushBack(value T) { //gd:Array.push_back
	array.Resize(array.Len() + 1)
	array.SetIndex(array.Len()-1, value)
}

// PushFront prepends an element at the beginning of the array. See also [PushBack].
//
// Note: This method shifts every other element's index forward, which may have a
// noticeable performance cost, especially on larger arrays.
func (array *Contains[T]) PushFront(value T) { //gd:Array.push_front
	array.Resize(array.Len() + 1)
	for i := array.Len() - 1; i > 0; i-- {
		array.SetIndex(i, array.Index(i-1))
	}
	array.SetIndex(0, value)
}

// Reduce Calls the given function for each element in array, accumulates the result in accum,
// then returns it.
//
// The method takes two arguments: the current value of accum and the current array element.
// If accum is null (as by default), the iteration will start from the second element, with
// the first one used as initial value of accum.
func Reduce[T, U any](array Contains[T], fn func(U, T) U, accum U) U { //gd:Array.reduce
	for i := range array.Len() {
		accum = fn(accum, array.Index(i))
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
func Remove[T any](array Contains[T], position int) { //gd:Array.remove_at
	index := Int.Posmod(position, array.Len())
	for i := index; i < array.Len()-1; i++ {
		array.SetIndex(i, array.Index(i+1))
	}
	array.Resize(array.Len() - 1)
}

// Resize changes the size of the array. If the new size is smaller than the current size,
// the array is truncated. If the new size is larger, the array is padded with the zero
// value for T.
func (array *Contains[T]) Resize(size int) { //gd:Array.resize
	if array.proxy == nil && size != 0 {
		array.proxy = new(localFirst[T])
	}
	array.proxy.Resize(array.state, size)
}

// Reverse reverses the order of elements in the array.
func Reverse[T any](array Contains[T]) { //gd:Array.reverse
	for i, j := 0, array.Len()-1; i < j; i, j = i+1, j-1 {
		v, w := array.Index(i), array.Index(j)
		array.SetIndex(i, w)
		array.SetIndex(j, v)
	}
}

// FindLast returns the index of the last occurrence of what in this array,
// or -1 if there are none.
func FindLast[T comparable](array Contains[T], what T) int { //gd:Array.rfind
	for i := array.Len() - 1; i >= 0; i-- {
		if array.Index(i) == what {
			return i
		}
	}
	return -1
}

// Shuffle shuffles all elements of the array in a random order.
func Shuffle[T any](array Contains[T]) { //gd:Array.shuffle
	for i := array.Len() - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		v, w := array.Index(i), array.Index(j)
		array.SetIndex(i, w)
		array.SetIndex(j, v)
	}
}

// Slice returns a new Array containing this array's elements, from index begin (inclusive) to end (exclusive).
// If either begin or end are negative, their value is relative to the end of the array.
func Slice[T any](array Contains[T], from, upto int) Contains[T] { //gd:Array.slice
	begin := from
	if begin < 0 {
		begin += array.Len()
	}
	end := upto
	if end < 0 {
		end += array.Len()
	}
	begin = min(begin, end)
	end = max(begin, end)
	begin = Int.Clamp(begin, 0, array.Len())
	end = Int.Clamp(end, 0, array.Len())
	if begin >= end {
		return Contains[T]{}
	}
	result := New[T]()
	result.Resize(end - begin)
	for i := begin; i < end; i++ {
		result.SetIndex(i-begin, array.Index(i))
	}
	return result
}

type sorter[T any] struct {
	array Contains[T]
	less  func(T, T) bool
}

func (s sorter[T]) Swap(i, j int) {
	v, w := s.array.Index(i), s.array.Index(j)
	s.array.SetIndex(i, w)
	s.array.SetIndex(j, v)
}
func (s sorter[T]) Less(i, j int) bool {
	return s.less(s.array.Index(i), s.array.Index(j))
}
func (s sorter[T]) Len() int { return s.array.Len() }

// Sort sorts the array in ascending order. The final order is dependent on the
// "less than" (<) comparison between elements.
func Sort[T cmp.Ordered](array Contains[T]) { //gd:Array.sort
	sort.Sort(sorter[T]{array: array, less: func(a, b T) bool { return a < b }})
}

// SortFunc sorts the array using a custom function.
//
// fn is called as many times as necessary, receiving two array elements as arguments.
// The function should return true if the first element should be moved before the
// second one, otherwise it should return false.
func SortFunc[T any](less func(a, b T) bool, array Contains[T]) { //gd:Array.sort_custom
	sort.Sort(sorter[T]{array: array, less: less})
}

// Iter returns a new iterator for this array.
func (a Contains[T]) Iter() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i := range a.Len() {
			if !yield(i, a.Index(i)) {
				return
			}
		}
	}
}
