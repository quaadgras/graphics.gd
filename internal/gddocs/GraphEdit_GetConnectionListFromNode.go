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

func GraphEdit_GetConnectionListFromNode() {
	type Connection struct {
		FromNode  string `json:"from_node"`
		FromPort  int    `json:"from_port"`
		ToNode    string `json:"to_node"`
		ToPort    int    `json:"to_port"`
		KeepAlive bool   `json:"keep_alive"`
	}
}
