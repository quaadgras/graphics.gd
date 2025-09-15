/*
{
    from_node: StringName,
    from_port: int,
    to_node: StringName,
    to_port: int,
    keep_alive: bool
}
*/

package main

func GraphEdit_GetConnectionsIntersectingWithRect() {
	type Connection struct {
		FromNode  string `gd:"from_node"`
		FromPort  int    `gd:"from_port"`
		ToNode    string `gd:"to_node"`
		ToPort    int    `gd:"to_port"`
		KeepAlive bool   `gd:"keep_alive"`
	}
}
