package game

import ()

const (
	seed              = 0xDEADBEEF
	chunk_x           = 16
	chunk_y           = 16
	world_x           = 3 // in chunks
	world_y           = 3 // in chunks
	depth             = 32
	defaultDepth      = 16
	max_ent_per_coord = 16
)

type wireConfig struct {
	Event   string `json:"event"`
	Chunk_x int    `json:"chunk_x"`
	Chunk_y int    `json:"chunk_y"`
	World_x int    `json:"world_x"`
	World_y int    `json:"world_y"`
	Id      string `json:"id"`
}
