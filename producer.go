package main

import (
	js "awesomeProject/Json_struct"
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"os"
	"time"
)

func ErrorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	jsonFile := make(chan js.Json_struct, 5)

	connect, err := nats.Connect("nats://127.0.0.1:4445")
	conn, err := nats.NewEncodedConn(connect, nats.JSON_ENCODER)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//conn.Publish("foo", jsonFile)
	conn.Subscribe("server-upload-model-json-start", func(json_struct js.Json_struct) {
		jsonFile <- json_struct
	})

	connect.Subscribe("server-upload-model-json-stop", func(msg *nats.Msg) {
		fmt.Println("try to STOP")
		if string(msg.Data) == "stop" {
			fmt.Println("STOPED IT!!!!")
			fmt.Println("with value", string(msg.Data))
			close(jsonFile)
		}
	})
	//Consumer()
	go func() {
		Consumer()
	}()
	for i := range jsonFile {
		fmt.Printf("received: %v\n", i.Delivery)
	}
	//fmt.Println("received:", <-jsonFile)
	conn.Close()
	connect.Close()
}

func Consumer() {
	var jsonFile js.Json_struct

	open, err := os.Open("model.json")
	ErrorHandler(err)
	decoder := json.NewDecoder(open)
	err = decoder.Decode(&jsonFile)
	ErrorHandler(err)
	connect, err := nats.Connect("nats://127.0.0.1:4445")
	ErrorHandler(err)
	connect.Subscribe("channel", func(msg *nats.Msg) { fmt.Println("message from server:", msg) })
	conn, err := nats.NewEncodedConn(connect, nats.JSON_ENCODER)
	ErrorHandler(err)
	conn.Publish("server-upload-model-json-start", jsonFile)
	conn.Publish("server-upload-model-json-start", jsonFile)
	conn.Publish("server-upload-model-json-start", jsonFile)
	time.Sleep(time.Second * 3)
	connect.Publish("server-upload-model-json-stop", []byte("stop"))
	conn.Close()
	connect.Close()
}
