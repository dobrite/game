package game

import (
	"log"
)

type controlledMap map[string]controlled

type controlled struct {
	ID string
}

func (db *db) addControlled(id string) {
	c := &controlled{
		ID: id,
	}

	if err := db.dbmap.Insert(c); err != nil {
		log.Fatal(err)
	}

	log.Println(c)
}
