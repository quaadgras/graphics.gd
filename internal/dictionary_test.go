package gd_test

import (
	"testing"

	"graphics.gd/classdb"
	"graphics.gd/classdb/GDScript"
	"graphics.gd/internal/pointers"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/String"
)

type StringEater struct {
	Field string
}

type BuggyNode struct {
	Object.Extension[BuggyNode]
}

func (b *BuggyNode) EatString(value StringEater) {}

func init() {
	classdb.Register[BuggyNode]()
}

// TestDictionaryStringFields tests that string fields within a dictionary are not inappropriately freed.
func TestDictionaryStringFields(t *testing.T) {
	const source string = `extends BuggyNode

func test_dictionary_string_fields() -> String:
	var the_string = "Hello world!"
	print(the_string)
	eat_string({
		"Field": the_string,
	})
	return the_string
`
	var runner = new(BuggyNode)
	var script = GDScript.New().AsScript()
	script.SetSourceCode(source)
	script.Reload()
	Object.Instance(runner.AsObject()).SetScript(script)
	engine := Object.Call(runner, "test_dictionary_string_fields").(String.Readable)
	if engine.String() != "Hello world!" {
		t.Fatalf("Expected 'Hello world!', got '%s'", engine)
	}
	pointers.Cycle()
	pointers.Cycle()
	engine = Object.Call(runner, "test_dictionary_string_fields").(String.Readable)
	if engine.String() != "Hello world!" {
		t.Fatalf("Expected 'Hello world!', got '%s'", engine)
	}
}
