package game

import (
	"game/tetris/block"
	"game/tetris/grid"
	"math/rand"

	"slices"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type game struct {
	grid grid.Grid

	blocks       []block.IBlock
	currentBlock block.IBlock
	nextBlock    block.IBlock

	score  int32
	isOver bool

	music rl.Music

	rotateSound rl.Sound
	clearSound  rl.Sound
}

func NewTetris() *game {
	tetrisPointer := &game{
		grid:   *grid.NewGrid(20, 10, 30),
		blocks: getAllBlocks(),

		score:  0,
		isOver: false,
	}
	tetrisPointer.currentBlock = tetrisPointer.getRandomBlock()
	tetrisPointer.nextBlock = tetrisPointer.getRandomBlock()

	rl.InitAudioDevice()

	tetrisPointer.music = rl.LoadMusicStream("sounds/music.mp3")
	tetrisPointer.rotateSound = rl.LoadSound("sounds/rotate.mp3")
	tetrisPointer.clearSound = rl.LoadSound("sounds/clear.mp3")

	rl.PlayMusicStream(tetrisPointer.music)

	return tetrisPointer
}

func (t *game) GetGameStatus() bool {
	return t.isOver
}

func (t *game) GetScore() int32 {
	return t.score
}

func (t *game) GetMusic() rl.Music {
	return t.music
}

var lastUpdateTime float64 = 0.0

func EventTriggered(interval float64) bool {
	currentTime := rl.GetTime()
	if currentTime-lastUpdateTime >= interval {
		lastUpdateTime = currentTime
		return true
	}
	return false
}

func (t *game) UnloadAll() {
	rl.UnloadMusicStream(t.music)
	rl.UnloadSound(t.rotateSound)
	rl.UnloadSound(t.clearSound)
}

func (t *game) Start() {

}

func (t *game) getRandomBlock() block.IBlock {
	if len(t.blocks) == 0 {
		t.blocks = getAllBlocks()
	}

	blockIndex := rand.Intn(len(t.blocks))
	block := t.blocks[blockIndex]

	t.blocks = slices.Delete(t.blocks, blockIndex, blockIndex+1)

	return block
}

func (t *game) Draw() {
	t.grid.Draw()
	t.currentBlock.Draw(11, 11)

	switch t.nextBlock.GetId() {
	case 3:
		t.nextBlock.Draw(255, 290)
	case 4:
		t.nextBlock.Draw(255, 280)
	default:
		t.nextBlock.Draw(270, 270)

	}
}

func (t *game) HandleInput() {
	keyPressed := rl.GetKeyPressed()
	if t.isOver && keyPressed != 0 {
		t.isOver = false
		t.Reset()
	}
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
		t.MoveBlockDown()
		t.updateScore(0, 1)
	case rl.KeyDown:
		t.MoveBlockDown()
		t.updateScore(0, 1)
	}
}

func (t *game) moveBlockLeft() {
	if !t.isOver {
		t.currentBlock.Move(0, -1)
		if t.isBlockOutside() || !t.blockFits() {
			t.currentBlock.Move(0, 1)
		}
	}

}

func (t *game) moveBlockRight() {
	if !t.isOver {
		t.currentBlock.Move(0, 1)
		if t.isBlockOutside() || !t.blockFits() {
			t.currentBlock.Move(0, -1)
		}
	}
}

func (t *game) MoveBlockDown() {
	if !t.isOver {
		t.currentBlock.Move(1, 0)
		if t.isBlockOutside() || !t.blockFits() {
			t.currentBlock.Move(-1, 0)
			t.lockBlock()
		}
	}
}

func (t *game) isBlockOutside() bool {
	tiles := t.currentBlock.GetCellPositions()

	for _, tile := range tiles {
		if t.grid.IsCellOutside(tile.GetRow(), tile.GetCol()) {
			return true
		}
	}

	return false
}

func (t *game) rotateBlock() {
	if !t.isOver {
		t.currentBlock.Rotate()
		if t.isBlockOutside() || !t.blockFits() {
			t.currentBlock.UndoRotation()
		} else {
			rl.PlaySound(t.rotateSound)
		}
	}
}

func (t *game) lockBlock() {
	tiles := t.currentBlock.GetCellPositions()

	for _, tile := range tiles {
		t.grid.GridMatrix[int(tile.GetRow())][int(tile.GetCol())] = t.currentBlock.GetId()
	}

	t.currentBlock = t.nextBlock
	if !t.blockFits() {
		t.isOver = true
	}
	t.nextBlock = t.getRandomBlock()
	rowsCleared := t.grid.ClearFullRows()

	if rowsCleared > 0 {
		rl.PlaySound(t.clearSound)
		t.updateScore(rowsCleared, 0)
	}
}

func (t *game) blockFits() bool {
	tiles := t.currentBlock.GetCellPositions()

	for _, tile := range tiles {
		if !t.grid.IsCellEmpty(tile.GetRow(), tile.GetCol()) {
			return false
		}
	}

	return true
}

func (t *game) updateScore(linesCleared, moveDownPoints int32) {
	switch linesCleared {
	case 1:
		t.score += 100
	case 2:
		t.score += 300
	case 3:
		t.score += 500
	}

	t.score += moveDownPoints
}

func (t *game) Reset() {
	t.grid.Initialize()
	t.blocks = getAllBlocks()
	t.currentBlock = t.getRandomBlock()
	t.nextBlock = t.getRandomBlock()
	t.score = 0
}

func getAllBlocks() []block.IBlock {
	return []block.IBlock{block.NewITypeBlock(), block.NewJTypeBlock(), block.NewLTypeBlock(), block.NewOTypeBlock(), block.NewSTypeBlock(), block.NewTTypeBlock(), block.NewZTypeBlock()}
}
