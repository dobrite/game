package game

import (
	"math"
)

const (
	seed           = 0xDEADBEEF
	chunkZ         = 16
	chunkX         = 16
	chunkY         = 16
	worldZ         = math.MaxUint32 // 4294967295
	worldX         = math.MaxUint32 // 4294967295
	worldY         = 256
	spawnZ         = 11
	spawnX         = 11
	spawnY         = 11 // TODO decide on height
	losZ           = 3
	losX           = 3
	losY           = 3
	defaultDepth   = 64
	maxEntPerCoord = 16
	tickTime       = 200
)
