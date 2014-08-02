package game

type positionsMap map[entity]position

type position struct {
	x int
	y int
	z int
}

func (p positionsMap) add(ent entity, x int, y int) {
	p[ent] = position{
		x: x,
		y: y,
		z: default_depth,
	}
}

func (p positionsMap) run() {}
