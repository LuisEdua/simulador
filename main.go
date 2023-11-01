package main

import (
	"Simulador/models"
	"Simulador/views"
)

func main() {
	go models.Start()
	views.Show()
}
