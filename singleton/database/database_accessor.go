package database

import (
	"fmt"
	"sync"
)

type accessor struct {
	name string
}

var instance *accessor = nil
var once sync.Once

func GetInstance() *accessor {
	once.Do(func() {
		instance = &accessor{name: "db"}
	})
	return instance
}

func (a accessor) Query(q string) {
	fmt.Println(fmt.Sprintf("querying %s from the database", q))
}
