/*
@export var energy = 0:
	set(value):
		energy = value
		update_configuration_warnings()

func _get_configuration_warnings():
	if energy < 0:
		return ["Energy must be 0 or greater."]
	else:
		return []
*/

package main

import (
	"graphics.gd/classdb/Node"
	"graphics.gd/variant/Float"
)

type ConfigurableNode struct {
	Node.Extension[ConfigurableNode]

	energy Float.X
}

func (cn ConfigurableNode) SetEnergy(energy Float.X) {
	cn.energy = energy
	cn.AsNode().UpdateConfigurationWarnings()
}

func (cn ConfigurableNode) GetConfigurationWarnings() []string {
	if cn.energy < 0 {
		return []string{"Energy must be 0 or greater."}
	}
	return nil
}
