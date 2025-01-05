package gd_test

import (
	"errors"
	"testing"

	"graphics.gd/classdb"
	"graphics.gd/classdb/GDScript"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/variant/Object"

	_ "embed"
)

//go:embed convert_types_test.gd
var convert_types_test string

var doneConversionsTest = make(chan error, 1)

type Converter struct {
	classdb.Extension[Converter, Node.Instance]
}

func (c Converter) Ready()                { doneConversionsTest <- errors.New("script didn't call Done") }
func (c Converter) AsNode() Node.Instance { return c.Super() }

func (c Converter) Done() { doneConversionsTest <- nil }
func (c Converter) Fail() { doneConversionsTest <- errors.New("test failed") }

func (c Converter) HelloWorld() string   { return "Hello, World!" }
func (c Converter) Echo(s string) string { return s }
func (c Converter) Int() int             { return 22 }
func (c Converter) Float() float64       { return 2.2 }
func (c Converter) Bool() bool           { return true }
func (c Converter) Int8() int8           { return 8 }
func (c Converter) Int16() int16         { return 16 }
func (c Converter) Int32() int32         { return 32 }
func (c Converter) Int64() int64         { return 64 }
func (c Converter) Uint() uint           { return 22 }
func (c Converter) Uint8() uint8         { return 8 }
func (c Converter) Uint16() uint16       { return 16 }
func (c Converter) Uint32() uint32       { return 32 }
func (c Converter) Uint64() uint64       { return 64 }

func TestConversions(t *testing.T) {
	converter := &Converter{}
	var script = GDScript.New().AsScript()
	script.SetSourceCode(convert_types_test)
	script.Reload()
	Object.Instance{converter.AsObject()}.SetScript(script)
	SceneTree.Add(converter)
	if err := <-doneConversionsTest; err != nil {
		t.Fatal(err)
	}
}
