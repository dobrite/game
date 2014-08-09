package game

import (
	"github.com/nu7hatch/gouuid"
)

type tile struct {
	position
	material
}

func makeTile(y, x, cy, cx int, t materialType) *uuid.UUID {
	ent := newUUID()

	positionsSet.add(ent, y, x, cy, cx)
	materialsSet.add(ent, t)

	return ent
}
