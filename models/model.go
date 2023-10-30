package models

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Vehiculo struct {
	ID                int
	TiempoEstacionado time.Duration
}

type Estacionamiento struct {
	capacidadMaxima int
	capacidadActual int
	vehiculos       []*Vehiculo
	mu              sync.Mutex
	semCajon        chan bool
	semPorton       chan bool
	ultimoID        int
}

func NewEstacionamiento(capacidad int) *Estacionamiento {
	e := &Estacionamiento{
		capacidadMaxima: capacidad,
		vehiculos:       make([]*Vehiculo, capacidad),
		semCajon:        make(chan bool, capacidad),
		semPorton:       make(chan bool, 1),
	}
	for i := 0; i < capacidad; i++ {
		e.semCajon <- true
	}
	e.semPorton <- true
	return e
}

func (e *Estacionamiento) obtenerNuevoID() int {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.ultimoID++
	return e.ultimoID
}

func (e *Estacionamiento) intentarEntrar(v *Vehiculo) bool {
	e.mu.Lock()
	defer e.mu.Unlock()
	if e.capacidadActual < e.capacidadMaxima {
		e.capacidadActual++
		return true
	}
	return false
}

func (e *Estacionamiento) estacionarVehiculo(v *Vehiculo) {
	<-e.semCajon
	index := e.buscarCajonDisponible()
	if index != -1 {
		e.vehiculos[index] = v
		time.Sleep(v.TiempoEstacionado)
		e.vehiculos[index] = nil
		e.semCajon <- true
	}
}

func (e *Estacionamiento) salidaVehiculo(v *Vehiculo) {
	<-e.semPorton
	e.mu.Lock()
	e.capacidadActual--
	e.mu.Unlock()
	fmt.Printf("Vehículo %d ha salido del estacionamiento.\n", v.ID)
	e.semPorton <- true
}

func procesarVehiculo(estacionamiento *Estacionamiento, v *Vehiculo) {
	<-estacionamiento.semPorton
	if estacionamiento.intentarEntrar(v) {
		estacionamiento.semPorton <- true // Liberar la entrada inmediatamente después de entrar
		fmt.Printf("Vehículo %d ha entrado del estacionamiento.\n", v.ID)
		estacionamiento.estacionarVehiculo(v)
		estacionamiento.salidaVehiculo(v)
	} else {
		estacionamiento.semPorton <- true
		time.Sleep(500 * time.Millisecond) // Esperar y volver a intentar entrar
	}
}

func (e *Estacionamiento) buscarCajonDisponible() int {
	for i, v := range e.vehiculos {
		if v == nil {
			return i
		}
	}
	return -1
}

func Simular(estacionamiento *Estacionamiento) {
	for {
		v := &Vehiculo{
			ID:                estacionamiento.obtenerNuevoID(),
			TiempoEstacionado: time.Duration(rand.Intn(5)+1) * time.Second,
		}
		go procesarVehiculo(estacionamiento, v)
		time.Sleep(time.Duration(rand.Intn(5)+1) * time.Second)
	}
}
