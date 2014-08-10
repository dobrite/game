package game

type positionsMap map[string]*position

type position struct {
	y, x, z, cy, cx int
}

func (p positionsMap) add(ent string, y, x, cy, cx int) {
	p[ent] = &position{
		y:  y,
		x:  x,
		z:  defaultDepth,
		cy: cy,
		cx: cx,
	}
}

func (p positionsMap) remove(ent string) {
	delete(p, ent)
}

func (p positionsMap) byEnt(ent string) *position {
	return p[ent]
}

func (p *position) toWorldCoords() *worldCoords {
	return &worldCoords{
		ChunkCoords: coords{p.cy, p.cx},
		Coords:      coords{p.y, p.x},
	}
}

func (p *position) move(y, x int) {
	if p.x+x > chunkX {
		p.cx = p.cx + 1
		p.x = 0
	} else if p.x+x < 0 {
		p.cx = p.cx - 1
		p.x = chunkX
	} else {
		p.x = p.x + x
	}

	if p.y+y > chunkY {
		p.cy = p.cy + 1
		p.y = 0
	} else if p.y+y < 0 {
		p.cy = p.cy - 1
		p.y = chunkY
	} else {
		p.y = p.y + y
	}
}
