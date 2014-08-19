package game

import (
	"encoding/json"
)

type chunk struct {
	a [chunkZ][chunkX][chunkY]string // entity
	c chunkCoords
	json.Marshaler
}

type chunkCoords coords

type messageChunk struct {
	Coords    chunkCoords                          `json:"coords"`
	Materials [chunkZ][chunkX][chunkY]materialType `json:"m"`
}

func (c *chunk) toArray() [chunkZ][chunkX][chunkY]materialType {
	var arr [chunkZ][chunkX][chunkY]materialType
	for z := 0; z < chunkZ; z++ {
		for x := 0; x < chunkX; x++ {
			arr[z][x][0] = d.getMaterial(c.a[z][x][0]).MaterialType
		}
	}
	return arr
}

func (c *chunk) toJSON() *messageChunk {
	return &messageChunk{
		Coords:    c.c,
		Materials: c.toArray(),
	}
}

func (c *chunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.toJSON())
}

func makeChunk(cz int, cx int, cy int) {
	for z := 0; z < chunkZ; z++ {
		for x := 0; x < chunkX; x++ {
			// "tile"
			id := d.newUUID()
			d.addPosition(id, z, x, 0, cz, cx, defaultDepth/chunkY)
			d.addMaterial(id, materialType(die(2)+2))
		}
	}
}
