package graph_representation

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddAndViewAdjacentMatrix(t *testing.T) {
	adjacentMatrix := NewAdjacentMatrix(4)
	adjacentMatrix.AddEdge(1, 0)
	adjacentMatrix.AddEdge(2, 3)
	//adjacentMatrix.RemoveEdge(2, 3)
	matrixArr := adjacentMatrix.GetAdjacentMatrixRepresentation()
	assert.Equal(t, (int)(matrixArr[1][0]), 1)
	assert.Equal(t, (int)(matrixArr[0][1]), 1)
	assert.Equal(t, (int)(matrixArr[2][3]), 1)
	assert.Equal(t, (int)(matrixArr[3][2]), 1)
}
