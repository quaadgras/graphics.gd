package Transform2D_test

import (
	"fmt"
	"testing"

	"grow.graphics/gd/gdmaths/Transform2D"
	"grow.graphics/gd/internal/gdtests"
)

func TestIdentity(t *testing.T) {
	var transform = Transform2D.Identity
	gdtests.Print(t, "| X | Y | Origin", "| X | Y | Origin")
	gdtests.Print(t, "| 1 | 0 | 0", fmt.Sprintf("| %v | %v | %v", transform.X.X, transform.Y.X, transform.Origin.X))
	gdtests.Print(t, "| 0 | 1 | 0", fmt.Sprintf("| %v | %v | %v", transform.X.Y, transform.Y.Y, transform.Origin.Y))
	// Prints:
	// | X | Y | Origin
	// | 1 | 0 | 0
	// | 0 | 1 | 0
}
