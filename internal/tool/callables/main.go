package main

import (
	"fmt"
	"os"

	"grow.graphics/gd/internal/gdjson"
	"runtime.link/api/xray"
)

func main() {
	if err := work(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func check(class gdjson.Class, method gdjson.Method, name string, gdtype string) {
	if gdtype == "Callable" || gdtype == "Signal" {
		key := (class.Name + "." + method.Name + "." + name)
		if _, ok := gdjson.Callables[key]; !ok {
			fmt.Printf("%q\n", key)
		}
	}
}

func work() error {
	spec, err := gdjson.LoadSpecification()
	if err != nil {
		return xray.New(err)
	}
	for _, class := range spec.Classes {
		for _, method := range class.Methods {
			if method.IsVirtual {
				continue
			}
			for _, arg := range method.Arguments {
				check(class, method, arg.Name, arg.Type)
			}
			check(class, method, "", method.ReturnValue.Type)
		}
	}
	return nil
}
