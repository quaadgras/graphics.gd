package main

import (
	"graphics.gd/classdb"
	"graphics.gd/startup"
	"graphics.gd/variant/RefCounted"
)

type Summator struct {
	RefCounted.Extension[Summator]

	count int
}

func (s *Summator) Add(n int)     { s.count += n }
func (s *Summator) Reset()        { s.count = 0 }
func (s *Summator) GetTotal() int { return s.count }

func main() {
	classdb.Register[Summator]()
	startup.AsExtension()
}
