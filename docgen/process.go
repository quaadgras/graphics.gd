package docgen

import "fmt"

func Process(packagePath string) error {
	gobPath, initPath := getDocPaths(packagePath)
	err := cleanFiles(packagePath)
	if err != nil {
		return fmt.Errorf("failed to clean documentation embedding files: %w", err)
	}
	pkgsDocs, err := parseDocumentation(packagePath)
	if err != nil {
		return fmt.Errorf("failed to parse documentation from Go source: %w", err)
	}
	err = serializeDocs(pkgsDocs, gobPath)
	if err != nil {
		return fmt.Errorf("failed to serialize documentation: %w", err)
	}
	err = writeInit(initPath)
	if err != nil {
		return fmt.Errorf("failed to write documentation init file: %w", err)
	}
	return nil
}
