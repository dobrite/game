package game

type item struct {
	position position
	material material
}

type itemJSON struct {
	Event        string       `json:"event"`
	WorldCoords  *worldCoords `json:"world_coords"`
	MaterialType materialType `json:"material_type"`
}

func (i *item) toJSON() *itemJSON {
	return &itemJSON{
		Event:        "game:item",
		WorldCoords:  i.position.toWorldCoords(),
		MaterialType: i.material.materialType,
	}
}
