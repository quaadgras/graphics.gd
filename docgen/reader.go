package docgen

import (
	"bytes"
	_ "embed"
	"encoding/gob"
	"fmt"
	"reflect"
)

var Docs AllPackages

func RegisterDocs(b []byte) error {
	buf := bytes.NewReader(b)
	var all AllPackages
	if err := gob.NewDecoder(buf).Decode(&all); err != nil {
		return fmt.Errorf("failed to decode docs: %w", err)
	}
	Docs = all
	return nil
}

func LoadFor(typ reflect.Type) (*PackageDoc, bool) {
	path := typ.PkgPath()
	pd, ok := Docs[path]
	return pd, ok
}
