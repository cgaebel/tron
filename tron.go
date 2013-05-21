package main

import "flag"
import "github.com/nsf/termbox-go"

func doServerStuff() {
}

func doClientStuff() {
    termbox.Init()
}

func main() {
    var server = flag.Bool("server", false, "use this flag to start a server, as opposed to a client.")
    flag.Parse();

    if(*server) {
        doServerStuff()
    } else {
        doClientStuff()
    }
}
