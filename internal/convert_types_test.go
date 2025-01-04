package gd_test

import (
	"fmt"
	"testing"

	"graphics.gd/classdb"
	"graphics.gd/classdb/Node"
	"graphics.gd/classdb/SceneTree"
)

var doneConversionsTest = make(chan error)

type Converter struct {
	classdb.Extension[Converter, Node.Instance]

	test *testing.T
}

func (c Converter) Ready() {
	fmt.Println("Ready")
	close(doneConversionsTest)
}

func (c Converter) AsNode() Node.Instance { return c.Super() }

func TestConversions(t *testing.T) {
	converter := &Converter{
		test: t,
	}
	SceneTree.Add(converter)
	if err := <-doneConversionsTest; err != nil {
		t.Fatal(err)
	}
}
