package main

import (
	"Simulador/views"

	"github.com/faiface/pixel/pixelgl"
)

func run() {
	views.View()
}

func main() {
	pixelgl.Run(run)
}
