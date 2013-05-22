package main

import (
	"net"
	"sync/atomic"
)

type Direction int32

const (
	Stopped Direction = 0
	North             = 65
	East              = 67
	South             = 66
	West              = 68
)

type Controller struct {
	conn             net.Conn
	currentDirection Direction
}

func (c Controller) Send(g GridT) {
    for line := range g {
        c.conn.Write([]byte(string(line)))
        c.conn.Write([]byte("\r\n"))
    }
    c.conn.Write([]byte("\r\n"))
}

func (c Controller) read() {
	buf := make([]byte, 1)
	for {
		_, err := c.conn.Read(buf)
		if err != nil {
			break
		}
		d := Direction(buf[0])
		switch d {
		case North, South, East, West:
			c.SetCurrentDirection(d)
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
	// SB IAC WILL SGA
	_, err = c.conn.Write([]byte{250, 255, 251, 3, 240})
	if err != nil {
		return
	}
	// SB IAC WILL ECHO
	_, err = c.conn.Write([]byte{250, 255, 251, 1, 240})
	if err != nil {
		return
	}

	go c.read()
	return
}
