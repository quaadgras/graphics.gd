package main

import (
	"fmt"
	"os"
	"strings"

	"graphics.gd/internal/gdjson"
	"graphics.gd/variant/String"
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
			var hasDefault bool
			for _, arg := range method.Arguments {
				if arg.DefaultValue != nil {
					hasDefault = true
				}
			}
			nodefault, ok := gdjson.NoDefaultName[class.Name+"."+method.Name]
			if hasDefault && !ok {
				fmt.Printf("%v.%v https://docs.godotengine.org/en/stable/classes/class_%[3]v.html#class-%[3]v-method-%[4]v\n",
					class.Name, method.Name, strings.ToLower(class.Name), strings.Replace(method.Name, "_", "-", -1))
			} else if nodefault == String.ToPascalCase(method.Name) {
				panic("no default for " + class.Name + "." + method.Name)
			}
		}
	}
	return nil
}
