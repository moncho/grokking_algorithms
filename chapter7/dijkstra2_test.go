package chapter7

import "testing"

func TestDijkstra2(t *testing.T) {
	solution, path := dijkstra(testGraph, StartNode)
	if solution[EndNode] != 6 {
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

func TestDijkstra2Again(t *testing.T) {
	solution, path := dijkstra(anotherGraph, StartNode)

	if solution[EndNode] != 35 {
		t.Errorf("Dijkstra solution to graph is %d, expected %d", solution[EndNode], 35)
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
