package main

import (
	"graphics.gd/classdb"
	"graphics.gd/startup"
)

func main() {
	classdb.Register[Coin]()
	classdb.Register[Bullet]()
	classdb.Register[Enemy](NewEnemy)
	classdb.Register[Player](NewPlayer)
	startup.Scene()
}
