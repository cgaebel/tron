package tronlistener

import (
    "net"   
    "sync/atomic"
    "fmt"
)

type Direction int32

const (
    Stopped Direction = iota
    North
    East
    South
    West
)

type Controller struct {
    conn net.Conn
    currentDirection Direction
}

func (c Controller) read() {
    buf := make([]byte, 1)
    for {
        _, err := c.conn.Read(buf)
        if err != nil {
            break
        }
        switch Direction(buf[0]) {
        case North:
            fmt.Println("North")
        case South:
            fmt.Println("South")
        case East:
            fmt.Println("East")
        case West:
            fmt.Println("West")
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
        conn:conn,
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