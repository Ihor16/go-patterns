package factory

type BigCarFactory struct {
}

func NewBigCarFactory() *BigCarFactory {
	return &BigCarFactory{}
}

func (b BigCarFactory) Create() Drivable {
	return newBigCar(b.createWheels(), b.createChassis(), b.createParts())
}

func (b BigCarFactory) createWheels() []string {
	return []string{"wheel1", "wheel2", "wheel3", "wheel4"}
}

func (b BigCarFactory) createChassis() []string {
	return []string{"big-chassis1", "big-chassis2"}
}

func (b BigCarFactory) createParts() []string {
	return []string{"big-part1", "big-part2", "big-part3"}
}
