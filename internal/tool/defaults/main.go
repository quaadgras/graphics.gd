package main

import (
	"fmt"
	"os"

	"graphics.gd/internal/gdjson"
	"runtime.link/api/xray"
)

func main() {
	if err := work(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func work() error {
	spec, err := gdjson.LoadSpecification()
	if err != nil {
		return xray.New(err)
	}
	var unique = make(map[string]bool)
	for _, class := range spec.Classes {
		for _, method := range class.Methods {
			if !method.IsStatic && !class.IsSingleton {
				continue
			}
			var hasDefault bool
			for _, arg := range method.Arguments {
				if arg.DefaultValue != nil {
					if !gdjson.IsTheDefaultValueZero(*arg.DefaultValue) {
						hasDefault = true
						unique[*arg.DefaultValue] = true
					}
				}
			}
			if hasDefault {
				fmt.Printf("%v.%vOptions\n",
					class.Name, gdjson.ConvertName(method.Name))
			}

		}
	}
	//fmt.Printf("%#v\n", sorted)
	/*for name := range unique {
	fmt.Println(name)
	}*/
	return nil
}
