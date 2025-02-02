package Packed

import (
	GenericArray "graphics.gd/variant/Array"
)

// Sort sorts the array in ascending order. The final order is dependent on the
// "less than" (<) comparison between elements.
func Sort[T ~byte | ~float32 | ~float64 | ~int32 | ~int64](array Array[T]) { //gd:PackedArray.sort
	GenericArray.Sort((GenericArray.Contains[T])(array))
}

// BinarySearch returns the index of value in the sorted array. If it cannot be found, returns where value should
// be inserted to keep the array sorted. The algorithm used is binary search. The returned index comes before all
// existing elements equal to value in the array.
//
// Note: Calling BinarySearch on an unsorted array will result in unexpected behavior. Use [Sort] before calling
// this method.
func BinarySearch[T ~byte | ~float32 | ~float64 | ~int32 | ~int64](array Array[T], value T, before bool) int { //gd:PackedArray.bsearch
	return GenericArray.BinarySearch((GenericArray.Contains[T])(array), value, before)
}
