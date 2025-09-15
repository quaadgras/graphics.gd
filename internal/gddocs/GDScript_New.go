/*
var MyClass = load("myclass.gd")
var instance = MyClass.new()
print(instance.get_script() == MyClass) # Prints true
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/GDScript"
	"graphics.gd/classdb/Resource"
	"graphics.gd/classdb/Script"
	"graphics.gd/variant/Object"
)

func GDScript_New() {
	var MyClass = Resource.Load[GDScript.Instance]("myclass.gd")
	var instance = MyClass.New()
	script, _ := Script.Get(instance.(Object.Instance))
	fmt.Println(Object.Aliases(MyClass, script)) // Prints true
}
