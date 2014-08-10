package game

import (
	"encoding/json"
)

type chunk struct {
	a [chunkY][chunkY][]string
	c chunkCoords
	json.Marshaler
}

type chunkCoords coords

type wireChunk struct {
	Coords    chunkCoords                  `json:"coords"`
	Materials [chunkY][chunkX]materialType `json:"m"`
}

func (c *chunk) buildChunk(cy int, cx int) *chunk {
	for y := 0; y < chunkY; y++ {
		for x := 0; x < chunkX; x++ {
			c.a[y][x] = make([]string, maxEntPerCoord)
			c.a[y][x][0] = makeTile(y, x, cy, cx, materialType(d(2)+2))
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
	}
}

func (c *chunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.toWire())
}
