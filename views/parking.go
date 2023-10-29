package views

import (
	"fmt"
	"Simulador/models"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

type EstacionamientoView struct {
	Estacionamiento *models.Estacionamiento
	BaseSprite      pixel.Sprite
	VehiculoSprite  pixel.Sprite
}

func (ev *EstacionamientoView) Dibujar(win *pixelgl.Window) {
	mat := pixel.IM.Moved(win.Bounds().Center())
	ev.BaseSprite.Draw(win, mat)

	for i, vehiculo := range ev.Estacionamiento.Vehiculos {
		position := pixel.V(100+float64(i)*60, win.Bounds().Center().Y)
		ev.DibujarVehiculo(win, vehiculo, position)
	}
}

func (ev *EstacionamientoView) DibujarVehiculo(win *pixelgl.Window, v *models.Vehiculo, position pixel.Vec) {
	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(position, basicAtlas)
	txt.Color = pixel.RGB(1, 0, 0) // Color rojo
	_, _ = fmt.Fprintf(txt, "%d", v.ID)
	txt.Draw(win, pixel.IM.Scaled(txt.Orig, 2))
}

func View() {
	win := cargarVentana()

	estacionamiento := models.NuevoEstacionamiento(20)
	estacionamientoView := &EstacionamientoView{Estacionamiento: estacionamiento}

	for !win.Closed() {
		win.Clear(pixel.RGB(0.9, 0.9, 0.9)) // fondo
		estacionamientoView.Dibujar(win)
		win.Update()
	}
}

func cargarVentana() *pixelgl.Window {
	winCfg := pixelgl.WindowConfig{
		Title:  "Estacionamiento",
		Bounds: pixel.R(0, 0, 800, 600),
	}
	win, err := pixelgl.NewWindow(winCfg)
	if err != nil {
		panic(err)
	}
	return win
}
