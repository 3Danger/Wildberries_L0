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
	"strings"
	"time"
)

func main() {
	sigInterrupt := make(chan os.Signal)
	signal.Notify(sigInterrupt, os.Interrupt)

	config := ut.ParseArgs()
	backend := Backend.BackEnd(config)
	defer backend.Close()

	go producer("client-1", config.ClusterID, config.ModelSubj)
	go DebugHandler(sigInterrupt, backend)

	<-sigInterrupt
	fmt.Println("\rGood bye!")
}

func DebugHandler(sigQuit chan<- os.Signal, backend *Backend.CommonBackend) {
	var input string
	for {
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if strings.Compare(input, "get") == 0 {
			for i, v := range backend.JModelSlice.GetSlice() {
				fmt.Println(i, v.Locale)
			}
		} else if strings.Compare(input, "stop") == 0 {
			sigQuit <- os.Interrupt
			return
		}
	}
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
	fmt.Println("producers work is done")
}
