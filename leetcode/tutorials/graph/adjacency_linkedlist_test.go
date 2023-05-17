package graph

import (
	"fmt"
	"testing"
)

func Test_graphAdjList_removeEdge(t *testing.T) {
	type fields struct {
		edges [][]Vertex
	}
	type args struct {
		vet1 Vertex
		vet2 Vertex
	}
	v1 := Vertex{Val: 1}
	v2 := Vertex{Val: 2}
	v3 := Vertex{Val: 3}
	// not exist node in the graph
	v4 := Vertex{Val: 4}
	var edges [][]Vertex
	edges = append(edges, []Vertex{v1, v2}, []Vertex{v1, v3}, []Vertex{v2, v3})
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{name: "delete exist edge", fields: fields{edges}, args: args{v1, v2}},
		{name: "delete not exist edge", fields: fields{edges}, args: args{v1, v3}},
		{name: "delete not exist node edge", fields: fields{edges}, args: args{v1, v4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := newGraphAdjList(tt.fields.edges)
			g.print()
			g.removeEdge(tt.args.vet1, tt.args.vet2)
			fmt.Printf("删除边: %v - %v,之后:\n", tt.args.vet1, tt.args.vet2)
			g.print()
		})
	}
}
