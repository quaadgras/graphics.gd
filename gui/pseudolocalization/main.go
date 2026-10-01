package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/CheckButton"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/LineEdit"
	"graphics.gd/classdb/ProjectSettings"
	"graphics.gd/classdb/SpinBox"
	"graphics.gd/classdb/TextEdit"
	"graphics.gd/classdb/TranslationServer"
	"graphics.gd/startup"
	"graphics.gd/variant/Float"
)

// Pseudolocalization lets you play with the engine's pseudolocalization
// settings at runtime and preview the result on arbitrary text.
type Pseudolocalization struct {
	Control.Extension[Pseudolocalization] `gd:"Pseudolocalization"`

	Toggle           CheckButton.Instance `gd:"Main/Pseudolocalization_options/toggle"`
	Accents          CheckButton.Instance `gd:"Main/Pseudolocalization_options/accents"`
	DoubleVowels     CheckButton.Instance `gd:"Main/Pseudolocalization_options/doublevowels"`
	FakeBidi         CheckButton.Instance `gd:"Main/Pseudolocalization_options/fakebidi"`
	Override         CheckButton.Instance `gd:"Main/Pseudolocalization_options/override"`
	SkipPlaceholders CheckButton.Instance `gd:"Main/Pseudolocalization_options/skipplaceholders"`
	ExpansionRatio   SpinBox.Instance     `gd:"Main/Pseudolocalization_options/exp_ratio/SpinBox"`
	Prefix           LineEdit.Instance    `gd:"Main/Pseudolocalization_options/prefix/TextEdit"`
	Suffix           LineEdit.Instance    `gd:"Main/Pseudolocalization_options/suffix/TextEdit"`
	Key              TextEdit.Instance    `gd:"Main/Pseudolocalizer/Key"`
	Result           TextEdit.Instance    `gd:"Main/Pseudolocalizer/Result"`
}

func boolSetting(name string) bool {
	value, _ := ProjectSettings.GetSetting(name, false).(bool)
	return value
}

func stringSetting(name string) string {
	value, _ := ProjectSettings.GetSetting(name, "").(string)
	return value
}

func floatSetting(name string) Float.X {
	switch value := ProjectSettings.GetSetting(name, 0.0).(type) {
	case int:
		return Float.X(value)
	case int64:
		return Float.X(value)
	case float64:
		return Float.X(value)
	}
	return 0
}

func (p *Pseudolocalization) Ready() {
	p.Accents.AsBaseButton().SetButtonPressed(boolSetting("internationalization/pseudolocalization/replace_with_accents"))
	p.Toggle.AsBaseButton().SetButtonPressed(TranslationServer.PseudolocalizationEnabled())
	p.FakeBidi.AsBaseButton().SetButtonPressed(boolSetting("internationalization/pseudolocalization/fake_bidi"))
	p.DoubleVowels.AsBaseButton().SetButtonPressed(boolSetting("internationalization/pseudolocalization/double_vowels"))
	p.Override.AsBaseButton().SetButtonPressed(boolSetting("internationalization/pseudolocalization/override"))
	p.SkipPlaceholders.AsBaseButton().SetButtonPressed(boolSetting("internationalization/pseudolocalization/skip_placeholders"))
	p.Prefix.SetText(stringSetting("internationalization/pseudolocalization/prefix"))
	p.Suffix.SetText(stringSetting("internationalization/pseudolocalization/suffix"))
	p.ExpansionRatio.AsRange().SetValue(floatSetting("internationalization/pseudolocalization/expansion_ratio"))
}

func (p *Pseudolocalization) OnAccentsToggled(buttonPressed bool) {
	ProjectSettings.SetSetting("internationalization/pseudolocalization/replace_with_accents", buttonPressed)
	TranslationServer.ReloadPseudolocalization()
}

func (p *Pseudolocalization) OnToggleToggled(buttonPressed bool) {
	TranslationServer.SetPseudolocalizationEnabled(buttonPressed)
}

func (p *Pseudolocalization) OnFakeBidiToggled(buttonPressed bool) {
	ProjectSettings.SetSetting("internationalization/pseudolocalization/fake_bidi", buttonPressed)
	TranslationServer.ReloadPseudolocalization()
}

func (p *Pseudolocalization) OnPrefixChanged(newText string) {
	ProjectSettings.SetSetting("internationalization/pseudolocalization/prefix", newText)
	TranslationServer.ReloadPseudolocalization()
}

func (p *Pseudolocalization) OnSuffixChanged(newText string) {
	ProjectSettings.SetSetting("internationalization/pseudolocalization/suffix", newText)
	TranslationServer.ReloadPseudolocalization()
}

func (p *Pseudolocalization) OnPseudolocalizePressed() {
	p.Result.SetText(TranslationServer.Pseudolocalize(p.Key.Text()))
}

func (p *Pseudolocalization) OnDoubleVowelsToggled(buttonPressed bool) {
	ProjectSettings.SetSetting("internationalization/pseudolocalization/double_vowels", buttonPressed)
	TranslationServer.ReloadPseudolocalization()
}

func (p *Pseudolocalization) OnExpansionRatioValueChanged(value Float.X) {
	ProjectSettings.SetSetting("internationalization/pseudolocalization/expansion_ratio", value)
	TranslationServer.ReloadPseudolocalization()
}

func (p *Pseudolocalization) OnOverrideToggled(buttonPressed bool) {
	ProjectSettings.SetSetting("internationalization/pseudolocalization/override", buttonPressed)
	TranslationServer.ReloadPseudolocalization()
}

func (p *Pseudolocalization) OnSkipPlaceholdersToggled(buttonPressed bool) {
	ProjectSettings.SetSetting("internationalization/pseudolocalization/skip_placeholders", buttonPressed)
	TranslationServer.ReloadPseudolocalization()
}

func main() {
	classdb.Register[Pseudolocalization]()
	startup.Scene()
}
