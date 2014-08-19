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
	for _, v := range d.allBrains() {
		// TODO join brain to position to get these directly
		p := d.getPosition(v.ID)
		b.strategyFunctions[v.Strategy](p)
		// TODO batch this then send when done
		reg.broadcast(buildMessageItem(v.ID))
	}
}
