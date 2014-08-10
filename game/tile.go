package game

func makeTile(y, x, cy, cx int, t materialType) string {
	ent := newUUID()

	positionsSet.add(ent, y, x, cy, cx)
	materialsSet.add(ent, t)

	return ent
}
