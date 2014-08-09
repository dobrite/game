package game

import (
	"github.com/nu7hatch/gouuid"
)

var controllableSystem controllable

type controllable struct {
	queue map[string]func()
	system
}

func (c *controllable) enqueue(ent *uuid.UUID, msg messageMove) {
	entStr := ent.String()
	c.queue[entStr] = func() {
		p := positionsSet[entStr]
		p.move(msg.Y, msg.X)
		positionsSet[entStr] = p
		reg.publish(buildMessageItem(ent))
	}
}

func (c *controllable) run() {
	for k, v := range c.queue {
		v()
		delete(c.queue, k)
	}
}
