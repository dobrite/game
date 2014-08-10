package game

import (
	"encoding/json"
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
	ChunkCoords coords `json:"chunkCoords"`
	Coords      coords `json:"coords"`
}

func (w *world) init() {
	w.chunks = make(map[chunkCoords]*chunk)
	w.buildSpawn(spawnX, spawnY)
}

func (w *world) buildSpawn(spawnY, spawnX int) {
	offsetY := div2(spawnY)
	offsetX := div2(spawnX)

	for y := 0; y < spawnY; y++ {
		for x := 0; x < spawnX; x++ {
			var c chunk
			cc := chunkCoords{y - offsetY, x - offsetX}
			w.chunks[cc] = c.buildChunk(y-offsetY, x-offsetX)
		}
	}
}

// TODO make w.chunks[cc] getter which will init chunk if not in map
func (w *world) los(cc chunkCoords) [][]*chunk {
	py := cc[0]
	px := cc[1]

	offsetY := div2(losY)
	offsetX := div2(losX)

	straight := make([]*chunk, losY*losX)
	grid := make([][]*chunk, losY)
	for y := range grid {
		grid[y] = straight[y*losX : (y+1)*losX]
		for x := range grid[y] {
			cc := chunkCoords{y + py - offsetY, x + px - offsetX}
			grid[y][x] = w.chunks[cc]
		}
	}
	return grid
}

func (w *world) toJSON(cc chunkCoords) *worldJSON {
	return &worldJSON{
		Event: "game:los",
		Data:  w.los(cc),
	}
}
