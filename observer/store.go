package observer

type Store struct {
	status       string
	observerList []Observer
}

func NewStore() *Store {
	return &Store{status: ""}
}

func (s *Store) Register(customer *Customer) {
	s.observerList = append(s.observerList, customer)
}

func (s *Store) UpdateStatus(status string) {
	s.status = status
	s.notifyObservers()
}

func (s *Store) notifyObservers() {
	for _, o := range s.observerList {
		o.update(s.status)
	}
}
