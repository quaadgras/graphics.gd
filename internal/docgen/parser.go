package docgen

import (
	"fmt"
	"go/ast"
	"go/doc"
	"go/types"
	"reflect"
	"strings"

	"golang.org/x/tools/go/packages"
	"graphics.gd/variant/String"
)

func parseDocumentation(path string) (XML, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedImports,
		Dir:  path,
	}
	pkgs, err := packages.Load(cfg, "graphics.gd/internal/gdclass", ".")
	if err != nil {
		return XML{}, fmt.Errorf("load error: %w", err)
	}
	var gdclassInterface *types.Interface
	for _, pkg := range pkgs {
		if pkg.Name == "gdclass" {
			gdclassInterface = pkg.Types.Scope().Lookup("Interface").Type().(*types.Named).Underlying().(*types.Interface)
			break
		}
	}
	var seen = make(map[string]bool)
	seen["graphics.gd/internal/gdclass"] = true
	var docs XML
	for _, pkg := range pkgs {
		docs.generateFromPackage(pkg, seen, gdclassInterface)
	}
	return docs, nil
}

func documentMemberSignals(members *[]Member, signals *[]Signal, ttype *types.Named, dtype *doc.Type) {
	asStruct, ok := ttype.Underlying().(*types.Struct)
	if !ok {
		return
	}
	var comments = make(map[string]string)
	if dtype != nil {
		for _, spec := range dtype.Decl.Specs {
			ts, _ := spec.(*ast.TypeSpec)
			st, ok := ts.Type.(*ast.StructType)
			if !ok {
				continue
			}
			for _, field := range st.Fields.List {
				for _, name := range field.Names {
					if field.Doc != nil {
						comments[name.Name] = strings.TrimSpace(field.Doc.Text())
					} else if field.Comment != nil {
						comments[name.Name] = strings.TrimSpace(field.Comment.Text())
					}
				}
			}
		}
	}
	for i := range asStruct.NumFields() {
		field := asStruct.Field(i)
		if field.Anonymous() {
			asNamed, ok := field.Type().(*types.Named)
			if ok {
				documentMemberSignals(members, signals, asNamed, nil) // TODO/FIXME get the doc from the package that defines the named type.
			}
			continue
		}
		if !field.Exported() {
			continue
		}
		name := String.ToSnakeCase(field.Name())
		if rename, ok := reflect.StructTag(asStruct.Tag(i)).Lookup("gd"); ok {
			if rename == "-" {
				continue // skip this field
			}
			name = rename
		}
		*members = append(*members, Member{
			Name:        name,
			Description: comments[field.Name()],
		})
	}
}

func (docs *XML) generateFromPackage(pkg *packages.Package, seen map[string]bool, gdclassInterface *types.Interface) error {
	if strings.HasPrefix(pkg.PkgPath, "graphics.gd") || !strings.Contains(pkg.PkgPath, ".") {
		return nil
	}
	if seen[pkg.PkgPath] {
		return nil
	}
	seen[pkg.PkgPath] = true
	dp, err := doc.NewFromFiles(pkg.Fset, pkg.Syntax, pkg.PkgPath)
	if err != nil {
		return fmt.Errorf("failed to create doc from files for package %q: %w", pkg.PkgPath, err)
	}
	for _, documented := range dp.Types {
		obj := pkg.Types.Scope().Lookup(documented.Name)
		if typeName, ok := obj.(*types.TypeName); ok {
			named, ok := typeName.Type().(*types.Named)
			if !ok {
				continue
			}
			underlying := named.Underlying()
			structType, isStruct := underlying.(*types.Struct)
			if !isStruct || structType.NumFields() == 0 || !structType.Field(0).Anonymous() {
				continue
			}
			if types.Implements(typeName.Type(), gdclassInterface) {
				var members []Member
				var methods []Method
				var signals []Signal
				documentMemberSignals(&members, &signals, named, documented)
				for _, method := range documented.Methods {
					name := String.ToSnakeCase(method.Name)
					doc := strings.TrimSpace(method.Doc)
					if strings.HasPrefix(doc, method.Name) {
						doc = name + strings.TrimPrefix(doc, method.Name)
					}
					methods = append(methods, Method{
						Name:        name,
						Description: doc,
					})
				}
				*docs = append(*docs, Class{
					Name:        typeName.Name(),
					Version:     "4.0",
					Description: strings.TrimSpace(documented.Doc),
					Members:     members,
					Methods:     methods,
				})
			}
		}
	}
	for _, imp := range pkg.Imports {
		if err := docs.generateFromPackage(imp, seen, gdclassInterface); err != nil {
			return err
		}
	}
	return nil
}
