//go:build !generate

package gd_test

import (
	"testing"

	"graphics.gd/classdb/GraphEdit"
)

// TestGraphEditConnectionListFromNode covers issue #328: the engine returns an
// Array of Dictionaries here, so the generated Go type has to be []Connection.
// It used to be [][]Connection, which no value the engine can return will ever
// convert into.
func TestGraphEditConnectionListFromNode(t *testing.T) {
	var graph = GraphEdit.New()
	defer graph.AsNode().QueueFree()
	if err := graph.ConnectNode("StartNode", 0, "EndNode", 1); err != nil {
		t.Fatal(err)
	}
	var connections = graph.GetConnectionListFromNode("StartNode")
	if len(connections) != 1 {
		t.Fatalf("expected 1 connection, got %d", len(connections))
	}
	var connection = connections[0]
	if connection.FromNode != "StartNode" || connection.FromPort != 0 {
		t.Errorf("unexpected source %q port %d", connection.FromNode, connection.FromPort)
	}
	if connection.ToNode != "EndNode" || connection.ToPort != 1 {
		t.Errorf("unexpected target %q port %d", connection.ToNode, connection.ToPort)
	}
}
