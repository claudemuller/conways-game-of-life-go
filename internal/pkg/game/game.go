package game

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	ScreenWidth  int32
	ScreenHeight int32

	rows     int32
	cols     int32
	cellSize int32

	gameOver bool
	pause    bool

	cells []bool
}

func NewGame(cellSize, cols, rows int32) Game {
	return Game{
		ScreenWidth:  cellSize * cols,
		ScreenHeight: cellSize * rows,
		rows:         rows,
		cols:         cols,
		cellSize:     cellSize,
		cells:        make([]bool, rows*cols),
	}
}

func (g *Game) Init() {
	// Random cell fill.
	// rand.Seed(time.Now().UnixNano())
	// for i := int32(0); i < g.rows*g.cols; i++ {
	// 	if rand.Intn(3) == 0 {
	// 		g.cells[i] = true
	// 	}
	// }
}

func (g *Game) Update() {
	if !g.gameOver {
		if rl.IsKeyPressed(rl.KeyP) {
			g.pause = !g.pause
		}

		if !g.pause {
			if rl.IsMouseButtonDown(rl.MouseLeftButton) {
				x := rl.GetMouseX() / g.cellSize
				y := rl.GetMouseY() / g.cellSize
				i := x*g.cols + y
				g.cells[i] = true
			}
		}
	} else {
		if rl.IsKeyPressed(rl.KeyEnter) {
			g.Init()
			g.gameOver = false
		}
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.RayWhite)

	if !g.gameOver {
		for i := int32(0); i < g.rows*g.cols; i++ {
			x := int32(g.cellSize * (i / g.cols))
			y := int32(g.cellSize * (i % g.cols))
			if g.cells[i] {
				rl.DrawRectangle(x, y, g.cellSize, g.cellSize, rl.Black)
				continue
			}
			// rl.DrawRectangleLines(x, y, cellSize, cellSize, rl.LightGray)
			// rl.DrawText(strconv.Itoa(i), x, y, 10, rl.Black)
		}
	} else {
		// Play again screen...
	}

	rl.EndDrawing()
}
