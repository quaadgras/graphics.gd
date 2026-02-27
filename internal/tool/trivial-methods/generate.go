//go:build ignore

// generate.go reads the clang trivial_methods output (results.jsonl)
// and the Godot extension API specification (extension_api.json), then
// produces trivial_methods.json: a map of class → list of method names
// that are both trivial (zero function calls in the C++ body) and
// actually bound in the extension API.
//
// Usage: go run generate.go [-results results.jsonl] [-api extension_api.json] [-o trivial_methods.json]
package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"sort"
)

type clangResult struct {
	Class   string `json:"class"`
	Method  string `json:"method"`
	File    string `json:"file"`
	Line    int    `json:"line"`
	Trivial bool   `json:"trivial"`
}

type apiSpec struct {
	Classes []struct {
		Name    string `json:"name"`
		Methods []struct {
			Name string `json:"name"`
		} `json:"methods"`
	} `json:"classes"`
}

func main() {
	resultsPath := flag.String("results", "results.jsonl", "path to clang trivial_methods output")
	apiPath := flag.String("api", "extension_api.json", "path to Godot extension_api.json")
	outPath := flag.String("o", "trivial_methods.json", "output path")
	flag.Parse()

	// Load clang results.
	trivialSet := make(map[string]bool) // "Class::Method" → true
	f, err := os.Open(*resultsPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening results: %v\n", err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 1<<20), 1<<20)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		var r clangResult
		if err := json.Unmarshal([]byte(line), &r); err != nil {
			fmt.Fprintf(os.Stderr, "warning: skipping malformed line: %v\n", err)
			continue
		}
		key := r.Class + "::" + r.Method
		if !r.Trivial {
			trivialSet[key] = false
		} else if _, seen := trivialSet[key]; !seen {
			trivialSet[key] = true
		}
	}
	f.Close()
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading results: %v\n", err)
		os.Exit(1)
	}

	// Load extension API.
	apiFile, err := os.Open(*apiPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error opening api: %v\n", err)
		os.Exit(1)
	}
	var api apiSpec
	if err := json.NewDecoder(apiFile).Decode(&api); err != nil {
		fmt.Fprintf(os.Stderr, "error parsing api: %v\n", err)
		os.Exit(1)
	}
	apiFile.Close()

	// Cross-reference: only keep methods that are both trivial and bound.
	output := make(map[string][]string) // class → sorted method names
	total := 0
	for _, cls := range api.Classes {
		for _, m := range cls.Methods {
			key := cls.Name + "::" + m.Name
			if trivialSet[key] {
				output[cls.Name] = append(output[cls.Name], m.Name)
				total++
			}
		}
	}

	// Sort methods within each class for stable output.
	for k := range output {
		sort.Strings(output[k])
	}

	// Write output.
	out, err := os.Create(*outPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating output: %v\n", err)
		os.Exit(1)
	}
	enc := json.NewEncoder(out)
	enc.SetIndent("", "  ")
	if err := enc.Encode(output); err != nil {
		fmt.Fprintf(os.Stderr, "error writing output: %v\n", err)
		os.Exit(1)
	}
	out.Close()

	fmt.Fprintf(os.Stderr, "wrote %d trivial methods across %d classes to %s\n", total, len(output), *outPath)
}
