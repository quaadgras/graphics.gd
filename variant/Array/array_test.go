package Array_test

import (
	"testing"

	"grow.graphics/gd/gdvalue/Array"
	"grow.graphics/gd/internal/gdtests"
)

func TestAll(t *testing.T) {
	greater_than_5 := func(number int) bool {
		return number > 5
	}
	gdtests.Print(t, "true", Array.All(greater_than_5, Array.New(6, 10, 6)))  // Prints true (3/3 elements evaluate to true).
	gdtests.Print(t, "false", Array.All(greater_than_5, Array.New(4, 10, 4))) // Prints false (1/3 elements evaluate to true).
	gdtests.Print(t, "false", Array.All(greater_than_5, Array.New(4, 4, 4)))  // Prints false (0/3 elements evaluate to true).
	gdtests.Print(t, "true", Array.All(greater_than_5, Array.New[int]()))     // Prints true (0/0 elements evaluate to true).
}

func TestAppendArray(t *testing.T) {
	var numbers = Array.New(1, 2, 3)
	var extra = Array.New(4, 5, 6)
	Array.AppendTo(numbers, extra)
	gdtests.Print(t, "[1, 2, 3, 4, 5, 6]", numbers) // Prints [1, 2, 3, 4, 5, 6]
}

func TestBinarySearch(t *testing.T) {
	var numbers = Array.New(2, 4, 8, 10)
	var idx = Array.BinarySearch(numbers, 7, true)

	Array.Insert(numbers, idx, 7)
	gdtests.Print(t, "[2, 4, 7, 8, 10]", numbers) // Prints [2, 4, 7, 8, 10]

	var fruits = Array.New("Apple", "Lemon", "Lemon", "Orange")
	gdtests.Print(t, "1", (Array.BinarySearch(fruits, "Lemon", true)))  // Prints 1, points at the first "Lemon".
	gdtests.Print(t, "3", (Array.BinarySearch(fruits, "Lemon", false))) // Prints 3, points after the last "Lemon".
}

func TestBinarySearchFunc(t *testing.T) {
	var sort_by_amount = func(a, b [2]any) bool {
		return a[1].(int) < b[1].(int)
	}
	var my_items = Array.New([2]any{"Tomato", 2}, [2]any{"Kiwi", 5}, [2]any{"Rice", 9})

	var apple = [2]any{"Apple", 5}
	// "Apple" is inserted before "Kiwi".
	Array.Insert(my_items, Array.BinarySearchFunc(sort_by_amount, my_items, apple, true), apple)

	var banana = [2]any{"Banana", 5}
	// "Banana" is inserted after "Kiwi".
	Array.Insert(my_items, Array.BinarySearchFunc(sort_by_amount, my_items, banana, false), banana)

	// Prints [["Tomato", 2], ["Apple", 5], ["Kiwi", 5], ["Banana", 5], ["Rice", 9]]
	gdtests.Print(t, `[["Tomato", 2], ["Apple", 5], ["Kiwi", 5], ["Banana", 5], ["Rice", 9]]`, my_items)
}

func TestFill(t *testing.T) {
	var array = Array.New[int]()
	array.Resize(5)
	Array.Fill(array, 2)
	gdtests.Print(t, "[2, 2, 2, 2, 2]", array) // Prints [2, 2, 2, 2, 2]
}

func TestHas(t *testing.T) {
	gdtests.Print(t, "true", Array.Has(Array.New[any]("inside", 7), "inside"))   // Prints true
	gdtests.Print(t, "false", Array.Has(Array.New[any]("inside", 7), "outside")) // Prints false
	gdtests.Print(t, "true", Array.Has(Array.New[any]("inside", 7), 7))          // Prints true
	gdtests.Print(t, "false", Array.Has(Array.New[any]("inside", 7), "7"))       // Prints false
}

func TestIsTyped(t *testing.T) {
	var numbers = Array.New(0.2, 4.2, -2.0)
	gdtests.Print(t, "true", Array.IsTyped(numbers)) // Prints true
}

func TestMap(t *testing.T) {
	var double = func(number int) int {
		return number * 2
	}
	gdtests.Print(t, "[2, 4, 6]", Array.Map(double, Array.New(1, 2, 3))) // Prints [2, 4, 6]
}

func TestReduce(t *testing.T) {
	var sum = func(accum, number int) int {
		return accum + number
	}
	gdtests.Print(t, "6", Array.Reduce(Array.New(1, 2, 3), sum, 0))   // Prints 6
	gdtests.Print(t, "16", Array.Reduce(Array.New(1, 2, 3), sum, 10)) // Prints 16
}

func TestSlice(t *testing.T) {
	var letters = Array.New("A", "B", "C", "D", "E", "F")

	gdtests.Print(t, `["A", "B"]`, Array.Slice(letters, 0, 2))  // Prints ["A", "B"]
	gdtests.Print(t, `["C", "D"]`, Array.Slice(letters, 2, -2)) // Prints ["C", "D"]
	gdtests.Print(t, `["E", "F"]`, Array.Slice(letters, -2, 6)) // Prints ["E", "F"]
}

func TestSort(t *testing.T) {
	var numbers = Array.New(10, 5, 2, 8)
	Array.Sort(numbers)
	gdtests.Print(t, "[2, 5, 8, 10]", numbers) // Prints [2, 5, 8, 10]
}

func TestSortFunc(t *testing.T) {
	sort_ascending := func(a, b [2]any) bool {
		return a[1].(int) < b[1].(int)
	}
	var my_items = Array.New([2]any{"Tomato", 5}, [2]any{"Apple", 9}, [2]any{"Rice", 4})
	Array.SortFunc(sort_ascending, my_items)
	gdtests.Print(t, `[["Rice", 4], ["Tomato", 5], ["Apple", 9]]`, my_items) // Prints [["Rice", 4], ["Tomato", 5], ["Apple", 9]]
}
