package tetris

import (
	"game/tetris/block"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type tetris struct {
	grid   grid
	blocks []block.InterfaceBlock
}

func NewTetris() *tetris {
	grid := newGrid(20, 10, 30)

	return &tetris{}
}

func (t *tetris) Start() {
	rl.InitWindow(300, 600, "Tetris")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	grid := newGrid(20, 10, 30)

	block := block.NewIBlock()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Gray)

		grid.Draw()
		block.Block.Draw()

		rl.EndDrawing()
	}
}
