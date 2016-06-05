package chapter7

import "testing"

var testGraph = WeightedGraph{
	StartNode: {
		"A": 6,
		"B": 2,
	},
	"A": {
		EndNode: 1,
	},
	"B": {
		"A":     3,
		EndNode: 6,
	},
	EndNode: {},
}

func TestDijkstra(t *testing.T) {
	solution, path := Dijkstra(testGraph)

	if solution != 6 {
		t.Errorf("Dijkstra solution to graph is %d, expected %d", solution, 6)
	}
	if path[EndNode] != "A" {
		t.Errorf("Parent of End is not B")
	}
	if path[path[EndNode]] != "B" {
		t.Errorf("Parent of A is not B")
	}
	if path[path[path[EndNode]]] != StartNode {
		t.Errorf("Parent of B is not the start of the graph")
	}
}

func TestDijkstraAgain(t *testing.T) {
	solution, path := Dijkstra(anotherGraph)

	if solution != 35 {
		t.Errorf("Dijkstra solution to graph is %d, expected %d", solution, 35)
	}
	if path[EndNode] != "Drums" {
		t.Errorf("Parent of End is not Drums")
	}
	if path[path[EndNode]] != "LP" {
		t.Errorf("Parent of Drums is not LP")
	}
	if path[path[path[EndNode]]] != StartNode {
		t.Errorf("Parent of LP is not the start of the graph")
	}
}
