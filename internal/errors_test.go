package gd_test

import (
	"testing"

	"graphics.gd/classdb/Expression"
)

func TestErrors(t *testing.T) {
	var expr = Expression.New()
	if err := expr.Parse("2 + 2"); err != nil {
		t.Error(err)
	}
	if err := expr.Parse("2+++"); err == nil {
		t.Error("expected error")
	}
}
