package main

import "github.com/nsf/termbox-go"

const (
	width  = 80
	height = 24
)

type Client struct {
}

func drawWallPanel(x, y int) {
	termbox.SetCell(x, y, '#', termbox.ColorWhite, termbox.ColorDefault)
}

func drawVertWalls() {
	for i := 0; i < height; i++ {
		drawWallPanel(0, i)
		drawWallPanel(width, i)
	}
}

func drawHorizWalls() {
	for i := 0; i < width; i++ {
		drawWallPanel(i, 0)
		drawWallPanel(i, height)
	}
}

func drawGrid() {
	drawVertWalls()
	drawHorizWalls()
}

func (client *Client) Tick(event termbox.Event) {
	drawGrid()
}
