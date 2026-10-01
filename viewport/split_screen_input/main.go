package main

import (
	"slices"
	"strconv"
	"strings"

	"graphics.gd/classdb"
	"graphics.gd/classdb/CharacterBody2D"
	"graphics.gd/classdb/Input"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/InputEventJoypadButton"
	"graphics.gd/classdb/InputEventKey"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/OptionButton"
	"graphics.gd/classdb/SubViewport"
	"graphics.gd/classdb/SubViewportContainer"
	"graphics.gd/classdb/VBoxContainer"
	"graphics.gd/classdb/Viewport"
	"graphics.gd/classdb/World2D"
	"graphics.gd/startup"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/Vector2"
)

// KeyboardOption is a named keyboard set for moving players around.
type KeyboardOption struct {
	Name string
	Keys []Input.Key
}

// KeyboardOptions are the 4 keyboard sets for moving players around.
var KeyboardOptions = []KeyboardOption{
	{"wasd", []Input.Key{Input.KeyW, Input.KeyA, Input.KeyS, Input.KeyD}},
	{"ijkl", []Input.Key{Input.KeyI, Input.KeyJ, Input.KeyK, Input.KeyL}},
	{"arrows", []Input.Key{Input.KeyLeft, Input.KeyRight, Input.KeyUp, Input.KeyDown}},
	{"numpad", []Input.Key{Input.KeyKp4, Input.KeyKp5, Input.KeyKp6, Input.KeyKp8}},
}

// PlayerColors are the modulate colors of each player.
var PlayerColors = []Color.RGBA{
	Color.X11.White,
	Color.Hex("ff8f02"),
	Color.Hex("05ff5a"),
	Color.Hex("ff05a0"),
}

// SplitConfig is the split screen configuration.
type SplitConfig struct {
	Keyboard []KeyboardOption
	Joypads  int
	World    World2D.Instance
	Position Vector2.XY
	Index    int
	Color    Color.RGBA
}

// SplitPlayer implements the player.
type SplitPlayer struct {
	CharacterBody2D.Extension[SplitPlayer] `gd:"SplitPlayer"`

	movement Vector2.XY // Current movement rate of node.
}

// Factor to multiply the movement.
const Factor = 200.0

// UnhandledInput updates the movement variable based on input that reaches this SubViewport.
func (p *SplitPlayer) UnhandledInput(event InputEvent.Instance) {
	switch {
	case event.IsActionPressed("ux_up") || event.IsActionReleased("ux_down"):
		p.movement.Y--
		Viewport.Get(p.AsNode()).SetInputAsHandled()
	case event.IsActionPressed("ux_down") || event.IsActionReleased("ux_up"):
		p.movement.Y++
		Viewport.Get(p.AsNode()).SetInputAsHandled()
	case event.IsActionPressed("ux_left") || event.IsActionReleased("ux_right"):
		p.movement.X--
		Viewport.Get(p.AsNode()).SetInputAsHandled()
	case event.IsActionPressed("ux_right") || event.IsActionReleased("ux_left"):
		p.movement.X++
		Viewport.Get(p.AsNode()).SetInputAsHandled()
	}
}

// PhysicsProcess moves the node based on the content of the movement variable.
func (p *SplitPlayer) PhysicsProcess(delta Float.X) {
	p.AsPhysicsBody2D().MoveAndCollide(Vector2.MulX(p.movement, Factor*delta))
}

// InputRoutingViewportContainer routes input for different SubViewports.
// Based on the provided input configuration, ensures only the correct
// events reach the SubViewport.
type InputRoutingViewportContainer struct {
	SubViewportContainer.Extension[InputRoutingViewportContainer] `gd:"InputRoutingViewportContainer"`

	currentKeyboardSet  []Input.Key // Currently used keyboard set.
	currentJoypadDevice int         // Currently used joypad device id.
}

func NewInputRoutingViewportContainer() *InputRoutingViewportContainer {
	return &InputRoutingViewportContainer{currentJoypadDevice: -1}
}

// PropagateInputEvent makes sure that only the events that are allowed
// via the OptionButton selection are sent to the SubViewport.
func (c *InputRoutingViewportContainer) PropagateInputEvent(event InputEvent.Instance) bool {
	if key, ok := Object.As[InputEventKey.Instance](event); ok {
		if slices.Contains(c.currentKeyboardSet, key.Keycode()) {
			return true
		}
	} else if button, ok := Object.As[InputEventJoypadButton.Instance](event); ok {
		if c.currentJoypadDevice > -1 && button.AsInputEvent().Device() == c.currentJoypadDevice {
			return true
		}
	}
	return false
}

// SetInputConfig sets a new config for input handling.
func (c *InputRoutingViewportContainer) SetInputConfig(keyboard []Input.Key, joypad int) {
	c.currentKeyboardSet = keyboard
	c.currentJoypadDevice = joypad
}

const JoypadPrefix = "Joypad"

// SplitScreen is the interface for a split screen.
type SplitScreen struct {
	VBoxContainer.Extension[SplitScreen] `gd:"SplitScreen"`

	InitPosition Vector2.XY

	Opt         OptionButton.Instance          `gd:"OptionButton"`
	InputRouter *InputRoutingViewportContainer `gd:"InputRoutingViewportContainer"`
	Viewport    SubViewport.Instance           `gd:"InputRoutingViewportContainer/SubViewport"`
	Play        *SplitPlayer                   `gd:"InputRoutingViewportContainer/SubViewport/Player"`

	keyboardOptions []KeyboardOption // Copy of all keyboard options.
}

// SetConfig sets the configuration of this split screen and performs OptionButton initialization.
func (s *SplitScreen) SetConfig(config SplitConfig) {
	s.keyboardOptions = config.Keyboard
	s.Play.AsNode2D().SetPosition(config.Position)
	s.Play.AsCanvasItem().SetModulate(config.Color)
	s.Opt.Clear()
	for _, option := range s.keyboardOptions {
		s.Opt.AddItem(option.Name)
	}
	for index := range config.Joypads {
		s.Opt.AddItem(JoypadPrefix + " " + strconv.Itoa(index+1))
	}
	s.Opt.Select(config.Index)
	s.OnOptionButtonItemSelected(config.Index)
	s.Viewport.AsViewport().SetWorld2d(config.World) // Connect all Split Screens to the same World2D.
}

// OnOptionButtonItemSelected updates keyboard settings after selecting them in the OptionButton.
func (s *SplitScreen) OnOptionButtonItemSelected(index int) {
	text := s.Opt.GetItemText(index)
	if strings.HasPrefix(text, JoypadPrefix) {
		device, _ := strconv.Atoi(text[len(text)-1:])
		s.InputRouter.SetInputConfig(nil, device)
	} else {
		for _, option := range s.keyboardOptions {
			if option.Name == text {
				s.InputRouter.SetInputConfig(option.Keys, -1)
				break
			}
		}
	}
}

// SplitScreenDemo sets up the different split screens, provides input
// configuration and connects the split screens to the play area.
type SplitScreenDemo struct {
	Node.Extension[SplitScreenDemo] `gd:"SplitScreenDemo"`

	PlayArea SubViewport.Instance `gd:"PlayArea"` // The central Viewport all Split Screens are sharing.
}

// Ready initializes each split screen and each player node.
func (d *SplitScreenDemo) Ready() {
	world := d.PlayArea.AsViewport().World2d()
	index := 0
	for _, child := range d.AsNode().GetChildren() {
		split, ok := Object.As[*SplitScreen](child)
		if !ok {
			continue
		}
		split.SetConfig(SplitConfig{
			Keyboard: KeyboardOptions,
			Joypads:  4,
			World:    world,
			Position: Vector2.Add(Vector2.MulX(Vector2.XY{X: Float.X(index % 2), Y: Float.X(index / 2)}, 132), Vector2.XY{X: 132, Y: 0}),
			Index:    index,
			Color:    PlayerColors[index],
		})
		index++
	}
}

func main() {
	classdb.Register[SplitPlayer]()
	classdb.Register[InputRoutingViewportContainer](NewInputRoutingViewportContainer)
	classdb.Register[SplitScreen]()
	classdb.Register[SplitScreenDemo]()
	startup.Scene()
}
