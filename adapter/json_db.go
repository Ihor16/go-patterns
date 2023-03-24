package adapter

import "fmt"

type JsonDB struct {
}

func (J JsonDB) GetRecordsAsJSON() string {
	fmt.Println("returning music records as JSON")
	return ""
}
