/*
# Create a ZIP archive with a single file at its root.
func write_zip_file():
    var writer = ZIPPacker.new()
    var err = writer.open("user://archive.zip")
    if err != OK:
        return err
    writer.start_file("hello.txt")
    writer.write_file("Hello World".to_utf8_buffer())
    writer.close_file()

    writer.close()
    return OK
*/

package main

import "graphics.gd/classdb/ZIPPacker"

func WriteZipFile() error {
	var writer = ZIPPacker.New()
	var err = writer.Open("user://archive.zip")
	if err != nil {
		return err
	}
	writer.StartFile("hello.txt")
	writer.WriteFile([]byte("Hello World"))
	writer.CloseFile()
	writer.Close()
	return nil
}
