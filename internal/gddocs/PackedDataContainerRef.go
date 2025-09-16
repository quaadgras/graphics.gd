/*
var packed = PackedDataContainer.new()
packed.pack([1, 2, 3, ["nested1", "nested2"], 4, 5, 6])

for element in packed:
	if element is PackedDataContainerRef:
		for subelement in element:
			print("::", subelement)
	else:
		print(element)
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/PackedDataContainer"
	"graphics.gd/classdb/PackedDataContainerRef"
	"graphics.gd/variant/Object"
)

func ExamplePackedDataContainerRef() {
	var packed = PackedDataContainer.New()
	packed.Pack([]any{1, 2, 3, []any{"nested1", "nested2"}, 4, 5, 6})
	for element := range Object.Iter(packed) {
		if ref, ok := element.(PackedDataContainerRef.Instance); ok {
			for subelement := range Object.Iter(ref) {
				fmt.Println("::", subelement)
			}
		} else {
			fmt.Println(element)
		}

	}
}
