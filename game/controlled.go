package game

type controlledMap map[string]controlled

type controlled struct {
	controller string
}

func (c controlledMap) add(ent string) {
	c[ent] = controlled{
		controller: ent,
	}
}

func (c controlledMap) remove(ent string) {
	delete(c, ent)
}
