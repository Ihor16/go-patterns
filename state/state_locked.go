package state

type stateLocked struct {
	phone *Phone
}

func newStateLocked(phone *Phone) *stateLocked {
	return &stateLocked{phone: phone}
}

func (s *stateLocked) pressHome() {
	s.phone.unlock()
	s.phone.home()
	s.phone.setState(newStateReady(s.phone))
}

func (s *stateLocked) pressPower() {
	s.phone.turnOff()
	s.phone.setState(newStateOff(s.phone))
}

func (s *stateLocked) getName() string {
	return "locked"
}
