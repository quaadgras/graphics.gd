/*
func get_connection_list_from_port(node, port):
	var connections = get_connection_list_from_node(node)
	var result = []
	for connection in connections:
		var dict = {}
		if connection["from_node"] == node and connection["from_port"] == port:
			dict["node"] = connection["to_node"]
			dict["port"] = connection["to_port"]
			dict["type"] = "left"
			result.push_back(dict)
		elif connection["to_node"] == node and connection["to_port"] == port:
			dict["node"] = connection["from_node"]
			dict["port"] = connection["from_port"]
			dict["type"] = "right"
			result.push_back(dict)
	return result
*/

package main

import "graphics.gd/classdb/GraphEdit"

func getConnectionListFromPort(self GraphEdit.Instance, node string, port int) []map[string]any {
	var connections = self.GetConnectionListFromNode(node)
	var result []map[string]any
	for _, group := range connections {
		for _, connection := range group {
			if connection.FromNode == node && connection.FromPort == port {
				result = append(result, map[string]any{"node": connection.ToNode, "port": connection.ToPort, "type": "left"})
			} else if connection.ToNode == node && connection.ToPort == port {
				result = append(result, map[string]any{"node": connection.FromNode, "port": connection.FromPort, "type": "right"})
			}
		}
	}
	return result
}
