package adapter

import "fmt"

type CSVMusicDB struct {
}

func (C CSVMusicDB) GetRecordsAsCSV() string {
	fmt.Println("returning music records as CSV")
	return ""
}
