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
	p.Mu.Lock()
	p.Cars = append(p.Cars, c)

	models.Come(c)

	for _, c := range p.Cars {
		imd := imdraw.New(nil)
		imd.Color = colornames.Blue

		imd.Push(c.P1)
		imd.Push(c.P2)
		imd.Line(c.Width)
		imd.Draw(win)
	}

	win.Update()

	if p.Capacity > 0 {
		p.Capacity--
		i := models.In(p, c)
		p.Mu.Unlock()

		time.Sleep(time.Duration(c.T) * time.Second)

		models.Out(i, p, c)

		time.Sleep(500 * time.Millisecond)

		models.Go(c)
	} else {
		models.Go(c)
		p.Mu.Unlock()
	}
}

func Run(i int, p *models.Parking, win *pixelgl.Window) {
	c := models.NewCar(i, pixel.V(40, 180), pixel.V(40, 120))
	go carRoutine(c, p, win)
}
