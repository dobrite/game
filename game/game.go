package game

import (
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
	return u4
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

func pump() {
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
}
