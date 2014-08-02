package game

type movable map[entity]movement

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
