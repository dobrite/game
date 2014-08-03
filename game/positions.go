package game

import (
	"github.com/nu7hatch/gouuid"
)

type positionsMap map[string]position

type position struct {
	x, y, z, cx, cy int
}

func (p positionsMap) add(ent *uuid.UUID, x int, y int) {
	p[ent.String()] = position{
		x:  x,
		y:  y,
		z:  defaultDepth,
		cx: 0,
		cy: 0,
	}
}

func (p positionsMap) run() {}
