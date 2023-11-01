package models

import (
	"fmt"
	"github.com/faiface/pixel"
	"math/rand"
)

type Car struct {
	ID    int
	P1    pixel.Vec
	P2    pixel.Vec
	Width float64
	T     int
}

func NewCar(id int, p1, p2 pixel.Vec) *Car {
	return &Car{
		ID:    id,
		P1:    p1,
		P2:    p2,
		Width: 30,
		T:     rand.Intn(5) + 1,
	}
}

func Come(c *Car) {
	fmt.Println("El auto ", c.ID, " acaba de llegar a la posicion", c.P1)
}

func In(p *Parking, c *Car) int {
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

func Out(i int, p *Parking, c *Car) {
	<-p.gateFree
	p.slots[i].free = true
	p.Capacity++
	p.gateFree <- true
	c.P1 = pixel.V(40, 360)
	c.P2 = pixel.V(40, 300)
	fmt.Println("El auto ", c.ID, " ha salido")
}

func Go(c *Car) {
	c.P1 = pixel.V(-30, 0)
	c.P2 = pixel.V(-30, 0)
	c.Width = 0
	fmt.Println("El auto ", c.ID, " se ha ido")
}
