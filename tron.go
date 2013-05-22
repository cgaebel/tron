package main

import (
	"flag"
)

func doServerStuff() {
}

func doClientStuff() {
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
