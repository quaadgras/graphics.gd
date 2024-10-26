package Basis_test

import (
	"fmt"
	"testing"

	"grow.graphics/gd/gdmaths/Basis"
	"grow.graphics/gd/internal/gdtests"
)

func TestIdentity(t *testing.T) {
	var basis = Basis.Identity
	gdtests.Print(t, "| X | Y | Z", "| X | Y | Z")
	gdtests.Print(t, "| 1 | 0 | 0", fmt.Sprintf("| %v | %v | %v", basis.X.X, basis.Y.X, basis.Z.X))
	gdtests.Print(t, "| 0 | 1 | 0", fmt.Sprintf("| %v | %v | %v", basis.X.Y, basis.Y.Y, basis.Z.Y))
	gdtests.Print(t, "| 0 | 0 | 1", fmt.Sprintf("| %v | %v | %v", basis.X.Z, basis.Y.Z, basis.Z.Z))
	// Prints:
	// | X | Y | Z
	// | 1 | 0 | 0
	// | 0 | 1 | 0
	// | 0 | 0 | 1
}
