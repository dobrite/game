package game

import (
	"encoding/json"
	"math"
)

var w world

type world struct {
	chunks map[chunkCoords]*chunk
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

func (w *world) init() {
	w.chunks = make(map[chunkCoords]*chunk)
	w.buildSpawn(spawnX, spawnY)
}

func (w *world) buildSpawn(spawnY, spawnX int) {
	offsetY := int(math.Floor(float64(spawnY) / 2))
	offsetX := int(math.Floor(float64(spawnX) / 2))

	for y := 0; y < spawnY; y++ {
		for x := 0; x < spawnX; x++ {
			var c chunk
			cc := chunkCoords{y - offsetY, x - offsetX}
			w.chunks[cc] = c.buildChunk(y-offsetY, x-offsetX)
		}
	}
}

// TODO make w.chunks[cc] getter which will init chunk if not in map
func (w *world) dtodd(cc chunkCoords) [][]*chunk {
	py := cc[0]
	px := cc[1]

	offsetY := int(math.Floor(float64(losY) / 2))
	offsetX := int(math.Floor(float64(losX) / 2))

	base := make([]*chunk, losY*losX)
	outer := make([][]*chunk, losY)
	for y := range outer {
		outer[y] = base[y*losX : (y+1)*losX]
		for x := range outer[y] {
			cc := chunkCoords{y + py - offsetY, x + px - offsetX}
			outer[y][x] = w.chunks[cc]
		}
	}
	return outer
}

func (w *world) toJSON(cc chunkCoords) *worldJSON {
	return &worldJSON{
		Event: "game:world",
		Data:  w.dtodd(cc),
	}
}

//func (w *world) MarshalJSON() ([]byte, error) {
//	log.Println(w)
//	return json.Marshal(w.toJSON())
//}
