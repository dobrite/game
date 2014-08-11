package game

import (
	"math"
)

const (
	seed           = 0xDEADBEEF
	chunkY         = 16
	chunkX         = 16
	worldY         = math.MaxUint32 // 4294967295
	worldX         = math.MaxUint32 // 4294967295
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
	ChunkY int    `json:"chunkY"`
	ChunkX int    `json:"chunkX"`
	LosY   int    `json:"losY"`
	LosX   int    `json:"losX"`
	Id     string `json:"id"`
}
