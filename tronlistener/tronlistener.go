package tronlistener

import (
	"fmt"
	"net"
	"sync/atomic"
)

type Direction int32

const (
	Stopped Direction = iota
	North             = 65
	East              = 67
	South             = 66
	West              = 68
)

type Controller struct {
	conn             net.Conn
	currentDirection Direction
}

func (c Controller) read() {
	buf := make([]byte, 1)
	for {
		_, err := c.conn.Read(buf)
		if err != nil {
			break
		}
		switch buf[0] {
		// here, check if it's one of the arrow keys. if so, set direction.
		}
	}
}

func (c Controller) CurrentDirection() (d Direction) {
	d = Direction(atomic.LoadInt32((*int32)(&c.currentDirection)))
	return
}

func (c Controller) SetCurrentDirection(d Direction) {
	atomic.StoreInt32((*int32)(&c.currentDirection), int32(d))
}

func AcceptController(l net.Listener) (c Controller, err error) {
	var conn net.Conn
	conn, err = l.Accept()
	if err != nil {
		return
	}
	c = Controller{
		conn: conn,
	}
	go c.read()
	return
}
