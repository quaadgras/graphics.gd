package gd

type Transform2D [3]Vector2

func (t Transform2D) tdotx(v Vector2) float32 { return t[0][x]*v[x] + t[1][x]*v[y] }
func (t Transform2D) tdoty(v Vector2) float32 { return t[0][y]*v[x] + t[1][y]*v[y] }
