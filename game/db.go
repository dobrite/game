package game

import (
	"github.com/coopernurse/gorp"
	"github.com/nu7hatch/gouuid"
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

//func (db *db) newUUID() string {
//	e := &entity{}
//	if err := d.dbmap.Insert(e); err != nil {
//		panic(err)
//	}
//
//	return e.ID
//}

func (db *db) newUUID() string {
	u4, err := uuid.NewV4()
	if err != nil {
		panic("uuid failed")
	}

	e := &entity{ID: u4.String()}
	if err := d.dbmap.Insert(e); err != nil {
		panic(err)
	}

	return u4.String()
}

func (db *db) init() {
	db.dbmap.AddTableWithName(entity{}, "entities")
	//db.dbmap.Dialect.AutoIncrInsertSuffix(&gorp.ColumnMap{ColumnName: "id"})

	db.dbmap.AddTableWithName(position{}, "positions")
	db.dbmap.AddTableWithName(material{}, "materials")
	db.dbmap.AddTableWithName(brain{}, "brains")
	db.dbmap.AddTableWithName(controlled{}, "controlled")
}
