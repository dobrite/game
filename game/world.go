package game

import (
	"encoding/json"
	"math"
)

var w world

type world struct {
	chunks [worldY][worldX]*chunk
	json.Marshaler
}

type worldJSON struct {
	Event string     `json:"event"`
	Data  [][]*chunk `json:"data"`
}

type worldCoords struct {
	ChunkCoords coords `json:"chunk_coords"`
	Coords      coords `json:"coords"`
}

func (w *world) buildWorld() {
	offsetX := int(math.Floor(worldX / 2))
	offsetY := int(math.Floor(worldY / 2))

	for x := 0; x < worldX; x++ {
		for y := 0; y < worldY; y++ {
			var c chunk
			w.chunks[y][x] = c.buildChunk(y-offsetY, x-offsetX)
		}
	}
}

func (w *world) dtodd() [][]*chunk {
	base := make([]*chunk, worldY*worldX)
	outer := make([][]*chunk, worldY)
	for i := range outer {
		outer[i] = base[i*worldX : (i+1)*worldX]
		for j := range outer[i] {
			outer[i][j] = w.chunks[j][i]
		}
	}
	return outer
}

func (w *world) toJSON() *worldJSON {
	return &worldJSON{
		Event: "game:world",
		Data:  w.dtodd(),
	}
}

//func (w *world) MarshalJSON() ([]byte, error) {
//	log.Println(w)
//	return json.Marshal(w.toJSON())
//}
