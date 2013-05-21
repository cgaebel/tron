package main

import "flag"
import "github.com/nsf/termbox-go"

func doServerStuff() {
}

func doClientStuff() {
    err := termbox.Init()
    if err != nil {
        panic(err)
    }

    defer termbox.Close()
    termbox.HideCursor()

    for {
        // main loop here.
        termbox.Flush()
    }
}

func main() {
    var server = flag.Bool("server", false, "use this flag to start a server, as opposed to a client.")
    flag.Parse();

    if *server {
        doServerStuff()
    } else {
        doClientStuff()
    }
}
