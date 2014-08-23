package game

import (
	"github.com/coopernurse/gorp"
	"log"
	"math/rand"
	"time"
)

var w world

var trashRand *rand.Rand

var reg *registry

type coords [3]int

type Game struct{}

func pump() {
	for _ = range time.Tick(tickTime * time.Millisecond) {
		t := time.Now()
		controllableSystem.run()
		brainableSystem.run()
		for k, v := range reg.commands {
			v()
			delete(reg.commands, k)
		}
		log.Printf("%s", time.Since(t))
	}
}

func (g *Game) populate() {
	id := d.newUUID()
	d.addPosition(id, 0, 0, 0, 0, 0, defaultDepth/chunkY)
	d.addMaterial(id, cow)
	d.addBrain(id, random)

	id = d.newUUID()
	d.addPosition(id, 0, 0, 0, 2, 1, defaultDepth/chunkY)
	d.addMaterial(id, pig)
	d.addBrain(id, random)

	id = d.newUUID()
	d.addPosition(id, 0, 0, 0, 0, 0, defaultDepth/chunkY)
	d.addMaterial(id, stone)
	d.addBrain(id, random)

	id = d.newUUID()
	d.addPosition(id, 0, 0, 0, 1, 1, defaultDepth/chunkY)
	d.addMaterial(id, stone)
	d.addBrain(id, random)

	id = d.newUUID()
	d.addPosition(id, 0, 0, 0, 0, 1, defaultDepth/chunkY)
	d.addMaterial(id, stone)
	d.addBrain(id, random)

	id = d.newUUID()
	d.addPosition(id, 0, 0, 0, 1, 0, defaultDepth/chunkY)
	d.addMaterial(id, stone)
	d.addBrain(id, random)
}

func (g *Game) Init(dbmap *gorp.DbMap) {
	d = &db{dbmap: dbmap}
	d.init()

	log.Printf("Starting server with seed: %s", seed)
	rand.Seed(seed)
	trashRand = rand.New(rand.NewSource(rand.Int63()))

	g.populate()
	w.init()

	reg = newRegistry()

	controllableSystem.init()
	brainableSystem.init()

	go pump()
}
