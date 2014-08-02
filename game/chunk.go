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
	M []int `json:"m"`
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

//func (c *chunk) toProto() *game.Col {
//	var col [chunk_y]*game.Row
//	for y := 0; y < chunk_y; y++ {
//		var row [chunk_x]int32
//		for x := 0; x < chunk_x; x++ {
//			row[x] = materials[c[y][x][0]].type_
//		}
//		r := &game.Row{
//			X: row[:],
//		}
//		col[y] = r
//	}
//	cl := &game.Col{
//		Rows: col[:],
//	}
//	return cl
//}

func (c *chunk) toArray() []int {
	var arr [Chunk_y * Chunk_x]int
	for y := 0; y < Chunk_y; y++ {
		for x := 0; x < Chunk_x; x++ {
			arr[(y*Chunk_y)+x] = materials[c.a[y][x][0]].type_
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
