package game

import (
	"database/sql"
	"log"
	"strconv"
	"strings"
)

type sqlStraightChunk struct {
	Arr  []materialType
	Grid [][]materialType
	sql.Scanner
}

type chunkCoords coords

func (c *chunkCoords) toJSON() *messageChunk {
	var s *sqlStraightChunk
	ssc := s.New()
	ssc.straightChunk(*c)
	ssc.ddChunk()

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
			d.addPosition(id, z, x, 0, cz, cx, defaultDepth/chunkY)
			d.addMaterial(id, materialType(die(2)+2))
		}
	}
}

func (ssc *sqlStraightChunk) New() *sqlStraightChunk {
	return &sqlStraightChunk{
		Arr:  make([]materialType, chunkZ*chunkX),
		Grid: make([][]materialType, chunkZ),
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

func (ssc *sqlStraightChunk) ddChunk() {
	for z := range ssc.Grid {
		ssc.Grid[z] = ssc.Arr[z*chunkX : (z+1)*chunkX]
		//for x := range ssc.Grid[z] {
		//ssc.Grid[z][x] = ssc.Arr[z*chunkZ+x]
		//}
	}
}

func (ssc *sqlStraightChunk) straightChunk(cc chunkCoords) {
	row := d.dbmap.Db.QueryRow(`SELECT array_agg(material_type)
	    FROM materials WHERE id IN
		(SELECT id FROM positions WHERE
			cz = $1 AND
			cx = $2 AND
			cy = 4
			ORDER BY x,z,y)
        and material_type in (2,3)`, cc[0], cc[1])
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
