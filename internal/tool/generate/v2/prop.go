package main

import (
	"fmt"
	"io"

	"graphics.gd/internal/gdjson"
	"graphics.gd/internal/gdtype"
)

func (classDB ClassDB) new(file io.Writer, class gdjson.Class) {
	fmt.Fprintf(file, "func New() Instance {\n")
	fmt.Fprintf(file, "\tobject := gd.Global.ClassDB.ConstructObject(gd.NewStringName(%q))\n", class.Name)
	fmt.Fprintf(file, "\treturn Instance{*(*gdclass.%v)(unsafe.Pointer(&object))}\n", class.Name)
	fmt.Fprintf(file, "}\n")
}

func (classDB ClassDB) properties(file io.Writer, class gdjson.Class, singleton bool) {
	if len(class.Properties) == 0 {
		return
	}
	for _, prop := range class.Properties {
		ptype := classDB.convertTypeSimple(class, "", "", prop.Type)
		expert := classDB.convertType("", "", prop.Type)
		var foundGetter bool
		var foundSetter bool
		if prop.Getter != "" {
			for _, method := range class.Methods {
				if method.Name == prop.Getter {
					ptype = classDB.convertTypeSimple(class, "", method.ReturnValue.Meta, method.ReturnValue.Type)
					foundGetter = true
					expert = classDB.convertType("", method.ReturnValue.Meta, method.ReturnValue.Type)
					break
				}
			}
		}
		if prop.Setter != "" {
			for _, method := range class.Methods {
				if method.Name == prop.Setter {
					var i = 0
					if prop.Index != nil {
						i = 1
					}
					ptype = classDB.convertTypeSimple(class, "", method.Arguments[i].Meta, method.Arguments[i].Type)
					expert = classDB.convertType("", method.Arguments[i].Meta, method.Arguments[i].Type)
					foundSetter = true
					break
				}
			}
		}
		if !foundGetter && !foundSetter {
			continue
		}
		if foundGetter {
			if singleton {
				fmt.Fprintf(file, "\nfunc %s() %s {\n", convertName(prop.Name), ptype)
			} else {
				fmt.Fprintf(file, "\nfunc (self Instance) %s() %s {\n", convertName(prop.Name), ptype)
			}
			val := fmt.Sprintf("class(self).%s()", convertName(prop.Getter))
			if prop.Index != nil {
				val = fmt.Sprintf("class(self).%s(%d)", convertName(prop.Getter), *prop.Index)
			}
			fmt.Fprintf(file, "\t\treturn %s(%s)\n", ptype, gdtype.Name(expert).ConvertToGo(val))
			fmt.Fprintf(file, "}\n")
		}

		if prop.Setter != "" {
			var found = true
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
					ptype = classDB.convertTypeSimple(class, "", method.Arguments[i].Meta, method.Arguments[i].Type)
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
				fmt.Fprintf(file, "\nfunc (self Instance) Set%s(value %s) {\n", convertName(prop.Name), ptype)
			}
			if prop.Index != nil {
				fmt.Fprintf(file, "\tclass(self).%s(%d, %s)\n", convertName(prop.Setter), *prop.Index, gdtype.Name(expert).ConvertToSimple("value"))
			} else {
				fmt.Fprintf(file, "\tclass(self).%s(%s)\n", convertName(prop.Setter), gdtype.Name(expert).ConvertToSimple("value"))
			}
			fmt.Fprintf(file, "}\n")
		}
	}
}
