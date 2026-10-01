package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/DisplayServer"
	"graphics.gd/classdb/InputEvent"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/SubViewport"
	"graphics.gd/startup"
)

type Game struct {
	Node.Extension[Game]

	PauseMenu *PauseMenu `gd:"InterfaceLayer/PauseMenu"`
}

func (g *Game) UnhandledInput(event InputEvent.Instance) {
	g.handleInput(g.AsNode(), g.PauseMenu, event)
}

func (g *Game) handleInput(node Node.Instance, pauseMenu *PauseMenu, event InputEvent.Instance) {
	tree := SceneTree.Get(node)
	if event.IsActionPressed("toggle_fullscreen") {
		mode := DisplayServer.WindowGetMode(DisplayServer.MainWindowId)
		if mode == DisplayServer.WindowModeFullscreen || mode == DisplayServer.WindowModeExclusiveFullscreen {
			DisplayServer.WindowSetMode(DisplayServer.WindowModeWindowed, DisplayServer.MainWindowId)
		} else {
			DisplayServer.WindowSetMode(DisplayServer.WindowModeFullscreen, DisplayServer.MainWindowId)
		}
		tree.Root().AsViewport().SetInputAsHandled()
	} else if event.IsActionPressed("toggle_pause") {
		tree.SetPaused(!tree.Paused())
		if tree.Paused() {
			pauseMenu.Open()
		} else {
			pauseMenu.Close()
		}
		tree.Root().AsViewport().SetInputAsHandled()
	}
}

// GameSplitscreen renders two players into two SubViewports sharing one world.
type GameSplitscreen struct {
	Node.Extension[GameSplitscreen]

	PauseMenu *PauseMenu           `gd:"InterfaceLayer/PauseMenu"`
	Player2   *Player              `gd:"%Player2"`
	Viewport1 SubViewport.Instance `gd:"%Viewport1"`
	Viewport2 SubViewport.Instance `gd:"%Viewport2"`
}

func (g *GameSplitscreen) Ready() {
	g.Viewport2.AsViewport().SetWorld2d(g.Viewport1.AsViewport().World2d())
	g.Player2.Camera.SetCustomViewport(g.Viewport2.AsNode())
	g.Player2.Camera.MakeCurrent()
}

func (g *GameSplitscreen) UnhandledInput(event InputEvent.Instance) {
	(&Game{}).handleInput(g.AsNode(), g.PauseMenu, event)
}

func main() {
	classdb.Register[Game]()
	classdb.Register[GameSplitscreen]()
	classdb.Register[PauseMenu](NewPauseMenu)
	classdb.Register[CoinsCounter]()
	classdb.Register[Coin]()
	classdb.Register[Bullet]()
	classdb.Register[Gun]()
	classdb.Register[Enemy]()
	classdb.Register[Player]()
	startup.Scene()
}
