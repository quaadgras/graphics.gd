package binpatch

type sorter struct {
	p *PatchSet
}

func (s sorter) Len() int {
	return len(s.p.Patches)
}

func (s sorter) Less(i, j int) bool {
	return s.p.Patches[i].Offset < s.p.Patches[j].Offset
}

func (s sorter) Swap(i, j int) {
	s.p.Patches[i], s.p.Patches[j] = s.p.Patches[j], s.p.Patches[i]
	s.p.Blobs[i], s.p.Blobs[j] = s.p.Blobs[j], s.p.Blobs[i]
}
