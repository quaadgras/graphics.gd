/*
[gdscript]
func dir_contents(path):
    var dir = DirAccess.open(path)
    if dir:
        dir.list_dir_begin()
        var file_name = dir.get_next()
        while file_name != "":
            if dir.current_is_dir():
                print("Found directory: " + file_name)
            else:
                print("Found file: " + file_name)
            file_name = dir.get_next()
    else:
        print("An error occurred when trying to access the path.")
[/gdscript]
[csharp]
public void DirContents(string path)
{
    using var dir = DirAccess.Open(path);
    if (dir != null)
    {
        dir.ListDirBegin();
        string fileName = dir.GetNext();
        while (fileName != "")
        {
            if (dir.CurrentIsDir())
            {
                GD.Print($"Found directory: {fileName}");
            }
            else
            {
                GD.Print($"Found file: {fileName}");
            }
            fileName = dir.GetNext();
        }
    }
    else
    {
        GD.Print("An error occurred when trying to access the path.");
    }
}
[/csharp]
*/

package main

import "graphics.gd/classdb/DirAccess"

func ExampleDirectoryList(path string) {
	var dir = DirAccess.Open(path)
	if dir != DirAccess.Nil {
		dir.ListDirBegin()
		var fileName = dir.GetNext()
		for fileName != "" {
			if dir.CurrentIsDir() {
				println("Found directory: " + fileName)
			} else {
				println("Found file: " + fileName)
			}
			fileName = dir.GetNext()
		}
	} else {
		println("An error occurred when trying to access the path.")
	}
}
