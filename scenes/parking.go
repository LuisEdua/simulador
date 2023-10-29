package scenes

import (
	"Simulador/models"
	"Simulador/views"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
)

type EstacionamientoScene struct {
	View *views.EstacionamientoView
}

func NuevoEstacionamientoScene(win *pixelgl.Window) *EstacionamientoScene {
	estacionamiento := models.NuevoEstacionamiento(20)
	view := &views.EstacionamientoView{Estacionamiento: estacionamiento}

	return &EstacionamientoScene{View: view}
}

func (es *EstacionamientoScene) Iniciar(win *pixelgl.Window) {
	for !win.Closed() {
		win.Clear(pixel.RGB(0.9, 0.9, 0.9)) // fondo
		es.View.Dibujar(win)
		win.Update()
	}
}
