package main

import (
	"graphics.gd/classdb"
	"graphics.gd/startup"
)

func main() {
	classdb.Register[Outline]()
	classdb.Register[Aura]()
	classdb.Register[Blur]()
	classdb.Register[Fatty]()
	classdb.Register[DropShadow]()
	classdb.Register[OffsetShadow]()
	classdb.Register[Silhouette]()
	classdb.Register[Glow]()
	classdb.Register[Disintegrate]()
	startup.Scene()
}
