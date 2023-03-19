package factory

type SmallCarFactory struct {
}

func NewSmallCarFactory() *SmallCarFactory {
	return &SmallCarFactory{}
}

func (s SmallCarFactory) Create() Drivable {
	return newSmallCar(s.createWheels(), s.createChassis(), s.createParts())
}

func (s SmallCarFactory) createWheels() []string {
	return []string{"wheel1", "wheel2", "wheel3", "wheel4"}
}

func (s SmallCarFactory) createChassis() []string {
	return []string{"small-chassis1"}
}

func (s SmallCarFactory) createParts() []string {
	return []string{"small-part1"}
}
