package models

import (
	"github.com/faiface/pixel"
	"sync"
)

type Vehiculo struct {
	ID       int
	Posicion pixel.Vec
}

var id int

type Estacionamiento struct {
	Capacidad           int
	Vehiculos           []*Vehiculo
	Mutex               sync.Mutex
	Semaforo            chan struct{}
	EspaciosDisponibles int
}

func NuevoEstacionamiento(capacidad int) *Estacionamiento {
	return &Estacionamiento{
		Capacidad:           capacidad,
		Vehiculos:           make([]*Vehiculo, capacidad), // Reservar espacio para los vehículos
		Semaforo:            make(chan struct{}, capacidad),
		EspaciosDisponibles: capacidad,
	}
}

func GenerarNuevoID() int {
	id++
	return id
}

func (e *Estacionamiento) Entrar(v *Vehiculo) bool {
	select {
	case e.Semaforo <- struct{}{}:
		e.Mutex.Lock()
		e.Vehiculos = append(e.Vehiculos, v)
		e.Mutex.Unlock()
		return true
	default:
		return false
	}
}

func (e *Estacionamiento) Salir(v *Vehiculo) {
	for i, vehiculo := range e.Vehiculos {
		if vehiculo == v {
			e.Mutex.Lock()
			e.Vehiculos = append(e.Vehiculos[:i], e.Vehiculos[i+1:]...)
			e.Mutex.Unlock()
			<-e.Semaforo
			return
		}
	}
}
