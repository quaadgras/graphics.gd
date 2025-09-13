/*
[gdscript]
var packer = PCKPacker.new()
packer.pck_start("test.pck")
packer.add_file("res://text.txt", "text.txt")
packer.flush()
[/gdscript]
[csharp]
var packer = new PckPacker();
packer.PckStart("test.pck");
packer.AddFile("res://text.txt", "text.txt");
packer.Flush();
[/csharp]
*/

package main

import "graphics.gd/classdb/PCKPacker"

func ExamplePCKPacker() {
	var packer = PCKPacker.New()
	packer.PckStart("test.pck")
	packer.AddFile("res://text.txt", "text.txt")
	packer.Flush()
}
