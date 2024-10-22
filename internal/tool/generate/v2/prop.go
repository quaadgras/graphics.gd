package main

import (
	"fmt"
	"io"

	"grow.graphics/gd/internal/gdjson"
	"grow.graphics/gd/internal/gdtype"
)

func (classDB ClassDB) new(file io.Writer, class gdjson.Class) {
	fmt.Fprintf(file, "func New() Go {\n")
	fmt.Fprintf(file, "\tgc := gd.GarbageCollector()\n")
	fmt.Fprintf(file, "\tobject := gc.API.ClassDB.ConstructObject(gc, gc.StringName(%q))\n", class.Name)
	fmt.Fprintf(file, "\treturn *(*Go)(unsafe.Pointer(&object))\n")
	fmt.Fprintf(file, "}\n")
}

func (classDB ClassDB) properties(file io.Writer, class gdjson.Class, singleton bool) {
	if len(class.Properties) == 0 {
		return
	}
	for _, prop := range class.Properties {
		ptype := classDB.convertTypeSimple("", prop.Type)
		expert := classDB.convertType("", "", prop.Type)
		var found bool
		var gc = ""
		if prop.Getter != "" {
			for _, method := range class.Methods {
				if method.Name == prop.Getter {
					ptype = classDB.convertTypeSimple(method.ReturnValue.Meta, method.ReturnValue.Type)
					found = true
					expert = classDB.convertType("", method.ReturnValue.Meta, method.ReturnValue.Type)
					if _, ok := classDB.isPointer(classDB.convertType("", method.ReturnValue.Meta, method.ReturnValue.Type)); ok {
						gc = "gc"
					}
					break
				}
			}
		}
		if !found && prop.Setter != "" {
			for _, method := range class.Methods {
				if method.Name == prop.Setter {
					var i = 0
					if prop.Index != nil {
						i = 1
					}
					ptype = classDB.convertTypeSimple(method.Arguments[i].Meta, method.Arguments[i].Type)
					expert = classDB.convertType("", method.Arguments[i].Meta, method.Arguments[i].Type)
					break
				}
			}
		}
		if !found {
			continue
		}
		if singleton {
			fmt.Fprintf(file, "\nfunc %s() %s {\n", convertName(prop.Name), ptype)
		} else {
			fmt.Fprintf(file, "\nfunc (self Go) %s() %s {\n", convertName(prop.Name), ptype)
		}
		fmt.Fprintf(file, "\tgc := gd.GarbageCollector(); _ = gc\n")
		val := fmt.Sprintf("class(self).%s(%s)", convertName(prop.Getter), gc)
		if prop.Index != nil {
			if gc != "" {
				gc = "gc,"
			}
			val = fmt.Sprintf("class(self).%s(%s%d)", convertName(prop.Getter), gc, *prop.Index)
		}
		fmt.Fprintf(file, "\t\treturn %s(%s)\n", ptype, gdtype.Name(expert).ConvertToGo(val))
		fmt.Fprintf(file, "}\n")

		if prop.Setter != "" {
			for _, method := range class.Methods {
				if convertName(method.Name) == convertName(prop.Setter) && method.Name != prop.Setter {
					found = false
					break
				}
			}
			if !found {
				continue
			}
			found = false
			for _, method := range class.Methods {
				if method.Name == prop.Setter {
					var i = 0
					if prop.Index != nil {
						i = 1
					}
					ptype = classDB.convertTypeSimple(method.Arguments[i].Meta, method.Arguments[i].Type)
					expert = classDB.convertType("", method.Arguments[i].Meta, method.Arguments[i].Type)
					found = true
					break
				}
			}
			if !found {
				continue
			}
			if singleton {
				fmt.Fprintf(file, "\nfunc Set%s(value %s) {\n", convertName(prop.Name), ptype)
			} else {
				fmt.Fprintf(file, "\nfunc (self Go) Set%s(value %s) {\n", convertName(prop.Name), ptype)
			}
			fmt.Fprintf(file, "\tgc := gd.GarbageCollector(); _ = gc\n")
			if prop.Index != nil {
				fmt.Fprintf(file, "\tclass(self).%s(%d, %s)\n", convertName(prop.Setter), *prop.Index, gdtype.Name(expert).ConvertToSimple("value"))
			} else {
				fmt.Fprintf(file, "\tclass(self).%s(%s)\n", convertName(prop.Setter), gdtype.Name(expert).ConvertToSimple("value"))
			}
			fmt.Fprintf(file, "}\n")
		}
	}
}
