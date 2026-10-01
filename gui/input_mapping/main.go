// A demo showing how to build an input key remapping screen.
//
//   - Click the buttons to change the bound keys.
//   - Persists the keys to disk, so they are preserved
//     after the project is restarted.
package main

import (
	"encoding/json"
	"fmt"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/FileAccess"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventKey"
	"graphics.gd/classdb/InputMap"
	"graphics.gd/startup"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Object"
)

// ActionRemapButton rebinds its action to the next key pressed while it is toggled on.
type ActionRemapButton struct {
	Button.Extension[ActionRemapButton] `gd:"ActionRemapButton"`

	Action string
}

func NewActionRemapButton() *ActionRemapButton {
	return &ActionRemapButton{Action: "ui_up"}
}

func (b *ActionRemapButton) Ready() {
	KeyPersistence.init()
	if !InputMap.HasAction(b.Action) {
		fmt.Println("input_mapping: unknown action:", b.Action)
		return
	}
	b.AsNode().SetProcessUnhandledKeyInput(false)
	b.displayCurrentKey()
}

func (b *ActionRemapButton) Toggled(isButtonPressed bool) {
	b.AsNode().SetProcessUnhandledKeyInput(isButtonPressed)
	if isButtonPressed {
		b.AsButton().SetText("<press a key>")
		b.AsCanvasItem().SetModulate(Color.X11.Yellow)
		b.AsControl().ReleaseFocus()
	} else {
		b.displayCurrentKey()
		b.AsCanvasItem().SetModulate(Color.X11.White)
		// Grab focus after assigning a key, so that you can go to the next
		// key using the keyboard.
		b.AsControl().GrabFocus()
	}
}

// NOTE: You can use the `Input()` callback instead, especially if
// you want to work with gamepads.
func (b *ActionRemapButton) UnhandledKeyInput(inputEvent InputEvent.Instance) {
	// Skip if pressing Enter, so that the input mapping GUI can be navigated
	// with the keyboard. The downside of this approach is that the Enter
	// key can't be bound to actions.
	if key, ok := Object.As[InputEventKey.Instance](inputEvent); ok && key.Keycode() != Input.KeyEnter {
		b.remapActionTo(key)
		b.AsBaseButton().SetButtonPressed(false)
	}
}

func (b *ActionRemapButton) remapActionTo(inputEvent InputEventKey.Instance) {
	// We first change the event in this game instance.
	InputMap.ActionEraseEvents(b.Action)
	InputMap.ActionAddEvent(b.Action, inputEvent.AsInputEvent())
	// And then save it to the keymaps file.
	KeyPersistence.keymaps[b.Action] = keymap{
		Keycode:  int(inputEvent.Keycode()),
		Physical: int(inputEvent.PhysicalKeycode()),
		Unicode:  inputEvent.Unicode(),
	}
	KeyPersistence.save()
	b.AsButton().SetText(inputEvent.AsInputEvent().AsText())
}

func (b *ActionRemapButton) displayCurrentKey() {
	if events := InputMap.ActionGetEvents(b.Action); len(events) > 0 {
		b.AsButton().SetText(events[0].AsText())
	}
}

// KeyPersistence saves the key maps in a simple way. It replaces the GDScript
// autoload singleton of the same name: instead of serializing InputEvent objects
// with store_var, it stores the key event's codes as JSON and rebuilds the
// InputEventKey on load.
var KeyPersistence = &keyPersistence{keymaps: map[string]keymap{}}

const keymapsPath = "user://keymaps.dat"

type keymap struct {
	Keycode  int `json:"keycode"`
	Physical int `json:"physical_keycode"`
	Unicode  int `json:"unicode"`
}

type keyPersistence struct {
	loaded  bool
	keymaps map[string]keymap
}

func (k *keyPersistence) init() {
	if k.loaded {
		return
	}
	k.loaded = true
	// First we create the keymap dictionary on startup with all
	// the keymap actions we have.
	for _, action := range InputMap.GetActions() {
		events := InputMap.ActionGetEvents(action)
		if len(events) == 0 {
			continue
		}
		if key, ok := Object.As[InputEventKey.Instance](events[0]); ok {
			k.keymaps[action] = keymap{
				Keycode:  int(key.Keycode()),
				Physical: int(key.PhysicalKeycode()),
				Unicode:  key.Unicode(),
			}
		}
	}
	k.load()
}

func (k *keyPersistence) load() {
	if !FileAccess.FileExists(keymapsPath) {
		// There is no save file yet, so let's create one.
		k.save()
		return
	}
	file := FileAccess.Open(keymapsPath, FileAccess.Read)
	if file == FileAccess.Nil {
		return
	}
	var stored map[string]keymap
	if err := json.Unmarshal([]byte(file.GetAsText()), &stored); err != nil {
		fmt.Println("input_mapping: failed to load keymaps:", err)
		return
	}
	// We don't just replace the keymaps dictionary, because if you
	// updated your game and removed/added keymaps, the data of this
	// save file may have invalid actions. So we check one by one to
	// make sure that the keymap dictionary really has all current actions.
	for action := range k.keymaps {
		saved, ok := stored[action]
		if !ok {
			continue
		}
		k.keymaps[action] = saved
		// Whilst setting the keymap dictionary, we also set the
		// correct InputMap event.
		event := InputEventKey.New()
		event.SetKeycode(Input.Key(saved.Keycode))
		event.SetPhysicalKeycode(Input.Key(saved.Physical))
		event.SetUnicode(saved.Unicode)
		InputMap.ActionEraseEvents(action)
		InputMap.ActionAddEvent(action, event.AsInputEvent())
	}
}

func (k *keyPersistence) save() {
	data, err := json.Marshal(k.keymaps)
	if err != nil {
		fmt.Println("input_mapping: failed to save keymaps:", err)
		return
	}
	file := FileAccess.Open(keymapsPath, FileAccess.Write)
	if file == FileAccess.Nil {
		return
	}
	file.StoreString(string(data))
}

func main() {
	classdb.Register[ActionRemapButton](NewActionRemapButton)
	startup.Scene()
}
