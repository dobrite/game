package game

import (
	"math"
)

var w world

type world [worldY][worldX]*chunk

type wireWorld struct {
	Event string     `json:"event"`
	Data  [][]*chunk `json:"data"`
}

func (w *world) buildWorld() {
	offsetX := int(math.Floor(worldX / 2))
	offsetY := int(math.Floor(worldY / 2))

	for x := 0; x < worldX; x++ {
		for y := 0; y < worldY; y++ {
			var c chunk
			w[y][x] = c.buildChunk(y-offsetY, x-offsetX)
		}
	}
}

//func (w *world) buildWorld() {
//	offsetX := math.Floor(worldX / 2)
//	offsetY := math.Floor(worldY / 2)
//
//	for x := -offsetX; x < (worldX - offsetX); x++ {
//		for y := -offsetY; y < (worldX - offsetY); y++ {
//			var c chunk
//			w[y][x] = c.buildChunk()
//		}
//	}
//}
