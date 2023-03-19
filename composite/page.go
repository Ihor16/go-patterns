package composite

import "fmt"

type Page struct {
}

func NewPage() *Page {
	return &Page{}
}

func (p *Page) Traverse() {
	fmt.Println("I'm a page")
}
