package game

import (
	"github.com/nu7hatch/gouuid"
	"log"
	"math/rand"
	"time"
)

var positions positionsMap
var materials materialsMap
var entities []*uuid.UUID
var reg *registry

type coords [2]int

type Game struct {
	positionsMap
	materialsMap
	movable
	strategy
}

func newUUID() *uuid.UUID {
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

func makeTile(y, x, cy, cx int, t materialType) *uuid.UUID {
	ent := newUUID()

	positions.add(ent, y, x, cy, cx)
	materials.add(ent, t)

	return ent
}

func pump() {
	for _ = range time.Tick(1000 * time.Millisecond) {
		for k, v := range reg.commands {
			v()
			delete(reg.commands, k)
		}
		reg.publish(buildMessageWorld())
		log.Println("tick")
	}
}

func init() {
	positions = make(positionsMap)
	materials = make(materialsMap)
	entities = make([]*uuid.UUID, 200000)
}

func (g *Game) Init() {
	log.Printf("Starting server with seed: %s", seed)
	rand.Seed(seed)
	w.buildWorld()
	reg = newRegistry()
	go pump()
}
