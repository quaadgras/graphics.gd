package Path_test

import (
	"testing"

	"graphics.gd/internal/gdtests"
	"graphics.gd/variant/Path"
	"graphics.gd/variant/String"
)

func TestBaseDir(t *testing.T) {
	var dir_path = Path.ToFile(String.New("/path/to/file.txt")).Directory() // dir_path is "/path/to"

	gdtests.That(t, dir_path.String(), "/path/to")
}

func TestFileExtension(t *testing.T) {
	var a = Path.ToFile(String.New("/path/to/file.txt")).FileExtension() // a is "txt"
	var b = Path.ToFile(String.New("cool.txt")).FileExtension()          // b is "txt"
	var c = Path.ToFile(String.New("cool.font.tres")).FileExtension()    // c is "tres"
	var d = Path.ToFile(String.New(".pack1")).FileExtension()            // d is "pack1"

	var e = Path.ToFile(String.New("file.txt.")).FileExtension()  // e is ""
	var f = Path.ToFile(String.New("file.txt..")).FileExtension() // f is ""
	var g = Path.ToFile(String.New("txt")).FileExtension()        // g is ""
	var h = Path.ToFile(String.New("")).FileExtension()           // h is ""

	gdtests.That(t, a.String(), "txt")
	gdtests.That(t, b.String(), "txt")
	gdtests.That(t, c.String(), "tres")
	gdtests.That(t, d.String(), "pack1")
	gdtests.That(t, e.String(), "")
	gdtests.That(t, f.String(), "")
	gdtests.That(t, g.String(), "")
	gdtests.That(t, h.String(), "")
}
