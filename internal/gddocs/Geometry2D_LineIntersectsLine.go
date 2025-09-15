/*
[gdscript]
var from_a = Vector2.ZERO
var dir_a = Vector2.RIGHT
var from_b = Vector2.DOWN

# Returns Vector2(1, 0)
Geometry2D.line_intersects_line(from_a, dir_a, from_b, Vector2(1, -1))
# Returns Vector2(-1, 0)
Geometry2D.line_intersects_line(from_a, dir_a, from_b, Vector2(-1, -1))
# Returns null
Geometry2D.line_intersects_line(from_a, dir_a, from_b, Vector2.RIGHT)
[/gdscript]
[csharp]
var fromA = Vector2.Zero;
var dirA = Vector2.Right;
var fromB = Vector2.Down;

// Returns new Vector2(1, 0)
Geometry2D.LineIntersectsLine(fromA, dirA, fromB, new Vector2(1, -1));
// Returns new Vector2(-1, 0)
Geometry2D.LineIntersectsLine(fromA, dirA, fromB, new Vector2(-1, -1));
// Returns null
Geometry2D.LineIntersectsLine(fromA, dirA, fromB, Vector2.Right);
[/csharp]
*/

package main

import (
	"graphics.gd/classdb/Geometry2D"
	"graphics.gd/variant/Vector2"
)

func Geometry2D_LineIntersectsLine() {
	var fromA = Vector2.Zero
	var dirA = Vector2.Right
	var fromB = Vector2.Down

	// Returns new Vector2(1, 0)
	Geometry2D.LineIntersectsLine(fromA, dirA, fromB, Vector2.New(1, -1))
	// Returns new Vector2(-1, 0)
	Geometry2D.LineIntersectsLine(fromA, dirA, fromB, Vector2.New(-1, -1))
	// Returns null
	Geometry2D.LineIntersectsLine(fromA, dirA, fromB, Vector2.Right)
}
