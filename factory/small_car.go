package factory

import "fmt"

type smallCar struct {
	wheels  []string
	chassis []string
	parts   []string
}

func newSmallCar(wheels []string, chassis []string, parts []string) *smallCar {
	return &smallCar{wheels: wheels, chassis: chassis, parts: parts}
}

func (c smallCar) Drive() {
	fmt.Println("driving a small car")
}
