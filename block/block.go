package block

import (
	"game/tetris/colors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type IBlock interface {
	Draw(offsetX int32, offsetY int32)
	Move(y, x int32)
	GetCellPositions() []position
	Rotate()
	UndoRotation()
	GetId() int32
}

type block struct {
	id int32

	cells map[int][]position

	colors []rl.Color

	cellSize      int32
	rotationState int

	rowOffset int32
	colOffset int32
}

func newBlock() *block {
	return &block{
		cellSize:      30,
		rotationState: 0,
		colors:        colors.GetColors(),
		rowOffset:     0,
		colOffset:     0,
	}
}

func (b *block) Draw(offsetX int32, offsetY int32) {
	tiles := b.GetCellPositions()
	for _, tile := range tiles {
		rl.DrawRectangle(tile.col*b.cellSize+offsetX, tile.row*b.cellSize+offsetY, b.cellSize-1, b.cellSize-1, b.colors[b.id])
	}
}

func (b *block) Move(rows, cols int32) {
	b.colOffset += cols
	b.rowOffset += rows
}

func (b *block) GetCellPositions() []position {
	tiles := b.cells[b.rotationState]
	var movedTiles []position

	for _, v := range tiles {
		movedTiles = append(movedTiles, *newPosition(v.row+b.rowOffset, v.col+b.colOffset))
	}

	return movedTiles
}

func (b *block) Rotate() {
	b.rotationState++
	if b.rotationState == len(b.cells) {
		b.rotationState = 0
	}
}

func (b *block) UndoRotation() {
	b.rotationState--
	if b.rotationState == -1 {
		b.rotationState = len(b.cells) - 1
	}
}

func (b *block) GetId() int32 {
	return b.id
}
