package main

import (
	"time"

	"github.com/miketheprogrammer/go-thrust"
	"github.com/miketheprogrammer/go-thrust/tutorial/provisioner"
)

func main() {
	thrust.InitLogger()
	// Set any Custom Provisioners before Start
	thrust.SetProvisioner(tutorial.NewTutorialProvisioner())
	// thrust.Start() must always come before any bindings are created.
	thrust.Start()
	thrustWindow := thrust.NewWindow("http://breach.cc/", nil)
	thrustWindow.Show()
	thrustWindow.Maximize()
	thrustWindow.Focus()

	thrustWindow.OpenDevtools()
	// Lets do a window timeout
	go func() {
		<-time.After(time.Second * 10)
		thrustWindow.CloseDevtools()
	}()
	// BLOCKING - Dont run before youve excuted all commands you want first
	thrust.LockThread()
}
