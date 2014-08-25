package game

import (
	"log"
)

type strategy int

type brain struct {
	ID       string
	Strategy strategy
}

const (
	rock = iota
	random
)

func (db *db) addBrain(id string, strat strategy) {
	b := &brain{
		ID:       id,
		Strategy: strat,
	}

	if err := db.dbmap.Insert(b); err != nil {
		log.Fatal(err)
	}
}

func (db *db) allBrains() []brain {
	var brains []brain

	if _, err := d.dbmap.Select(&brains, "select id, strategy from brains;"); err != nil {
		log.Fatal(err)
	}

	return brains
}

func randomBrain(delay int) func(*position) {
	return func(p *position) {
		if trashD(delay) == 0 {
			p.move(trashD(3)-1, trashD(3)-1)
			d.setPosition(p)
		}
	}
}

func rockBrain() func(*position) {
	return func(p *position) {}
}
