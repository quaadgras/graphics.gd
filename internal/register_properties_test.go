//go:build !generate

package gd_test

import (
	"testing"

	"graphics.gd/classdb"
	"graphics.gd/classdb/GDScript"
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Enum"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector2"
)

type TestingExportedProperties struct {
	Node.Extension[TestingExportedProperties]

	TestInt    int
	TestString string
	TestFloat  Float.X
	TestVector Vector2.XY
	TestEnum   CustomEnum

	setCalls   []kvPair
	onSetCalls []kvPair
}

func (ep *TestingExportedProperties) Set(k string, v any) bool {
	// Keep track of all Set() invocations
	ep.setCalls = append(ep.setCalls, kvPair{k, v})
	return true
}

func (ep *TestingExportedProperties) OnSet(k string, v any) {
	// Keep track of all OnSet() invocations
	ep.onSetCalls = append(ep.onSetCalls, kvPair{k, v})
}

type CustomEnum Enum.Int[struct {
	One,
	Two,
	Three CustomEnum
}]

var CustomEnums = Enum.Values[CustomEnum]()

type kvPair struct {
	k string
	v any
}

func init() {
	classdb.Register[TestingExportedProperties]()
}

const exportedPropertiesScript = `extends TestingExportedProperties

func set_values() -> void:
	self.test_int = 5
	self.test_string = "test_string"
	self.test_float = 2.2
	self.test_vector = Vector2(100, 200)
	self.test_enum = CustomEnum.THREE
`

func TestRegisterExportedProperties(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		obj := &TestingExportedProperties{}
		var script = GDScript.New().AsScript()
		script.SetSourceCode(exportedPropertiesScript)
		script.Reload()
		Object.Instance(obj.AsObject()).SetScript(script)
		wanted := &TestingExportedProperties{
			TestInt:    5,
			TestString: "test_string",
			TestFloat:  2.2,
			TestVector: Vector2.New(100, 200),
			TestEnum:   CustomEnums.Three,
		}
		wantedOnSet := []kvPair{
			kvPair{"test_int", wanted.TestInt},
			kvPair{"test_string", wanted.TestString},
			kvPair{"test_float", wanted.TestFloat},
			kvPair{"test_vector", wanted.TestVector},
			kvPair{"test_enum", wanted.TestEnum.Int()},
		}

		Object.Call(obj, "set_values")

		if obj.TestInt != wanted.TestInt {
			t.Errorf("obj.TestInt = %v, wanted %v", obj.TestInt, wanted.TestInt)
		}
		if obj.TestString != wanted.TestString {
			t.Errorf("obj.TestString = %v, wanted %v", obj.TestString, wanted.TestString)
		}
		if obj.TestFloat != wanted.TestFloat {
			t.Errorf("obj.TestFloat = %v, wanted %v", obj.TestFloat, wanted.TestFloat)
		}
		if obj.TestVector != wanted.TestVector {
			t.Errorf("obj.TestVector = %v, wanted %v", obj.TestVector, wanted.TestVector)
		}
		if len(obj.setCalls) != 0 {
			t.Errorf("obj.Set() called %v times, wanted 0", len(obj.setCalls))
		}
		if len(obj.onSetCalls) != len(wantedOnSet) {
			t.Errorf("obj.OnSet() called %v times, wanted %v", len(obj.onSetCalls), len(wantedOnSet))
		}
		for i, _ := range wantedOnSet {
			if obj.onSetCalls[i].k != wantedOnSet[i].k {
				t.Errorf("obj.OnSet() called with key = %v, wanted %v", obj.onSetCalls[i].k, wantedOnSet[i].k)
			}

			// Floats don't retain perfect precision, so instead of requiring equality, check that they're "close enough."
			if floatVal, ok := obj.onSetCalls[i].v.(float64); ok {
				if !Float.IsApproximatelyEqual(floatVal, float64(wantedOnSet[i].v.(float32))) {
					t.Errorf("obj.OnSet() called with value = %v, wanted %v", obj.onSetCalls[i].v, wantedOnSet[i].v)
				}
			} else {
				if obj.onSetCalls[i].v != wantedOnSet[i].v {
					t.Errorf("obj.OnSet() called with value = %v, wanted %v", obj.onSetCalls[i].v, wantedOnSet[i].v)
				}
			}
		}
	})
}

const unrecognizedPropertySetter = `extends TestingExportedProperties

func set_values() -> void:
	self.unrecognized_int = 10
	self.unrecognized_string = "Hello world"
	self.unrecognized_float = 3.14
	self.unrecognized_vector = Vector2(5.1, 5.3)
`

func TestRegisterUnrecognizedPropertySetter(t *testing.T) {
	runOnMain(t, func(t testing.TB) {
		obj := &TestingExportedProperties{}
		var script = GDScript.New().AsScript()
		script.SetSourceCode(unrecognizedPropertySetter)
		script.Reload()
		Object.Instance(obj.AsObject()).SetScript(script)
		wanted := []kvPair{
			kvPair{"unrecognized_int", 10},
			kvPair{"unrecognized_string", "Hello world"},
			kvPair{"unrecognized_float", 3.14},
			kvPair{"unrecognized_vector", Vector2.New(5.1, 5.3)},
		}

		Object.Call(obj, "set_values")

		if len(obj.setCalls) != len(wanted) {
			t.Errorf("obj.Set() called %v times, wanted %v", len(obj.setCalls), len(wanted))
		}
		for i, _ := range wanted {
			if obj.setCalls[i].k != wanted[i].k {
				t.Errorf("obj.Set() called with key = %v, wanted %v", obj.setCalls[i].k, wanted[i].k)
			}
			if obj.setCalls[i].v != wanted[i].v {
				t.Errorf("obj.Set() called with value = %v, wanted %v", obj.setCalls[i].v, wanted[i].v)
			}
		}
		if len(obj.onSetCalls) != 0 {
			t.Errorf("obj.OnSet() called %v times, wanted 0", len(obj.setCalls))
		}
	})
}
