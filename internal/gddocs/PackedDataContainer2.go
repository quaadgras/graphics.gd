/*
var container = load("packed_data.res")
for key in container:
    prints(key, container[key])
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/PackedDataContainer"
	"graphics.gd/classdb/Resource"
	"graphics.gd/variant/Object"
)

func ExamplePackedDataContainerLoad() {
	var container = Resource.Load[PackedDataContainer.Instance]("packed_data.res")
	for _, prop := range Object.GetPropertyList(container) {
		fmt.Println(prop.Name, Object.Get(container, prop.Name))
	}
}
