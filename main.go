package main

import (
	"Simulador/models"
)

func main() {
	estacionamiento := models.NewEstacionamiento(20)
	models.Simular(estacionamiento)
}