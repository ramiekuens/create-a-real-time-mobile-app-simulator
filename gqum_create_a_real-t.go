// Package main is the entry point for our real-time mobile app simulator
package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
	"time"
)

// Simulator represents the real-time mobile app simulator
type Simulator struct {
	app fyne.App
 win fyne.Window
}

// NewSimulator creates a new instance of the simulator
func NewSimulator() *Simulator {
	a := app.New()
	w := a.NewWindow("Real-time Mobile App Simulator")
	return &Simulator{
		app: a,
		win: w,
	}
}

// Run starts the simulator
func (s *Simulator) Run() {
	s.win.Resize(fyne.NewSize(360, 640)) // iPhone 11 Pro size

	// Create a mobile app UI
	ui := container.NewVBox(
		widget.NewLabel("Mobile App Simulator"),
		widget.NewButton("Button 1", func() {}),
		widget.NewButton("Button 2", func() {}),
	)

	s.win.SetContent(ui)

	// Start the simulator
	go func() {
		for {
			s.win.Canvas().Refresh()
			time.Sleep(16 * time.Millisecond) // 60Hz refresh rate
		}
	}()

	s.win.ShowAndRun()
}

func main() {
	sim := NewSimulator()
	sim.Run()
	log.Println("Simulator started")
}