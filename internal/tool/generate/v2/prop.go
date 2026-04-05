package main

import (
	"fmt"
	"io"

	"graphics.gd/internal/gdjson"
	"graphics.gd/internal/tool/generate/gdtype"
)

func (classDB ClassDB) new(file io.Writer, class gdjson.Class) {
	fmt.Fprintf(file, "func New() Instance {\n")
	fmt.Fprintf(file, `if !gd.Linked {
		var placeholder = Instance([1]gdclass.%[1]s{gdclass.New%[1]s(gdreference.NewObject())})
		gd.StartupFunctions = append(gd.StartupFunctions, func() {
			if gd.Linked {
				raw, _ := gdreference.EndObject(New().AsObject()[0])
				gdreference.SetObject(gdclass.Get%[1]s(placeholder[0])[0], raw)
				gd.RegisterCleanup(func() {
					if raw := gdreference.GetObject(placeholder.AsObject()[0]); raw != 0 {
						gdunsafe.Object(raw).Free()
					}
				})
			}
		})
		return placeholder
	}
`, class.Name)
	fmt.Fprintf(file, "\tcasted := Instance([1]gdclass.%[1]s{gdclass.New%[1]s(gdreference.OwnObject(gdextension.Object(gdunsafe.MakeObject(gdunsafe.StringName(sname[0]))), gd.Free))})\n", class.Name)
	if class.IsRefcounted {
		fmt.Fprintf(file, "\tcasted.AsRefCounted()[0].InitRef()\n")
	}
	fmt.Fprintf(file, "\tgd.ObjectNotification(casted.AsObject()[0], 0, false)\n")
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
					ptype = classDB.convertTypeSimple(class, class.Name+"."+prop.Getter+".", method.ReturnValue.Meta, method.ReturnValue.Type)
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
					ptype = classDB.convertTypeSimple(class, class.Name+"."+prop.Setter+"."+prop.Name, method.Arguments[i].Meta, method.Arguments[i].Type)
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
			if prop.Description != "" {
				fmt.Fprintln(file, "\n/*")
				fmt.Fprint(file, gdjson.DocsToGoDoc(prop.Description, classDB, class.Name, class.Name+"_"+convertName(prop.Name)))
				fmt.Fprint(file, "\n*/")
			}
			if singleton {
				fmt.Fprintf(file, "\nfunc %s() %s { //gd:%s.%s\n", convertName(prop.Name), ptype, class.Name, prop.Name)
				fmt.Fprintf(file, "once.Do(singleton)\n\t")
			} else {
				fmt.Fprintf(file, "\nfunc (self Instance) %s() %s { //gd:%s.%s\n", convertName(prop.Name), ptype, class.Name, prop.Name)
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
					ptype = classDB.convertTypeSimple(class, class.Name+"."+prop.Setter+"."+prop.Name, method.Arguments[i].Meta, method.Arguments[i].Type)
					expert = gdtype.EngineTypeAsGoType(class.Name, method.Arguments[i].Meta, method.Arguments[i].Type)
					found = true
					break
				}
			}
			if !found {
				continue
			}
			if !foundGetter {
				if prop.Description != "" {
					fmt.Fprintln(file, "\n/*")
					fmt.Fprint(file, gdjson.DocsToGoDoc(prop.Description, classDB, class.Name, class.Name+"_"+convertName(prop.Name)))
					if !singleton {
						fmt.Fprintf(file, "\nReturns the instance, so that property settings can be chained.")
					}
					fmt.Fprint(file, "\n*/")
				}
			} else {
				fmt.Fprintf(file, "\n// Set%s sets the property returned by [%s].", convertName(prop.Name), convertName(prop.Getter))
				if !singleton {
					fmt.Fprintf(file, " Returns the instance, so that property settings can be chained.")
				}
			}
			if singleton {
				fmt.Fprintf(file, "\nfunc Set%s(value %s) { //gd:%s.%s\n", convertName(prop.Name), ptype, class.Name, prop.Name)
				fmt.Fprintf(file, "once.Do(singleton)\n\t")
			} else {
				fmt.Fprintf(file, "\nfunc (self Instance) Set%s(value %s) Instance { //gd:%s.%s\n", convertName(prop.Name), ptype, class.Name, prop.Name)
			}
			if prop.Index != nil {
				fmt.Fprintf(file, "\tclass(self).%s(%d, %s)\n", convertName(prop.Setter), *prop.Index, gdtype.Name(expert).ConvertToSimple("value", ptype))
			} else {
				fmt.Fprintf(file, "\tclass(self).%s(%s)\n", convertName(prop.Setter), gdtype.Name(expert).ConvertToSimple("value", ptype))
			}
			if !singleton {
				fmt.Fprintf(file, "\treturn self\n")
			}
			fmt.Fprintf(file, "}\n")
		}
	}
}
