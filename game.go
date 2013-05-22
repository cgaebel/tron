package main

const ClockSpeed = 30

type Pos struct {
	X int
	Y int
}

func (p *Pos) Add(other Pos) {
	p.X += other.X
	p.Y += other.Y
}

type Player struct {
	Symbol     rune
	HeadPos    Pos
	HeadDir    Pos
	DeathCount int
	TickCount  int
	Name       string
}

type GridT [Height][Width]rune

func (g *GridT) IsEmpty(p Pos) bool {
	return g[p.Y][p.X] == Empty
}

func (g *GridT) SetCellValue(p Pos, val rune) {
	g[p.Y][p.X] = val
}

// 80x24
const Empty = rune(32)
const Width = 80
const Height = 24

var Grid GridT

var Players []Player

func Init() {
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			Grid[y][x] = Empty
		}
	}
}

func Step() {
	for _, p := range Players {
		p.TickCount += 1

		p.HeadPos.Add(p.HeadDir)
		if Grid.IsEmpty(p.HeadPos) {
			Grid.SetCellValue(p.HeadPos, p.Symbol)
		} else {
			p.DeathCount += 1
		}
	}
}
