package main

import (
	"game/tetris/game"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(500, 620, "Tetris")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	t := game.NewTetris()

	backgroundColor := rl.NewColor(30, 30, 30, 255)
	uiBackground := rl.NewColor(20, 20, 20, 255)

	for !rl.WindowShouldClose() {
		t.HandleInput()

		if game.EventTriggered(0.2) {
			t.MoveBlockDown()
		}

		rl.BeginDrawing()

		rl.ClearBackground(backgroundColor)

		rl.DrawText("Score", 365, 15, 30, rl.White)
		rl.DrawText("Next", 370, 175, 30, rl.White)
		if t.GetGameStatus() {
			rl.DrawText("GAME OVER", 320, 450, 27, rl.White)
		}
		rl.DrawRectangleRounded(rl.Rectangle{X: 320, Y: 55, Width: 170, Height: 60}, 0.3, 6, uiBackground)
		rl.DrawText(strconv.Itoa(int(t.GetScore())), 365, 72, 30, rl.White)
		rl.DrawRectangleRounded(rl.Rectangle{X: 320, Y: 215, Width: 170, Height: 180}, 0.3, 6, uiBackground)

		t.Draw()

		rl.EndDrawing()
	}
	game.NewTetris().Start()
}
