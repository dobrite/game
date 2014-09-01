package game

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
)

type sqlStraightChunk struct {
	Arr  []materialType
	Grid [][][]materialType
	sql.Scanner
}

type chunkCoords coords

func (c *chunkCoords) toJSON() *messageChunk {
	var s *sqlStraightChunk
	ssc := s.New()
	ssc.straightChunk(*c)
	ssc.dddChunk()

	return &messageChunk{
		Event:     "game:chunk",
		Coords:    *c,
		Materials: ssc.Grid,
	}
}

func makeChunk(cz int, cx int, cy int) {
	log.Printf("Making chunk: %d, %d, %d", cz, cx, cy)
	for z := 0; z < chunkZ; z++ {
		for x := 0; x < chunkX; x++ {
			// "tile"
			id := d.newUUID()
			n := noise.simplexFBM2(float64((cz*chunkZ)+z), float64((cx*chunkX)+x))
			d.addPosition(id, z, x, int(getHeightmap2(n)), cz, cx, defaultDepth/chunkY)
			d.addMaterial(id, 2)
		}
	}
}

func (ssc *sqlStraightChunk) New() *sqlStraightChunk {
	return &sqlStraightChunk{
		Arr:  make([]materialType, chunkZ*chunkX*chunkY),
		Grid: make([][][]materialType, chunkY),
	}
}

func (ssc *sqlStraightChunk) Convert(str string) error {
	ss := strings.Split(str[1:len(str)-1], ",")
	for i, s := range ss {
		n, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		ssc.Arr[i] = materialType(n)
	}
	return nil
}

func (ssc *sqlStraightChunk) dddChunk() {
	// z -> x -> y
	for z := range ssc.Grid {
		ssc.Grid[z] = make([][]materialType, chunkX)
		for x := range ssc.Grid[z] {
			ssc.Grid[z][x] = ssc.Arr[(z*chunkZ*chunkX)+(x*chunkX) : (z*chunkZ*chunkX)+((x+1)*chunkX)]
		}
	}
}

func (ssc *sqlStraightChunk) straightChunk(cc chunkCoords) {
	row := d.dbmap.Db.QueryRow(`
    SELECT array_agg(material_type)
    FROM
    (SELECT
      (CASE
        WHEN positions.id IS NULL
        THEN 0
        ELSE materials.material_type
      END)
    FROM positions
    JOIN materials
      ON (positions.id = materials.id
      OR positions.id IS NULL)
      AND materials.material_type IN (2, 3)
    RIGHT JOIN empty_chunk
      ON  positions.z  = empty_chunk.z
      AND positions.x  = empty_chunk.x
      AND positions.y  = empty_chunk.y
      AND positions.cz = $1
      AND positions.cx = $2
      AND positions.cy = $3
    ORDER BY empty_chunk.y,
             empty_chunk.x,
             empty_chunk.z
  ) as t;`, cc[0], cc[1], cc[2])
	var b []byte
	err := row.Scan(&b)
	switch {
	case err == sql.ErrNoRows:
		log.Println("no rows")
	case err != nil:
		log.Fatal(err)
	}
	ssc.Convert(string(b))
}
