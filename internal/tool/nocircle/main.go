package main

import (
	"fmt"
	"sort"

	"graphics.gd/internal/gdjson"
	"runtime.link/api/xray"
)

// Graph represents a directed graph using an adjacency list.
type Graph map[string][]string

// AddEdge adds a directed edge from u to v, ensuring no duplicates.
func (g Graph) AddEdge(u, v string) {
	if g[u] == nil {
		g[u] = []string{}
	}
	// Check for duplicates
	for _, neighbor := range g[u] {
		if neighbor == v {
			return
		}
	}
	g[u] = append(g[u], v)
}

// getOutdegree returns the number of outgoing edges for a node.
func (g Graph) getOutdegree(node string) int {
	return len(g[node])
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
					if len(cycle) == 0 || cycle[0] != node {
						cycle = append(cycle, neighbor)
					}
					return true
				}
			} else if recStack[neighbor] {
				// Cycle detected
				cycle = append([]string{neighbor, node}, cycle...)
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

// ResolveCircularDependencies finds the edges to remove to eliminate cycles,
// prioritizing edges where the source node has fewer outgoing edges.
func (g Graph) ResolveCircularDependencies() []string {
	edgesToRemove := []string{}

	for {
		_, cycle := g.topologicalSort()
		if cycle == nil {
			// Graph is acyclic, no more edges to remove
			break
		}

		// Find the edge in the cycle with the source node having the smallest outdegree
		var edgeToRemove struct {
			u, v   string
			degree int
		}
		minDegree := -1

		// Cycle is a list of nodes forming a loop (e.g., [A, B, C, A])
		for i := 0; i < len(cycle)-1; i++ {
			u := cycle[i+1]
			v := cycle[i]
			degree := g.getOutdegree(u)
			if minDegree == -1 || degree < minDegree {
				minDegree = degree
				edgeToRemove = struct {
					u, v   string
					degree int
				}{u, v, degree}
			}
		}

		// Remove the chosen edge
		u, v := edgeToRemove.u, edgeToRemove.v
		edgesToRemove = append(edgesToRemove, fmt.Sprintf("%s -> %s", u, v))

		// Update the graph by removing the edge
		for i, neighbor := range g[u] {
			if neighbor == v {
				g[u] = append(g[u][:i], g[u][i+1:]...)
				break
			}
		}
	}

	// Sort edges for consistent output
	sort.Strings(edgesToRemove)
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
			if gdjson.Relocations[class.Name+"."+method.Name] != "" {
				continue
			}
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
