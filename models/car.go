package models

import (
	"fmt"
	"github.com/faiface/pixel"
	"math/rand"
	"sync"
	"time"
)

type Car struct {
	ID    int
	P1    pixel.Vec
	P2    pixel.Vec
	Width int
	t     int
}

func NewCar(id int, p1, p2 pixel.Vec) *Car {
	return &Car{
		ID:    id,
		P1:    p1,
		P2:    p2,
		Width: 30,
		t:     rand.Intn(5) + 1,
	}
}

func come(c *Car) {
	fmt.Println("El auto ", c.ID, " acaba de llegar a la posicion", c.P1)
}

func in(p *Parking, c *Car) int {
	var freeNum int
	<-p.gateFree

	for i := range p.slots {
		if p.slots[i].free {
			c.P1 = p.slots[i].p
			c.P2 = pixel.V(c.P1.X, c.P1.Y-60)
			freeNum = i
			break
		}
	}
	fmt.Println("El auto ", c.ID, "esta estacionado en la posicion ", p.slots[freeNum].p)
	p.slots[freeNum].free = false
	p.gateFree <- true

	return freeNum
}

func out(i int, p *Parking, c *Car) {
	<-p.gateFree
	c.P1 = pixel.V(-30, 0)
	c.P2 = c.P1
	c.Width = 0
	p.slots[i].free = true
	p.gateFree <- true
	fmt.Println("El auto ", c.ID, " se ha ido")
}

func notSpace(c *Car) {
	c.P1 = pixel.V(-30, 0)
	c.P2 = c.P1
	c.Width = 0
	fmt.Println("El auto ", c.ID, " se ha ido")
}

func CarRoutine(c *Car, p *Parking) {
	p.mu.Lock()
	come(c)
	if p.capacity > 0 {
		i := in(p, c)
		p.mu.Unlock()
		time.Sleep(time.Duration(c.t) * time.Second)
		out(i, p, c)
	} else {
		notSpace(c)
		p.mu.Unlock()
	}
}

func Start() {
	gateFree := make(chan bool, 1)
	mu := sync.Mutex{}

	p := NewParking(20, mu, gateFree)

	gateFree <- true

	i := 1

	for {
		c := NewCar(i, pixel.V(40, 180), pixel.V(40, 120))
		go CarRoutine(c, p)
		time.Sleep(500 * time.Millisecond)
		i++
	}
}
