package game

var w world

type world [world_y][world_x]*chunk

type wireWorld struct {
	Event string `json:"event"`
	Data  *chunk `json:"data"`
}

func (w *world) buildWorld() {
	for x := 0; x < world_x; x++ {
		for y := 0; y < world_x; y++ {
			var c chunk
			w[y][x] = c.buildChunk()
		}
	}
}
