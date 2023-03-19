package factory

import "fmt"

type bigCar struct {
	wheels  []string
	chassis []string
	parts   []string
}

func newBigCar(wheels []string, chassis []string, parts []string) *bigCar {
	return &bigCar{wheels: wheels, chassis: chassis, parts: parts}
}

func (c bigCar) Drive() {
	fmt.Println("driving a big car")
}
