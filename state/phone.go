package state

import "fmt"

const LOCKED = "locked"
const OFF = "off"
const READY = "ready"

type Phone struct {
	state string
}

func NewPhone() *Phone {
	return &Phone{state: OFF}
}

func (p *Phone) PressHome() {
	switch p.state {
	case OFF:
		p.turnOn()
		p.setState(LOCKED)
	case LOCKED:
		p.unlock()
		p.home()
		p.setState(READY)
	case READY:
		p.setState(READY)
	default:
		fmt.Println("error: invalid transition")
	}
}

func (p *Phone) PressPower() {
	switch p.state {
	case OFF:
		p.turnOn()
		p.lock()
		p.setState(LOCKED)
	case LOCKED:
		p.turnOff()
		p.setState(OFF)
	case READY:
		p.turnOff()
		p.setState(OFF)
	default:
		fmt.Println("error: invalid transition")
	}
}

func (p *Phone) lock() {
	fmt.Println("Locking phone")
}

func (p *Phone) home() {
	fmt.Println("Showing home screen")
}

func (p *Phone) unlock() {
	fmt.Println("Unlocking phone")
}

func (p *Phone) turnOn() {
	fmt.Println("Turning on phone")
}

func (p *Phone) turnOff() {
	fmt.Println("Turning off phone")
}

func (p *Phone) setState(s string) {
	fmt.Println("=> TRANSITION: FROM", p.state, "TO", s)
	p.state = s
}
