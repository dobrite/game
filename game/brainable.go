package game

var brainableSystem brainable

type brainable struct {
	strategyFunctions map[strategy]func(string) bool
	system
}

func (b *brainable) init() {
	b.strategyFunctions = make(map[strategy](func(id string) bool))
	b.strategyFunctions[rock] = rockBrain()
	b.strategyFunctions[random] = randomBrain(10)
}

func (b *brainable) run() {
	for _, v := range d.allBrains() {
		if b.strategyFunctions[v.Strategy](v.ID) {
			// TODO batch this then send when done
			reg.broadcast(buildMessageItem(v.ID))
		}
	}
}
