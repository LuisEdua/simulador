package models

import "sync"

type Parking struct {
	capacity int
	mu       sync.Mutex
	gateFree chan bool
	slots    []slot
}

func NewParking(c int, mu sync.Mutex, gateFree chan bool) *Parking {
	return &Parking{
		capacity: c,
		mu:       mu,
		gateFree: gateFree,
		slots:    generarSlots(),
	}
}
