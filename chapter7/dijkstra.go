package chapter7

import "math"

type cost map[Node]int
type parents map[Node]Node
type nodeSet map[Node]struct{}

//Node a graph node
type Node string

//WeightedGraph a graph whose edges have been assigned weights
type WeightedGraph map[Node]map[Node]int

const (
	//StartNode must be the first node of the graph
	StartNode = "start"
	//EndNode must be the end node of the graph
	EndNode = "fin"
)

//Dijkstra calculates the shortest path for a weighted graph.
func Dijkstra(graph WeightedGraph) (int, map[Node]Node) {
	_ = "breakpoint"
	costs, parentNodes := createCostsParents(graph)
	var processed nodeSet = make(map[Node]struct{})

	for node := findLowestCostNode(costs, processed); node != EndNode; node = findLowestCostNode(costs, processed) {
		cost := costs[node]
		neighbors := graph[node]
		for k, v := range neighbors {
			newCost := cost + v
			currentCost, present := costs[k]
			if !present || currentCost > newCost {
				costs[k] = newCost
				parentNodes[k] = node
			}
			processed[node] = struct{}{}
		}
	}
	_ = "breakpoint"
	return costs[EndNode], parentNodes
}

//createCostsParents initializes helper structs for tracking
//node costs and parents
func createCostsParents(graph WeightedGraph) (cost, parents) {
	costs := make(cost)
	parentNodes := make(parents)

	for node, cost := range graph[StartNode] {
		costs[node] = cost
		parentNodes[node] = StartNode
	}
	costs[EndNode] = math.MaxInt32
	return costs, parentNodes
}

func findLowestCostNode(cost cost, processedNodes nodeSet) Node {
	var lowestCostNode Node
	lowestCost := math.MaxInt32
	for node, cost := range cost {
		if cost < lowestCost && !processed(processedNodes, node) {
			lowestCost = cost
			lowestCostNode = node
		}
	}
	return lowestCostNode
}

func processed(processedNodes nodeSet, node Node) bool {
	_, isPresent := processedNodes[node]
	return isPresent
}
