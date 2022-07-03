package graph_representation

import (
	"fmt"
	"github.com/myawesomeproject1/my_utils"
	"testing"
)

func TestPrintEdgesAndTheWeight(t *testing.T) {
	edge_list := []*my_utils.GraphEdge{}
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "A", End: "B", Weight: 5})
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "A", End: "C", Weight: 7})
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "A", End: "D", Weight: 3})
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "B", End: "E", Weight: 2})
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "B", End: "F", Weight: 10})
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "C", End: "G", Weight: 1})
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "D", End: "H", Weight: 11})
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "E", End: "H", Weight: 9})
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "F", End: "H", Weight: 4})
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "G", End: "H", Weight: 6})
	PrintEdgesAndTheWeight(edge_list)
}

func TestGetVerticesFromEdgeList(t *testing.T) {
	edge_list := []*my_utils.GraphEdge{}
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "A", End: "B", Weight: 5})
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "A", End: "C", Weight: 7})
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "A", End: "D", Weight: 3})
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "B", End: "E", Weight: 2})
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "B", End: "F", Weight: 10})
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "C", End: "G", Weight: 1})
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "D", End: "H", Weight: 11})
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "E", End: "H", Weight: 9})
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "F", End: "H", Weight: 4})
	edge_list = append(edge_list, &my_utils.GraphEdge{Start: "G", End: "H", Weight: 6})
	vertices := GetVerticesFromEdgeList(edge_list)
	for key, _ := range vertices {
		fmt.Println(key)
	}
}
