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

//type item struct {
//	Coords       chunkCoords `json:"coords"`
//	materialType `json:"mt"`
//}

type wireChunk struct {
	Coords    chunkCoords                  `json:"coords"`
	Materials [chunkY][chunkX]materialType `json:"m"`
	//	Items     []item                       `json:"i"`
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
	c.c = chunkCoords{cy, cx}
	return c
}

func (c *chunk) toArray() [chunkY][chunkX]materialType {
	var arr [chunkY][chunkX]materialType
	for y := 0; y < chunkY; y++ {
		for x := 0; x < chunkX; x++ {
			arr[y][x] = materialsSet[c.a[y][x][0]].materialType
		}
	}
	return arr
}

func (c *chunk) toWire() *wireChunk {
	return &wireChunk{
		Coords:    c.c,
		Materials: c.toArray(),
		//Items:     c.allItems(),
	}
}

func (c *chunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.toWire())
}
