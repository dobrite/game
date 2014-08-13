package game

type positionsMap map[string]*position

type position struct {
	z, x, y, cz, cx int
}

func (p positionsMap) add(ent string, z, x, cz, cx int) {
	p[ent] = &position{
		z:  z,
		x:  x,
		y:  defaultDepth,
		cz: cz,
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
		ChunkCoords: coords{p.cz, p.cx},
		Coords:      coords{p.z, p.x},
	}
}

func (p *position) move(z, x int) {
	if p.x+x >= chunkX {
		p.cx = p.cx + 1
		p.x = 0
	} else if p.x+x < 0 {
		p.cx = p.cx - 1
		p.x = chunkX - 1
	} else {
		p.x = p.x + x
	}

	if p.z+z >= chunkZ {
		p.cz = p.cz + 1
		p.z = 0
	} else if p.z+z < 0 {
		p.cz = p.cz - 1
		p.z = chunkZ - 1
	} else {
		p.z = p.z + z
	}
}
