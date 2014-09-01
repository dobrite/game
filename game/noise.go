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

func noise1(x float64) float64 {
	return simplexnoise.Noise1(x)
}

func noise2(x, y float64) float64 {
	return simplexnoise.Noise2(x, y)
}

func noise3(x, y, z float64) float64 {
	return simplexnoise.Noise3(x, y, z)
}

//func main() {
//	pixels := Pic(dx, dy)
//	pixels = setMinToZero(pixels)
//	img := Create(pixels)
//	WriteImage("img.jpg", img)
//}

func setMinToZero(pixels [][]uint8) [][]uint8 {
	min := uint8(255)
	for y := range pixels {
		for x := range pixels[y] {
			if pixels[y][x] < min {
				min = pixels[y][x]
			}

		}
	}
	for y := range pixels {
		for x := range pixels[y] {
			pixels[y][x] = pixels[y][x] - min
		}
	}
	return pixels
}

func getSimplexFBM(x, y float64, octaves int, lacunarity float64) float64 {
	result := 0.0
	f := 1.0
	var exponents [129]float64

	for i := 0; i <= octaves; i++ {
		exponents[i] = 1.0 / f
		f *= lacunarity
		result += noise2(x, y) * exponents[i]
		x *= lacunarity
		y *= lacunarity
	}

	// clamp -1.0 to 1.0
	ret := math.Max(-1, math.Min(1, result))
	return ret
}

func gen2DSimplexFBM(dx, dy int) [][]uint8 {
	pixels := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		pixels[y] = make([]uint8, dx)
		for x := 0; x < dx; x++ {
			scaledx := float64(x) / scale
			scaledy := float64(y) / scale
			val := getSimplexFBM(scaledx, scaledy, octaves, lacunarity)
			val = (1 + val) / 2
			pixels[y][x] = uint8(val * 255)
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

func writeImage(n string, m image.Image) {
	f, err := os.OpenFile(n, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	err = jpeg.Encode(f, m, nil)
	if err != nil {
		panic(err)
	}
}
