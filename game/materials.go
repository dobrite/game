package game

import (
	"log"
)

type materialsMap map[string]material

type material struct {
	ID           string
	MaterialType materialType `db:"material_type"`
}

type materialType int

const (
	nothing materialType = iota
	air
	dirt
	grass
	water
	stone
	flesh
	cow
	pig
)

func (db *db) addMaterial(id string, t materialType) {
	m := &material{
		ID:           id,
		MaterialType: t,
	}

	if err := db.dbmap.Insert(m); err != nil {
		panic(err)
	}

	log.Println(m)
}

func (db *db) getMaterial(id string) *material {
	obj, _ := db.dbmap.Get(material{}, id)
	return obj.(*material)
}
