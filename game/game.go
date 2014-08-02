package game

import (
	"encoding/json"
	"github.com/nu7hatch/gouuid"
	"log"
	"math/rand"
)

func newUUID() entity {
	u4, err := uuid.NewV4()
	if err != nil {
		panic("uuid failed")
	}
	return entity{u4}
}

func coinFlip() bool {
	if rand.Float32() <= 0.5 {
		return false
	} else {
		return true
	}
}

func makeTile(x int, y int, t int) entity {
	ent := newUUID()

	positions[ent] = position{
		x: x,
		y: y,
		z: default_depth,
	}

	materials[ent] = material{
		type_: t,
	}

	return ent
}

// TODO turn this into init?
func initialize() {
	positions = make(map[entity]position)
	materials = make(map[entity]material)
	entities = make([]entity, 2000000)
}

func (g *Game) Init() {
	log.Printf("Starting server with seed: %s", seed)
	rand.Seed(seed)
	initialize()
	w.buildWorld()

	wc := &wireConfig{
		Chunk_x: Chunk_x,
		Chunk_y: Chunk_y,
	}

	c, _ := json.Marshal(wc)
	log.Println(string(c))
	j, _ := json.Marshal(w[0][0])
	log.Println(string(j))
}
