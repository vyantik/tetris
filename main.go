package main

import (
	"game/tetris/game"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(500, 620, "Tetris")

	defer rl.CloseWindow()
	defer rl.CloseAudioDevice()

	rl.SetTargetFPS(60)

	t := game.NewTetris()

	defer t.UnloadAll()

	backgroundColor := rl.NewColor(30, 30, 30, 255)
	uiBackground := rl.NewColor(20, 20, 20, 255)

	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream(t.GetMusic())

		t.HandleInput()

		if game.EventTriggered(0.4) {
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

		score := strconv.Itoa(int(t.GetScore()))
		mesureScore := rl.MeasureText(score, 38)

		rl.DrawText(score, 320+(170-mesureScore)/2, 68, 38, rl.White)

		rl.DrawRectangleRounded(rl.Rectangle{X: 320, Y: 215, Width: 170, Height: 180}, 0.3, 6, uiBackground)

		t.Draw()

		rl.EndDrawing()
	}
	game.NewTetris().Start()
}
