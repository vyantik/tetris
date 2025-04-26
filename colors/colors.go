package colors

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	figuresColors = []rl.Color{
		rl.DarkGray,
		rl.Red,
		rl.DarkBlue,
		rl.DarkGreen,
		rl.Yellow,
		rl.Orange,
		rl.Pink,
		rl.Purple,
	}
)

func GetColors() []rl.Color {
	return figuresColors
}
