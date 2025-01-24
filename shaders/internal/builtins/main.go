// builtins checks gdmaths and gdvalue packages for builtin class methods and reports
// any builtin methods that are missing from gd or any duplicates.
package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"maps"
	"os"
	"slices"
	"sort"
	"strings"

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
		tagIdx := strings.Index(line, "//glsl:")
		if tagIdx == -1 {
			continue
		}
		tag := strings.TrimSpace(line[tagIdx+7:])
		if tag == "" {
			continue
		}
		splits := strings.Split(tag, " ")
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
	var tags = make(map[string]FileLine)
	var checked = make(map[string]bool)
	if err := extractFromPackages("./", tags); err != nil {
		return xray.New(err)
	}
	keys := slices.Collect(maps.Keys(builtins))
	sort.Strings(keys)
	for _, name := range keys {
		args := builtins[name]
		for _, sets := range args {
			var done = make(map[string]bool)
			replacements := [][2]string{}
			result, func_name, _ := strings.Cut(name, " ")
			if expansions, ok := expansions[result]; ok {
				for _, expansion := range expansions {
					if !done[expansion] {
						replacements = append(replacements, [2]string{result, expansion})
						done[expansion] = true
					}
				}
			}
			for _, arg := range sets {
				if expansions, ok := expansions[arg]; ok {
					for _, expansion := range expansions {
						if !done[expansion] {
							replacements = append(replacements, [2]string{arg, expansion})
							done[expansion] = true
						}
					}
				}
			}
			if len(replacements) == 0 {
				replacements = append(replacements, [2]string{result, result})
			}
			for _, replacement := range replacements {
				var tag strings.Builder
				tag.WriteString(func_name)
				if len(sets) > 0 {
					tag.WriteString("(")
					for i, arg := range sets {
						tag.WriteString(strings.ReplaceAll(arg, replacement[0], replacement[1]))
						if i < len(sets)-1 {
							tag.WriteString(",")
						}
					}
					tag.WriteString(")")
					tag.WriteString(strings.ReplaceAll(result, replacement[0], replacement[1]))
				}
				if _, ok := tags[tag.String()]; !ok && !checked[tag.String()] {
					fmt.Printf("//glsl:%s\n", tag.String())
				} else {
					checked[tag.String()] = true
					delete(tags, tag.String())
				}
			}
		}
	}
	if len(tags) > 0 {
		fmt.Println("unused tags:")
		for tag, fl := range tags {
			fmt.Printf("%s:%d: %s\n", fl.File, fl.Line, tag)
		}
	}
	return nil
}
