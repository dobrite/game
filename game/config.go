package game

const (
	seed              = 0xDEADBEEF
	Chunk_x           = 16
	Chunk_y           = 16
	world_x           = 1 // in chunks
	world_y           = 1 // in chunks
	depth             = 32
	default_depth     = 16
	max_ent_per_coord = 16
)

type wireConfig struct {
	Chunk_x int `json:"chunk_x"`
	Chunk_y int `json:"chunk_y"`
}
