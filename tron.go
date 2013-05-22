package main

import (
	"net"
)

func main() {
	laddr, err := net.ResolveTCPAddr("tcp", "localhost:5678")
	if err != nil {
		panic(err)
	}

	listener, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		panic(err)
	}

	c, err := AcceptController(listener)
	if err != nil {
		panic(err)
	}

	go Run()
	p := AddPlayer()

	for {
		<-Tick
		Step()
		dir := c.CurrentDirection()
		p.SetDirection(dir.ToBetterDir())
		c.Send(Grid)
		Grid.Debug()
	}
}
