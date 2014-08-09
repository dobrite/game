package game

import ()

const (
	seed              = 0xDEADBEEF
	chunkX            = 16
	chunkY            = 16
	worldX            = 3 // in chunks, odd
	worldY            = 3 // in chunks, odd
	depth             = 32
	defaultDepth      = 16
	max_ent_per_coord = 16
	tickTime          = 50
)

type wireConfig struct {
	Event  string `json:"event"`
	ChunkX int    `json:"chunk_x"`
	ChunkY int    `json:"chunk_y"`
	WorldX int    `json:"world_x"`
	WorldY int    `json:"world_y"`
	Id     string `json:"id"`
}
