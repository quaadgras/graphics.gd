/*
var dummy_text_server = TextServerManager.find_interface("Dummy")
if dummy_text_server != null:
    TextServerManager.set_primary_interface(dummy_text_server)
    # If the other text servers are unneeded, they can be removed:
    for i in TextServerManager.get_interface_count():
        var text_server = TextServerManager.get_interface(i)
        if text_server != dummy_text_server:
            TextServerManager.remove_interface(text_server)
*/

package main

import (
	"graphics.gd/classdb/TextServer"
	"graphics.gd/classdb/TextServerManager"
)

func ExampleTextServerDummy() {
	var dummy_text_server = TextServerManager.FindInterface("Dummy")
	if dummy_text_server != TextServer.Nil {
		TextServerManager.SetPrimaryInterface(dummy_text_server)
		// If the other text servers are unneeded, they can be removed:
		for i := range TextServerManager.GetInterfaceCount() {
			var text_server = TextServerManager.GetInterface(i)
			if text_server != dummy_text_server {
				TextServerManager.RemoveInterface(text_server)
			}
		}
	}
}
