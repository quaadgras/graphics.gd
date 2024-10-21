package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"grow.graphics/gd/internal/gdjson"
	"runtime.link/api/xray"
)

func main() {
	if err := generate(); err != nil {
		log.Fatal(err)
	}
}

func generate() error {
	spec, err := LoadSpecification()
	if err != nil {
		return xray.New(err)
	}
	if err := os.MkdirAll("./object", 0755); err != nil {
		return xray.New(err)
	}
	var classDB = make(ClassDB)
	for _, enum := range spec.GlobalEnums {
		classDB[strings.Replace(enum.Name, ".", "", -1)] = gdjson.Class{
			IsEnum:  true,
			Name:    strings.Replace(enum.Name, ".", "", -1),
			Package: "internal",
		}
	}
	for _, class := range spec.BuiltinClasses {
		for _, enum := range class.Enums {
			classDB[class.Name+strings.Replace(enum.Name, ".", "", -1)] = gdjson.Class{
				IsEnum:  true,
				Name:    class.Name + strings.Replace(enum.Name, ".", "", -1),
				Package: "internal",
			}
		}
	}
	for i, class := range spec.Classes {
		var pkg = "internal"
		if !inCore(class.Name) {
			pkg = "classdb"
		}
		class.Package = pkg
		spec.Classes[i] = class
		classDB[class.Name] = class
	}
	file, err := os.Create("./object/objects.go")
	if err != nil {
		return xray.New(err)
	}
	defer file.Close()
	fmt.Fprintln(file, `package object`)
	fmt.Fprintln(file)
	fmt.Fprintln(file, `import classdb "grow.graphics/gd/internal/classdb"`)
	fmt.Fprintln(file)
	for _, class := range spec.Classes {
		if inCore(class.Name) {
			continue
		}
		fmt.Fprintf(file, "type %s = [1]classdb.%[1]s\n", class.Name)
		if err := classDB.generateObjectPackage(class); err != nil {
			return xray.New(err)
		}
	}
	return nil
}

func generateEnum(code io.Writer, prefix string, enum gdjson.Enum) {
	fmt.Fprintf(code, "type %v = classdb.%s%s\n\n", enum.Name, prefix, enum.Name)
	fmt.Fprintf(code, "const (\n")
	for _, value := range enum.Values {
		n := convertName(value.Name)
		if n == enum.Name {
			n += "Default"
		}
		if value.Description != "" {
			fmt.Fprint(code, "/*")
			fmt.Fprint(code, value.Description)
			fmt.Fprintln(code, "*/")
		}
		fmt.Fprintf(code, "\t%v %v = %v\n", n, enum.Name, value.Value)
	}
	fmt.Fprintf(code, ")\n")
}

func (classDB ClassDB) generateObjectPackage(class gdjson.Class) error {
	if err := os.MkdirAll("./object/"+class.Name, 0755); err != nil {
		return xray.New(err)
	}
	file, err := os.Create("./object/" + class.Name + "/class.go")
	if err != nil {
		return xray.New(err)
	}
	defer file.Close()
	fmt.Fprintf(file, "package %s\n\n", class.Name)
	fmt.Fprintln(file, `import "unsafe"`)
	fmt.Fprintln(file, `import "reflect"`)
	fmt.Fprintln(file, `import "grow.graphics/gd/internal/mmm"`)
	fmt.Fprintln(file, `import "grow.graphics/gd/internal/callframe"`)
	fmt.Fprintln(file, `import gd "grow.graphics/gd/internal"`)
	fmt.Fprintln(file, `import object "grow.graphics/gd/object"`)
	fmt.Fprintln(file, `import classdb "grow.graphics/gd/internal/classdb"`)
	var imported = make(map[string]bool)
	if class.Inherits != "" {
		super := classDB[class.Inherits]
		for super.Name != "" && super.Name != "Object" && super.Name != "RefCounted" {
			imported[super.Name] = true
			fmt.Fprintf(file, "import \"grow.graphics/gd/object/%s\"\n", super.Name)
			super = classDB[super.Inherits]
		}
	}
	fmt.Fprintln(file)
	fmt.Fprintln(file, "var _ unsafe.Pointer")
	fmt.Fprintln(file, "var _ object.Engine")
	fmt.Fprintln(file, "var _ reflect.Type")
	fmt.Fprintln(file, "var _ callframe.Frame")
	fmt.Fprintln(file, "var _ mmm.Lifetime")
	fmt.Fprintln(file)
	if class.Description != "" {
		fmt.Fprintln(file, "/*")
		fmt.Fprint(file, strings.Replace(class.Description, "*/", "", -1))
		fmt.Fprintln(file)
		var hasVirtual bool
		for _, method := range class.Methods {
			if method.IsVirtual {
				hasVirtual = true
				break
			}
		}
		if hasVirtual {
			fmt.Fprintf(file, "\t// %s methods that can be overridden by a [Class] that extends it.\n", class.Name)
			fmt.Fprintf(file, "\ttype %s interface {\n", class.Name)
			for _, method := range class.Methods {
				if method.IsVirtual {
					if method.Description != "" {
						description := strings.Replace(method.Description, "*/", "", -1)
						description = strings.TrimSpace(description)
						description = strings.Replace(description, "\n", "\n\t\t//", -1)
						fmt.Fprintln(file, "\t\t//"+description)
					}
					fmt.Fprintf(file, "\t\t%s(", convertName(method.Name))
					for i, arg := range method.Arguments {
						if i > 0 {
							fmt.Fprint(file, ", ")
						}
						fmt.Fprint(file, fixReserved(arg.Name), " ", classDB.convertType(class.Package, arg.Meta, arg.Type))
					}
					fmt.Fprint(file, ") ", classDB.convertType(class.Package, method.ReturnValue.Meta, method.ReturnValue.Type))
					fmt.Fprintln(file)
				}
			}
			fmt.Fprintln(file, "\t}")
		}
		fmt.Fprintln(file, "\n*/")
	}
	fmt.Fprintf(file, "type Simple [1]classdb.%s\n", class.Name)
	for _, method := range class.Methods {
		classDB.simpleCall(file, class, method)
	}
	fmt.Fprintf(file, `// Expert 1:1 low-level instance of the class, undocumented, for those who know what they are doing.`)
	fmt.Fprintf(file, "\ntype Expert = class\n")
	fmt.Fprintf(file, "type class [1]classdb.%s\n", class.Name)
	fmt.Fprintln(file, "func (self class) AsObject() gd.Object { return self[0].AsObject() }")
	fmt.Fprintln(file, "func (self Simple) AsObject() gd.Object { return self[0].AsObject() }")
	fmt.Fprintf(file, "\n\n//go:nosplit\nfunc (self *Simple) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }\n")
	fmt.Fprintf(file, "\n\n//go:nosplit\nfunc (self *class) SetPointer(ptr gd.Pointer) { self[0].SetPointer(ptr) }\n")
	for _, method := range class.Methods {
		classDB.methodCall(file, class.Package, class, method, callDefault)
	}
	if class.Inherits != "" {
		var i = 1
		super := classDB[class.Inherits]
		fmt.Fprintf(file, "\n\n//go:nosplit\nfunc (self class) As%[1]v() Expert { return self[0].As%[1]v() }\n", class.Name)
		fmt.Fprintf(file, "\n\n//go:nosplit\nfunc (self Simple) As%[1]v() Simple { return self[0].As%[1]v() }\n", class.Name)
		for super.Name != "" && super.Name != "Object" {
			if super.Name == "RefCounted" {
				fmt.Fprintf(file, "\n\n//go:nosplit\nfunc (self class) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }\n")
				fmt.Fprintf(file, "\n\n//go:nosplit\nfunc (self Simple) AsRefCounted() gd.RefCounted { return self[0].AsRefCounted() }\n")
			} else {
				fmt.Fprintf(file, "\n\n//go:nosplit\nfunc (self class) As%[2]v() %[2]v.Expert { return self[0].As%[2]v() }\n", class.Name, super.Name)
				fmt.Fprintf(file, "\n\n//go:nosplit\nfunc (self Simple) As%[2]v() %[2]v.Simple { return self[0].As%[2]v() }\n", class.Name, super.Name)
			}
			i++
			super = classDB[super.Inherits]
		}
	}
	for _, self := range []string{"class", "Simple"} {
		fmt.Fprintf(file, "\nfunc (self %s) Virtual(name string) reflect.Value {\n", self)
		fmt.Fprintf(file, "\tswitch name {\n")
		for _, method := range class.Methods {
			if method.IsVirtual {
				fmt.Fprintf(file, "\tcase \"%v\": return reflect.ValueOf(self.%v);\n", method.Name, method.Name)
			}
		}
		if class.Inherits != "" {
			fmt.Fprintf(file, "\tdefault: return gd.VirtualByName(self[0].Super()[0], name)\n")
		} else {
			fmt.Fprintf(file, "\tdefault: return reflect.Value{}\n")
		}
		fmt.Fprintf(file, "\t}\n")
		fmt.Fprintf(file, "}\n")
	}
	fmt.Fprintf(file, `func init() {`)
	fmt.Fprintf(file, `classdb.Register("%s", func(ptr gd.Pointer) any {var class Expert; class[0].SetPointer(ptr); return class })`, class.Name)
	fmt.Fprintf(file, "}\n")
	for _, enum := range class.Enums {
		generateEnum(file, class.Name, enum)
	}
	return nil
}
