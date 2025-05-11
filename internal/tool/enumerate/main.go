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
	for _, enum := range spec.GlobalEnums {
		if gdjson.Renumeration[enum.Name] != "" {
			continue
		}
		fmt.Println(enum.Name)
	}
	return nil
}
