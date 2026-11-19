// Command deprecated writes deprecated.json, the classdb classes, methods and
// properties that Godot's class reference marks as deprecated, for the
// generator to leave out of builds with the gd_omit_deprecated tag (whole
// packages for classes, each package's deprecated.go for the rest). extension_api.json does not carry deprecation, so it
// is read from the doc XML of a Godot source tree, which should be checked out
// at the version that extension_api.json was dumped from:
//
//	go run ./internal/tool/deprecated ~/git/godot
package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type classXML struct {
	Name       string  `xml:"name,attr"`
	Deprecated *string `xml:"deprecated,attr"`
	Methods []struct {
		Name       string  `xml:"name,attr"`
		Qualifiers string  `xml:"qualifiers,attr"`
		Deprecated *string `xml:"deprecated,attr"`
	} `xml:"methods>method"`
	Members []struct {
		Name       string  `xml:"name,attr"`
		Deprecated *string `xml:"deprecated,attr"`
	} `xml:"members>member"`
}

// Class lists what is deprecated in one class, by engine name, with the
// deprecation message (in the class reference's BBCode). Deprecated is set
// when the whole class is.
type Class struct {
	Deprecated *string           `json:"deprecated,omitempty"`
	Methods    map[string]string `json:"methods,omitempty"`
	Properties map[string]string `json:"properties,omitempty"`
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("usage: deprecated <godot source directory>")
	}
	root := os.Args[1]
	var classes = make(map[string]*Class)
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			switch d.Name() {
			case ".git", "thirdparty", "bin", "tests":
				return filepath.SkipDir
			}
			return nil
		}
		dir := filepath.Base(filepath.Dir(path))
		if !strings.HasSuffix(path, ".xml") || (dir != "doc_classes" && !(dir == "classes" && filepath.Base(filepath.Dir(filepath.Dir(path))) == "doc")) {
			return nil
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		var class classXML
		if err := xml.Unmarshal(data, &class); err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}
		if class.Name == "" || strings.HasPrefix(class.Name, "@") {
			return nil
		}
		entry := &Class{Deprecated: class.Deprecated}
		for _, method := range class.Methods {
			// Virtual methods are part of the Interface that extensions
			// implement, so they stay where they are.
			if method.Deprecated == nil || strings.Contains(method.Qualifiers, "virtual") {
				continue
			}
			if entry.Methods == nil {
				entry.Methods = make(map[string]string)
			}
			entry.Methods[method.Name] = *method.Deprecated
		}
		for _, member := range class.Members {
			if member.Deprecated == nil {
				continue
			}
			if entry.Properties == nil {
				entry.Properties = make(map[string]string)
			}
			entry.Properties[member.Name] = *member.Deprecated
		}
		if entry.Deprecated != nil || entry.Methods != nil || entry.Properties != nil {
			classes[class.Name] = entry
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.MarshalIndent(classes, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("deprecated.json", append(data, '\n'), 0644); err != nil {
		log.Fatal(err)
	}
}
