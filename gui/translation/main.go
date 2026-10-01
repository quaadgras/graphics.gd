// This is a demo of how a project can be internationalized.
//
// For more information on localization,
// search Godot's online documentation "Internationalization", or visit
// https://docs.godotengine.org/en/latest/tutorials/i18n/index.html
package main

import (
	"fmt"
	"strconv"
	"strings"

	"graphics.gd/classdb"
	"graphics.gd/classdb/AudioStreamPlayer2D"
	"graphics.gd/classdb/Engine"
	"graphics.gd/classdb/Panel"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/TranslationServer"
	"graphics.gd/startup"
	"graphics.gd/variant/Int"
)

func printLanguage() {
	Engine.PrintRich(fmt.Sprintf("\n[b]Language:[/b] %s (%s)",
		TranslationServer.GetLocaleName(TranslationServer.GetLocale()), TranslationServer.GetLocale()))
}

// TranslationDemoCSV demonstrates translations imported from a CSV file.
type TranslationDemoCSV struct {
	Panel.Extension[TranslationDemoCSV] `gd:"TranslationDemoCSV"`

	Audio AudioStreamPlayer2D.Instance
}

func (t *TranslationDemoCSV) Ready() { t.printIntro() }

func (t *TranslationDemoCSV) OnEnglishPressed() {
	TranslationServer.SetLocale("en")
	t.printIntro()
}

func (t *TranslationDemoCSV) OnSpanishPressed() {
	TranslationServer.SetLocale("es")
	t.printIntro()
}

func (t *TranslationDemoCSV) OnJapanesePressed() {
	TranslationServer.SetLocale("ja")
	t.printIntro()
}

func (t *TranslationDemoCSV) OnRussianPressed() {
	TranslationServer.SetLocale("ru")
	t.printIntro()
}

func (t *TranslationDemoCSV) OnPlayAudioPressed() { t.Audio.Play() }

func (t *TranslationDemoCSV) printIntro() {
	printLanguage()

	// In CSV translation, use the appropriate key to fetch the corresponding translation.
	// This is the same for scene nodes containing user-facing texts to be translated.
	fmt.Println(TranslationServer.Translate("KEY_INTRO", ""))

	// CSV does not support plural translations. If you need pluralization, you must use PO instead.
}

func (t *TranslationDemoCSV) OnGoToPoTranslationDemoPressed() {
	_ = SceneTree.Get(t.AsNode()).ChangeSceneToFile("res://translation_demo_po.tscn")
}

// TranslationDemoPO demonstrates translations loaded from gettext PO files.
//
// For more information on translation using PO files,
// search Godot's online documentation "Localization using gettext", or visit
// https://docs.godotengine.org/en/latest/tutorials/i18n/localization_using_gettext.html
type TranslationDemoPO struct {
	Panel.Extension[TranslationDemoPO] `gd:"TranslationDemoPO"`

	Audio AudioStreamPlayer2D.Instance
}

func (t *TranslationDemoPO) Ready() { t.printIntro() }

func (t *TranslationDemoPO) OnEnglishPressed() {
	TranslationServer.SetLocale("en")
	t.printIntro()
}

func (t *TranslationDemoPO) OnSpanishPressed() {
	TranslationServer.SetLocale("es")
	t.printIntro()
}

func (t *TranslationDemoPO) OnJapanesePressed() {
	TranslationServer.SetLocale("ja")
	t.printIntro()
}

func (t *TranslationDemoPO) OnPlayAudioPressed() { t.Audio.Play() }

func (t *TranslationDemoPO) printIntro() {
	printLanguage()

	// In PO translation, you would use the source string as the 'key' to translate.
	// This is the same for scene nodes containing user-facing texts to be translated.
	fmt.Println(TranslationServer.Translate("Hello, this is a translation demo project.", ""))

	// PO plural translation example.
	// The difference with CSV is that you must add the "plural_message" argument, because PO files
	// expect the data (else undefined behaviour might occur).
	daysPassed := Int.RandomBetween(1, 3)
	fmt.Println(strings.ReplaceAll(
		TranslationServer.TranslatePlural("One day ago.", "{days} days ago.", daysPassed, ""),
		"{days}", strconv.Itoa(daysPassed)))
}

func (t *TranslationDemoPO) OnGoToCsvTranslationDemoPressed() {
	_ = SceneTree.Get(t.AsNode()).ChangeSceneToFile("res://translation_demo_csv.tscn")
}

func main() {
	classdb.Register[TranslationDemoCSV]()
	classdb.Register[TranslationDemoPO]()
	startup.Scene()
}
