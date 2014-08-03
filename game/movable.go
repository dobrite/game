package game

import (
	"github.com/nu7hatch/gouuid"
)

type movable map[*uuid.UUID]movement

type movement struct {
	direction
	speed int
}

type direction int

const (
	up direction = iota
	down
	left
	right
)
