package game

type materialsMap map[string]material

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
	cow
)

func (m materialsMap) add(ent string, t materialType) {
	m[ent] = material{
		materialType: t,
	}
}

func (m materialsMap) remove(ent string) {
	delete(m, ent)
}

func (m materialsMap) byType(t materialType) []string {
	var ret []string
	for k, v := range materialsSet {
		if v.materialType == t {
			ret = append(ret, k)
		}
	}
	return ret
}

func (m materialsMap) byEnt(ent string) material {
	return m[ent]
}
