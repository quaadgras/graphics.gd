package main

import (
	"fmt"

	"graphics.gd/internal/gdjson"
	"runtime.link/api/xray"
)

// Graph represents a directed graph using an adjacency list.
type Graph map[string][]string

// AddEdge adds a directed edge from u to v.
func (g Graph) AddEdge(u, v string) {
	if g[u] == nil {
		g[u] = []string{}
	}
	g[u] = append(g[u], v)
}

// topologicalSort attempts a topological sort and detects cycles.
// Returns the sorted order if acyclic, or nil and a cycle if a cycle exists.
func (g Graph) topologicalSort() ([]string, []string) {
	visited := make(map[string]bool)
	recStack := make(map[string]bool) // Tracks nodes in the current recursion stack
	order := []string{}
	cycle := []string{}

	var dfs func(node string) bool
	dfs = func(node string) bool {
		visited[node] = true
		recStack[node] = true

		for _, neighbor := range g[node] {
			if !visited[neighbor] {
				if dfs(neighbor) {
					cycle = append(cycle, neighbor)
					return true
				}
			} else if recStack[neighbor] {
				// Cycle detected
				cycle = append(cycle, neighbor, node)
				return true
			}
		}

		recStack[node] = false
		order = append([]string{node}, order...) // Prepend to order
		return false
	}

	for node := range g {
		if !visited[node] {
			if dfs(node) {
				return nil, cycle
			}
		}
	}

	return order, nil
}

// ResolveCircularDependencies finds the edges to remove to eliminate cycles.
func (g Graph) ResolveCircularDependencies() []string {
	edgesToRemove := []string{}

	for {
		_, cycle := g.topologicalSort()
		if cycle == nil {
			// Graph is acyclic, no more edges to remove
			break
		}

		// Choose an edge to remove from the cycle (e.g., first edge)
		u := cycle[1] // Node in cycle pointing to...
		v := cycle[0] // ...the node it loops back to
		edgesToRemove = append(edgesToRemove, fmt.Sprintf("%s -> %s", u, v))

		// Remove the edge from the graph
		for i, neighbor := range g[u] {
			if neighbor == v {
				g[u] = append(g[u][:i], g[u][i+1:]...)
				break
			}
		}
	}

	return edgesToRemove
}

func main() {
	var dependency = make(Graph)
	spec, err := gdjson.LoadSpecification()
	if err != nil {
		panic(xray.New(err))
	}
	var classdb = make(map[string]bool)
	for _, class := range spec.Classes {
		classdb[class.Name] = true
	}
	for _, class := range spec.Classes {
		for _, method := range class.Methods {
			for _, arg := range method.Arguments {
				if classdb[arg.Type] && arg.Type != class.Name {
					dependency.AddEdge(class.Name, arg.Type)
				}
			}
			if classdb[method.ReturnValue.Type] && method.ReturnValue.Type != class.Name {
				dependency.AddEdge(class.Name, method.ReturnValue.Type)
			}
		}
	}
	// Resolve circular dependencies
	edgesToRemove := dependency.ResolveCircularDependencies()

	fmt.Println("Edges to remove to eliminate circular dependencies:")
	for _, edge := range edgesToRemove {
		fmt.Println(edge)
	}
}
