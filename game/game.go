package game

import (
	"github.com/nu7hatch/gouuid"
	"log"
	"math/rand"
	"time"
)

var positionsSet positionsMap
var materialsSet materialsMap
var controlledSet controlledMap
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

func pump() {
	for _ = range time.Tick(tickTime * time.Millisecond) {
		controllableSystem.run()
		for k, v := range reg.commands {
			v()
			delete(reg.commands, k)
		}
		//reg.publish(buildMessageWorld())
		log.Println("tick")
	}
}

func init() {
	positionsSet = make(positionsMap)
	materialsSet = make(materialsMap)
	entities = make([]*uuid.UUID, 200000)
	controlledSet = make(controlledMap)
	controllableSystem.queue = make(map[string]func())
}

func (g *Game) Init() {
	log.Printf("Starting server with seed: %s", seed)
	rand.Seed(seed)
	w.buildWorld()
	reg = newRegistry()
	//init systems
	go pump()
}
