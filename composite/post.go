package composite

import "fmt"

type Post struct {
}

func NewPost() *Post {
	return &Post{}
}

func (p *Post) Traverse() {
	fmt.Println("I'm a post")
}
