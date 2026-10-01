package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/TranslationServer"
	"graphics.gd/startup"
)

// UIMirroring demonstrates right-to-left UI mirroring: swapping the locale
// between English and Arabic flips the layout direction of controls that
// derive their direction from the locale.
type UIMirroring struct {
	Control.Extension[UIMirroring] `gd:"UIMirroring"`

	Label Label.Instance
}

func (u *UIMirroring) Ready() {
	u.Label.SetText(TranslationServer.GetLocale())
}

func (u *UIMirroring) OnButtonPressed() {
	if TranslationServer.GetLocale() != "ar" {
		TranslationServer.SetLocale("ar")
	} else {
		TranslationServer.SetLocale("en")
	}
	u.Label.SetText(TranslationServer.GetLocale())
}

func main() {
	classdb.Register[UIMirroring]()
	startup.Scene()
}
