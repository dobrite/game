package game

import (
	"math"
)

const (
	seed     = 0xDEADBEEF
	chunkZ   = 16
	chunkX   = 16
	chunkY   = 16
	worldZ   = math.MaxInt32 // 2147483647
	worldX   = math.MaxInt32 // 2147483647
	worldY   = 64
	spawnZ   = 3
	spawnX   = 3
	spawnY   = worldY / chunkY
	losZ     = 3
	losX     = 3
	losY     = worldY / chunkY
	tickTime = 200
)
