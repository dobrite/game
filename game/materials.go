package game

type materialsMap map[entity]material

type material struct {
	materialType
}

type materialType int

const (
	grass materialType = iota
	dirt
	flesh
)

func (m materialsMap) add(ent entity, t materialType) {
	m[ent] = material{
		materialType: t,
	}
}
