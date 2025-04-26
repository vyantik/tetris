package tetris

import (
	"game/tetris/block"
	"math/rand"

	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type tetris struct {
	grid grid

	blocks       []block.IBlock
	currentBlock block.IBlock
	nextBlock    block.IBlock
}

func NewTetris() *tetris {
	tetrisPointer := &tetris{
		grid:   *newGrid(20, 10, 30),
		blocks: getAllBlocks(),
	}
	tetrisPointer.currentBlock = tetrisPointer.getRandomBlock()
	tetrisPointer.nextBlock = tetrisPointer.getRandomBlock()

	return tetrisPointer
}

func (t *tetris) Start() {
	rl.InitWindow(300, 600, "Tetris")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		t.handleInput()

		rl.BeginDrawing()

		rl.ClearBackground(rl.Gray)
		t.draw()

		rl.EndDrawing()
	}
}

func (t *tetris) getRandomBlock() block.IBlock {
	if len(t.blocks) == 0 {
		t.blocks = getAllBlocks()
	}

	blockIndex := rand.Intn(len(t.blocks))
	block := t.blocks[blockIndex]

	t.blocks = slices.Delete(t.blocks, blockIndex, blockIndex+1)

	return block
}

func (t *tetris) draw() {
	t.grid.draw()
	t.currentBlock.Draw()
}

func (t *tetris) handleInput() {
	keyPressed := rl.GetKeyPressed()
	switch keyPressed {
	case rl.KeyA:
		t.moveBlockLeft()
	case rl.KeyLeft:
		t.moveBlockLeft()
	case rl.KeyD:
		t.moveBlockRight()
	case rl.KeyRight:
		t.moveBlockRight()
	case rl.KeyW:
		t.rotateBlock()
	case rl.KeyUp:
		t.rotateBlock()
	case rl.KeyS:
		t.moveBlockDown()
	}
}

func (t *tetris) moveBlockLeft() {
	t.currentBlock.Move(0, -1)
	if t.isBlockOutside() {
		t.currentBlock.Move(0, 1)
	}
}

func (t *tetris) moveBlockRight() {
	t.currentBlock.Move(0, 1)
	if t.isBlockOutside() {
		t.currentBlock.Move(0, -1)
	}
}

func (t *tetris) moveBlockDown() {
	t.currentBlock.Move(1, 0)
	if t.isBlockOutside() {
		t.currentBlock.Move(-1, 0)
	}
}

func (t *tetris) isBlockOutside() bool {
	tiles := t.currentBlock.GetCellPositions()
	for _, tile := range tiles {
		if t.grid.IsCellOutside(tile.GetRow(), tile.GetCol()) {
			return true
		}
	}
	return false
}

func (t *tetris) rotateBlock() {
	t.currentBlock.Rotate()
	if t.isBlockOutside() {
		t.currentBlock.UndoRotation()
	}
}

func getAllBlocks() []block.IBlock {
	return []block.IBlock{block.NewITypeBlock(), block.NewJTypeBlock(), block.NewLTypeBlock(), block.NewOTypeBlock(), block.NewSTypeBlock(), block.NewTTypeBlock(), block.NewZTypeBlock()}
}
