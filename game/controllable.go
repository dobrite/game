package game

var controllableSystem controllable

type controllable struct {
	queue map[string]func()
	system
}

func (c *controllable) init() {
	c.queue = make(map[string]func())
}

func (c *controllable) enqueue(ent string, msg messageMove) {
	c.queue[ent] = func() {
		positionsSet[ent].move(msg.Z, msg.X)
		reg.broadcast(buildMessageItem(ent))
	}
}

func (c *controllable) run() {
	for k, v := range c.queue {
		v()
		delete(c.queue, k)
	}
}
