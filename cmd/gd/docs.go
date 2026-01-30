package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"graphics.gd/cmd/gd/internal/tooling"
)

func doc(args ...string) error {
	if len(args) == 0 {
		return tooling.Go.Exec(append([]string{"doc"}, args...)...)
	}

	query := args[0]
	remaining := args[1:]

	// Search for matching //gd: comments in the classdb packages
	matches, err := findGdDocMatches(query)
	if err != nil {
		return err
	}

	if len(matches) == 0 {
		// No gd: matches found, fall back to regular go doc
		return tooling.Go.Exec(append([]string{"doc"}, args...)...)
	}

	// Display matching go doc entries
	for i, match := range matches {
		if i > 0 {
			fmt.Println()
			fmt.Println("---")
			fmt.Println()
		}
		docArgs := append([]string{"doc", match.goDocPath}, remaining...)
		if err := tooling.Go.Exec(docArgs...); err != nil {
			fmt.Fprintf(os.Stderr, "warning: could not get doc for %s: %v\n", match.goDocPath, err)
		}
	}

	return nil
}

type gdDocMatch struct {
	gdTag     string // e.g., "Viewport.set_input_as_handled"
	goDocPath string // e.g., "graphics.gd/classdb/Viewport.SetInputAsHandled"
}

func findGdDocMatches(query string) ([]gdDocMatch, error) {
	// Get the module root for graphics.gd
	goPath, err := tooling.Go.Lookup()
	if err != nil {
		return nil, err
	}

	// Find the classdb directory
	modRoot, err := exec.Command(goPath, "list", "-m", "-f", "{{.Dir}}", "graphics.gd").Output()
	if err != nil {
		return nil, fmt.Errorf("could not find graphics.gd module: %w", err)
	}
	classdbDir := filepath.Join(strings.TrimSpace(string(modRoot)), "classdb")

	var matches []gdDocMatch

	// Walk through all class.go files in classdb/*/
	entries, err := os.ReadDir(classdbDir)
	if err != nil {
		return nil, fmt.Errorf("could not read classdb directory: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		classFile := filepath.Join(classdbDir, entry.Name(), "class.go")
		file, err := os.Open(classFile)
		if err != nil {
			continue // Skip if class.go doesn't exist
		}

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()

			if strings.Contains(line, "(self class)") {
				continue
			}

			// Look for //gd: comments
			_, gdTag, ok := strings.Cut(line, "//gd:")
			if !ok {
				continue
			}

			// Parse the gd tag: ClassName.method_name
			parts := strings.SplitN(gdTag, ".", 2)
			if len(parts) != 2 {
				continue
			}

			className := parts[0]
			methodName := parts[1]

			// Check if this matches our query (case-insensitive substring match)
			if methodName != query {
				continue
			}

			// Convert method_name to MethodName (snake_case to PascalCase)
			goMethodName := snakeToPascal(methodName)

			// Build the go doc path
			goDocPath := fmt.Sprintf("graphics.gd/classdb/%s.%s", className, goMethodName)

			matches = append(matches, gdDocMatch{
				gdTag:     gdTag,
				goDocPath: goDocPath,
			})
		}
		file.Close()
	}

	return matches, nil
}

func snakeToPascal(s string) string {
	parts := strings.Split(s, "_")
	for i := range parts {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}
