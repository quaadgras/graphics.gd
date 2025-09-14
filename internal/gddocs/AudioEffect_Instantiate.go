/*
extends AudioEffect

@export var strength = 4.0

func _instantiate():
    var effect = CustomAudioEffectInstance.new()
    effect.base = self

    return effect
*/

package main

import (
	"graphics.gd/classdb/AudioEffect"
	"graphics.gd/classdb/AudioEffectAmplify"
	"graphics.gd/variant/Object"
)

type MyAudioEffect struct {
	AudioEffect.Extension[MyAudioEffect]
}

func (e *MyAudioEffect) Instantiate() AudioEffect.Instance {
	effect := AudioEffectAmplify.New()
	Object.Set(effect, "base", e)
	return effect.AsAudioEffect()
}
