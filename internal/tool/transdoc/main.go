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
		docs := class.Description
		var blocks []string
		for codeblock := range codeblocks(docs) {
			blocks = append(blocks, codeblock)
		}
		if len(blocks) > 0 {
			if existing, err := os.ReadFile("./internal/gddocs/" + class.Name + ".go"); err == nil {
				if !bytes.HasPrefix(existing, []byte("/*"+blocks[0]+"*/")) {
					fmt.Println(class.Name + ".go")
					fmt.Print("/*" + blocks[0] + "*/")
					os.Exit(0)
				}
			} else {
				fmt.Println(class.Name + ".go")
				fmt.Print("/*" + blocks[0] + "*/")
				os.Exit(0)
			}
		}
		if len(blocks) > 1 {
			for i, block := range blocks[1:] {
				if existing, err := os.ReadFile(fmt.Sprintf("./internal/gddocs/%s%d.go", class.Name, i+2)); err == nil {
					if bytes.HasPrefix(existing, []byte("/*"+block+"*/")) {
						continue
					}
				}
				fmt.Printf("%s%d.go", class.Name, i+2)
				fmt.Print(block)
				os.Exit(0)
			}
		}
	}
	return nil
}
