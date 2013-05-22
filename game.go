package main

import (
	"math/rand"
)

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
}

func (p *Player) Die() {
	Grid.ClearSymbol(symbol)
	p.DeathCount += 1
}

type GridT [Height][Width]rune

func (g *GridT) IsEmpty(p Pos) bool {
	return g[p.Y][p.X] == Empty
}

func (g *GridT) SetCellValue(p Pos, val rune) {
	g[p.Y][p.X] = val
}

func (g *GridT) ClearSymbol(symbol rune) {
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			if Grid[y][x] == symbol {
				Grid[y][x] = Empty
			}
		}
	}
}

func (g *GridT) GetStartingVector() (pos Pos, dir Pos) {
	// pick random location
	pos.X = rand.Intn(Width)
	pos.Y = rand.Intn(Height)

	// pick a random direction
	// chosen by fair dice roll
	dir.X = 1
	dir.Y = 0

	// increment until safe
	for !g.IsEmpty(pos) {
		pos.X += 1
		if pos.X >= Width {
			pos.X = 0
			pos.Y += 1
			if pos.Y >= Height {
				pos.Y = 0
			}
		}
	}
	return
}

// 80x24
const Empty = ' '
const Width = 80
const Height = 24

var Grid GridT

var Players map[rune]Player

var NextSymbol = 'A'

func Init() {
	Players = make(map[rune]Player)

	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			Grid[y][x] = Empty
		}
	}
}

func AddPlayer() rune {
	p := new(Player)
	p.Symbol = NextSymbol
	NextSymbol += 1

	p.HeadPos, p.HeadDir = Grid.GetStartingVector()

	Players[p.Symbol] = *p
	return p.Symbol
}

func RemovePlayer(symbol rune) {
	Grid.ClearSymbol(symbol)
	delete(Players, symbol)
}

func Step() {
	for _, p := range Players {
		p.TickCount += 1

		p.HeadPos.Add(p.HeadDir)

		// wrap the pos
		p.HeadPos.X = p.HeadPos.X % Width
		p.HeadPos.Y = p.HeadPos.Y % Height

		if Grid.IsEmpty(p.HeadPos) {
			Grid.SetCellValue(p.HeadPos, p.Symbol)
		} else {
			p.Die()
		}
	}
}
