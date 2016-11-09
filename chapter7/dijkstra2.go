package chapter7

// dijkstra's algorithm as described https://en.wikipedia.org/wiki/Dijkstra%27s_algorithm
func dijkstra(graph WeightedGraph, start Node) (map[Node]int, map[Node]Node) {
	var unvisited []Node
	dist := map[Node]int{}
	prev := map[Node]Node{}

	for k := range graph {
		unvisited = append(unvisited, k)
		dist[k] = max
	}
	dist[start] = 0

	for len(unvisited) > 0 {
		min := findMin(unvisited, dist)
		unvisited = remove(min, unvisited)
		neighbors := graph[min]
		for neighbor := range neighbors {
			newDistance := dist[min] + graph[min][neighbor]

			if newDistance < dist[neighbor] {
				dist[neighbor] = newDistance
				prev[neighbor] = min
			}
		}
	}
	return dist, prev
}

func findMin(unvisited []Node, dist map[Node]int) Node {
	min := max
	var result Node
	for _, v := range unvisited {
		if dist[v] < min {
			result = v
			min = dist[v]
		}
	}
	return result
}

func remove(Node Node, vertices []Node) []Node {
	index := -1
	for i, v := range vertices {
		if Node == v {
			index = i
			break
		}
	}
	return append(vertices[:index], vertices[index+1:]...)
}
