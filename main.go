package main

import (
	pq "awesomeProject/Postgresql"
	ss "awesomeProject/StanType"
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
)

func main() {
	jsonFile := make(chan string)
	var connect ss.StanType
	var db pq.Postgresql
	err := db.Connect("csamuro", "irGJg$3.5.7", "localhost")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = connect.ConnectStan("TEST-CLUSTER-ID", "client-1")
	if err != nil {
		fmt.Println(err)
		return
	}
	subscribe, err := connect.GetStan().Subscribe("test-consumer-jsonModel", func(msg *stan.Msg) {
		jsonFile <- string(msg.Data)
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	producer()

	fmt.Println("-------Hello from JSON!!!!!! ***********\n", <-jsonFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	close(jsonFile)
	err = subscribe.Close()
}

func producer() {
	jsonByte, err := ioutil.ReadFile("model.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	var connect ss.StanType
	err = connect.ConnectStan("TEST-CLUSTER-ID", "client-2")
	if err != nil {
		fmt.Println(err)
		return
	}

	cont := connect.GetStan()
	err = cont.Publish("test-consumer-jsonModel", jsonByte)
	if err != nil {
		fmt.Println(err)
		return
	}
	//err = connect.GetStan().Publish("test-consumer", []byte("Hello world"))
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

}
