package main

import (
	"bytes"
	"fmt"
	"iter"
	"os"
	"strings"

	"graphics.gd/internal/gdjson"
	"runtime.link/api/xray"
)

func main() {
	if err := transdoc(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

// codeblocks extracts [codeblocks] tags from the given bbcode documentation.
func codeblocks(docs string) iter.Seq[string] {
	return func(yield func(string) bool) {
		for {
			_, codeblock, hasCodeBlock := strings.Cut(docs, "[codeblock]")
			_, codeblocks, hasCodeBlocks := strings.Cut(docs, "[codeblocks]")
			if !hasCodeBlock && !hasCodeBlocks {
				return
			}
			switch {
			case hasCodeBlock:
				codeblock, docs, hasCodeBlock = strings.Cut(codeblock, "[/codeblock]")
				if !hasCodeBlock {
					return
				}
				if !yield(codeblock) {
					return
				}
			case hasCodeBlocks:
				codeblocks, docs, hasCodeBlocks = strings.Cut(codeblocks, "[/codeblocks]")
				if !hasCodeBlocks {
					return
				}
				if !yield(codeblocks) {
					return
				}
			}
		}
	}
}

func transdoc() error {
	spec, err := gdjson.LoadSpecification()
	if err != nil {
		return xray.New(err)
	}
	for _, class := range spec.Classes {
		if class.Name == "Object" {
			continue
		}
		docs := class.Description
		var blocks []string
		for codeblock := range codeblocks(docs) {
			blocks = append(blocks, codeblock)
		}
		if len(blocks) > 0 {
			if existing, err := os.ReadFile("./internal/gddocs/" + class.Name + ".go"); err == nil {
				if bytes.HasPrefix(existing, []byte("/*"+blocks[0]+"*/")) {
					continue
				}
			}
			fmt.Println(class.Name + ".go")
			fmt.Println("/*" + blocks[0] + "*/")
			fmt.Println("\npackage main")
			os.Exit(0)
		}
		if len(blocks) > 1 {
			for i, block := range blocks[1:] {
				if existing, err := os.ReadFile(fmt.Sprintf("./internal/gddocs/%s%d.go", class.Name, i+2)); err == nil {
					if bytes.HasPrefix(existing, []byte("/*"+block+"*/")) {
						continue
					}
				}
				fmt.Printf("%s%d.go", class.Name, i+2)
				fmt.Println("/*" + block + "*/")
				fmt.Println("\npackage main")
				os.Exit(0)
			}
		}
		for _, method := range class.Methods {
			docs := method.Description
			var blocks []string
			for codeblock := range codeblocks(docs) {
				blocks = append(blocks, codeblock)
			}
			if len(blocks) > 0 {
				name := fmt.Sprintf("%s_%s.go", class.Name, gdjson.ConvertName(method.Name))
				if existing, err := os.ReadFile("./internal/gddocs/" + name); err == nil {
					if bytes.HasPrefix(existing, []byte("/*"+blocks[0]+"*/")) {
						continue
					}
				}
				fmt.Println(name)
				fmt.Println("/*" + blocks[0] + "*/")
				fmt.Println("\npackage main")
				os.Exit(0)
			}
			if len(blocks) > 1 {
				for i, block := range blocks[1:] {
					name := fmt.Sprintf("%s_%s%d.go", class.Name, gdjson.ConvertName(method.Name), i+2)
					if existing, err := os.ReadFile("./internal/gddocs/" + name); err == nil {
						if bytes.HasPrefix(existing, []byte("/*"+block+"*/")) {
							continue
						}
					}
					fmt.Println(name)
					fmt.Println("/*" + block + "*/")
					fmt.Println("\npackage main")
					os.Exit(0)
				}
			}
		}
	}

	return nil
}
