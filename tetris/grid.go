package tetris

import (
	"game/tetris/colors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type grid struct {
	numRows    int32
	numColumns int32
	cellSize   int32

	figuresColors []rl.Color

	gridMatrix [][]int32
}

func newGrid(numRows, numColumns, cellSize int32) *grid {
	gridMatrix := make([][]int32, numRows)
	for i := range gridMatrix {
		gridMatrix[i] = make([]int32, numColumns)
	}

	return &grid{
		numRows:       numRows,
		numColumns:    numColumns,
		cellSize:      cellSize,
		figuresColors: colors.GetColors(),
		gridMatrix:    gridMatrix,
	}
}

func (g *grid) Draw() {
	for row := range g.numRows {
		for col := range g.numColumns {
			cellValue := g.gridMatrix[row][col]
			rl.DrawRectangle(col*g.cellSize+1, row*g.cellSize+1, g.cellSize-1, g.cellSize-1, g.figuresColors[cellValue])
		}
	}
}
