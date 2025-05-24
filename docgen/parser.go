package docgen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/doc"
	"go/parser"
	"go/printer"
	"go/token"
	"golang.org/x/tools/go/packages"
	"path/filepath"
	"reflect"
	"strings"
)

var fset = token.NewFileSet()

type AllPackages map[string]*PackageDoc

type PackageDoc struct {
	Name  string
	Doc   string
	Funcs map[string]*FuncDoc
	Types map[string]*TypeDoc
}

type FuncDoc struct {
	Name string
	Doc  string
}

func (pd *PackageDoc) GetTypeDoc(typ reflect.Type) *TypeDoc {
	typeDoc, ok := pd.Types[typ.Name()]
	if !ok {
		return &TypeDoc{}
	}
	return typeDoc
}

func (pd *PackageDoc) GetFuncDoc(fn reflect.Value) *FuncDoc {
	fnName := fn.Type().Name()
	fnDoc, ok := pd.Funcs[fnName]
	if !ok {
		return &FuncDoc{}
	}
	return fnDoc
}

type TypeDoc struct {
	Name    string
	Doc     string
	Fields  map[string]*FieldDoc
	Methods map[string]*FuncDoc
}

type FieldDoc struct {
	Names []string
	Type  string
	Doc   string
}

func (td *TypeDoc) GetFieldDoc(fld reflect.StructField) *FieldDoc {
	if fld.Anonymous {
		return &FieldDoc{}
	}
	if fldDoc, ok := td.Fields[fld.Name]; ok {
		return fldDoc
	}
	return &FieldDoc{}
}

func (td *TypeDoc) GetMethodDoc(fn reflect.Method) *FuncDoc {
	if fnDoc, ok := td.Methods[fn.Name]; ok {
		return fnDoc
	}
	return &FuncDoc{}
}

// extractTypes extracts the types (structs) from the provided package documentation.
func extractTypes(docPkg *doc.Package) map[string]*TypeDoc {
	out := make(map[string]*TypeDoc)
	for _, t := range docPkg.Types {
		td := &TypeDoc{
			Name:    t.Name,
			Doc:     strings.TrimSpace(t.Doc),
			Fields:  make(map[string]*FieldDoc),
			Methods: make(map[string]*FuncDoc),
		}

		// get the fields
		for _, spec := range t.Decl.Specs {
			ts, _ := spec.(*ast.TypeSpec)
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				continue
			}
			for _, f := range st.Fields.List {
				// collect the field names (anonymous fields get no Names)
				names := make([]string, 0, len(f.Names))
				for _, n := range f.Names {
					names = append(names, n.Name)
				}
				// pretty-print the type expression
				var buf bytes.Buffer
				err := printer.Fprint(&buf, fset, f.Type)
				if err != nil {
					continue
				}

				// get any comment or doc on the field
				fieldDoc := ""
				if f.Doc != nil {
					fieldDoc = strings.TrimSpace(f.Doc.Text())
				} else if f.Comment != nil {
					fieldDoc = strings.TrimSpace(f.Comment.Text())
				}
				for _, name := range names {
					td.Fields[name] = &FieldDoc{
						Names: names,
						Type:  buf.String(),
						Doc:   fieldDoc,
					}
				}
			}
		}

		// get the methods
		for _, m := range t.Methods {
			td.Methods[m.Name] = &FuncDoc{
				Name: m.Name,
				Doc:  strings.TrimSpace(m.Doc),
			}
		}

		out[td.Name] = td
	}
	return out
}

func parseDocumentation(path string) (AllPackages, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax,
		Dir:  path,
	}
	// load all the packages in the provided path
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return nil, fmt.Errorf("load error: %w", err)
	}

	all := make(AllPackages)
	for _, p := range pkgs {
		if len(p.GoFiles) == 0 {
			continue
		}
		// get the directory of the package
		pkgDir := filepath.Dir(p.GoFiles[0])

		astPkgs, err := parser.ParseDir(fset, pkgDir, nil, parser.ParseComments)
		if err != nil {
			return nil, fmt.Errorf("parse error for %q: %w", p.PkgPath, err)
		}

		astPkg, ok := astPkgs[p.Name]
		if !ok {
			return nil, fmt.Errorf("AST package %q not found in %s", p.Name, pkgDir)
		}
		dp := doc.New(astPkg, p.PkgPath, 0)

		// construct the package documentation
		pd := &PackageDoc{
			Name:  dp.Name,
			Doc:   dp.Doc,
			Types: extractTypes(dp),
		}
		// construct the docs for top-level functions
		for _, f := range dp.Funcs {
			pd.Funcs[f.Name] = &FuncDoc{
				Name: f.Name,
				Doc:  f.Doc,
			}
		}
		all[p.PkgPath] = pd
	}

	return all, nil
}
