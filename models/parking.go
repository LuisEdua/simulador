package models

import "sync"

type Estacionamiento struct {
    Capacidad int
    Ocupados  int
    mu        sync.Mutex
}

func (e *Estacionamiento) Entrar() bool {
    e.mu.Lock()
    defer e.mu.Unlock()

    if e.Ocupados < e.Capacidad {
        e.Ocupados++
        return true
    }
    return false
}

func (e *Estacionamiento) Salir() {
    e.mu.Lock()
    e.Ocupados--
    e.mu.Unlock()
}
