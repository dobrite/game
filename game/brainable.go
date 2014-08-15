package game

var brainableSystem brainable

type brainable struct {
	strategyFunctions map[strategy]func(*position)
	system
}

func (b *brainable) init() {
	b.strategyFunctions = make(map[strategy]func(*position))
	b.strategyFunctions[rock] = rockBrain()
	b.strategyFunctions[random] = randomBrain(10)
}

func (b *brainable) run() {
	for k, v := range brainSet {
		p := positionsSet.byEnt(k)
		b.strategyFunctions[v.strategy](p)
		// TODO batch this then send when done
		reg.broadcast(buildMessageItem(k))
	}
}
