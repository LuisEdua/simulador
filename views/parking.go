package views

import (
	"Simulador/models"
	"fmt"
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"math/rand"
	"time"
)

const (
	ESPACIO_ANCHO    = 40
	ESPACIO_ALTO     = 80
	ESPACIO_MARGEN   = 30
	PARED_GROSOR     = 10
	PORTON_ANCHO     = 10
	PORTON_ALTO      = 60
	MARGEN_IZQUIERDO = 90
	MARGEN_DERECHO   = 10
	MARGEN_SUPERIOR  = 100
	MARGEN_INFERIOR  = 100
	PASILLO          = 60
)

type EstacionamientoView struct {
	Estacionamiento *models.Estacionamiento
	SemaforoEspacio chan struct{}
	SemaforoPuerta  chan struct{}
}

func (ev *EstacionamientoView) Dibujar(win *pixelgl.Window) {
	anchoTotal := float64(10)*(ESPACIO_ANCHO+ESPACIO_MARGEN) - ESPACIO_MARGEN // restamos el último margen
	altoTotal := 2*ESPACIO_ALTO + PASILLO
	baseX := PARED_GROSOR + ESPACIO_MARGEN + 65

	fmt.Println("X:", anchoTotal)
	fmt.Println("Y:", altoTotal)

	mitadAlto := win.Bounds().H() / 2
	pasilloMitad := PASILLO / 2
	distanciaHastaPasillo := mitadAlto - float64(pasilloMitad) - ESPACIO_ALTO

	DibujarPared(win, MARGEN_IZQUIERDO, MARGEN_INFERIOR, MARGEN_IZQUIERDO+PARED_GROSOR, win.Bounds().Max.Y-MARGEN_SUPERIOR)

	DibujarPared(win, win.Bounds().Max.X-PARED_GROSOR-MARGEN_DERECHO, MARGEN_INFERIOR, win.Bounds().Max.X-MARGEN_DERECHO, win.Bounds().Max.Y-MARGEN_SUPERIOR)

	DibujarPared(win, MARGEN_IZQUIERDO, win.Bounds().Max.Y-PARED_GROSOR-MARGEN_SUPERIOR, win.Bounds().Max.X-MARGEN_DERECHO, win.Bounds().Max.Y-MARGEN_SUPERIOR)

	DibujarPared(win, MARGEN_IZQUIERDO, MARGEN_INFERIOR, win.Bounds().Max.X-MARGEN_DERECHO, MARGEN_INFERIOR+PARED_GROSOR)

	DibujarPorton(win, MARGEN_IZQUIERDO, win.Bounds().Center().Y-(PORTON_ALTO/2))

	for i := 0; i < 10; i++ {
		position := pixel.V(float64(baseX+i*(ESPACIO_ANCHO+ESPACIO_MARGEN)), distanciaHastaPasillo)
		ev.DibujarEspacio(win, position)
	}

	for i := 0; i < 10; i++ {
		position := pixel.V(float64(baseX+i*(ESPACIO_ANCHO+ESPACIO_MARGEN)), distanciaHastaPasillo+ESPACIO_ALTO+PASILLO)
		ev.DibujarEspacio(win, position)
	}

	go ev.IniciarSimulacion()

}

func (ev *EstacionamientoView) IniciarSimulacion() {
	models.NuevoEstacionamiento(20)
	for {
		time.Sleep(time.Duration(rand.Intn(10)+1) * time.Second)
		vehiculo := &models.Vehiculo{ID: models.GenerarNuevoID()}
		ev.AñadirVehiculo(vehiculo)
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

func (ev *EstacionamientoView) AñadirVehiculo(v *models.Vehiculo) {
	ev.SemaforoPuerta <- struct{}{}

	ev.SemaforoEspacio <- struct{}{}
	hayEspacio := ev.Estacionamiento.EspaciosDisponibles > 0
	if hayEspacio {
		ev.Estacionamiento.EspaciosDisponibles--
	}
	<-ev.SemaforoEspacio

	<-ev.SemaforoPuerta

	if hayEspacio {
		espacio := ev.BuscarEspacioDisponible()
		ev.AnimarEntrada(v, espacio)
		ev.EstacionarVehiculo(v, espacio)
	} else {
		// Si no hay espacio, el vehículo se va
		// ... (lógica para hacer que el vehículo se vaya, si es necesario)
	}
}

func (ev *EstacionamientoView) BuscarEspacioDisponible() int {
	for i, v := range ev.Estacionamiento.Vehiculos {
		if v == nil {
			return i
		}
	}
	return -1
}

func (ev *EstacionamientoView) AnimarEntrada(v *models.Vehiculo, espacio pixel.Vec, win *pixelgl.Window) {
	const velocidad = 2.0 // pixels por frame, puedes ajustar según necesites

	// Posición inicial del vehículo (en el portón)
	posicionVehiculo := pixel.V(MARGEN_IZQUIERDO, win.Bounds().Center().Y)

	for {
		// Dibuja el estacionamiento y el vehículo en su posición actual
		win.Clear(pixel.RGB(255, 255, 255)) // fondo
		ev.Dibujar(win)
		DibujarVehiculo(win, posicionVehiculo, pixel.RGB(0, 1, 0)) // Vehículo color verde

		// Actualiza la ventana
		win.Update()

		// Si el vehículo ya llegó a su espacio, termina la animación
		if posicionVehiculo.X >= espacio.X {
			break
		}

		// Mueve el vehículo horizontalmente
		posicionVehiculo.X += velocidad

		// Retardo para controlar la velocidad de animación (opcional)
		time.Sleep(16 * time.Millisecond) // aprox. 60 FPS
	}
}

func DibujarVehiculo(win *pixelgl.Window, position pixel.Vec, color pixel.RGBA) {
	imd := imdraw.New(nil)
	imd.Color = color
	imd.Push(position, position.Add(pixel.V(ESPACIO_ANCHO, ESPACIO_ALTO)))
	imd.Rectangle(0)
	imd.Draw(win)
}

func (ev *EstacionamientoView) EstacionarVehiculo(v *models.Vehiculo, espacio int) {
	ev.Estacionamiento.Vehiculos[espacio] = v
}

func (ev *EstacionamientoView) VehiculoSale(v *models.Vehiculo) {
	// Lógica para la salida del vehículo (animación, etc.)
	// ...

	// Actualiza el número de espacios disponibles
	ev.SemaforoEspacio <- struct{}{}
	ev.Estacionamiento.EspaciosDisponibles++
	<-ev.SemaforoEspacio
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
