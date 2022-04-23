package Backend

import (
	"awesomeProject/srcs/Backend/JsonStruct"
	pq "awesomeProject/srcs/Backend/Postgresql"
	"awesomeProject/srcs/Backend/Utils"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"time"
)

type CommonBackend struct {
	DataBase    pq.Postgresql
	ConnectStan *StanManager
	JModelSlice JsonStruct.JsonSlice
}

func BackEnd(Configs *Utils.Configs) *CommonBackend {
	var backend CommonBackend
	backend.DataBase.Connect(Configs, time.Second*3)
	backend.JModelSlice = JsonStruct.NewJsonSlice()
	ReadFromDataBase(&backend)
	backend.ConnectStan = NewConnect(Configs, "server-1")
	ModelSubscribe(&backend, Configs.ModelSubj)
	return &backend
}

func (c *CommonBackend) Close() {
	c.DataBase.Disconnect()
	c.ConnectStan.UnscribeAll()
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
		err := bk.JModelSlice.AddFromData(jsonData)
		bk.JModelSlice.Unlock()
		if err != nil {
			log.Println(err)
			return
		}
	}
}

//TODO Проверка на наличие элемента в кеше
//func Find(jsonSlice []*JsonStruct.JsonStruct, jsonStruct *JsonStruct.JsonStruct) int {
//	for i, v := range jsonSlice {
//		if v.OrderUid == jsonStruct.OrderUid {
//			if reflect.DeepEqual(*v, *jsonStruct) {
//				return i
//			}
//		}
//	}
//	return -1
//}

// ModelSubscribe Можно гораздо ускорить принятие данных путем создания -
//- отдельных структур (для каждого клиента) со своими мютексами и слайсами
func ModelSubscribe(bk *CommonBackend, subject string) {
	fmt.Println("\033[36m"+"Count elem in cache before:", len(bk.JModelSlice.GetSlice()), "\033[0m")
	bk.ConnectStan.NewSubscribe(&subject, func(msg *stan.Msg) {
		js, ok := JsonStruct.ParseBytes(msg.Data)
		if ok != nil {
			fmt.Println("\u001B[31m" + "incoming json model is invalid" + "\033[0m")
			return
		}
		//TODO при необходимости исклчать повторяющиеся данные
		//if Find(bk.JModelSlice.GetSlice(), js) >= 0 {
		//	fmt.Println("\033[34m"+"skipped"+"\033[0m", js.OrderUid, "\033[34m"+"because it exist in cache"+"\033[0m")
		//	return
		//}
		_, ok = bk.DataBase.GetRaw().Exec("INSERT INTO models (model) VALUES ($1)", msg.Data)
		if ok != nil {
			log.Println(ok)
			return
		}
		bk.JModelSlice.Lock()
		defer bk.JModelSlice.Unlock()
		bk.JModelSlice.Add(js)
		fmt.Println("\033[32m"+"Count elem in cache after: ", len(bk.JModelSlice.GetSlice()), "\b"+"\033[0m")
	})
}
