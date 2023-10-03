package factory

type warrior struct {
	speed   int
	attack  int
	defense int
}

func (w warrior) getSpeed() int {
	return w.speed
}

func (w warrior) getAttack() int {
	return w.attack
}

func (w warrior) getDefense() int {
	return w.defense
}
