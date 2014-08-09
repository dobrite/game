package game

import (
	"github.com/nu7hatch/gouuid"
)

type controlledMap map[*uuid.UUID]controlled

type controlled struct {
	controller controller
}

type controller *uuid.UUID

func (c controlledMap) add(ent *uuid.UUID) {
	c[ent] = controlled{
		controller: ent,
	}
}

func (c controlledMap) remove(ent *uuid.UUID) {
	delete(c, ent)
}
