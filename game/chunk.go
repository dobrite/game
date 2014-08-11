package game

import (
	"encoding/json"
)

type chunk struct {
	a [chunkZ][chunkX][]string
	c chunkCoords
	json.Marshaler
}

type chunkCoords coords

type wireChunk struct {
	Coords    chunkCoords                  `json:"coords"`
	Materials [chunkZ][chunkX]materialType `json:"m"`
}

func (c *chunk) buildChunk(cz int, cx int) *chunk {
	for z := 0; z < chunkZ; z++ {
		for x := 0; x < chunkX; x++ {
			c.a[z][x] = make([]string, maxEntPerCoord)
			c.a[z][x][0] = makeTile(z, x, cz, cx, materialType(d(2)+2))
		}
	}
	c.c = chunkCoords{cz, cx}
	return c
}

func (c *chunk) toArray() [chunkZ][chunkX]materialType {
	var arr [chunkZ][chunkX]materialType
	for z := 0; z < chunkZ; z++ {
		for x := 0; x < chunkX; x++ {
			arr[z][x] = materialsSet[c.a[z][x][0]].materialType
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
