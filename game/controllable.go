package game

var controllableSystem controllable

type controllable struct {
	queue map[string]func()
	system
}

func (c *controllable) init() {
	c.queue = make(map[string]func())
}

func (c *controllable) enqueue(id string, msg messageMove) {
	c.queue[id] = func() {
		p := d.getPosition(id)
		p.move(msg.Z, msg.X)
		d.setPosition(p)
		reg.broadcast(buildMessageItem(id))
	}
}

func (c *controllable) run() {
	for k, v := range c.queue {
		v()
		delete(c.queue, k)
	}
}
