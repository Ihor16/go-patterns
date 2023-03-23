package state

type stateReady struct {
	phone *Phone
}

func newStateReady(phone *Phone) *stateReady {
	return &stateReady{phone: phone}
}

func (s *stateReady) pressHome() {
	s.phone.setState(newStateReady(s.phone))
}

func (s *stateReady) pressPower() {
	s.phone.turnOff()
	s.phone.setState(newStateOff(s.phone))
}

func (s *stateReady) getName() string {
	return "ready"
}
