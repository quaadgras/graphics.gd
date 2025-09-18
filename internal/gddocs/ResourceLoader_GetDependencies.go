/*
for dependency in ResourceLoader.get_dependencies(path):
	if dependency.contains("::"):
		print(dependency.get_slice("::", 0)) # Prints the UID.
		print(dependency.get_slice("::", 2)) # Prints the fallback path.
	else:
		print(dependency) # Prints the path.
*/

package main

import (
	"fmt"
	"strings"

	"graphics.gd/classdb/ResourceLoader"
)

var path string

func ResourceLoader_GetDependencies() {
	for _, dependency := range ResourceLoader.GetDependencies(path) {
		if strings.Contains(dependency, "::") {
			fmt.Println(strings.Split(dependency, "::")[0]) // Prints the UID.
			fmt.Println(strings.Split(dependency, "::")[2]) // Prints the fallback path.
		} else {
			fmt.Println(dependency) // Prints the path.
		}
	}
}
