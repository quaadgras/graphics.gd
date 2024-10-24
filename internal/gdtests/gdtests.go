package gdtests

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func That(t *testing.T, value, expectation any) {
	t.Helper()
	if !reflect.DeepEqual(value, expectation) {
		t.Fatalf("expected %v, got %v", expectation, value)
	}
}

func Sprint(value any) string {
	if value == nil {
		return "nil"
	}
	rvalue := reflect.ValueOf(value)
	if collection, ok := value.(interface{ Length() int }); ok {
		printed := "["
		for i := 0; i < collection.Length(); i++ {
			if i > 0 {
				printed += ", "
			}
			printed += Sprint(rvalue.MethodByName("Lookup").Call([]reflect.Value{reflect.ValueOf(i)})[0].Interface())
		}
		printed += "]"
		return printed
	}
	for rvalue.Kind() == reflect.Ptr && !rvalue.IsNil() {
		rvalue = rvalue.Elem()
	}
	switch rvalue.Kind() {
	case reflect.Slice, reflect.Array:
		printed := "["
		for i := range rvalue.Len() {
			if i > 0 {
				printed += ", "
			}
			printed += Sprint(rvalue.Index(i).Interface())
		}
		printed += "]"
		return printed
	case reflect.String:
		return fmt.Sprintf("%q", rvalue.Interface())
	case reflect.Float32, reflect.Float64:
		printed := fmt.Sprint(value)
		if !strings.Contains(printed, ".") {
			printed += ".0"
		}
		return printed
	default:
		return fmt.Sprintf("%v", rvalue.Interface())
	}
}

func Print(t *testing.T, output string, value any) {
	t.Helper()
	printed := Sprint(value)
	if printed != output {
		t.Fatalf("expected %q, got %q", output, printed)
	}
}
