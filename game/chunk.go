package game

import (
	"encoding/json"
)

//type chunk [chunk_y][chunk_x][max_ent_per_coord]entity
type chunk struct {
	a [Chunk_y][Chunk_x][max_ent_per_coord]entity
	json.Marshaler
}

type wireChunk struct {
	M []materialType `json:"m"`
}

func (c *chunk) buildChunk() *chunk {
	for y := 0; y < Chunk_y; y++ {
		for x := 0; x < Chunk_x; x++ {
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

func (c *chunk) toArray() []materialType {
	var arr [Chunk_y * Chunk_x]materialType
	for y := 0; y < Chunk_y; y++ {
		for x := 0; x < Chunk_x; x++ {
			arr[(y*Chunk_y)+x] = materials[c.a[y][x][0]].materialType
		}
	}
	return arr[:]
}

func (c *chunk) toWire() *wireChunk {
	// TODO send these at start of game
	return &wireChunk{
		M: c.toArray(),
	}
}

func (c *chunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.toWire())
}
