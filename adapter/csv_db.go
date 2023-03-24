package adapter

import "fmt"

type CsvDB struct {
}

func (C CsvDB) GetRecordsAsCSV() string {
	fmt.Println("returning music records as CSV")
	return ""
}
