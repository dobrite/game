package game

import (
	"github.com/coopernurse/gorp"
	"log"
	"os"
)

var d *db

type db struct {
	dbmap *gorp.DbMap
}

func (db *db) traceOn() {
	db.dbmap.TraceOn("", log.New(os.Stdout, "Gorp: ", log.Ldate|log.Ltime|log.Lshortfile))
}

func (db *db) traceOff() {
	db.dbmap.TraceOff()
}

//func (e *entity) Scan(src interface{}) error {
//	log.Println(src)
//	asBytes, ok := src.([]byte)
//	log.Println(asBytes)
//	if !ok {
//		log.Println("b")
//	}
//
//	asString := string(asBytes)
//	log.Println(asString)
//	*e = entity{ID: asString}
//	return nil
//}

func (db *db) newUUID() string {
	e := &entity{}
	if err := d.dbmap.Insert(e); err != nil {
		log.Fatal(err)
	}

	return e.ID
}

func (db *db) init() {
	db.dbmap.AddTableWithName(entity{}, "entities").SetKeys(true, "ID")
	db.dbmap.AddTableWithName(position{}, "positions").SetKeys(false, "ID")
	db.dbmap.AddTableWithName(material{}, "materials").SetKeys(false, "ID")
	db.dbmap.AddTableWithName(brain{}, "brains").SetKeys(false, "ID")
	db.dbmap.AddTableWithName(controlled{}, "controlled").SetKeys(false, "ID")
}
