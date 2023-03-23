package state

type State interface {
	pressHome()
	pressPower()
	getName() string
}
