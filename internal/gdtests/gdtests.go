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
	if collection, ok := value.(interface{ Size() int }); ok {
		var printed strings.Builder
		printed.WriteString("[")
		for i := 0; i < collection.Size(); i++ {
			if i > 0 {
				printed.WriteString(", ")
			}
			printed.WriteString(Sprint(rvalue.MethodByName("Index").Call([]reflect.Value{reflect.ValueOf(i)})[0].Interface()))
		}
		printed.WriteString("]")
		return printed.String()
	}
	if collection, ok := value.(interface{ Len() int }); ok {
		var printed strings.Builder
		printed.WriteString("[")
		for i := 0; i < collection.Len(); i++ {
			if i > 0 {
				printed.WriteString(", ")
			}
			printed.WriteString(Sprint(rvalue.MethodByName("Index").Call([]reflect.Value{reflect.ValueOf(i)})[0].Interface()))
		}
		printed.WriteString("]")
		return printed.String()
	}
	for rvalue.Kind() == reflect.Pointer && !rvalue.IsNil() {
		rvalue = rvalue.Elem()
	}
	switch rvalue.Kind() {
	case reflect.Slice, reflect.Array:
		var printed strings.Builder
		printed.WriteString("[")
		for i := range rvalue.Len() {
			if i > 0 {
				printed.WriteString(", ")
			}
			printed.WriteString(Sprint(rvalue.Index(i).Interface()))
		}
		printed.WriteString("]")
		return printed.String()
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
	if s, ok := value.(string); ok {
		printed = s
	}
	if printed != output {
		t.Fatalf("expected %q, got %q", output, printed)
	}
}
