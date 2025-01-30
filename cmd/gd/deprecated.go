package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"runtime"
	"slices"
	"strings"

	"golang.org/x/tools/go/packages"
	"graphics.gd/cmd/gd/internal/refactor/eg"
	"graphics.gd/variant/String"
	"runtime.link/api/xray"

	_ "embed"
)

//go:embed deprecated.txt
var fixes string

func checkForFixes(undefined []string) {
	for example := range String.Splits(fixes, "\n\n") {
		_, before, _ := strings.Cut(example, "func before(")
		_, name, _ := strings.Cut(before, "{")
		_, nameAfterReturn, ok := strings.Cut(name, "return")
		if ok {
			name = nameAfterReturn
		}
		name, _, _ = strings.Cut(name, "(")
		name = strings.TrimSpace(name)
		if slices.Contains(undefined, name) {
			fmt.Fprintln(os.Stderr)
			fmt.Fprintln(os.Stderr, "NOTE it looks like some of your compilation errors may be fixed by running `gd fix`")
			fmt.Fprintln(os.Stderr, "this will rewrite your project to refactor deprecated functions to use the new API.")
			fmt.Fprintln(os.Stderr, "(you should back up your code or use version control before running this command).")
			fmt.Fprintln(os.Stderr)
			return
		}
	}
}

func fix() error {
	cfg := &packages.Config{
		Fset:  token.NewFileSet(),
		Mode:  packages.NeedName | packages.NeedTypes | packages.NeedSyntax | packages.NeedImports | packages.NeedDeps | packages.NeedCompiledGoFiles,
		Tests: true,
	}
	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return xray.New(err)
	}
	var transformers []*eg.Transformer
	for example := range String.Splits(fixes, "\n\n") {
		f, err := parser.ParseFile(cfg.Fset, "/tmp/fixes.go", strings.NewReader(example), parser.ParseComments)
		if err != nil {
			return xray.New(err)
		}
		tInfo := types.Info{
			Types:      make(map[ast.Expr]types.TypeAndValue),
			Defs:       make(map[*ast.Ident]types.Object),
			Uses:       make(map[*ast.Ident]types.Object),
			Implicits:  make(map[ast.Node]types.Object),
			Selections: make(map[*ast.SelectorExpr]*types.Selection),
			Scopes:     make(map[ast.Node]*types.Scope),
		}
		conf := types.Config{Importer: pkgsImporter(pkgs), Sizes: types.SizesFor("gc", runtime.GOARCH)}
		tPkg, _ := conf.Check("egtemplate", cfg.Fset, []*ast.File{f}, &tInfo)
		xform, err := eg.NewTransformer(cfg.Fset, tPkg, f, &tInfo, false)
		if err != nil {
			return xray.New(err)
		}
		transformers = append(transformers, xform)
	}
	var hadErrors bool
	for _, pkg := range pkgs {
		for i, filename := range pkg.CompiledGoFiles {
			if filename == "/tmp/fixes.go" {
				continue // Don't rewrite the template file.
			}
			file := pkg.Syntax[i]
			var n int
			for _, xform := range transformers {
				n += xform.Transform(pkg.TypesInfo, pkg.Types, file)
			}
			if n == 0 {
				continue
			}
			fmt.Fprintf(os.Stderr, "=== %s (%d matches)\n", filename, n)
			if err := eg.WriteAST(cfg.Fset, filename, file); err != nil {
				fmt.Fprintf(os.Stderr, "eg: %s\n", err)
				hadErrors = true
			}
		}
	}
	if hadErrors {
		os.Exit(1)
	}
	return nil
}

type pkgsImporter []*packages.Package

func (p pkgsImporter) Import(path string) (tpkg *types.Package, err error) {
	packages.Visit([]*packages.Package(p), func(pkg *packages.Package) bool {
		if pkg.PkgPath == path {
			tpkg = pkg.Types
			return false
		}
		return true
	}, nil)
	if tpkg != nil {
		return tpkg, nil
	}
	return nil, fmt.Errorf("package %q not found", path)
}
