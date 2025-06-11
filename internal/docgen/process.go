package docgen

import (
	"encoding/xml"
	"fmt"
	"os"
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

func serializeDocs(docs XML, outPath string) error {
	file, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("failed to create xml documentation file: %w", err)
	}
	defer file.Close()
	enc := xml.NewEncoder(file)
	enc.Indent("", "\t")
	if err := enc.Encode(docs); err != nil {
		return fmt.Errorf("failed to encode docs: %w", err)
	}
	return nil
}
