/*
[gdscript]
func _ready():
	var monitor_value = Callable(self, "get_monitor_value")

	# Adds monitor with name "MyName" to category "MyCategory".
	Performance.add_custom_monitor("MyCategory/MyMonitor", monitor_value)

	# Adds monitor with name "MyName" to category "Custom".
	# Note: "MyCategory/MyMonitor" and "MyMonitor" have same name but different IDs, so the code is valid.
	Performance.add_custom_monitor("MyMonitor", monitor_value)

	# Adds monitor with name "MyName" to category "Custom".
	# Note: "MyMonitor" and "Custom/MyMonitor" have same name and same category but different IDs, so the code is valid.
	Performance.add_custom_monitor("Custom/MyMonitor", monitor_value)

	# Adds monitor with name "MyCategoryOne/MyCategoryTwo/MyMonitor" to category "Custom".
	Performance.add_custom_monitor("MyCategoryOne/MyCategoryTwo/MyMonitor", monitor_value)

func get_monitor_value():
	return randi() % 25
[/gdscript]
[csharp]
public override void _Ready()
{
	var monitorValue = new Callable(this, MethodName.GetMonitorValue);

	// Adds monitor with name "MyName" to category "MyCategory".
	Performance.AddCustomMonitor("MyCategory/MyMonitor", monitorValue);
	// Adds monitor with name "MyName" to category "Custom".
	// Note: "MyCategory/MyMonitor" and "MyMonitor" have same name but different ids so the code is valid.
	Performance.AddCustomMonitor("MyMonitor", monitorValue);

	// Adds monitor with name "MyName" to category "Custom".
	// Note: "MyMonitor" and "Custom/MyMonitor" have same name and same category but different ids so the code is valid.
	Performance.AddCustomMonitor("Custom/MyMonitor", monitorValue);

	// Adds monitor with name "MyCategoryOne/MyCategoryTwo/MyMonitor" to category "Custom".
	Performance.AddCustomMonitor("MyCategoryOne/MyCategoryTwo/MyMonitor", monitorValue);
}

public int GetMonitorValue()
{
	return GD.Randi() % 25;
}
[/csharp]
*/

package main

import (
	"math/rand/v2"

	"graphics.gd/classdb/Performance"
	"graphics.gd/variant/Callable"
)

func Performance_AddCustomMonitor() {
	var monitorValue = func() int {
		return rand.IntN(25)
	}
	// Adds monitor with name "MyName" to category "MyCategory".
	Performance.AddCustomMonitor("MyCategory/MyMonitor", Callable.New(monitorValue), nil)

	// Adds monitor with name "MyName" to category "Custom".
	// Note: "MyCategory/MyMonitor" and "MyMonitor" have same name but different ids so the code is valid.
	Performance.AddCustomMonitor("MyMonitor", Callable.New(monitorValue), nil)

	// Adds monitor with name "MyName" to category "Custom".
	// Note: "MyMonitor" and "Custom/MyMonitor" have same name and same category but different ids so the code is valid.
	Performance.AddCustomMonitor("Custom/MyMonitor", Callable.New(monitorValue), nil)

	// Adds monitor with name "MyCategoryOne/MyCategoryTwo/MyMonitor" to category "Custom".
	Performance.AddCustomMonitor("MyCategoryOne/MyCategoryTwo/MyMonitor", Callable.New(monitorValue), nil)
}
