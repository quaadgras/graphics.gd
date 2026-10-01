package main

import (
	"graphics.gd/classdb"
	"graphics.gd/classdb/Tree"
	"graphics.gd/startup"
)

// GalleryTree populates itself with a few example items.
type GalleryTree struct {
	Tree.Extension[GalleryTree] `gd:"GalleryTree"`
	classdb.Tool
}

func (t *GalleryTree) Ready() {
	tree := t.AsTree()
	root := tree.CreateItem()
	root.SetText(0, "Tree - Root")
	child1 := tree.MoreArgs().CreateItem(root, -1)
	child1.SetText(0, "Tree - Child 1")
	child2 := tree.MoreArgs().CreateItem(root, -1)
	child2.SetText(0, "Tree - Child 2")
	subchild1 := tree.MoreArgs().CreateItem(child1, -1)
	subchild1.SetText(0, "Tree - Subchild 1")
}

func main() {
	classdb.Register[GalleryTree]()
	startup.Scene()
}
