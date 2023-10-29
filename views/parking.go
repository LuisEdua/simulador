package views

import (
	"Simulador/models"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/font/basicfont"
)

const (
	ESPACIO_ANCHO  = 40
	ESPACIO_ALTO   = 80
	ESPACIO_MARGEN = 30
	PASILLO_ANCHO  = 100
	FILA_CAPACIDAD = 10
	PARED_GROSOR   = 10
	PORTON_ANCHO   = 10
	PORTON_ALTO    = 60
	PARED_MARGEN   = 10
	PASILLO        = 60
)

type EstacionamientoView struct {
	Estacionamiento *models.Estacionamiento
}

func (ev *EstacionamientoView) Dibujar(win *pixelgl.Window) {
	//anchoTotal := float64(10) * (ESPACIO_ANCHO + ESPACIO_MARGEN) - ESPACIO_MARGEN // restamos el último margen
	//altoTotal := 2*ESPACIO_ALTO + PASILLO
	baseX := PARED_GROSOR + ESPACIO_MARGEN

	mitadAlto := win.Bounds().H() / 2
	pasilloMitad := PASILLO / 2
	distanciaHastaPasillo := mitadAlto - float64(pasilloMitad) - ESPACIO_ALTO

	// Dibuja las paredes
	DibujarPared(win, 0, 0, PARED_GROSOR, win.Bounds().Max.Y)                                     // Pared izquierda
	DibujarPared(win, win.Bounds().Max.X-PARED_GROSOR, 0, win.Bounds().Max.X, win.Bounds().Max.Y) // Pared derecha
	DibujarPared(win, 0, win.Bounds().Max.Y-PARED_GROSOR, win.Bounds().Max.X, win.Bounds().Max.Y) // Pared superior
	DibujarPared(win, 0, 0, win.Bounds().Max.X, PARED_GROSOR)                                     // Pared inferior

	// Portón en la pared izquierda
	DibujarPorton(win, 0, win.Bounds().Center().Y-(PORTON_ALTO/2))

	for i := 0; i < 10; i++ {
		position := pixel.V(float64(baseX+i*(ESPACIO_ANCHO+ESPACIO_MARGEN)), distanciaHastaPasillo)
		ev.DibujarEspacio(win, position)
	}

	for i := 0; i < 10; i++ {
		position := pixel.V(float64(baseX+i*(ESPACIO_ANCHO+ESPACIO_MARGEN)), distanciaHastaPasillo+ESPACIO_ALTO+PASILLO)
		ev.DibujarEspacio(win, position)
	}
}

func DibujarPared(win *pixelgl.Window, x1, y1, x2, y2 float64) {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(0.3, 0.3, 0.3) // Color gris oscuro para las paredes
	imd.Push(pixel.V(x1, y1), pixel.V(x2, y2))
	imd.Rectangle(0)
	imd.Draw(win)
}

func DibujarPorton(win *pixelgl.Window, x, y float64) {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(255, 255, 255) // Color blanco para el portón
	imd.Push(pixel.V(x, y), pixel.V(x+PORTON_ANCHO, y+PORTON_ALTO))
	imd.Rectangle(0)
	imd.Draw(win)
}

func (ev *EstacionamientoView) DibujarEspacio(win *pixelgl.Window, position pixel.Vec) {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(0.5, 0.5, 0.5) // Cambio a un color gris más claro
	imd.Push(position, position.Add(pixel.V(ESPACIO_ANCHO, ESPACIO_ALTO)))
	imd.Rectangle(0)
	imd.Draw(win)
}

func (ev *EstacionamientoView) DibujarVehiculo(win *pixelgl.Window, v *models.Vehiculo, position pixel.Vec) {
	imd := imdraw.New(nil)
	imd.Color = pixel.RGB(0, 0, 1)
	imd.Push(position)
	imd.Push(position.Add(pixel.V(50, 30)))
	imd.Rectangle(0)

	imd.Draw(win)

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	txt := text.New(position.Add(pixel.V(5, 5)), basicAtlas)
	txt.Color = pixel.RGB(1, 0, 0)
	_, _ = fmt.Fprintf(txt, "%d", v.ID)
	txt.Draw(win, pixel.IM.Scaled(txt.Orig, 2))
}

func View() {
	win := cargarVentana()

	estacionamiento := models.NuevoEstacionamiento(20)
	estacionamientoView := &EstacionamientoView{Estacionamiento: estacionamiento}

	for !win.Closed() {
		win.Clear(pixel.RGB(255, 255, 255)) // fondo
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
		panic(fmt.Sprintf("Error al crear ventana: %v", err))
	}
	return win
}
