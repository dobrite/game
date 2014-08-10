package game

type brainMap map[string]*brain

type strategy int

type brain struct {
	strategy strategy
}

const (
	rock = iota
	random
)

func (b brainMap) add(ent string, strat strategy) {
	b[ent] = &brain{
		strategy: strat,
	}
}

func (b brainMap) remove(ent string) {
	delete(b, ent)
}

func (b brainMap) byEnt(ent string) *brain {
	return b[ent]
}

func randomBrain(delay int) func(*position) {
	return func(pos *position) {
		if trashD(delay) == 0 {
			pos.move(trashD(3)-1, trashD(3)-1)
		}
	}
}

func rockBrain() func(*position) {
	return func(pos *position) {
	}
}
