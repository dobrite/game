package game

type entity struct {
	ID string
}

func (db *db) remove(id string) {
	d.dbmap.Delete(&entity{ID: id})
}
