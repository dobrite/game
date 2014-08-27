package game

import (
	"github.com/larspensjo/Go-simplex-noise/simplexnoise"
)

func noise1(x float64) float64 {
	return simplexnoise.Noise1(x)
}

func noise2(x, y float64) float64 {
	return simplexnoise.Noise2(x, y)
}

func noise3(x, y, z float64) float64 {
	return simplexnoise.Noise3(x, y, z)
}
