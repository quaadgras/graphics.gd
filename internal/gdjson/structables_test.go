package gdjson_test

import (
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"

	"graphics.gd/internal/gdjson"
)

// TestStructablesTypedArrayElement checks that every [gdjson.Structables] entry
// for a 'typedarray::' method argument or return value names the element type
// rather than a slice of it. The generator prepends the "[]" itself, so a slice
// here produces a doubly nested Go type ([][]Connection instead of []Connection)
// that no conversion can ever satisfy, which is what broke
// GraphEdit.get_connection_list_from_node in issue #328.
func TestStructablesTypedArrayElement(t *testing.T) {
	file, err := os.Open("../extension_api.json")
	if err != nil {
		t.Skipf("no extension_api.json available: %v", err)
	}
	defer file.Close()
	var spec gdjson.Specification
	if err := json.NewDecoder(file).Decode(&spec); err != nil {
		t.Fatal(err)
	}
	check := func(key, gdType string) {
		element, ok := strings.CutPrefix(gdType, "typedarray::")
		if !ok || element == "Array" { // an Array element is itself a slice.
			return
		}
		rtype, ok := gdjson.Structables[key]
		if !ok {
			return
		}
		if rtype.Kind() == reflect.Slice {
			t.Errorf("Structables[%q] is %s, expected the element type %s (the generator adds the slice for typedarray:: types)",
				key, rtype, rtype.Elem())
		}
	}
	for _, class := range spec.Classes {
		for _, method := range class.Methods {
			for _, arg := range method.Arguments {
				check(class.Name+"."+method.Name+"."+arg.Name, arg.Type)
			}
			check(class.Name+"."+method.Name+".", method.ReturnValue.Type)
		}
	}
}
