package game

import (
	"encoding/json"
	"github.com/nu7hatch/gouuid"
)

type chunk struct {
	a [Chunk_y][Chunk_x][]*uuid.UUID
	json.Marshaler
}

type item struct {
	Coords       [2]int `json:"coords"`
	materialType `json:"mt"`
}

type wireChunk struct {
	M [Chunk_y][Chunk_x]materialType `json:"m"`
	I []item                         `json:"i"`
}

func (c *chunk) buildChunk() *chunk {
	for y := 0; y < Chunk_y; y++ {
		for x := 0; x < Chunk_x; x++ {
			c.a[y][x] = make([]*uuid.UUID, max_ent_per_coord)
			if coinFlip() {
				t := makeTile(x, y, grass)
				entities = append(entities, t)
				c.a[y][x][0] = t
			} else {
				t := makeTile(x, y, dirt)
				entities = append(entities, t)
				c.a[y][x][0] = t
			}
		}
	}
	return c
}

func (c *chunk) toArray() [Chunk_y][Chunk_x]materialType {
	var arr [Chunk_y][Chunk_x]materialType
	for y := 0; y < Chunk_y; y++ {
		for x := 0; x < Chunk_x; x++ {
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
	// TODO send these at start of game
	return &wireChunk{
		M: c.toArray(),
		I: c.allItems(),
	}
}

func (c *chunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.toWire())
}
