package game

import (
	"math"
)

const (
	seed           = 0xDEADBEEF
	chunkZ         = 16
	chunkX         = 16
	worldZ         = math.MaxUint32 // 4294967295
	worldX         = math.MaxUint32 // 4294967295
	spawnZ         = 11
	spawnX         = 11
	losZ           = 3
	losX           = 3
	depth          = 32
	defaultDepth   = 16
	maxEntPerCoord = 16
	tickTime       = 200
)

type wireConfig struct {
	Event  string `json:"event"`
	ChunkZ int    `json:"chunkZ"`
	ChunkX int    `json:"chunkX"`
	LosZ   int    `json:"losZ"`
	LosX   int    `json:"losX"`
	Id     string `json:"id"`
}
