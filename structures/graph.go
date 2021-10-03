package main

import "fmt"

type Vertex struct {
	name string
	age  int
	city string
}

type AdjacencyList struct {
	head *Vertex
}

type Graph struct {
	name  string
	size  int
	nodes map[Vertex][]*Vertex
}

// Constructor of the graph
func createGraph(name string, size int) *Graph {
	return &Graph{
		name:  name,
		size:  size,
		nodes: make(map[Vertex][]*Vertex),
	}
}

func (g *Graph) addVertex() {

}

func (g *Graph) addEdge() {

}

/*
	Visit children fo lower level then return to root
*/
func dfs() {

}

func depthSuccessors() {

}

/*
	Visit all children of root then descend to lower level
*/
func bfs() {

}

func breadthSuccessors() {

}

func main() {

	graph := createGraph("Employee Graph", 10)
	fmt.Println(graph)
}
