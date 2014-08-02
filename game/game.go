package game

import (
	"encoding/json"
	"github.com/nu7hatch/gouuid"
	"log"
	"math/rand"
)

var positions positionsMap
var materials materialsMap
var entities []entity

type Game struct {
	positionsMap
	materialsMap
	movable
	strategy
}

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

func makeTile(x int, y int, t materialType) entity {
	ent := newUUID()

	positions.add(ent, x, y)
	materials.add(ent, t)

	return ent
}

func init() {
	positions = make(positionsMap)
	materials = make(materialsMap)
	entities = make([]entity, 200000)
}

func (g *Game) Init() {
	log.Printf("Starting server with seed: %s", seed)
	rand.Seed(seed)
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
