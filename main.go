package main

import (
	"awesomeProject/srcs/Backend"
	ut "awesomeProject/srcs/Backend/Utils"
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	sigInterrupt := make(chan os.Signal)
	signal.Notify(sigInterrupt, os.Interrupt)

	config := ut.ParseArgs()
	backend := Backend.BackEnd(config)
	defer backend.Close()

	go producer("client-1", config.ClusterID, config.ModelSubj)

	<-sigInterrupt
	fmt.Println("\rGood bye!")
}

func producer(clientID, clusterID, subject string) {
	jsonByte, err := ioutil.ReadFile("model.json")
	if err != nil {
		log.Panic(err)
		return
	}
	connect, _ := stan.Connect(clusterID, clientID)
	for i := 0; i < 100000; i++ {
		err = ut.TryDoIt(time.Second, 10, func() (ok error) {
			ok = connect.Publish(subject, jsonByte)
			return ok
		})
		if err != nil {
			log.Panic("producer err:", err)
			return
		}
	}
	//fmt.Println("\r\nproducers work is done")
}
