package String_test

import (
	"slices"
	"testing"

	"graphics.gd/internal/gdtests"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/String"
)

func TestBigrams(t *testing.T) {
	gdtests.Print(t, `["Ge", "et", "t ", " u", "up", "p!"]`, slices.Collect(String.Bigrams("Get up!"))) // Prints ["Ge", "et", "t ", " u", "up", "p!"]
}

func TestBinaryToInteger(t *testing.T) {
	gdtests.Print(t, "5", String.BinaryToInteger("101"))    // Prints 5
	gdtests.Print(t, "5", String.BinaryToInteger("0b101"))  // Prints 5
	gdtests.Print(t, "-2", String.BinaryToInteger("-0b10")) // Prints -2
}

func TestCapitalize(t *testing.T) {
	gdtests.That(t, String.Capitalize("move_local_x"), "Move Local X")      // Returns "Move Local X"
	gdtests.That(t, String.Capitalize("sceneFile_path"), "Scene File Path") // Returns "Scene File Path"
	gdtests.That(t, String.Capitalize("2d, fps, png"), "2d, Fps, Png")      // Returns "2d, Fps, Png"
}

func TestContains(t *testing.T) {
	gdtests.Print(t, "true", String.Contains("de", "Node")) // Prints true
	gdtests.Print(t, "false", String.Contains("I", "team")) // Prints false
}

func TestCharacter(t *testing.T) {
	gdtests.Print(t, "A", String.Character(65))     // Prints "A"
	gdtests.Print(t, "🤖", String.Character(129302)) // Prints "🤖" (robot face emoji)
}

func TestFindIndex(t *testing.T) {
	gdtests.Print(t, "-1", String.FindIndex("Team", "I"))        // Prints -1
	gdtests.Print(t, "2", String.FindIndex("Potato", "t"))       // Prints 2
	gdtests.Print(t, "4", 3+String.FindIndex("Potato"[3:], "t")) // Prints 4
	gdtests.Print(t, "-1", String.FindIndex("Potato"[5:], "t"))  // Prints -1
}

func TestFormat(t *testing.T) {
	// Prints "Waiting for Godot is a play by Samuel Beckett, and Godot Engine is named after it."
	var use_array_values = "Waiting for {0} is a play by {1}, and {0} Engine is named after it."
	gdtests.Print(t, "Waiting for Godot is a play by Samuel Beckett, and Godot Engine is named after it.", String.Format(use_array_values, "Godot", "Samuel Beckett"))

	// Prints "User 42 is Godot."
	gdtests.Print(t, "User 42 is Godot.", String.Format("User 42 is Godot.", map[string]any{"id": 42, "name": "Godot"}))
}

func TestBaseDir(t *testing.T) {
	var dir_path = String.Directory("/path/to/file.txt") // dir_path is "/path/to"

	gdtests.That(t, dir_path, "/path/to")
}

func TestFileExtension(t *testing.T) {
	var a = String.FileExtension("/path/to/file.txt") // a is "txt"
	var b = String.FileExtension("cool.txt")          // b is "txt"
	var c = String.FileExtension("cool.font.tres")    // c is "tres"
	var d = String.FileExtension(".pack1")            // d is "pack1"

	var e = String.FileExtension("file.txt.")  // e is ""
	var f = String.FileExtension("file.txt..") // f is ""
	var g = String.FileExtension("txt")        // g is ""
	var h = String.FileExtension("")           // h is ""

	gdtests.That(t, a, "txt")
	gdtests.That(t, b, "txt")
	gdtests.That(t, c, "tres")
	gdtests.That(t, d, "pack1")
	gdtests.That(t, e, "")
	gdtests.That(t, f, "")
	gdtests.That(t, g, "")
	gdtests.That(t, h, "")
}

func TestFileName(t *testing.T) {
	var file = String.FileName("/path/to/icon.png") // file is "icon.png"

	gdtests.That(t, file, "icon.png")
}

func TestExtract(t *testing.T) {
	gdtests.Print(t, "i", String.Extract("i/am/example/hi", "/", 0))       // Prints "i"
	gdtests.Print(t, "am", String.Extract("i/am/example/hi", "/", 1))      // Prints "am"
	gdtests.Print(t, "example", String.Extract("i/am/example/hi", "/", 2)) // Prints "example"
	gdtests.Print(t, "hi", String.Extract("i/am/example/hi", "/", 3))      // Prints "hi"
}

func TestDecodeHex(t *testing.T) {
	var text = "hello world"
	var encoded = String.EncodeHex(text) // outputs "68656c6c6f20776f726c64"
	gdtests.Print(t, "hello world", String.DecodeHex(encoded))
}

func TestHexToInteger(t *testing.T) {
	gdtests.Print(t, "255", String.HexToInteger("0xff")) // Prints 255
	gdtests.Print(t, "171", String.HexToInteger("ab"))   // Prints 171
}

func TestIsSubsequenceOf(t *testing.T) {
	var text = "Wow, incredible!"

	gdtests.Print(t, "true", String.IsSubsequenceOf(text, "inedible")) // Prints true
	gdtests.Print(t, "true", String.IsSubsequenceOf(text, "Word!"))    // Prints true
	gdtests.Print(t, "false", String.IsSubsequenceOf(text, "Window"))  // Prints false
	gdtests.Print(t, "true", String.IsSubsequenceOf(text, ""))         // Prints true
}

func TestIsValidFloat(t *testing.T) {
	gdtests.Print(t, "true", String.IsValidFloat("1.7"))    // Prints true
	gdtests.Print(t, "true", String.IsValidFloat("24"))     // Prints true
	gdtests.Print(t, "true", String.IsValidFloat("7e3"))    // Prints true
	gdtests.Print(t, "false", String.IsValidFloat("Hello")) // Prints false
}

func TestIsValidHexNumber(t *testing.T) {
	gdtests.Print(t, "true", String.IsValidHexNumber("A08E"))    // Prints true
	gdtests.Print(t, "true", String.IsValidHexNumber("-AbCdEf")) // Prints true
	gdtests.Print(t, "false", String.IsValidHexNumber("2.5"))    // Prints false
}

func TestIsValidIdentifier(t *testing.T) {
	gdtests.Print(t, "true", String.IsValidIdentifier("node_2d"))     // Prints true
	gdtests.Print(t, "true", String.IsValidIdentifier("TYPE_FLOAT"))  // Prints true
	gdtests.Print(t, "false", String.IsValidIdentifier("1st_method")) // Prints false
	gdtests.Print(t, "false", String.IsValidIdentifier("MyMethod#2")) // Prints false
}

func TestIsValidInt(t *testing.T) {
	gdtests.Print(t, "true", String.IsValidInt("7"))     // Prints true
	gdtests.Print(t, "false", String.IsValidInt("1.65")) // Prints false
	gdtests.Print(t, "false", String.IsValidInt("Hi"))   // Prints false
	gdtests.Print(t, "true", String.IsValidInt("+3"))    // Prints true
	gdtests.Print(t, "true", String.IsValidInt("-12"))   // Prints true
}

func TestJoinedWith(t *testing.T) {
	var fruits = []string{"Apple", "Orange", "Pear", "Kiwi"}

	gdtests.Print(t, "Apple, Orange, Pear, Kiwi", String.JoinedWith(", ", fruits...))     // Prints "Apple, Orange, Pear, Kiwi"
	gdtests.Print(t, "Apple---Orange---Pear---Kiwi", String.JoinedWith("---", fruits...)) // Prints "Apple---Orange---Pear---Kiwi"
}

func TestLast(t *testing.T) {
	gdtests.Print(t, "ld!", String.Last(3, "Hello World!"))       // Prints "ld!"
	gdtests.Print(t, "o World!", String.Last(-4, "Hello World!")) // Prints "o World!"
}

func TestSimilarity(t *testing.T) {
	gdtests.Print(t, "1.0", String.Similarity("ABC123", "ABC123")) // Prints 1.0
	gdtests.Print(t, "0.0", String.Similarity("ABC123", "XYZ456")) // Prints 0.0
	gdtests.Print(t, "0.8", String.Similarity("ABC123", "123ABC")) // Prints 0.8
	gdtests.Print(t, "0.4", String.Similarity("ABC123", "abc123")) // Prints 0.4
}

func TestSimplifypath(t *testing.T) {
	var simple_path = String.SimplifyPath("./path/to///../file")
	gdtests.Print(t, "path/file", simple_path) // Prints "path/file"
}

func TestSplits(t *testing.T) {
	var some_array = slices.Collect(String.Splits("One,Two,Three,Four", ","))

	gdtests.Print(t, "4", len(some_array)) // Prints 4
	gdtests.Print(t, "One", some_array[0]) // Prints "One"
	gdtests.Print(t, "Two", some_array[1]) // Prints "Two"
}

func TestExtractFloats(t *testing.T) {
	var a = String.ExtractFloats(",", "1,2,4.5")  // a is [1.0, 2.0, 4.5]
	var c = String.ExtractFloats("|", "1| ||4.5") // c is [1.0, 0.0, 0.0, 4.5]
	var b = String.ExtractFloats("|", "1| ||4.5") // b is [1.0, 4.5]

	gdtests.That(t, a, []Float.X{1.0, 2.0, 4.5})
	gdtests.That(t, b, []Float.X{1.0, 0.0, 0.0, 4.5})
	gdtests.That(t, c, []Float.X{1.0, 0.0, 0.0, 4.5})
}

func TestToFloat(t *testing.T) {
	var a = String.ToFloat("12.35")  // a is 12.35
	var b = String.ToFloat("1.2.3")  // b is 1.2
	var c = String.ToFloat("12xy3")  // c is 12.0
	var d = String.ToFloat("1e3")    // d is 1000.0
	var e = String.ToFloat("Hello!") // e is 0.0

	gdtests.Print(t, "12.35", a)
	gdtests.Print(t, "1.2", b)
	gdtests.Print(t, "12.0", c)
	gdtests.Print(t, "1000.0", d)
	gdtests.Print(t, "0.0", e)
}

func TestToInt(t *testing.T) {
	var a = String.ToInt("123")    // a is 123
	var b = String.ToInt("x1y2z3") // b is 123
	var c = String.ToInt("-1.2.3") // c is -1
	var d = String.ToInt("Hello!") // d is 0

	gdtests.That(t, a, 123)
	gdtests.That(t, b, 123)
	gdtests.That(t, c, -1)
	gdtests.That(t, d, 0)
}

func TestToSnakeCase(t *testing.T) {
	gdtests.That(t, String.ToSnakeCase("Node2D"), "node_2d")                               // Returns "node_2d"
	gdtests.That(t, String.ToSnakeCase("2nd place"), "2_nd_place")                         // Returns "2_nd_place"
	gdtests.That(t, String.ToSnakeCase("Texture3DAssetFolder"), "texture_3d_asset_folder") // Returns "texture_3d_asset_folder"
}

func TestSplit(t *testing.T) {
	var a = slices.Collect(String.Splits("One,Two,Three,Four", ",")) // a is ["One", "Two", "Three", "Four"]

	gdtests.Print(t, "4", len(a)) // Prints 4
	gdtests.Print(t, "One", a[0]) // Prints "One"
	gdtests.Print(t, "Two", a[1]) // Prints "Two"
}
