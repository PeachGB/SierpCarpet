package main

import (
	"fmt"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const WIDTH = 1600
const HEIGHT = 1200

const ROWS = 4
const COLS = 4

const AREAY = HEIGHT / ROWS
const AREAX = WIDTH / COLS
const CENTERX = (WIDTH / 2)
const CENTERY = (HEIGHT / 2)

func SierpCarp(x int, y int, w int, h int, depth int, d int) {
	if d == depth {

	} else {
		rl.DrawRectangle(int32(x-w/2), int32(y-h/2), int32(w), int32(h), rl.Red)

		SierpCarp(x+1*w, y, w/3, h/3, depth, d+1)
		SierpCarp(x-1*w, y, w/3, h/3, depth, d+1)
		SierpCarp(x, y+1*h, w/3, h/3, depth, d+1)
		SierpCarp(x, y-1*h, w/3, h/3, depth, d+1)
		SierpCarp(x+1*w, y+1*h, w/3, h/3, depth, d+1)
		SierpCarp(x-1*w, y-1*h, w/3, h/3, depth, d+1)
		SierpCarp(x-1*h, y+1*h, w/3, h/3, depth, d+1)
		SierpCarp(x+1*h, y-1*h, w/3, h/3, depth, d+1)
	}
}
func main() {
	var i int
	fmt.Println("TYPE THE DEPTH NUMBER:")
	fmt.Scan(&i)

	rl.InitWindow(int32(WIDTH), int32(HEIGHT), "Sierpiński carpet")
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		SierpCarp(CENTERX, CENTERY, AREAX, AREAX, i, 0)
		rl.EndDrawing()
	}

}
