package models

import "github.com/faiface/pixel"

type slot struct {
	p    pixel.Vec
	free bool
}

func newSlot(p pixel.Vec) slot {
	return slot{
		p:    p,
		free: true,
	}
}

func generarSlots() []slot {
	var slots []slot
	width := 40.0
	space := 10.0

	for x := 105.0; x <= 595; x += width + space {

		xM := (x + (x + width)) / 2
		p := pixel.V(xM, 445)

		slots = append(slots, newSlot(p))
	}

	for x := 105.0; x <= 595; x += width + space {

		xM := (x + (x + width)) / 2
		p := pixel.V(xM, 15+80)

		slots = append(slots, newSlot(p))
	}
	return slots
}
