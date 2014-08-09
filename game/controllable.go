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
		if p.x+msg.X > chunkX {
			p.cx = p.cx + 1
			p.x = 0
		} else if p.x+msg.X < 0 {
			p.cx = p.cx - 1
			p.x = chunkX
		} else {
			p.x = p.x + msg.X
		}
		if p.y+msg.Y > chunkY {
			p.cy = p.cy + 1
			p.y = 0
		} else if p.y+msg.Y < 0 {
			p.cy = p.cy - 1
			p.y = chunkY
		} else {
			p.y = p.y + msg.Y
		}
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
