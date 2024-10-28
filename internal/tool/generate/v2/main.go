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
	if err := os.MkdirAll("./gdclass", 0755); err != nil {
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
	if err := os.MkdirAll("./gdenums", 0755); err != nil {
		return xray.New(err)
	}
	enums, err := os.Create("./gdenums/enums.go")
	if err != nil {
		return xray.New(err)
	}
	fmt.Fprintln(enums, `package gdenums`)
	fmt.Fprintln(enums)
	for _, enum := range spec.GlobalEnums {
		generateEnum(enums, "", enum, "")
		fmt.Fprintln(enums)
	}
	defer enums.Close()
	for i, class := range spec.Classes {
		var pkg = "internal"
		if !inCore(class.Name) {
			pkg = "classdb"
		}
		class.Package = pkg
		spec.Classes[i] = class
		classDB[class.Name] = class
	}
	var singletons = make(map[string]bool)
	for _, class := range spec.Singletons {
		singletons[class.Name] = true
		mod := classDB[class.Name]
		mod.IsSingleton = true
		classDB[class.Name] = mod
	}
	file, err := os.Create("./gdclass/gdclass.go")
	if err != nil {
		return xray.New(err)
	}
	defer file.Close()
	fmt.Fprintln(file, `package gdclass`)
	fmt.Fprintln(file)
	fmt.Fprintln(file, `import classdb "grow.graphics/gd/internal/classdb"`)
	fmt.Fprintln(file)
	for _, class := range spec.Classes {
		if inCore(class.Name) {
			continue
		}
		fmt.Fprintf(file, "type %s = [1]classdb.%[1]s\n", class.Name)
		if err := classDB.generateObjectPackage(class, singletons[class.Name]); err != nil {
			return xray.New(err)
		}
	}
	return nil
}

func generateEnum(code io.Writer, prefix string, enum gdjson.Enum, classdb string) {
	rename := enum.Name
	if enum.Name == "MouseMode" {
		rename = "MouseModeValue"
	}
	rename = strings.Replace(rename, ".", "", -1)
	enum.Name = strings.Replace(enum.Name, ".", "", -1)

	if classdb != "" {
		fmt.Fprintf(code, "type %v = %s%s%s\n\n", rename, classdb, prefix, enum.Name)
	} else {
		fmt.Fprintf(code, "type %v int\n\n", rename)
	}
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
		fmt.Fprintf(code, "\t%v %v = %v\n", n, rename, value.Value)
	}
	fmt.Fprintf(code, ")\n")
}

func (classDB ClassDB) generateObjectPackage(class gdjson.Class, singleton bool) error {
	if err := os.MkdirAll("./gdclass/"+class.Name, 0755); err != nil {
		return xray.New(err)
	}
	file, err := os.Create("./gdclass/" + class.Name + "/class.go")
	if err != nil {
		return xray.New(err)
	}
	defer file.Close()
	fmt.Fprintf(file, "package %s\n\n", class.Name)
	fmt.Fprintln(file, `import "unsafe"`)
	if singleton {
		fmt.Fprintln(file, `import "sync"`)
	}
	fmt.Fprintln(file, `import "reflect"`)
	fmt.Fprintln(file, `import "grow.graphics/gd/internal/discreet"`)
	fmt.Fprintln(file, `import "grow.graphics/gd/internal/callframe"`)
	fmt.Fprintln(file, `import gd "grow.graphics/gd/internal"`)
	fmt.Fprintln(file, `import "grow.graphics/gd/gdclass"`)
	fmt.Fprintln(file, `import classdb "grow.graphics/gd/internal/classdb"`)
	var imported = make(map[string]bool)
	if class.Inherits != "" {
		super := classDB[class.Inherits]
		for super.Name != "" && super.Name != "Object" && super.Name != "RefCounted" && !classDB[super.Name].IsSingleton {
			imported[super.Name] = true
			fmt.Fprintf(file, "import \"grow.graphics/gd/gdclass/%s\"\n", super.Name)
			super = classDB[super.Inherits]
		}
	}
	fmt.Fprintln(file)
	fmt.Fprintln(file, "var _ unsafe.Pointer")
	fmt.Fprintln(file, "var _ gdclass.Engine")
	fmt.Fprintln(file, "var _ reflect.Type")
	fmt.Fprintln(file, "var _ callframe.Frame")
	fmt.Fprintln(file, "var _ = discreet.Root")
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
						fmt.Fprint(file, fixReserved(arg.Name), " ", classDB.convertTypeSimple(arg.Meta, arg.Type))
					}
					fmt.Fprint(file, ") ", classDB.convertTypeSimple(method.ReturnValue.Meta, method.ReturnValue.Type))
					fmt.Fprintln(file)
				}
			}
			fmt.Fprintln(file, "\t}")
		}
		fmt.Fprintln(file, "\n*/")
	}
	if singleton {
		fmt.Fprintf(file, "var self gdclass.%s\n", class.Name)
		fmt.Fprintf(file, "var once sync.Once\n")
		fmt.Fprintf(file, "func singleton() {\n")
		fmt.Fprintf(file, "\tobj := gd.Global.Object.GetSingleton(gd.Global.Singletons.%s)\n", class.Name)
		fmt.Fprintf(file, "\tself = *(*gdclass.%s)(unsafe.Pointer(&obj))\n", class.Name)
		fmt.Fprintf(file, "}\n")
	} else {
		fmt.Fprintf(file, "type Go [1]classdb.%s\n", class.Name)
	}
	var getter_setters = make(map[string]bool)
	for _, property := range class.Properties {
		if property.Getter != "" {
			getter_setters[property.Getter] = true
		}
		if property.Setter != "" {
			getter_setters[property.Setter] = true
		}
	}
	for _, method := range class.Methods {
		if getter_setters[method.Name] {
			continue
		}
		classDB.simpleCall(file, class, method, singleton)
	}
	fmt.Fprintf(file, `// GD is a 1:1 low-level instance of the class, undocumented, for those who know what they are doing.`)
	if singleton {
		fmt.Fprintf(file, "\nfunc GD() class { once.Do(singleton); return self }\n")
	} else {
		fmt.Fprintf(file, "\ntype GD = class\n")
	}
	fmt.Fprintf(file, "type class [1]classdb.%s\n", class.Name)
	fmt.Fprintln(file, "func (self class) AsObject() gd.Object { return self[0].AsObject() }")
	if !singleton {
		fmt.Fprintln(file, "func (self Go) AsObject() gd.Object { return self[0].AsObject() }")
	}
	if !singleton {
		classDB.new(file, class)
	}
	classDB.properties(file, class, singleton)
	for _, method := range class.Methods {
		classDB.methodCall(file, class.Package, class, method, callDefault)
	}
	for _, signal := range class.Signals {
		classDB.signalCall(file, class, signal, singleton)
	}
	super := classDB[class.Inherits]
	if class.Inherits != "" {
		var i = 1
		if !singleton {
			fmt.Fprintf(file, "\nfunc (self class) As%[1]v() GD { return *((*GD)(unsafe.Pointer(&self))) }\n", class.Name)
			fmt.Fprintf(file, "func (self Go) As%[1]v() Go { return *((*Go)(unsafe.Pointer(&self))) }\n", class.Name)
		}
		super := classDB[class.Inherits]
		for super.Name != "" && super.Name != "Object" {
			if classDB[super.Name].IsSingleton {
				super = classDB[super.Inherits]
				continue
			}
			if super.Name == "RefCounted" {
				fmt.Fprintf(file, "func (self class) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }\n")
				if !singleton {
					fmt.Fprintf(file, "func (self Go) AsRefCounted() gd.RefCounted { return *((*gd.RefCounted)(unsafe.Pointer(&self))) }\n")
				}
			} else {
				fmt.Fprintf(file, "func (self class) As%[2]v() %[2]v.GD { return *((*%[2]v.GD)(unsafe.Pointer(&self))) }\n", class.Name, super.Name)
				if !singleton {
					fmt.Fprintf(file, "func (self Go) As%[2]v() %[2]v.Go { return *((*%[2]v.Go)(unsafe.Pointer(&self))) }\n", class.Name, super.Name)
				}
			}
			i++
			super = classDB[super.Inherits]
		}
	}
	for _, self := range []string{"class", "Go"} {
		if self == "Go" && singleton {
			continue
		}
		fmt.Fprintf(file, "\nfunc (self %s) Virtual(name string) reflect.Value {\n", self)
		fmt.Fprintf(file, "\tswitch name {\n")
		for _, method := range class.Methods {
			if method.IsVirtual {
				fmt.Fprintf(file, "\tcase \"%v\": return reflect.ValueOf(self.%v);\n", method.Name, method.Name)
			}
		}
		if class.Inherits != "" && !classDB[class.Inherits].IsSingleton {
			fmt.Fprintf(file, "\tdefault: return gd.VirtualByName(self.As%s(), name)\n", super.Name)
		} else {
			fmt.Fprintf(file, "\tdefault: return reflect.Value{}\n")
		}
		fmt.Fprintf(file, "\t}\n")
		fmt.Fprintf(file, "}\n")
	}
	fmt.Fprintf(file, `func init() {`)
	fmt.Fprintf(file, `classdb.Register("%s", func(ptr gd.Object) any { return classdb.%[1]s(ptr) })`, class.Name)
	fmt.Fprintf(file, "}\n")
	for _, enum := range class.Enums {
		generateEnum(file, class.Name, enum, "classdb.")
	}
	return nil
}
