package observer

import "fmt"

type Customer struct {
}

func (c Customer) update(message string) {
	fmt.Println("Message from store:", message)
}
