package game

type item struct {
	id       string
	position *position
	material *material
}

type itemJSON struct {
	Event        string       `json:"event"`
	Id           string       `json:"id"`
	WorldCoords  *worldCoords `json:"wc"`
	MaterialType materialType `json:"materialType"`
}

func (i *item) toJSON() *itemJSON {
	return &itemJSON{
		Event:        "game:item",
		Id:           i.id,
		WorldCoords:  i.position.toWorldCoords(),
		MaterialType: i.material.MaterialType,
	}
}
