package main

import (
	"flag"
	"fmt"

	"github.com/miketheprogrammer/go-thrust"
)

var (
	host = flag.String("host", "0.0.0.0", "IP address to bind to")
	port = flag.Int("port", 8000, "TCP port to listen on")
)

func main() {
	flag.Parse()
	thrust.InitLogger()
	thrust.Start()

	thrustWindow := thrust.NewWindow(fmt.Sprintf("http://127.0.0.1:%d", *port), nil)
	thrustWindow.Show()
	thrustWindow.Focus()
	// BLOCKING - Dont run before youve excuted all commands you want first.
	thrust.LockThread()
}
