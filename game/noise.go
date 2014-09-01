package game

import (
	"github.com/larspensjo/Go-simplex-noise/simplexnoise"
	"image"
	"image/jpeg"
	"math"
	"os"
)

const (
	dx         = 48
	dy         = 48
	scale      = 5000.0
	octaves    = 128
	lacunarity = 2.0
)

type noiseFuncs struct {
	simplexFBM2 func(float64, float64) float64
}

var noise *noiseFuncs

func init() {
	noise = &noiseFuncs{
		simplexFBM2: getSimplexFBM2(5000.0, 2.0, 1.0, 128),
	}
}

func noise1(x float64) float64 {
	return simplexnoise.Noise1(x)
}

func noise2(x, y float64) float64 {
	return simplexnoise.Noise2(x, y)
}

func noise3(x, y, z float64) float64 {
	return simplexnoise.Noise3(x, y, z)
}

// clamp -1.0 to 1.0
func clamp(val float64) float64 {
	return math.Max(-1, math.Min(1, val))
}

// convert to between 0 and 1
func scaleToOne(val float64) float64 {
	return (1 + val) / 2
}

func getSimplexFBM2(scale, lacunarity, exp float64, octaves int) func(float64, float64) float64 {
	return func(x, y float64) float64 {
		scaledx := x / scale
		scaledy := y / scale
		result := 0.0

		for i := 0; i <= octaves; i++ {
			exponent := 1.0 / exp
			exp *= lacunarity
			result += noise2(scaledx, scaledy) * exponent
			x *= lacunarity
			y *= lacunarity
		}

		return result
	}
}

func getHeightmap2(val float64) uint8 {
	val = clamp(val)
	val = scaleToOne(val)
	return uint8(val * 255)
}

func genSimplexFBM2(dx, dy int) [][]uint8 {
	simplexFBM2 := getSimplexFBM2(5000.0, 2.0, 1.0, 128)
	pixels := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		pixels[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			pixels[y][x] = getHeightmap2(simplexFBM2(float64(x), float64(y)))
		}
	}
	return pixels
}

func createImage(data [][]uint8) image.Image {
	m := image.NewRGBA(image.Rect(0, 0, dx, dy))
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			v := data[y][x]
			i := y*m.Stride + x*4
			m.Pix[i] = v
			m.Pix[i+1] = v
			m.Pix[i+2] = v
			m.Pix[i+3] = 255
		}
	}
	return m
}

func writeImage(filename string, m image.Image) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	err = jpeg.Encode(f, m, nil)
	if err != nil {
		panic(err)
	}
}
