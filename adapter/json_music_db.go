package adapter

import "fmt"

type JSONMusicDB struct {
}

func (J JSONMusicDB) GetRecordsAsJSON() string {
	fmt.Println("returning music records as JSON")
	return ""
}
