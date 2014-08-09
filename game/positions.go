package game

import (
	"github.com/nu7hatch/gouuid"
)

type positionsMap map[string]position

type position struct {
	y, x, z, cy, cx int
}

func (p positionsMap) add(ent *uuid.UUID, y, x, cy, cx int) {
	p[ent.String()] = position{
		y:  y,
		x:  x,
		z:  defaultDepth,
		cy: cy,
		cx: cx,
	}
}

func (p positionsMap) run() {}
