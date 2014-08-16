package game

import (
	"encoding/json"
	"log"
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
	w.buildSpawn(spawnZ, spawnX)
}

func (w *world) buildSpawn(spawnZ, spawnX int) {
	offsetZ := div2(spawnZ)
	offsetX := div2(spawnX)

	for z := 0; z < spawnZ; z++ {
		for x := 0; x < spawnX; x++ {
			var c chunk
			cc := chunkCoords{z - offsetZ, x - offsetX}
			w.chunks[cc] = c.buildChunk(z-offsetZ, x-offsetX)
		}
	}
}

// TODO make w.chunks[cc] getter which will init chunk if not in map
func (w *world) los(cc chunkCoords) [][]*chunk {
	pz := cc[0]
	px := cc[1]
	log.Println(pz, px)
	offsetZ := div2(losZ)
	offsetX := div2(losX)

	straight := make([]*chunk, losZ*losX)
	grid := make([][]*chunk, losZ)
	for z := range grid {
		grid[z] = straight[z*losX : (z+1)*losX]
		for x := range grid[z] {
			cc := chunkCoords{z + pz - offsetZ, x + px - offsetX}
			grid[z][x] = w.chunks[cc]
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
