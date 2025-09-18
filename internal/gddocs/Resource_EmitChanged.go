/*
var damage:
	set(new_value):
		if damage != new_value:
			damage = new_value
			emit_changed()
*/

package main

import "graphics.gd/classdb/Resource"

type SomeResource struct {
	Resource.Extension[SomeResource]

	damage int
}

func (r *SomeResource) SetDamage(newValue int) {
	if r.damage != newValue {
		r.damage = newValue
		r.AsResource().EmitChanged()
	}
}
