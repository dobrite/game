package game

import (
	"github.com/nu7hatch/gouuid"
)

type materialsMap map[*uuid.UUID]material

type material struct {
	materialType
}

type materialType int

const (
	nothing materialType = iota
	air
	dirt
	grass
	water
	flesh
)

func (m materialsMap) add(ent *uuid.UUID, t materialType) {
	m[ent] = material{
		materialType: t,
	}
}

func (m materialsMap) remove(ent *uuid.UUID) {
	delete(m, ent)
}

func (m materialsMap) byType(t materialType) []*uuid.UUID {
	var ret []*uuid.UUID
	for k, v := range materialsSet {
		if v.materialType == t {
			ret = append(ret, k)
		}
	}
	return ret
}

func (m materialsMap) byEnt(ent *uuid.UUID) material {
	if mat, ok := m[ent]; ok {
		return mat
	}
	return material{nothing}
}
