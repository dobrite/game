package game

func makeTile(z, x, cz, cx int, t materialType) string {
	ent := newUUID()

	positionsSet.add(ent, z, x, cz, cx)
	materialsSet.add(ent, t)

	return ent
}
