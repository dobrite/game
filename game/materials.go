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
	grass materialType = iota
	dirt
	flesh
)

func (m materialsMap) add(ent *uuid.UUID, t materialType) {
	m[ent] = material{
		materialType: t,
	}
}

func (m materialsMap) byType(t materialType) []*uuid.UUID {
	var ret []*uuid.UUID
	for k, v := range materials {
		if v.materialType == t {
			ret = append(ret, k)
		}
	}
	return ret
}
