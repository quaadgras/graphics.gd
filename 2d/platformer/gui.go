package main

import (
	"strconv"

	"graphics.gd/classdb/AnimatedSprite2D"
	"graphics.gd/classdb/Button"
	"graphics.gd/classdb/CenterContainer"
	"graphics.gd/classdb/Control"
	"graphics.gd/classdb/Label"
	"graphics.gd/classdb/Panel"
	"graphics.gd/classdb/PropertyTweener"
	"graphics.gd/classdb/SceneTree"
	"graphics.gd/classdb/Tween"
	"graphics.gd/variant/Float"
	"graphics.gd/variant/Rect2"
)

type CoinsCounter struct {
	Panel.Extension[CoinsCounter]

	Label          Label.Instance
	AnimatedSprite AnimatedSprite2D.Instance `gd:"AnimatedSprite2D"`

	coinsCollected int
}

func (c *CoinsCounter) Ready() {
	c.Label.SetText(strconv.Itoa(c.coinsCollected))
	c.AnimatedSprite.Play()
}

func (c *CoinsCounter) CollectCoin() {
	c.coinsCollected++
	c.Label.SetText(strconv.Itoa(c.coinsCollected))
}

type PauseMenu struct {
	Control.Extension[PauseMenu]

	FadeInDuration  Float.X
	FadeOutDuration Float.X

	CenterCont   CenterContainer.Instance `gd:"ColorRect/CenterContainer"`
	ResumeButton Button.Instance          `gd:"ColorRect/CenterContainer/VBoxContainer/ResumeButton"`
	CoinsCounter *CoinsCounter            `gd:"ColorRect/CoinsCounter"`
}

func NewPauseMenu() *PauseMenu {
	return &PauseMenu{FadeInDuration: 0.3, FadeOutDuration: 0.2}
}

func (p *PauseMenu) Ready() {
	p.AsCanvasItem().Hide()
}

func (p *PauseMenu) Close() {
	tween := p.AsNode().CreateTween()
	SceneTree.Get(p.AsNode()).SetPaused(false)
	PropertyTweener.Make(tween, p.AsObject(), "modulate:a", 0.0, p.FadeOutDuration).
		SetTrans(Tween.TransLinear).SetEase(Tween.EaseOut)
	PropertyTweener.Make(tween.Parallel(), p.CenterCont.AsObject(), "anchor_bottom", 0.5, p.FadeOutDuration).
		SetTrans(Tween.TransCubic).SetEase(Tween.EaseOut)
	tween.TweenCallback(p.AsCanvasItem().Hide)
}

func (p *PauseMenu) Open() {
	p.AsCanvasItem().Show()
	p.ResumeButton.AsControl().GrabFocus()

	modulate := p.AsCanvasItem().Modulate()
	modulate.A = 0
	p.AsCanvasItem().SetModulate(modulate)
	p.CenterCont.AsControl().SetAnchor(Rect2.SideBottom, 0.5)
	tween := p.AsNode().CreateTween()
	PropertyTweener.Make(tween, p.AsObject(), "modulate:a", 1.0, p.FadeInDuration).
		SetTrans(Tween.TransLinear).SetEase(Tween.EaseIn)
	PropertyTweener.Make(tween.Parallel(), p.CenterCont.AsObject(), "anchor_bottom", 1.0, p.FadeOutDuration).
		SetTrans(Tween.TransCubic).SetEase(Tween.EaseOut)
}

func (p *PauseMenu) OnCoinCollected() {
	p.CoinsCounter.CollectCoin()
}

func (p *PauseMenu) OnResumeButtonPressed() {
	p.Close()
}

func (p *PauseMenu) OnSingleplayerButtonPressed() {
	if p.AsCanvasItem().Visible() {
		tree := SceneTree.Get(p.AsNode())
		tree.SetPaused(false)
		tree.ChangeSceneToFile("res://game_singleplayer.tscn")
	}
}

func (p *PauseMenu) OnSplitscreenButtonPressed() {
	if p.AsCanvasItem().Visible() {
		tree := SceneTree.Get(p.AsNode())
		tree.SetPaused(false)
		tree.ChangeSceneToFile("res://game_splitscreen.tscn")
	}
}

func (p *PauseMenu) OnQuitButtonPressed() {
	if p.AsCanvasItem().Visible() {
		SceneTree.Get(p.AsNode()).Quit()
	}
}
