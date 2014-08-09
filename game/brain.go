package game

import (
	"github.com/nu7hatch/gouuid"
)

type brainMap map[string]*brain

type brain struct {
	strategy strategy
}

type strategy int

const (
	random = iota
)

func (b brainMap) add(ent *uuid.UUID, strat strategy) {
	b[ent.String()] = &brain{strat}
}

func (b brainMap) remove(ent *uuid.UUID) {
	delete(b, ent.String())
}

func (b brainMap) byEnt(ent *uuid.UUID) (*brain, bool) {
	br, ok := b[ent.String()]
	return br, ok
}

func randomBrain(pos position) position {
	return position{}
}
