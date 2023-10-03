package factory

type archer struct {
	speed   int
	attack  int
	defense int
}

func (a archer) getSpeed() int {
	return a.speed
}

func (a archer) getAttack() int {
	return a.attack
}

func (a archer) getDefense() int {
	return a.defense
}
