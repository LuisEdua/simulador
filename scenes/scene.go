// scenes/scenes.go
package scenes

import (
	"Simulador/models"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"time"
)

func carRoutine(c *models.Car, p *models.Parking, win *pixelgl.Window) {
	imd := imdraw.New(nil)
	imd.Color = colornames.Blue

	p.Mu.Lock()
	models.Come(c)
	imd.Push(c.P1)
	imd.Push(c.P2)
	imd.Line(c.Width)
	imd.Draw(win)
	win.Update()
	if p.Capacity > 0 {
		p.Capacity--
		i := models.In(p, c)
		imd.Clear()
		imd.Draw(win)
		imd.Push(c.P1)
		imd.Push(c.P2)
		imd.Line(c.Width)
		imd.Draw(win)
		win.Update()
		p.Mu.Unlock()
		time.Sleep(time.Duration(c.T) * time.Second)
		models.Out(i, p, c)
		imd.Clear()
		imd.Draw(win)
		imd.Push(c.P1)
		imd.Push(c.P2)
		imd.Line(c.Width)
		imd.Draw(win)
		win.Update()
		p.Mu.Unlock()
		time.Sleep(500 * time.Millisecond)
		imd.Clear()
		imd.Draw(win)
		win.Update()
		return
	} else {
		models.Go(c)
		p.Mu.Unlock()
		return
	}
}

func Start(win *pixelgl.Window, p *models.Parking, i int) {
	c := models.NewCar(i, pixel.V(40, 180), pixel.V(40, 120))
	go carRoutine(c, p, win)
	i++
}
