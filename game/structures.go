package game

import (
	"github.com/nu7hatch/gouuid"
)

type Game struct{}

type entity struct {
	*uuid.UUID
}

type tile struct {
	position
	material
}

var entities []entity
var positions map[entity]position
var materials map[entity]material
