package composite

import "fmt"

type Article struct {
	traversers []Traverser
}

func NewArticle(traversers []Traverser) *Article {
	return &Article{traversers: traversers}
}

func (a Article) Traverse() {
	fmt.Println("I'm an article")
	for _, c := range a.traversers {
		c.Traverse()
	}
}
