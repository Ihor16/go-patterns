package factory

type Creator interface {
	Create() Drivable
}
