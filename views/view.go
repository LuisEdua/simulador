package views

import (
	"Simulador/models"
	"Simulador/scenes"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"sync"
	"time"
)

func drawParking(win *pixelgl.Window) {
	imd := imdraw.New(nil)
	imd.Color = colornames.White

	// Define los puntos para las líneas del cuadro
	p1 := pixel.V(100, 460)
	p2 := pixel.V(100, 20)
	p3 := pixel.V(600, 20)
	p4 := pixel.V(600, 460)
	p5 := pixel.V(100, 300)
	p6 := pixel.V(100, 180)

	imd.Push(p1)
	imd.Push(p5)
	imd.Line(1)

	imd.Push(p6)
	imd.Push(p2)
	imd.Line(1)

	imd.Push(p2)
	imd.Push(p3)
	imd.Line(1)

	imd.Push(p3)
	imd.Push(p4)
	imd.Line(1)

	imd.Push(p4)
	imd.Push(p1)
	imd.Line(1)

	width := 40.0
	height := 80.0
	space := 10.0

	for x := 105.0; x <= 595; x += width + space {
		p1 := pixel.V(x, 455)
		p2 := pixel.V(x, 455-height)
		p3 := pixel.V(x+width, 455-height)
		p4 := pixel.V(x+width, 455)

		imd.Push(p1)
		imd.Push(p2)
		imd.Line(1)

		imd.Push(p3)
		imd.Push(p4)
		imd.Line(1)

		imd.Push(p4)
		imd.Push(p1)
		imd.Line(1)
	}

	for x := 105.0; x <= 595; x += width + space {
		p1 := pixel.V(x, 25)
		p2 := pixel.V(x, 25+height)
		p3 := pixel.V(x+width, 25+height)
		p4 := pixel.V(x+width, 25)

		imd.Push(p1)
		imd.Push(p2)
		imd.Line(1)

		imd.Push(p3)
		imd.Push(p4)
		imd.Line(1)

		imd.Push(p4)
		imd.Push(p1)
		imd.Line(1)
	}

	imd.Draw(win)

}

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Líneas en Pixel",
		Bounds: pixel.R(0, 0, 640, 480),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	gateFree := make(chan bool, 1)
	mu := sync.Mutex{}

	p := models.NewParking(20, mu, gateFree)

	gateFree <- true

	i := 1

	for !win.Closed() {
		win.Clear(colornames.Dimgray)

		drawParking(win)

		go scenes.Start(win, p, i)
		time.Sleep(500 * time.Millisecond)

		win.Update()
	}
}

func Show() {
	pixelgl.Run(run)
}
