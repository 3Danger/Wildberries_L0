package Backend

import (
	"awesomeProject/srcs/Backend/JsonStruct"
	pq "awesomeProject/srcs/Backend/Postgresql"
	"awesomeProject/srcs/Backend/Utils"
	"log"
	"time"
)

type CommonBackend struct {
	DataBase    pq.Postgresql
	Connect     *StanManager
	JModelSlice JsonStruct.JsonSlice
}

func BackEnd(Configs *Utils.Configs, StopQueueSelect chan bool) *CommonBackend {
	var backend CommonBackend
	backend.DataBase.Connect(Configs, time.Second*3)
	backend.JModelSlice = JsonStruct.NewJsonSlice()
	ReadFromDataBase(&backend)
	backend.Connect = NewConnect(Configs, "server-1")
	ModelSubscribe(&backend, Configs, StopQueueSelect)
	return &backend
}

func (c *CommonBackend) Close() {
	c.DataBase.Disconnect()
	c.Connect.UnscribeAll()
}

func ReadFromDataBase(bk *CommonBackend) {
	var jsonData []byte
	query, err := bk.DataBase.GetRaw().Query("SELECT model FROM models;")
	if err != nil {
		log.Panic(err)
		return
	}
	for query.Next() {
		err = query.Scan(&jsonData)
		if err != nil {
			log.Panic(err)
		}
		bk.JModelSlice.Lock()
		bk.JModelSlice.AddFromData(jsonData)
		bk.JModelSlice.Unlock()
	}
}
