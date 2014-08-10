package game

var brainableSystem brainable

type brainable struct {
	strategyFunctions map[strategy]func(*position)
	system
}

func (b *brainable) init() {
	b.strategyFunctions = make(map[strategy]func(*position))
	b.strategyFunctions[random] = randomBrain(10)
}

func (b *brainable) run() {
	for k, v := range brainSet {
		p := positionsSet.byEnt(k)
		b.strategyFunctions[v.strategy](p)
		reg.publish(buildMessageItem(k))
	}
}
