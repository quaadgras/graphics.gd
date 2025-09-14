/*
[gdscript]
print(Engine.has_singleton("OS"))          # Prints true
print(Engine.has_singleton("Engine"))      # Prints true
print(Engine.has_singleton("AudioServer")) # Prints true
print(Engine.has_singleton("Unknown"))     # Prints false
[/gdscript]
[csharp]
GD.Print(Engine.HasSingleton("OS"));          // Prints True
GD.Print(Engine.HasSingleton("Engine"));      // Prints True
GD.Print(Engine.HasSingleton("AudioServer")); // Prints True
GD.Print(Engine.HasSingleton("Unknown"));     // Prints False
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/Engine"
)

func Engine_HasSingleton() {
	fmt.Println(Engine.HasSingleton("OS"))          // Prints true
	fmt.Println(Engine.HasSingleton("Engine"))      // Prints true
	fmt.Println(Engine.HasSingleton("AudioServer")) // Prints true
	fmt.Println(Engine.HasSingleton("Unknown"))     // Prints false
}
