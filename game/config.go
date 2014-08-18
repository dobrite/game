package game

import (
	"math"
)

const (
	seed           = 0xDEADBEEF
	chunkZ         = 16
	chunkX         = 16
	chunkY         = 16
	worldZ         = math.MaxInt32 // 2147483647
	worldX         = math.MaxInt32 // 2147483647
	worldY         = 256
	spawnZ         = 3
	spawnX         = 3
	spawnY         = 3 // TODO decide on height
	losZ           = 3
	losX           = 3
	losY           = 3
	defaultDepth   = 64
	maxEntPerCoord = 16
	tickTime       = 200
)
