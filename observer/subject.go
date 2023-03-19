package observer

type Subject interface {
	notifyObservers()
}
