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
	// cases that we don't want to see in the Engine API.
	var cases = [][2]string{
		{"angle", "float"},
		{"rotation", "float"},
		{"angles", "Vector3"},
		{"rotation", "Vector3"},
	}
	for _, class := range spec.Classes {
		for _, method := range class.Methods {
			switch method.Name {
			case "get_rotation_smoothing_speed":
				continue
			}
			for _, arg := range method.Arguments {
			cases1:
				for _, c := range cases {
					if strings.Contains(arg.Name, c[0]) && arg.Type == c[1] {
						for _, d := range gdjson.Distinctions[class.Name] {
							matchFunc, matchArg, hasFuncMatch := strings.Cut(d[0], " ")
							if !hasFuncMatch {
								matchArg = matchFunc
							}
							if matchArg == "" || (strings.Contains(arg.Name, matchArg) && arg.Type == d[1] && (!hasFuncMatch || strings.Contains(method.Name, matchFunc))) {
								continue cases1
							}
						}
						fmt.Printf("Class: %s https://docs.godotengine.org/en/stable/classes/class_%s.html Method: %s Arg %s\n", class.Name, strings.ToLower(class.Name), method.Name, arg.Name)
					}
				}
			}
		cases2:
			for _, c := range cases {
				for _, d := range gdjson.Distinctions[class.Name] {
					matchFunc, matchArg, hasFuncMatch := strings.Cut(d[0], " ")
					if !hasFuncMatch {
						matchArg = matchFunc
					}
					if matchArg == "" && method.ReturnValue.Type == d[1] && (!hasFuncMatch || strings.Contains(method.Name, matchFunc)) {
						continue cases2
					}
				}
				if strings.Contains(method.Name, c[0]) && method.ReturnValue.Type == c[1] {
					fmt.Printf("Class: %s https://docs.godotengine.org/en/stable/classes/class_%s.html Method: %s Return Value\n", class.Name, strings.ToLower(class.Name), method.Name)
				}
			}
		}
	}
	return nil
}
