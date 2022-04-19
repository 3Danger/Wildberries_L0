package main

import (
	js "awesomeProject/JsonStruct"
	pq "awesomeProject/Postgresql"
	"awesomeProject/Utils"
	"fmt"
	stan "github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
	"time"
)

func ReadFromDataBase(jsSlice *js.JsonSlice, DataBase *pq.Postgresql) {
	var jsonData []byte
	query, err := DataBase.GetRaw().Query("SELECT model FROM models;")
	if err != nil {
		log.Panic(err)
		return
	}
	for query.Next() {
		err = query.Scan(&jsonData)
		if err != nil {
			log.Panic(err)
		}
		jsSlice.AddFromData(jsonData)
	}
}

func main() {

	var (
		//userDB, passDB, addrDB string
		//clusterID              string
		//modelSubj              string
		//stopSubj               string
		configs  Utils.Configs
		dataBase pq.Postgresql
		conn     stan.Conn
	)
	Utils.ParseArgs(&configs)
	jModelSlice := js.NewJsonSlice()

	err := dataBase.Connect(configs.UserDB, configs.PassDB, configs.AddrDB)
	if err != nil {
		log.Panic(err)
		return
	}
	ReadFromDataBase(&jModelSlice, &dataBase)
	conn, err = stan.Connect(configs.ClusterID, "server-1")
	if err != nil {
		log.Panic(err)
		return
	}
	subscribe, err := conn.Subscribe(configs.ModelSubj, func(msg *stan.Msg) {
		tmp, err := js.ParseBytes(msg.Data)
		if err != nil {
			log.Panic(err)
			return
		}
		jModelSlice.Lock()
		_, err = dataBase.GetRaw().Query("INSERT INTO models (model) VALUES ($1)", tmp)
		if err != nil {
			log.Panic(err)
			return
		}
		jModelSlice.Add(&tmp)
		jModelSlice.Unlock()
	})
	defer func(subscribe stan.Subscription) {
		err := subscribe.Close()
		if err != nil {
			log.Panic(err)
		}
	}(subscribe)

	//producer("client-1")
	time.Sleep(time.Second * 3)
	for _, x := range jModelSlice.GetSlice() {
		fmt.Println(x.DateCreated)
	}
	time.Sleep(time.Minute)
}

func producer(clientID string) {
	jsonByte, err := ioutil.ReadFile("model.json")
	if err != nil {
		log.Panic(err)
		return
	}
	connect, _ := stan.Connect("TEST-CLUSTER-ID", clientID)
	err = connect.Publish("jsonModel", jsonByte)
	if err != nil {
		log.Panic(err)
		return
	}
}
