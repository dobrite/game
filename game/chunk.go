package game

import (
	"encoding/json"
	"github.com/nu7hatch/gouuid"
)

type chunk struct {
	a [chunkY][chunkY][]*uuid.UUID
	c chunkCoords
	json.Marshaler
}

type chunkCoords coords

type item struct {
	Coords       chunkCoords `json:"coords"`
	materialType `json:"mt"`
}

type wireChunk struct {
	Coords    chunkCoords                  `json:"coords"`
	Materials [chunkY][chunkX]materialType `json:"m"`
	//I []item                         `json:"i"`
}

func (c *chunk) buildChunk(cy int, cx int) *chunk {
	for y := 0; y < chunkY; y++ {
		for x := 0; x < chunkX; x++ {
			c.a[y][x] = make([]*uuid.UUID, max_ent_per_coord)

			var mat materialType
			if coinFlip() {
				mat = grass
			} else {
				mat = dirt
			}

			t := makeTile(y, x, cy, cx, mat)
			entities = append(entities, t)
			c.a[y][x][0] = t
		}
	}
	c.c = [2]int{cy, cx}
	return c
}

func (c *chunk) toArray() [chunkY][chunkX]materialType {
	var arr [chunkY][chunkX]materialType
	for y := 0; y < chunkY; y++ {
		for x := 0; x < chunkX; x++ {
			arr[y][x] = materials[c.a[y][x][0]].materialType
		}
	}
	return arr
}

func (c *chunk) allItems() []item {
	var ret []item
	ents := materials.byType(flesh)
	for _, e := range ents {
		position := positions[e.String()]
		i := item{
			Coords:       [2]int{position.y, position.x},
			materialType: flesh,
		}
		ret = append(ret, i)
	}
	return ret
}

func (c *chunk) toWire() *wireChunk {
	return &wireChunk{
		Coords:    c.c,
		Materials: c.toArray(),
		//	I: c.allItems(),
	}
}

func (c *chunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.toWire())
}
