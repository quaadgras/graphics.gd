/*
{
	0: {
		"color": Color(1, 0, 0)
	},
	5: {
		"color": Color(0, 1, 0)
	}
}
*/

package main

import (
	"graphics.gd/classdb/SyntaxHighlighter"
	"graphics.gd/variant/Color"
)

func SyntaxHighlighter_GetLineSyntaxHighlighting() {
	var example = map[int]SyntaxHighlighter.Entry{
		0: {
			Color: Color.RGBA{1, 0, 0, 1},
		},
		5: {
			Color: Color.RGBA{0, 1, 0, 1},
		},
	}
	_ = example
}
