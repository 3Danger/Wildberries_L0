package JsonStruct

import (
	"encoding/json"
	"log"
	"sync"
)

type JsonSlice struct {
	sync.RWMutex
	counter   uint
	sliceJson []*JsonStruct
}

func NewJsonSlice() (j JsonSlice) {
	return JsonSlice{sync.RWMutex{}, 0, make([]*JsonStruct, 0, 10)}
}

func (j *JsonSlice) AddFromFile(fileName string) (ok error) {
	JsonModel, ok := NewFromFile(fileName)
	if ok != nil {
		return
	}
	j.sliceJson = append(j.sliceJson, JsonModel)
	return
}

func (j *JsonSlice) Add(jsonModel ...*JsonStruct) {
	j.sliceJson = append(j.sliceJson, jsonModel...)
}
func (j *JsonSlice) AddFromData(jsonData []byte) {
	var jsonModel JsonStruct
	err := json.Unmarshal(jsonData, &jsonModel)
	if err != nil {
		log.Panic(err)
	}
	j.sliceJson = append(j.sliceJson, &jsonModel)
}

func (j *JsonSlice) GetSlice() []*JsonStruct {
	return j.sliceJson
}
