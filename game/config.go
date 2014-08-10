package game

import (
	"math"
)

const (
	seed           = 0xDEADBEEF
	chunkX         = 16
	chunkY         = 16
	worldX         = math.MaxUint32 // 4294967295
	worldY         = math.MaxUint32 // 4294967295
	spawnY         = 11
	spawnX         = 11
	losY           = 3
	losX           = 3
	depth          = 32
	defaultDepth   = 16
	maxEntPerCoord = 16
	tickTime       = 200
)

type wireConfig struct {
	Event  string `json:"event"`
	ChunkY int    `json:"chunk_y"`
	ChunkX int    `json:"chunk_x"`
	LosY   int    `json:"los_y"`
	LosX   int    `json:"los_x"`
	Id     string `json:"id"`
}
