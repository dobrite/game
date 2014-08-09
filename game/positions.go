package game

import (
	"github.com/nu7hatch/gouuid"
)

type positionsMap map[string]position

type position struct {
	y, x, z, cy, cx int
	nil             bool
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

func (p positionsMap) remove(ent *uuid.UUID) {
	delete(p, ent.String())
}

func (p positionsMap) byEnt(ent *uuid.UUID) position {
	if pos, ok := p[ent.String()]; ok {
		return pos
	}
	return position{nil: true}
}

func (p *position) toWorldCoords() *worldCoords {
	return &worldCoords{
		ChunkCoords: coords{p.cy, p.cx},
		Coords:      coords{p.y, p.x},
	}
}
