package grid

import (
	"game/tetris/colors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Grid struct {
	numRows    int32
	numColumns int32
	cellSize   int32

	figuresColors []rl.Color

	GridMatrix [][]int32
}

func NewGrid(numRows, numColumns, cellSize int32) *Grid {
	g := &Grid{
		numRows:       numRows,
		numColumns:    numColumns,
		cellSize:      cellSize,
		figuresColors: colors.GetColors(),
	}

	g.Initialize()

	return g
}

func (g *Grid) Initialize() {
	gridMatrix := make([][]int32, g.numRows)
	for i := range gridMatrix {
		gridMatrix[i] = make([]int32, g.numColumns)
	}

	g.GridMatrix = gridMatrix
}

func (g *Grid) Draw() {
	for row := range g.numRows {
		for col := range g.numColumns {
			cellValue := g.GridMatrix[row][col]
			rl.DrawRectangle(col*g.cellSize+11, row*g.cellSize+11, g.cellSize-1, g.cellSize-1, g.figuresColors[cellValue])
		}
	}
}

func (g *Grid) IsCellOutside(row, col int32) bool {
	if row >= 0 && row < g.numRows && col >= 0 && col < g.numColumns {
		return false
	}
	return true
}

func (g *Grid) IsCellEmpty(row, col int32) bool {
	if g.GridMatrix[row][col] == 0 {
		return true
	}
	return false
}

func (g *Grid) isRowFull(row int32) bool {
	for col := range g.numColumns {
		if g.GridMatrix[row][col] == 0 {
			return false
		}
	}
	return true
}

func (g *Grid) clearRow(row int32) {
	for col := range g.numColumns {
		g.GridMatrix[row][col] = 0
	}
}

func (g *Grid) moveRowDown(row, numRows int32) {
	for col := range g.numColumns {
		g.GridMatrix[row+numRows][col] = g.GridMatrix[row][col]
		g.GridMatrix[row][col] = 0
	}
}

func (g *Grid) ClearFullRows() int32 {
	var completed int32 = 0

	for row := g.numRows - 1; row >= 0; row-- {
		if g.isRowFull(row) {
			g.clearRow(row)
			completed++
		} else if completed > 0 {
			g.moveRowDown(row, completed)
		}
	}

	return completed
}
