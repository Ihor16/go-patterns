package state

type stateOff struct {
	phone *Phone
}

func newStateOff(phone *Phone) *stateOff {
	return &stateOff{phone: phone}
}

func (s *stateOff) pressHome() {
	s.phone.turnOn()
	s.phone.setState(newStateLocked(s.phone))
}

func (s *stateOff) pressPower() {
	s.phone.turnOn()
	s.phone.lock()
	s.phone.setState(newStateLocked(s.phone))
}

func (s *stateOff) getName() string {
	return "off"
}
