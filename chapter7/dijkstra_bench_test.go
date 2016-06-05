package chapter7

import "testing"

var anotherGraph = WeightedGraph{
	StartNode: {
		"LP":     5,
		"Poster": 0,
	},
	"LP": {
		"Guitar": 15,
		"Drums":  20,
	},
	"Poster": {
		"Guitar": 30,
		"Drums":  35,
	},
	"Guitar": {
		EndNode: 30,
	},
	"Drums": {
		EndNode: 10,
	},
	EndNode: {},
}

//BenchmarkDijkstra benchmarks Dijkstra algorithm.
func BenchmarkDijkstra(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Dijkstra(anotherGraph)
	}
}
