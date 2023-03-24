package adapter

import "fmt"

type MusicStore struct {
	csvDB  CSVMusicDB
	jsonDB JSONMusicDB
}

func NewMusicStore(csvDB CSVMusicDB, jsonDB JSONMusicDB) *MusicStore {
	return &MusicStore{csvDB: csvDB, jsonDB: jsonDB}
}

func (m *MusicStore) GetRecords() []Record {
	csv := m.csvDB.GetRecordsAsCSV()
	csvRecords := m.convertCSVToRecords(csv)
	json := m.jsonDB.GetRecordsAsJSON()
	jsonRecords := m.convertJSONToRecords(json)
	combined := append(csvRecords, jsonRecords...)
	return combined
}

func (m *MusicStore) convertCSVToRecords(_ string) []Record {
	fmt.Println("converting csv to records slice")
	return nil
}

func (m *MusicStore) convertJSONToRecords(_ string) []Record {
	fmt.Println("converting json to records slice")
	return nil
}
