package adapter

import "fmt"

type CsvAdapter struct {
	csvDB CsvDB
}

func NewCSVAdapter(csvDB CsvDB) *CsvAdapter {
	return &CsvAdapter{csvDB: csvDB}
}

func (m CsvAdapter) GetRecords() []Record {
	_ = m.csvDB.GetRecordsAsCSV()
	fmt.Println("converting CSV to records slice")
	return nil
}
