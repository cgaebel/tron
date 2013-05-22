package main

import (
    "net"
    "github.com/wowus/tron/tronlistener"
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

    c, err := tronlistener.AcceptController(listener)
    if err != nil {
        panic(err)
    }

    _ = c

    select{}
}
