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

func work() error {
	spec, err := gdjson.LoadSpecification()
	if err != nil {
		return xray.New(err)
	}
	for _, class := range spec.Classes {
		for _, method := range class.Methods {
			for _, arg := range method.Arguments {
				if strings.Contains(arg.Name, os.Args[1]) {
					fmt.Printf("Class: %s Method: %s Argument: %s\n", class.Name, method.Name, arg.Name)
				}
			}
			if strings.Contains(method.Name, os.Args[1]) {
				fmt.Printf("Class: %s Method: %s\n", class.Name, method.Name)
			}

		}
	}
	return nil
}
