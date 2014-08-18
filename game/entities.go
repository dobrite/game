package game

type entity struct {
	ID string `db:"id"`
}

func (db *db) remove(id string) {
	db.dbmap.Delete(&entity{ID: id})
}
