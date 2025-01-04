// builtins checks gdmaths and gdvalue packages for builtin class methods and reports
// any builtin methods that are missing from gd or any duplicates.
package main

import (
	"bufio"
	"fmt"
	"io/fs"
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

func extractPkg(pkg string, tags map[string]FileLine) error {
	file, err := os.Open(pkg)
	if err != nil {
		return xray.New(err)
	}
	defer file.Close()
	files, err := file.ReadDir(0)
	if err != nil {
		return xray.New(err)
	}
	for _, f := range files {
		if f.IsDir() || !strings.HasSuffix(f.Name(), ".go") {
			continue
		}
		name := pkg + "/" + f.Name()
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
	return nil
}

func extractFromPackages(prefix string, tags map[string]FileLine) error {
	dir, err := list(prefix)
	if err != nil {
		return xray.New(err)
	}
	for _, pkg := range dir {
		if !pkg.IsDir() {
			continue
		}
		if err := extractPkg(prefix+pkg.Name(), tags); err != nil {
			return xray.New(err)
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
		splits := strings.Split(tag, " ")
		for _, tag := range splits {
			if strings.HasPrefix(tag, "PackedArray") {
				suffix := strings.TrimPrefix(tag, "PackedArray")
				splits = append(splits, "PackedByteArray"+suffix)
				splits = append(splits, "PackedColorArray"+suffix)
				splits = append(splits, "PackedFloat32Array"+suffix)
				splits = append(splits, "PackedFloat64Array"+suffix)
				splits = append(splits, "PackedInt32Array"+suffix)
				splits = append(splits, "PackedInt64Array"+suffix)
				splits = append(splits, "PackedStringArray"+suffix)
				splits = append(splits, "PackedVector2Array"+suffix)
				splits = append(splits, "PackedVector3Array"+suffix)
				splits = append(splits, "PackedVector4Array"+suffix)
			}
		}
		for _, tag := range splits {
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
	var tags = make(map[string]FileLine)
	if err := extractPkg("./variant", tags); err != nil {
		return xray.New(err)
	}
	if err := extractPkg(".", tags); err != nil {
		return xray.New(err)
	}
	if err := extractPkg("./classdb", tags); err != nil {
		return xray.New(err)
	}
	if err := extractFromPackages("./variant/", tags); err != nil {
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
	for _, utility := range spec.UtilityFunctions {
		tag := utility.Name
		if tag == "error_string" {
			continue
		}
		if _, ok := tags[tag]; !ok {
			fmt.Println("//gd:" + tag)
		}
	}
	return nil
}
