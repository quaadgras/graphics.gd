// This is a demo of UI accessibility in Godot.
//
// Demo shows basic UI accessibility features and making custom nodes accessible.
package main

import (
	"strconv"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/GUI"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/LineEdit"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/TextServer"
	"graphics.gd/startup"
	"graphics.gd/variant/Color"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Int"
	"graphics.gd/variant/Object"
	"graphics.gd/variant/RID"
	"graphics.gd/variant/Rect2"
	"graphics.gd/variant/Vector2"
)

// Accessibility is the root control of the demo.
type Accessibility struct {
	Control.Extension[Accessibility] `gd:"Accessibility"`

	LineEditName    LineEdit.Instance
	LineEditLiveReg LineEdit.Instance
	LabelRegion     Label.Instance `gd:"Panel/LabelRegion"`
}

func (a *Accessibility) Ready() {
	// Accessible UI should always have keyboard focus, since it is a main way of interacting with UI.
	a.LineEditName.AsControl().GrabFocus()
}

func (a *Accessibility) OnButtonSetPressed() {
	a.LabelRegion.SetText(a.LineEditLiveReg.Text()) // Set live region text.
}

// CustomControl shows how accessibility support for a node can be fully customized by
// handling the ACCESSIBILITY_* notifications and associated virtual methods.
type CustomControl struct {
	Control.Extension[CustomControl] `gd:"CustomControl"`

	itemAes    [3]RID.AccessibilityElement
	itemNames  [3]string
	itemValues [3]int
	itemRects  [3]Rect2.PositionSize
	selected   int
}

func NewCustomControl() *CustomControl {
	return &CustomControl{
		itemNames: [3]string{"Item 1", "Item 2", "Item 3"},
		itemRects: [3]Rect2.PositionSize{
			Rect2.New(0, 0, 40, 40),
			Rect2.New(40, 0, 40, 40),
			Rect2.New(80, 0, 40, 40),
		},
	}
}

// Input:

func (c *CustomControl) GuiInput(inputEvent InputEvent.Instance) {
	if inputEvent.IsActionPressed("ui_left") {
		c.selected = (c.selected + len(c.itemAes) - 1) % len(c.itemAes)
		c.refresh()
		c.AsControl().AcceptEvent()
	}
	if inputEvent.IsActionPressed("ui_right") {
		c.selected = (c.selected + 1) % len(c.itemAes)
		c.refresh()
		c.AsControl().AcceptEvent()
	}
	if inputEvent.IsActionPressed("ui_up") {
		c.itemValues[c.selected] = Int.Clamp(c.itemValues[c.selected]-1, -100, 100)
		c.refresh()
		c.AsControl().AcceptEvent()
	}
	if inputEvent.IsActionPressed("ui_down") {
		c.itemValues[c.selected] = Int.Clamp(c.itemValues[c.selected]+1, -100, 100)
		c.refresh()
		c.AsControl().AcceptEvent()
	}
}

func (c *CustomControl) refresh() {
	c.AsCanvasItem().QueueRedraw()
	// Request node accessibility information update. Similar to "queue_redraw" for drawing.
	c.AsNode().QueueAccessibilityUpdate()
}

// Accessibility actions and focus callback:

func (c *CustomControl) accessibilityActionAdjust(item, delta int) {
	// Numeric value increment/decrement.
	c.itemValues[item] = Int.Clamp(c.itemValues[item]+delta, -100, 100)
	c.refresh()
}

func (c *CustomControl) accessibilityActionSetNumValue(data any, item int) {
	// Numeric value set, "data" is a new value.
	switch value := data.(type) {
	case int:
		c.itemValues[item] = Int.Clamp(value, -100, 100)
	case int64:
		c.itemValues[item] = Int.Clamp(int(value), -100, 100)
	case float64:
		c.itemValues[item] = Int.Clamp(int(value), -100, 100)
	}
	c.refresh()
}

// GetFocusedAccessibilityElement returns the focused sub-element, if no sub-element is
// focused, it should return the base element (the value returned by GetAccessibilityElement)
// instead.
func (c *CustomControl) GetFocusedAccessibilityElement() RID.AccessibilityElement {
	return c.itemAes[c.selected]
}

// Notifications handler:

func (c *CustomControl) Notification(what Object.Notification) {
	if what == Node.NotificationAccessibilityInvalidate {
		// Accessibility cleanup:
		//
		// Called when the existing main element is destroyed.

		// Note: since item sub-elements are children of the main element, there's no need to
		// destroy them manually. But we should keep track when handles are invalidated.
		for i := range c.itemAes {
			c.itemAes[i] = RID.AccessibilityElement(0)
		}
	}

	if what == Node.NotificationAccessibilityUpdate {
		// Accessibility update handler:
		//
		// This function acts as an alternative "draw" for the screen reader, and provides
		// information about this node.

		// Get handle to the accessibility element, the accessibility element is created and
		// destroyed automatically.
		ae := c.AsNode().GetAccessibilityElement()

		// Set role of the element.
		DisplayServer.AccessibilityUpdateSetRole(ae, DisplayServer.RoleListBox)

		// Set other properties.
		DisplayServer.AccessibilityUpdateSetListItemCount(ae, len(c.itemAes))
		DisplayServer.AccessibilityUpdateSetName(ae, "List")

		for i := range c.itemAes {
			// Create a new sub-element for the item if it doesn't exist.
			if c.itemAes[i] == RID.AccessibilityElement(0) {
				c.itemAes[i] = DisplayServer.AccessibilityCreateSubElement(ae, DisplayServer.RoleListBoxOption)
			}

			// Sub-element properties.
			DisplayServer.AccessibilityUpdateSetListItemIndex(c.itemAes[i], i)
			DisplayServer.AccessibilityUpdateSetListItemSelected(c.itemAes[i], c.selected == i)
			DisplayServer.AccessibilityUpdateSetName(c.itemAes[i], c.itemNames[i])                 // Readable name.
			DisplayServer.AccessibilityUpdateSetValue(c.itemAes[i], strconv.Itoa(c.itemValues[i])) // Readable value.

			// Numeric value info for the actions.
			DisplayServer.AccessibilityUpdateSetNumValue(c.itemAes[i], Float.X(c.itemValues[i]))
			DisplayServer.AccessibilityUpdateSetNumRange(c.itemAes[i], -100.0, 100.0)
			DisplayServer.AccessibilityUpdateSetNumStep(c.itemAes[i], 1.0)

			// Sub-element bounding box, relative to the parent element.
			DisplayServer.AccessibilityUpdateSetBounds(c.itemAes[i], c.itemRects[i])

			// Set supported actions for the item, actions can be invoked directly by the
			// screen-reader (e.g, via global keyboard shortcuts).
			item := i
			DisplayServer.AccessibilityUpdateAddAction(c.itemAes[i], DisplayServer.ActionDecrement, func(_ any) {
				c.accessibilityActionAdjust(item, -1)
			})
			DisplayServer.AccessibilityUpdateAddAction(c.itemAes[i], DisplayServer.ActionIncrement, func(_ any) {
				c.accessibilityActionAdjust(item, +1)
			})
			DisplayServer.AccessibilityUpdateAddAction(c.itemAes[i], DisplayServer.ActionSetValue, func(data any) {
				c.accessibilityActionSetNumValue(data, item)
			})
		}
	}

	if what == Control.NotificationFocusEnter || what == Control.NotificationFocusExit {
		c.AsCanvasItem().QueueRedraw()
	}
}

// Draw:

func (c *CustomControl) Draw() {
	// Draw, provided for convenience and NOT required for screen-reader support.
	canvas := c.AsCanvasItem()
	font := c.AsControl().GetThemeFont("font")
	for i := range c.itemAes {
		canvas.MoreArgs().DrawRect(c.itemRects[c.selected], Color.RGBA{R: 0.8, G: 0.8, B: 0.8, A: 0.5}, false, 1.0, false)
		canvas.MoreArgs().DrawString(font,
			Vector2.Add(c.itemRects[i].Position, Vector2.XY{X: 0, Y: 30}),
			strconv.Itoa(c.itemValues[i]),
			GUI.HorizontalAlignmentCenter, 40.0, 16, Color.X11.White,
			TextServer.JustificationKashida|TextServer.JustificationWordBound, 0, 0, 0.0)
	}

	if c.AsControl().HasFocus() {
		size := c.AsControl().Size()
		canvas.MoreArgs().DrawRect(Rect2.New(0, 0, size.X, size.Y), Color.RGBA{R: 0, G: 0, B: 1, A: 0.5}, false, 3.0, false)
		canvas.MoreArgs().DrawRect(c.itemRects[c.selected], Color.RGBA{R: 0, G: 1, B: 0, A: 0.5}, false, 2.0, false)
	}
}

func main() {
	classdb.Register[Accessibility]()
	classdb.Register[CustomControl](NewCustomControl)
	startup.Scene()
}
