package main

const (
	width  = 80
	height = 24
)

type Client struct {
	toDisplay [width][height]rune
}

func (client *Client) drawWallPanel(x, y int) {
	client.toDisplay[x][y] = '#'
}

func (client *Client) drawVertWalls() {
	for i := 0; i < height; i++ {
		client.drawWallPanel(0, i)
		client.drawWallPanel(width, i)
	}
}

func (client *Client) drawHorizWalls() {
	for i := 0; i < width; i++ {
		client.drawWallPanel(i, 0)
		client.drawWallPanel(i, height)
	}
}

func (client *Client) drawGrid() {
	client.drawVertWalls()
	client.drawHorizWalls()
}
