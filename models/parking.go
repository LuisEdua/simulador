package models

import "sync"

type Parking struct {
	Capacity int
	Mu       sync.Mutex
	gateFree chan bool
	slots    []slot
	Cars     []*Car
}

func NewParking(c int, mu sync.Mutex, gateFree chan bool) *Parking {
	return &Parking{
		Capacity: c,
		Mu:       mu,
		gateFree: gateFree,
		slots:    generarSlots(),
		Cars:     make([]*Car, 0),
	}
}
