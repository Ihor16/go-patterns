package state

import "fmt"

var state State

type Phone struct {
}

func NewPhone() *Phone {
	p := &Phone{}
	state = newStateOff(p)
	return p
}

func (p *Phone) PressHome() {
	state.pressHome()
}

func (p *Phone) PressPower() {
	state.pressPower()
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

func (p *Phone) setState(s State) {
	fmt.Println("=> TRANSITION: FROM", state.getName(), "TO", s.getName())
	state = s
}
