package main

import (
	"fmt"
)

/*
Adjacency lists slice implementation
Depth first: O(Vertex+Edges)
Breadth first: O(Vertex+Edges)
*/

type Vertex struct {
	name    string
	age     int
	city    string
	visited bool
}

type Graph struct {
	name  string
	size  int
	nodes []*Vertex
	edges map[*Vertex][]*Vertex
}

// Constructor for vertexes
func createVertex(name string, age int, city string) *Vertex {
	return &Vertex{
		name:    name,
		age:     age,
		city:    city,
		visited: false,
	}
}

// Constructor of the graph
func createGraph(name string, size int) *Graph {
	return &Graph{
		name:  name,
		size:  size,
		nodes: make([]*Vertex, 0),
		edges: make(map[*Vertex][]*Vertex),
	}
}

// Adds a vertex to the graph
func (g *Graph) addVertex(v *Vertex) {
	g.nodes = append(g.nodes, v)
}

// Adds an edge to the graph (directed)
func (g *Graph) addDirectedEdge(v1, v2 *Vertex) {
	fmt.Println("Adding directed edge: ", v1.name, "->", v2.name)
	g.edges[v1] = append(g.edges[v1], v2)
}

// Adds an edge to the graph (undirected)
func (g *Graph) addUndirectedEdge(v1, v2 *Vertex) {
	fmt.Println("Adding undirected edge: ", v1.name, "<->", v2.name)
	g.edges[v1] = append(g.edges[v1], v2)
	g.edges[v2] = append(g.edges[v2], v1)
}

// Depth first search algorithm entry point
func (g *Graph) dfs() {
	fmt.Println("Starting Depth First Search")
	for _, v := range g.nodes {
		v.visited = false
	}
	for _, v := range g.nodes {
		if !v.visited {
			depthSuccessors(g, v)
		}
	}
}

// Explore successors of root vertex recursively
func depthSuccessors(g *Graph, root *Vertex) {
	fmt.Println("Visiting:", root.name)
	root.visited = true
	g.printEdges(root)
	// Explore all first child vertex
	for _, neighbour := range g.edges[root] {
		if !neighbour.visited {
			depthSuccessors(g, neighbour)
		}
	}
}

// Breadth first search algorithm entry point
func (g *Graph) bfs() {
	fmt.Println("Starting Breadth First Search")
	for _, v := range g.nodes {
		v.visited = false
	}
	for _, v := range g.nodes {
		if !v.visited {
			breadthSuccessors(g, v)
		}
	}
}

// Vertex Queue for BFS algorithm
var queue []*Vertex

// Explore successors of root vertex recursively
func breadthSuccessors(g *Graph, root *Vertex) {
	fmt.Println("Visiting:", root.name)
	root.visited = true
	g.printEdges(root)

	queue = append(queue, root)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		// Explore all child vertexes
		for _, neighbour := range g.edges[curr] {
			if !neighbour.visited {
				fmt.Println("Visiting:", neighbour.name)
				neighbour.visited = true
				queue = append(queue, neighbour)
				g.printEdges(neighbour)
			}
		}
	}
}

// Utility method to print edges of a vertex
func (g *Graph) printEdges(v *Vertex) {
	fmt.Println("Edges ->")
	for _, edge := range g.edges[v] {
		fmt.Printf("\t%v\n", *edge)
	}
}

// Main entry point
func main() {

	graph := createGraph("Employee Graph", 10)
	node1 := createVertex("Vince", 23, "Montreal")
	node2 := createVertex("Enrique", 27, "Mexico")
	node3 := createVertex("Francis", 38, "Paris")
	node4 := createVertex("Julius", 76, "Rome")
	node5 := createVertex("Pietra", 51, "Russia")

	graph.addVertex(node1)
	graph.addVertex(node2)
	graph.addVertex(node3)
	graph.addVertex(node4)
	graph.addVertex(node5)

	graph.addDirectedEdge(node1, node2)
	graph.addDirectedEdge(node2, node3)
	graph.addDirectedEdge(node3, node4)

	graph.addUndirectedEdge(node5, node1)

	fmt.Printf("\n=================================\n\n")
	graph.dfs()
	fmt.Printf("\n=================================\n\n")
	graph.bfs()
}
