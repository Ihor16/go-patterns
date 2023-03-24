package adapter

import "fmt"

type JsonAdapter struct {
	jsonDB JsonDB
}

func NewJSONAdapter(jsonDB JsonDB) *JsonAdapter {
	return &JsonAdapter{jsonDB: jsonDB}
}

func (j *JsonAdapter) GetRecords() []Record {
	_ = j.jsonDB.GetRecordsAsJSON()
	fmt.Println("converting JSON to records slice")
	return nil
}
