package JsonStruct

import (
	"encoding/json"
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
func (j *JsonSlice) AddFromData(jsonData []byte) (ok error) {
	var jsonModel JsonStruct
	ok = json.Unmarshal(jsonData, &jsonModel)
	if ok == nil {
		j.sliceJson = append(j.sliceJson, &jsonModel)
	}
	return ok
}

func (j *JsonSlice) GetSlice() []*JsonStruct {
	return j.sliceJson
}
