/*
[gdscript]
while file.get_position() < file.get_length():
	# Read data
[/gdscript]
[csharp]
while (file.GetPosition() < file.GetLength())
{
	// Read data
}
[/csharp]
*/

package main

import "graphics.gd/classdb/FileAccess"

func ExampleFileAccessEofReached(file FileAccess.Instance) {
	for file.GetPosition() < file.GetLength() {
		// Read data
	}
}
