package game

type world struct{}

type chunkJSON struct {
	Event  string           `json:"event"`
	Array  [][]materialType `json:"array"`
	Coords coords           `json:"coords"`
}

type worldCoords struct {
	ChunkCoords coords `json:"cc"`
	Coords      coords `json:"c"`
}

func (w *world) init() {
	w.buildSpawn(spawnZ, spawnX)
}

func (w *world) buildSpawn(spawnZ, spawnX int) {
	offsetZ := div2(spawnZ)
	offsetX := div2(spawnX)

	for z := 0; z < spawnZ; z++ {
		for x := 0; x < spawnX; x++ {
			makeChunk(z-offsetZ, x-offsetX, defaultDepth/chunkY)
		}
	}
}

func sendLos(session *session, cc chunkCoords) {
	offsetZ := div2(losZ)
	offsetX := div2(losX)
	for z := 0; z < losZ; z++ {
		for x := 0; x < losX; x++ {
			cc := chunkCoords{cc[0] + z - offsetZ, cc[1] + x - offsetX, defaultDepth / chunkY}
			reg.send(session, buildMessageChunk(cc))
		}
	}
}
