/*
for key in translation.get_message_list():
	var p = key.find("\u0004")
	if p == -1:
		var untranslated = key
		print("Message %s" % untranslated)
	else:
		var context = key.substr(0, p)
		var untranslated = key.substr(p + 1)
		print("Message %s with context %s" % [untranslated, context])
*/

package main

import (
	"fmt"
	"strings"

	"graphics.gd/classdb/Translation"
)

var translation Translation.Instance

func Translation_GetMessageList() {
	for _, key := range translation.GetMessageList() {
		var p = strings.Index(key, "\u0004")
		if p == -1 {
			var untranslated = key
			fmt.Printf("Message %s", untranslated)
		} else {
			var context = key[0:p]
			var untranslated = key[p+1:]
			fmt.Printf("Message %s with context %s", untranslated, context)
		}
	}
}
