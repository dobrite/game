package game

const (
	seed              = 0xDEADBEEF
	Chunk_x           = 4
	Chunk_y           = 4
	world_x           = 1 // in chunks
	world_y           = 1 // in chunks
	depth             = 32
	default_depth     = 16
	max_ent_per_coord = 16
)

const (
	grass = iota
	dirt
)

type wireConfig struct {
	Chunk_x int `json:"chunk_x"`
	Chunk_y int `json:"chunk_y"`
}
