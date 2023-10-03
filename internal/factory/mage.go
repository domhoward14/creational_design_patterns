package factory

type mage struct {
	speed   int
	attack  int
	defense int
}

func (m mage) getSpeed() int {
	return m.speed
}

func (m mage) getAttack() int {
	return m.attack
}

func (m mage) getDefense() int {
	return m.defense
}
