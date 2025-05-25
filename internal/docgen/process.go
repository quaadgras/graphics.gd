package docgen

import (
	"fmt"
	"path/filepath"
)

func Process(packagePath string) error {
	pkgsDocs, err := parseDocumentation(packagePath)
	if err != nil {
		return fmt.Errorf("failed to parse documentation from Go source: %w", err)
	}
	err = serializeDocs(pkgsDocs, filepath.Join(packagePath, "graphics", "library_documentation.xml"))
	if err != nil {
		return fmt.Errorf("failed to serialize documentation: %w", err)
	}
	return nil
}
