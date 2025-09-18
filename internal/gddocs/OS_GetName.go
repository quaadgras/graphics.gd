/*
[gdscript]
match OS.get_name():
	"Windows":
		print("Welcome to Windows!")
	"macOS":
		print("Welcome to macOS!")
	"Linux", "FreeBSD", "NetBSD", "OpenBSD", "BSD":
		print("Welcome to Linux/BSD!")
	"Android":
		print("Welcome to Android!")
	"iOS":
		print("Welcome to iOS!")
	"Web":
		print("Welcome to the Web!")
[/gdscript]
[csharp]
switch (OS.GetName())
{
	case "Windows":
		GD.Print("Welcome to Windows");
		break;
	case "macOS":
		GD.Print("Welcome to macOS!");
		break;
	case "Linux":
	case "FreeBSD":
	case "NetBSD":
	case "OpenBSD":
	case "BSD":
		GD.Print("Welcome to Linux/BSD!");
		break;
	case "Android":
		GD.Print("Welcome to Android!");
		break;
	case "iOS":
		GD.Print("Welcome to iOS!");
		break;
	case "Web":
		GD.Print("Welcome to the Web!");
		break;
}
[/csharp]
*/

package main

import (
	"fmt"

	"graphics.gd/classdb/OS"
)

func OS_GetName() {
	switch OS.GetName() {
	case "Windows":
		fmt.Println("Welcome to Windows!")
	case "macOS":
		fmt.Println("Welcome to macOS!")
	case "Linux", "FreeBSD", "NetBSD", "OpenBSD", "BSD":
		fmt.Println("Welcome to Linux/BSD!")
	case "Android":
		fmt.Println("Welcome to Android!")
	case "iOS":
		fmt.Println("Welcome to iOS!")
	case "Web":
		fmt.Println("Welcome to the Web!")
	default:
		fmt.Println("Unknown OS")
	}
}
