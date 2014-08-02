package game

type brain struct {
	strategy
}

type strategy int

const (
	player strategy = iota
	random
)
