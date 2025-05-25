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
		if len(class.Constants) == 0 {
			continue
		}
		if class.Name == "Object" {
			continue
		}
		var first = true
	constants:
		for _, constant := range class.Constants {
			for prefix := range gdjson.Consitution[class.Name] {
				if strings.HasPrefix(constant.Name, prefix) {
					continue constants
				}
			}
			if first {
				fmt.Printf("Class: %s https://docs.godotengine.org/en/stable/classes/class_%s.html\n", class.Name, strings.ToLower(class.Name))
				first = false
			}
			fmt.Println(constant.Name, "=", constant.Value)
		}
	}
	return nil
}
