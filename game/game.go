package game

import (
	"github.com/nu7hatch/gouuid"
	"log"
	"math/rand"
	"time"
)

var trashRand *rand.Rand

var positionsSet positionsMap
var materialsSet materialsMap
var controlledSet controlledMap
var brainSet brainMap
var entitiesSet []string
var reg *registry

type coords [2]int

type Game struct{}

func newUUID() string {
	u4, err := uuid.NewV4()
	if err != nil {
		panic("uuid failed")
	}
	entitiesSet = append(entitiesSet, u4.String())
	return u4.String()
}

func d(n int) int {
	return rand.Intn(n)
}

func trashD(n int) int {
	return trashRand.Intn(n)
}

func pump() {
	for _ = range time.Tick(tickTime * time.Millisecond) {
		controllableSystem.run()
		brainableSystem.run()
		for k, v := range reg.commands {
			v()
			delete(reg.commands, k)
		}
		//reg.publish(buildMessageWorld())
		log.Println("tick")
	}
}

func init() {
	entitiesSet = make([]string, 200000)
	positionsSet = make(positionsMap)
	materialsSet = make(materialsMap)
	controlledSet = make(controlledMap)
	brainSet = make(brainMap)
}

func (g *Game) Init() {
	log.Printf("Starting server with seed: %s", seed)
	rand.Seed(seed)
	trashRand = rand.New(rand.NewSource(rand.Int63()))
	w.buildWorld()
	reg = newRegistry()

	controllableSystem.init()
	brainableSystem.init()

	id := newUUID()
	positionsSet.add(id, 12, 12, 0, 0)
	materialsSet.add(id, cow)
	brainSet.add(id, random)

	// TODO init systems
	go pump()
}
