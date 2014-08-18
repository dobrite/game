package game

import (
	"math"
	"math/rand"
)

func div2(op int) int {
	return int(math.Floor(float64(op) / 2))
}

func die(n int) int {
	return rand.Intn(n)
}

func trashD(n int) int {
	return trashRand.Intn(n)
}
