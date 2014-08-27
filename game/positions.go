package game

import (
	"log"
)

type positionsMap map[string]*position

type position struct {
	ID                  string
	Z, X, Y, Cz, Cx, Cy int
}

func (db *db) addPosition(id string, z, x, y, cz, cx, cy int) {
	p := &position{
		ID: id,
		Z:  z,
		X:  x,
		Y:  y,
		Cz: cz,
		Cx: cx,
		Cy: cy,
	}

	if err := db.dbmap.Insert(p); err != nil {
		log.Fatal(err)
	}
}

func (db *db) getPosition(id string) *position {
	obj, err := d.dbmap.Get(position{}, id)
	if err != nil {
		log.Fatal(err)
	}
	return obj.(*position)
}

func (db *db) setPosition(p *position) {
	d.dbmap.Update(p)
}

func (p *position) toWorldCoords() *worldCoords {
	return &worldCoords{
		ChunkCoords: coords{p.Cz, p.Cx},
		Coords:      coords{p.Z, p.X},
	}
}

func (p *position) move(z, x int) {
	if p.X+x >= chunkX {
		p.Cx += 1
		p.X = 0
	} else if p.X+x < 0 {
		p.Cx -= 1
		p.X = chunkX - 1
	} else {
		p.X += x
	}

	if p.Z+z >= chunkZ {
		p.Cz += 1
		p.Z = 0
	} else if p.Z+z < 0 {
		p.Cz -= 1
		p.Z = chunkZ - 1
	} else {
		p.Z += z
	}
}
