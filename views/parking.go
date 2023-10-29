package views

import (
	"fmt"
	"Simulador/models"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/font/basicfont"
)

type EstacionamientoView struct {
	Estacionamiento *models.Estacionamiento
}

func (ev *EstacionamientoView) Dibujar(win *pixelgl.Window) {
    // Dibujar todos los espacios del estacionamiento
    for i := 0; i < len(ev.Estacionamiento.Vehiculos); i++ {
        position := pixel.V(100+float64(i)*60, win.Bounds().Center().Y)
        ev.DibujarEspacio(win, position)
    }

    // Dibujar vehículos
    for i, vehiculo := range ev.Estacionamiento.Vehiculos {
        position := pixel.V(100+float64(i)*60, win.Bounds().Center().Y)
        ev.DibujarVehiculo(win, vehiculo, position)
    }
}

func (ev *EstacionamientoView) DibujarEspacio(win *pixelgl.Window, position pixel.Vec) {
    imd := imdraw.New(nil)
    imd.Color = pixel.RGB(0.7, 0.7, 0.7) // Color gris para el espacio
    imd.Push(position)
    imd.Push(position.Add(pixel.V(50, 30)))
    imd.Rectangle(0)

    imd.Draw(win)
}


func (ev *EstacionamientoView) DibujarVehiculo(win *pixelgl.Window, v *models.Vehiculo, position pixel.Vec) {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(0, 0, 1) // Color azul para el vehículo
	imd.Push(position)
	imd.Push(position.Add(pixel.V(50, 30)))
	imd.Rectangle(0)

	imd.Draw(win)

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(position.Add(pixel.V(5, 5)), basicAtlas) // Texto ligeramente desplazado para que no se superponga con el vehículo
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