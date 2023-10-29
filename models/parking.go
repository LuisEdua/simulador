package models

import (
	"sync"
)

type Vehiculo struct {
	ID int
}

type Estacionamiento struct {
	Capacidad      int
	Vehiculos      []*Vehiculo
	Mutex          sync.Mutex
	Semaforo       chan struct{}
}

func NuevoEstacionamiento(capacidad int) *Estacionamiento {
	return &Estacionamiento{
		Capacidad: capacidad,
		Semaforo:  make(chan struct{}, capacidad),
	}
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
