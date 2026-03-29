package main

import (
	"fmt"
	"os"
	"strings"

	"graphics.gd/internal/gdjson"
	"runtime.link/api/xray"
)

func main() {
	if err := work(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func isPointer(typ string) bool {
	return strings.Contains(typ, "*")
}

// methodKey returns "Class.method".
func methodKey(className, methodName string) string {
	return className + "." + methodName
}

func work() error {
	spec, err := gdjson.LoadSpecification()
	if err != nil {
		return xray.New(err)
	}
	// Build a lookup of all methods by "Class.method" -> arguments.
	methods := make(map[string][]gdjson.Argument)
	collect := func(className string, ms []gdjson.Method) {
		for _, m := range ms {
			methods[methodKey(className, m.Name)] = m.Arguments
		}
	}
	for _, class := range spec.BuiltinClasses {
		collect(class.Name, class.Methods)
	}
	for _, class := range spec.Classes {
		collect(class.Name, class.Methods)
	}

	// Report unmapped pointer parameters and return values.
	check := func(className string, ms []gdjson.Method) {
		for _, method := range ms {
			retType := method.ReturnValue.Type
			if retType == "" {
				retType = method.ReturnType
			}
			if isPointer(retType) {
				key := className + "." + method.Name + "."
				if gdjson.Addressables[key] == "" {
					fmt.Printf("%s returns %s\n", key, retType)
				}
			}
			for _, arg := range method.Arguments {
				if isPointer(arg.Type) {
					key := className + "." + method.Name + "." + arg.Name
					if gdjson.Addressables[key] == "" {
						fmt.Printf("%s: %s\n", key, arg.Type)
					}
				}
			}
		}
	}
	for _, class := range spec.BuiltinClasses {
		check(class.Name, class.Methods)
	}
	for _, class := range spec.Classes {
		check(class.Name, class.Methods)
	}
	for _, fn := range spec.UtilityFunctions {
		if isPointer(fn.ReturnType) {
			key := "." + fn.Name + "."
			if gdjson.Addressables[key] == "" {
				fmt.Printf("%s returns %s\n", key, fn.ReturnType)
			}
		}
		for _, arg := range fn.Arguments {
			if isPointer(arg.Type) {
				key := "." + fn.Name + "." + arg.Name
				if gdjson.Addressables[key] == "" {
					fmt.Printf("%s: %s\n", key, arg.Type)
				}
			}
		}
	}

	// Validate sliceables: check that both the pointer param and count param exist.
	var errs []string
	for key, sliceable := range gdjson.Sliceables {
		parts := strings.SplitN(key, ".", 3)
		if len(parts) != 3 {
			errs = append(errs, fmt.Sprintf("sliceable %q: expected Class.method.param format", key))
			continue
		}
		className, methodName, paramName := parts[0], parts[1], parts[2]
		args, ok := methods[methodKey(className, methodName)]
		if !ok {
			errs = append(errs, fmt.Sprintf("sliceable %q: method %s.%s not found", key, className, methodName))
			continue
		}
		hasParam := false
		hasCount := false
		for _, arg := range args {
			if arg.Name == paramName {
				hasParam = true
			}
			if arg.Name == sliceable.Count {
				hasCount = true
			}
		}
		if !hasParam {
			errs = append(errs, fmt.Sprintf("sliceable %q: param %q not found in %s.%s", key, paramName, className, methodName))
		}
		if !hasCount {
			errs = append(errs, fmt.Sprintf("sliceable %q: count %q not found in %s.%s", key, sliceable.Count, className, methodName))
		}
	}
	if len(errs) > 0 {
		for _, e := range errs {
			fmt.Fprintf(os.Stderr, "error: %s\n", e)
		}
		return fmt.Errorf("%d sliceable validation error(s)", len(errs))
	}
	return nil
}
