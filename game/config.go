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
	worldY   = 256
	spawnZ   = 9
	spawnX   = 9
	spawnY   = worldY / chunkY
	losZ     = 9
	losX     = 9
	losY     = worldY / chunkY
	tickTime = 200
)
