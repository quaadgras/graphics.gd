// builtins checks gdmaths and gdvalue packages for builtin class methods and reports
// any builtin methods that are missing from gd or any duplicates.
package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"grow.graphics/gd/internal/gdjson"
	"runtime.link/api/xray"
)

func main() {
	if err := work(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func list(dir string) ([]fs.DirEntry, error) {
	folder, err := os.Open(dir)
	if err != nil {
		return nil, xray.New(err)
	}
	defer folder.Close()
	return folder.ReadDir(0)
}

type FileLine struct {
	File string
	Line int
}

func extractTags(prefix string, packages []fs.DirEntry, tags map[string]FileLine) error {
	for _, pkg := range packages {
		if !pkg.IsDir() {
			continue
		}
		file, err := os.Open(prefix + pkg.Name())
		if err != nil {
			return xray.New(err)
		}
		defer file.Close()
		files, err := file.ReadDir(0)
		if err != nil {
			return xray.New(err)
		}
		for _, f := range files {
			if f.IsDir() {
				continue
			}
			name := prefix + pkg.Name() + "/" + f.Name()
			f, err := os.Open(name)
			if err != nil {
				return xray.New(err)
			}
			if err := extractTagsFrom(name, bufio.NewReader(f), tags); err != nil {
				return xray.New(err)
			}
			if err := f.Close(); err != nil {
				return xray.New(err)
			}
		}
	}
	return nil
}

func extractTagsFrom(name string, lines *bufio.Reader, tags map[string]FileLine) error {
	for n := 1; true; n++ {
		line, err := lines.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return xray.New(err)
		}
		if !strings.HasPrefix(line, "func ") {
			continue
		}
		tagIdx := strings.Index(line, "//gd:")
		if tagIdx == -1 {
			continue
		}
		tag := strings.TrimSpace(line[tagIdx+5:])
		if tag == "" {
			continue
		}
		for _, tag := range strings.Split(tag, " ") {
			if _, ok := tags[tag]; ok {
				return fmt.Errorf("duplicate tag %q in %s:%d and %s:%d\n",
					tag, tags[tag].File, tags[tag].Line, name, n)
			}
			tags[tag] = FileLine{
				File: name,
				Line: n,
			}
		}
	}
	return nil
}

func work() error {
	spec, err := gdjson.LoadSpecification()
	if err != nil {
		return xray.New(err)
	}
	gdvalue, err := list("./gdvalue")
	if err != nil {
		return xray.New(err)
	}
	gdmaths, err := list("./gdmaths")
	if err != nil {
		return xray.New(err)
	}
	var tags = make(map[string]FileLine)
	if err := extractTags("./gdvalue/", gdvalue, tags); err != nil {
		return xray.New(err)
	}
	if err := extractTags("./gdmaths/", gdmaths, tags); err != nil {
		return xray.New(err)
	}
	for _, builtin := range spec.BuiltinClasses {
		if builtin.Name == "StringName" {
			builtin.Name = "String"
		}
		if builtin.Name == "RID" {
			continue // handled by gdclass/Resource
		}
		for _, method := range builtin.Methods {
			tag := builtin.Name + "." + method.Name
			if _, ok := tags[tag]; !ok {
				fmt.Println("//gd:" + tag)
			}
		}
	}
	return nil
}
