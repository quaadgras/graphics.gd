package main

import (
	"fmt"
	"strings"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/ColorRect"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/HSlider"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/LineEdit"
	"graphics.gd/classdb/OS"
	"graphics.gd/classdb/RichTextLabel"
	"graphics.gd/classdb/TextEdit"
	"graphics.gd/classdb/Tree"
	"graphics.gd/classdb/TreeItem"
	"graphics.gd/startup"
	"graphics.gd/variant/Callable"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
)

type VoiceList struct {
	Control.Extension[VoiceList] `gd:"VoiceList"`

	Tree          Tree.Instance
	Log           TextEdit.Instance
	RichTextLabel RichTextLabel.Instance
	Utterance     TextEdit.Instance
	ButtonPause   Button.Instance
	ColorRect     ColorRect.Instance

	HSliderRate   HSlider.Instance
	HSliderPitch  HSlider.Instance
	HSliderVolume HSlider.Instance
	RateValue     Label.Instance `gd:"HSliderRate/Value"`
	PitchValue    Label.Instance `gd:"HSliderPitch/Value"`
	VolumeValue   Label.Instance `gd:"HSliderVolume/Value"`

	LineEditFilterName LineEdit.Instance
	LineEditFilterLang LineEdit.Instance

	id    int // The utterance ID to use for text to speech.
	utMap map[int]string
	vs    []DisplayServer.TextToSpeechVoice
}

func (v *VoiceList) Ready() {
	v.utMap = map[int]string{}

	// Get voice data.
	v.vs = DisplayServer.TtsGetVoices()
	root := v.Tree.CreateItem()
	v.Tree.SetHideRoot(true)
	v.Tree.SetColumnTitle(0, "Name")
	v.Tree.SetColumnTitle(1, "Language")
	v.Tree.SetColumnTitlesVisible(true)
	for _, voice := range v.vs {
		child := v.Tree.MoreArgs().CreateItem(root, -1)
		child.SetText(0, voice.Name)
		child.SetMetadata(0, voice.ID)
		child.SetText(1, voice.Language)
	}
	v.appendLog(fmt.Sprintf("%d voices available.\n", len(v.vs)))
	v.appendLog("=======\n")

	// Ensure the first voice added to the list is preselected.
	if first := v.Tree.GetRoot().GetChild(0); first != TreeItem.Nil {
		first.Select(0)
	}

	// Add callbacks. The started/ended/canceled callbacks take a single int
	// (the utterance ID), so they are registered through the Advanced API
	// rather than the two-int convenience wrapper.
	DisplayServer.Advanced().TtsSetUtteranceCallback(DisplayServer.TtsUtteranceStarted, Callable.New(v.onUtteranceStart))
	DisplayServer.Advanced().TtsSetUtteranceCallback(DisplayServer.TtsUtteranceEnded, Callable.New(v.onUtteranceEnd))
	DisplayServer.Advanced().TtsSetUtteranceCallback(DisplayServer.TtsUtteranceCanceled, Callable.New(v.onUtteranceError))
	DisplayServer.Advanced().TtsSetUtteranceCallback(DisplayServer.TtsUtteranceBoundary, Callable.New(v.onUtteranceBoundary))
}

func (v *VoiceList) Process(_ Float.X) {
	v.ButtonPause.AsBaseButton().SetPressedNoSignal(DisplayServer.TtsIsPaused())
	if DisplayServer.TtsIsSpeaking() {
		v.ColorRect.SetColor(Color.RGBA{0.9, 0.3, 0.1, 1})
	} else {
		v.ColorRect.SetColor(Color.RGBA{1, 1, 1, 1})
	}
}

func (v *VoiceList) appendLog(text string) {
	v.Log.SetText(v.Log.Text() + text)
}

func (v *VoiceList) onUtteranceBoundary(pos, utID int) {
	text := []rune(v.utMap[utID])
	pos = min(pos, len(text))
	v.RichTextLabel.SetText("[bgcolor=yellow][color=black]" + string(text[:pos]) + "[/color][/bgcolor]" + string(text[pos:]))
}

func (v *VoiceList) onUtteranceStart(utID int) {
	v.appendLog(fmt.Sprintf("Utterance %d started.\n", utID))
}

func (v *VoiceList) onUtteranceEnd(utID int) {
	v.RichTextLabel.SetText("[bgcolor=yellow][color=black]" + v.utMap[utID] + "[/color][/bgcolor]")
	v.appendLog(fmt.Sprintf("Utterance %d ended.\n", utID))
	delete(v.utMap, utID)
}

func (v *VoiceList) onUtteranceError(utID int) {
	v.RichTextLabel.SetText("")
	v.appendLog(fmt.Sprintf("Utterance %d canceled/failed.\n", utID))
	delete(v.utMap, utID)
}

func (v *VoiceList) OnStopPressed() {
	DisplayServer.TtsStop()
}

func (v *VoiceList) OnPausePressed() {
	if v.ButtonPause.AsBaseButton().ButtonPressed() {
		DisplayServer.TtsPause()
	} else {
		DisplayServer.TtsResume()
	}
}

func (v *VoiceList) speak(action string, interrupt bool) {
	selected := v.Tree.GetSelected()
	if selected == TreeItem.Nil {
		OS.Alert("No voice selected.\nSelect a voice in the list, then try using Speak again.")
		return
	}
	v.appendLog(fmt.Sprintf("Utterance %d %s.\n", v.id, action))
	text := v.Utterance.Text()
	v.utMap[v.id] = text
	voice, _ := selected.GetMetadata(0).(string)
	DisplayServer.TtsSpeakOptions(text, voice,
		int(v.HSliderVolume.AsRange().Value()),
		v.HSliderPitch.AsRange().Value(),
		v.HSliderRate.AsRange().Value(),
		v.id, interrupt)
	v.id++
}

func (v *VoiceList) OnSpeakPressed()    { v.speak("queried", false) }
func (v *VoiceList) OnIntSpeakPressed() { v.speak("interrupted", true) }

func (v *VoiceList) OnClearLogPressed() {
	v.Log.SetText("")
}

func (v *VoiceList) OnRateChanged(value Float.X) {
	v.RateValue.SetText(fmt.Sprintf("%.2fx", value))
}

func (v *VoiceList) OnPitchChanged(value Float.X) {
	v.PitchValue.SetText(fmt.Sprintf("%.2fx", value))
}

func (v *VoiceList) OnVolumeChanged(value Float.X) {
	v.VolumeValue.SetText(fmt.Sprintf("%d%%", int(value)))
}

func (v *VoiceList) OnDemoPressed() {
	v.demo("en", "Beware the Jabberwock, my son!", "The jaws that bite, the claws that catch!")
	v.demo("es", "¡Cuidado, hijo, con el Fablistanón!", "¡Con sus dientes y garras, muerde, apresa!")
	v.demo("ru", "О, бойся Бармаглота, сын!", "Он так свирлеп и дик!")
}

func (v *VoiceList) demo(language, first, second string) {
	vc := DisplayServer.TtsGetVoicesForLanguage(language)
	if len(vc) == 0 {
		return
	}
	volume := int(v.HSliderVolume.AsRange().Value())
	pitch := v.HSliderPitch.AsRange().Value()
	rate := v.HSliderRate.AsRange().Value()
	v.utMap[v.id] = first
	v.utMap[v.id+1] = second
	DisplayServer.TtsSpeakOptions(first, vc[0], volume, pitch, rate, v.id, false)
	DisplayServer.TtsSpeakOptions(second, vc[0], volume, pitch, rate, v.id+1, false)
	v.id += 2
}

func (v *VoiceList) OnFilterChanged(_ string) {
	v.Tree.Clear()
	root := v.Tree.CreateItem()
	nameFilter := strings.ToLower(v.LineEditFilterName.Text())
	langFilter := strings.ToLower(v.LineEditFilterLang.Text())
	for _, voice := range v.vs {
		if (nameFilter == "" || strings.Contains(strings.ToLower(voice.Name), nameFilter)) &&
			(langFilter == "" || strings.Contains(strings.ToLower(voice.Language), langFilter)) {
			child := v.Tree.MoreArgs().CreateItem(root, -1)
			child.SetText(0, voice.Name)
			child.SetMetadata(0, voice.ID)
			child.SetText(1, voice.Language)
		}
	}
}

func main() {
	classdb.Register[VoiceList]()
	startup.Scene()
}
