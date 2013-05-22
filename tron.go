package main

import (
	"flag"
	"github.com/nsf/termbox-go"
)

func doServerStuff() {
}

func doClientStuff() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}

	defer termbox.Close()
	termbox.HideCursor()

	client := new(Client)

	for {
		event := termbox.PollEvent()

		if event.Type == termbox.EventError {
			panic(event.Err)
		}

		client.Tick(event)
		termbox.Flush()
	}
}

func main() {
	var server = flag.Bool("server", false, "use this flag to start a server, as opposed to a client.")
	flag.Parse()

	if *server {
		doServerStuff()
	} else {
		doClientStuff()
	}
}
