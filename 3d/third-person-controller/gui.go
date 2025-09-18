package main

import (
	"fmt"

	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/GridContainer"
	"graphics.gd/classdb/HBoxContainer"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/OS"
	"graphics.gd/classdb/PanelContainer"
	"graphics.gd/classdb/PropertyTweener"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/TextureButton"
	"graphics.gd/classdb/TextureRect"
	"graphics.gd/classdb/Timer"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Enum"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
)

type InstructionType Enum.Int[struct {
	Keyboard InstructionType
	Joypad   InstructionType
}]

var InstructionTypes = Enum.Values[InstructionType]()

type DemoPage struct {
	Node.Extension[DemoPage] `gd:"DemoPage"`

	DemoPageRoot          Control.Instance       `gd:"CanvasLayer/DemoPageRoot"`
	ResumeButton          Button.Instance        `gd:"CanvasLayer/DemoPageRoot/Content/MarginContainer/Buttons/Resume"`
	ExitButton            Button.Instance        `gd:"CanvasLayer/DemoPageRoot/Content/MarginContainer/Buttons/Exit"`
	KeyboardButton        Button.Instance        `gd:"%KeyboardButton"`
	JoypadButton          Button.Instance        `gd:"%JoypadButton"`
	GridContainerKeyboard GridContainer.Instance `gd:"%GridContainerKeyboard"`
	GridContainerJoypad   GridContainer.Instance `gd:"%GridContainerJoypad"`

	demoMouseMode Input.MouseModeValue
}

func (page *DemoPage) Ready() {
	var tree = SceneTree.Get(page.AsNode())
	tree.SetPaused(true)
	page.demoMouseMode = Input.MouseMode()
	Input.SetMouseMode(Input.MouseModeVisible)
	page.ResumeButton.AsBaseButton().OnPressed(page.resume_demo)
	page.ExitButton.AsBaseButton().OnPressed(func() {
		SceneTree.Get(page.AsNode()).Quit()
	})
	page.KeyboardButton.AsBaseButton().OnPressed(func() {
		page.change_instruction(InstructionTypes.Keyboard)
	})
	page.JoypadButton.AsBaseButton().OnPressed(func() {
		page.change_instruction(InstructionTypes.Joypad)
	})
	if len(Input.GetConnectedJoypads()) > 0 {
		page.change_instruction(InstructionTypes.Joypad)
	} else {
		page.change_instruction(InstructionTypes.Keyboard)
	}
}

func (page *DemoPage) resume_demo() {
	SceneTree.Get(page.AsNode()).SetPaused(false)
	var tween = page.AsNode().CreateTween()
	PropertyTweener.Make(tween, page.DemoPageRoot.AsObject(), "modulate", Color.Transparent, 0.3)
	tween.TweenCallback(page.DemoPageRoot.AsCanvasItem().Hide)
	Input.SetMouseMode(page.demoMouseMode)
}

func (page *DemoPage) change_instruction(itype InstructionType) {
	withAlpha := func(alpha Float.X, c Color.RGBA) Color.RGBA {
		c.A = alpha
		return c
	}
	switch itype {
	case InstructionTypes.Keyboard:
		page.KeyboardButton.AsCanvasItem().SetModulate(withAlpha(1, page.KeyboardButton.AsCanvasItem().Modulate()))
		page.JoypadButton.AsCanvasItem().SetModulate(withAlpha(0.3, page.JoypadButton.AsCanvasItem().Modulate()))
		page.GridContainerKeyboard.AsCanvasItem().Show()
		page.GridContainerJoypad.AsCanvasItem().Hide()
	case InstructionTypes.Joypad:
		page.KeyboardButton.AsCanvasItem().SetModulate(withAlpha(0.3, page.KeyboardButton.AsCanvasItem().Modulate()))
		page.JoypadButton.AsCanvasItem().SetModulate(withAlpha(1, page.JoypadButton.AsCanvasItem().Modulate()))
		page.GridContainerKeyboard.AsCanvasItem().Hide()
		page.GridContainerJoypad.AsCanvasItem().Show()
	}
	page.KeyboardButton.AsControl().ReleaseFocus()
	page.JoypadButton.AsControl().ReleaseFocus()
}

func (page *DemoPage) Input(event InputEvent.Instance) {
	if event.IsActionPressed("pause") && !event.IsEcho() {
		if SceneTree.Get(page.AsNode()).Paused() {
			page.resume_demo()
		} else {
			page.pause_demo()
		}
	}
}

func (page *DemoPage) pause_demo() {
	page.demoMouseMode = Input.MouseMode()
	SceneTree.Get(page.AsNode()).SetPaused(true)
	page.DemoPageRoot.AsCanvasItem().Show()
	var tween = page.AsNode().CreateTween()
	PropertyTweener.Make(tween, page.DemoPageRoot.AsObject(), "modulate", Color.X11.White, 0.3)
	Input.SetMouseMode(Input.MouseModeVisible)
}

type DemoLinkButton struct {
	TextureButton.Extension[DemoLinkButton] `gd:"DemoLinkButton"`

	Link string
}

func (button *DemoLinkButton) Ready() {
	button.AsBaseButton().OnPressed(func() {
		OS.ShellOpen(button.Link)
	})
}

type Icone struct {
	TextureRect.Extension[Icone] `gd:"Icone"`

	DisabledAlpha Float.X
}

func NewIcone() *Icone {
	return &Icone{DisabledAlpha: 0.2}
}

func (icon *Icone) Ready() {
	modulate := icon.AsCanvasItem().Modulate()
	modulate.A = icon.DisabledAlpha
	icon.AsCanvasItem().SetModulate(modulate)
}

func (icon *Icone) SetState(state bool) {
	var from, to = Color.RGBA{1, 1, 1, icon.DisabledAlpha}, Color.W3C.White
	if state {
		from, to = to, from
	}
	var tween = icon.AsNode().CreateTween()
	PropertyTweener.Make(tween, icon.AsObject(), "modulate", to, 0.3).From(from)
}

type WeaponUI struct {
	PanelContainer.Extension[WeaponUI] `gd:"WeaponUI"`

	Nodes        map[string]*Icone
	SelectedNode string
}

func (ui *WeaponUI) Ready() {
	ui.Nodes = map[string]*Icone{
		"DEFAULT": Object.To[*Icone](ui.AsNode().GetNode("%Flash")),
		"GRENADE": Object.To[*Icone](ui.AsNode().GetNode("%Grenade")),
	}
}

func (ui *WeaponUI) SwitchTo(node_name string) {
	if node_name == ui.SelectedNode {
		return
	}
	if ui.SelectedNode != "" {
		ui.Nodes[ui.SelectedNode].SetState(false)
	}
	ui.Nodes[node_name].SetState(true)
	ui.SelectedNode = node_name
}

type CoinsContainer struct {
	HBoxContainer.Extension[CoinsContainer] `gd:"CoinsContainer"`

	DisplayTimer Timer.Instance `gd:"Timer"`
	CoinsLabel   Label.Instance `gd:"CoinsLabel"`
}

func (container *CoinsContainer) Ready() {
	const Hidden_Y_Pos = -100.0
	container.DisplayTimer.AsTimer().OnTimeout(func() {
		var tween = container.AsNode().CreateTween()
		PropertyTweener.Make(tween, container.AsObject(), "position:y", Hidden_Y_Pos, 0.5)
	})
}

func (container *CoinsContainer) UpdateCoinsAmount(amount int) {
	const Display_Y_Pos = 20.0
	if container.DisplayTimer.IsStopped() {
		var tween = container.AsNode().CreateTween()
		PropertyTweener.Make(tween, container.AsObject(), "position:y", Display_Y_Pos, 0.5)
	}
	container.DisplayTimer.Start()
	container.CoinsLabel.SetText(fmt.Sprintf("%d", amount))
}
