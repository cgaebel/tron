package main

import (
	"fmt"
	"math/rand"
	"time"
)

const TickMs = 500

var Tick = make(chan bool)

// 80x24
const Empty = ' '
const Width = 80
const Height = 24

var Grid GridT

// this is a set. bool is a dummy value
var Players map[*Player]bool

var NextSymbol = 'A'

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
	Grid.ClearSymbol(p.Symbol)
	p.DeathCount += 1
}

type GridT [Height][Width]rune

func (g *GridT) Debug() {
	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			fmt.Printf("%c", Grid[y][x])
		}
		fmt.Print("\n")
	}
}

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

	dirBits := rand.Int()

	// pick a random direction
	// chosen by fair dice roll
	dir.X = dirBits & 1
	dir.Y = ^(dirBits & 1)

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

func init() {
	Players = make(map[*Player]bool)

	for y := 0; y < Height; y++ {
		for x := 0; x < Width; x++ {
			Grid[y][x] = Empty
		}
	}
}

func AddPlayer() *Player {
	p := new(Player)
	p.Symbol = NextSymbol
	NextSymbol += 1

	p.HeadPos, p.HeadDir = Grid.GetStartingVector()

	Players[p] = true
	return p
}

func (p *Player) Remove() {
	Grid.ClearSymbol(p.Symbol)
	delete(Players, p)
}

func (p *Player) SetDirection(dir Pos) {
	p.HeadDir = dir
}

func Step() {
	for p := range Players {
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

func Run() {
	for {
		time.Sleep(TickMs)
		Tick <- true
	}
}
