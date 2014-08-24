package game

import (
	"database/sql"
	"encoding/json"
	"log"
	"strings"
)

type chunk struct {
	a [chunkZ][chunkX][chunkY]string // entity
	c chunkCoords
	json.Marshaler
}

type chunkCoords coords

type messageChunk struct {
	Coords    chunkCoords                          `json:"coords"`
	Materials [chunkZ][chunkX][chunkY]materialType `json:"m"`
}

func (c *chunk) toArray() [chunkZ][chunkX][chunkY]materialType {
	var arr [chunkZ][chunkX][chunkY]materialType
	for z := 0; z < chunkZ; z++ {
		for x := 0; x < chunkX; x++ {
			arr[z][x][0] = d.getMaterial(c.a[z][x][0]).MaterialType
		}
	}
	return arr
}

func (c *chunk) toJSON() *messageChunk {
	return &messageChunk{
		Coords:    c.c,
		Materials: c.toArray(),
	}
}

func (c *chunk) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.toJSON())
}

func makeChunk(cz int, cx int, cy int) {
	log.Printf("Making chunk: %d, %d, %d", cz, cx, cy)
	for z := 0; z < chunkZ; z++ {
		for x := 0; x < chunkX; x++ {
			// "tile"
			id := d.newUUID()
			d.addPosition(id, z, x, 0, cz, cx, defaultDepth/chunkY)
			d.addMaterial(id, materialType(die(2)+2))
		}
	}
}

type sqlStraightChunk struct {
	Arr string
	sql.Scanner
}

func (ssc *sqlStraightChunk) Scan(src interface{}) error {
	s := string(src.([]uint8))
	s = strings.Replace(s, "{", "[", -1)
	s = strings.Replace(s, "}", "]", -1)
	ssc.Arr = s
	return nil
}

func (ssc *sqlStraightChunk) straightChunk() {
	row := d.dbmap.Db.QueryRow(`SELECT array_agg(material_type)
	    FROM materials WHERE id IN
		(SELECT id FROM positions WHERE
			cx = 1 AND
			cz = 1 AND
			cy = 4
			ORDER BY x,z,y);`)
	err := row.Scan(&ssc)
	switch {
	case err == sql.ErrNoRows:
		log.Println("no rows")
	case err != nil:
		log.Fatal(err)
	default:
		log.Println(ssc.Arr)
	}
}
