package main

import (
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"net/http"
	"os"

	"grow.graphics/gd/internal/gdjson"
)

// LoadSpecification, either from a local file or by downloading
// it from the Godot Github repository.
func LoadSpecification() (*gdjson.Specification, error) {
	file, err := os.Open("./extension_api.json")
	if os.IsNotExist(err) {
		req, err := http.NewRequest("GET", "https://raw.githubusercontent.com/godotengine/godot-headers/master/extension_api.json", nil)
		if err != nil {
			return nil, err
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, err
		}
		file, err = os.Create("./extension_api.json")
		if err != nil {
			return nil, err
		}
		_, err = io.Copy(file, resp.Body)
		if err != nil {
			return nil, err
		}
		file.Seek(0, 0)
		resp.Body.Close()
	}
	var spec gdjson.Specification
	if err := json.NewDecoder(file).Decode(&spec); err != nil {
		return nil, err
	}
	return &spec, nil
}

func main() {
	spec, err := LoadSpecification()
	if err != nil {
		panic(err)
	}
	var parent = make(map[string]string)
	var isRefCounted = make(map[string]bool)
	for _, class := range spec.Classes {
		parent[class.Name] = class.Inherits
	}
	for _, class := range spec.Classes {
		for name := class.Inherits; name != ""; {
			if name == "" {
				isRefCounted[class.Name] = false
				break
			}
			if name == "RefCounted" {
				isRefCounted[class.Name] = true
			}
			name = parent[name]
		}
	}
	ClassMethodOwnership := maps.Clone(gdjson.ClassMethodOwnership)
	for _, class := range spec.Classes {
		for _, method := range class.Methods {
			if method.IsVirtual {
				continue
			}
			var isUnsafe bool
			for _, arg := range method.Arguments {
				_, ok := ClassMethodOwnership[class.Name][method.Name][arg.Name]
				if ok {
					delete(ClassMethodOwnership[class.Name][method.Name], arg.Name)
					continue
				}
				if parent[arg.Type] != "" && !isRefCounted[arg.Type] {
					isUnsafe = true
					break
				}
			}
			_, ok := ClassMethodOwnership[class.Name][method.Name]["return value"]
			if ok {
				delete(ClassMethodOwnership[class.Name][method.Name], "return value")
			}
			if parent[method.ReturnValue.Type] != "" && !isRefCounted[method.ReturnValue.Type] && !ok {
				isUnsafe = true
			}
			if isUnsafe {
				fmt.Print(class.Name, "::", method.Name, "(")
				for i, arg := range method.Arguments {
					if i > 0 {
						fmt.Print(",")
					}
					fmt.Print(arg.Name)
					if parent[arg.Type] != "" && !isRefCounted[arg.Type] {
						fmt.Print("?")
					}
				}
				fmt.Print(")")
				fmt.Print(method.ReturnValue.Type)
				if parent[method.ReturnValue.Type] != "" && !isRefCounted[method.ReturnValue.Type] {
					fmt.Print("?")
				}
				fmt.Println()
			}
			if len(ClassMethodOwnership[class.Name][method.Name]) == 0 {
				delete(ClassMethodOwnership[class.Name], method.Name)
			}
		}

	}
	var unused bool
	for class, methods := range ClassMethodOwnership {
		for method, params := range methods {
			for param := range params {
				fmt.Println("UNUSED! ", class, "::", method, "(", param, ")")
				unused = true
			}
		}
	}
	if unused {
		os.Exit(1)
	}
}
