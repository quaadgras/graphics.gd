package main

import (
	"fmt"
	"io"

	"graphics.gd/internal/gdjson"
	"graphics.gd/internal/gdtype"
)

func (classDB ClassDB) new(file io.Writer, class gdjson.Class) {
	fmt.Fprintf(file, "func New() Instance {\n")
	fmt.Fprintln(file, `
	if !gd.Linked {
		var placeholder Instance
		*(*gd.Object)(unsafe.Pointer(&placeholder)) = pointers.Add[gd.Object]([3]uint64{})
		gd.StartupFunctions = append(gd.StartupFunctions, func() {
			if gd.Linked {
				raw, _ := pointers.End(New().AsObject()[0])
				pointers.Set(*(*gd.Object)(unsafe.Pointer(&placeholder)), raw)
				gd.RegisterCleanup(func() {
					if raw := pointers.Get[gd.Object](placeholder.AsObject()[0]); raw[0] != 0 && raw[1] == 0 {
						gdextension.Host.Objects.Unsafe.Free(gdextension.Object(raw[0]))
					}
				})
			}
		})
		return placeholder
	}`)
	fmt.Fprintf(file, "\tobject := [1]gd.Object{pointers.New[gd.Object]([3]uint64{uint64(gdextension.Host.Objects.Make(sname))})}\n")
	fmt.Fprintf(file, "\tcasted := Instance{*(*gdclass.%v)(unsafe.Pointer(&object))}\n", class.Name)
	if class.IsRefcounted {
		//fmt.Fprintf(file, "\tcasted.AsRefCounted()[0].Reference()\n")
	}
	fmt.Fprintf(file, "\tobject[0].Notification(0, false)\n")
	fmt.Fprintf(file, "\treturn casted\n")
	fmt.Fprintf(file, "}\n")
}

func (classDB ClassDB) properties(file io.Writer, class gdjson.Class, singleton bool) {
	if len(class.Properties) == 0 {
		return
	}
	for _, prop := range class.Properties {
		ptype := classDB.convertTypeSimple(class, class.Name+"."+prop.Name, "", prop.Type)
		expert := gdtype.EngineTypeAsGoType(class.Name, "", prop.Type)
		var foundGetter bool
		var foundSetter bool
		if prop.Getter != "" {
			for _, method := range class.Methods {
				if gdjson.Relocations[class.Name+"."+method.Name] != "" {
					continue
				}
				if method.Name == prop.Getter {
					ptype = classDB.convertTypeSimple(class, class.Name+"."+prop.Name, method.ReturnValue.Meta, method.ReturnValue.Type)
					foundGetter = true
					expert = gdtype.EngineTypeAsGoType(class.Name, method.ReturnValue.Meta, method.ReturnValue.Type)
					break
				}
			}
		}
		if prop.Setter != "" {
			for _, method := range class.Methods {
				if gdjson.Relocations[class.Name+"."+method.Name] != "" {
					continue
				}
				if method.Name == prop.Setter {
					var i = 0
					if prop.Index != nil {
						i = 1
					}
					ptype = classDB.convertTypeSimple(class, class.Name+"."+prop.Name, method.Arguments[i].Meta, method.Arguments[i].Type)
					expert = gdtype.EngineTypeAsGoType(class.Name, method.Arguments[i].Meta, method.Arguments[i].Type)
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
				fmt.Fprintf(file, "once.Do(singleton)\n\t")
			} else {
				fmt.Fprintf(file, "\nfunc (self Instance) %s() %s {\n", convertName(prop.Name), ptype)
			}
			val := fmt.Sprintf("class(self).%s()", convertName(prop.Getter))
			if prop.Index != nil {
				val = fmt.Sprintf("class(self).%s(%d)", convertName(prop.Getter), *prop.Index)
			}
			fmt.Fprintf(file, "\t\treturn %s(%s)\n", ptype, gdtype.Name(expert).ConvertToGo(val, ptype))
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
					ptype = classDB.convertTypeSimple(class, class.Name+"."+prop.Name, method.Arguments[i].Meta, method.Arguments[i].Type)
					expert = gdtype.EngineTypeAsGoType(class.Name, method.Arguments[i].Meta, method.Arguments[i].Type)
					found = true
					break
				}
			}
			if !found {
				continue
			}
			if singleton {
				fmt.Fprintf(file, "\nfunc Set%s(value %s) {\n", convertName(prop.Name), ptype)
				fmt.Fprintf(file, "once.Do(singleton)\n\t")
			} else {
				fmt.Fprintf(file, "\nfunc (self Instance) Set%s(value %s) {\n", convertName(prop.Name), ptype)
			}
			if prop.Index != nil {
				fmt.Fprintf(file, "\tclass(self).%s(%d, %s)\n", convertName(prop.Setter), *prop.Index, gdtype.Name(expert).ConvertToSimple("value", ptype))
			} else {
				fmt.Fprintf(file, "\tclass(self).%s(%s)\n", convertName(prop.Setter), gdtype.Name(expert).ConvertToSimple("value", ptype))
			}
			fmt.Fprintf(file, "}\n")
		}
	}
}
