package main

import (
	"awesomeProject/srcs/Backend"
	"awesomeProject/srcs/Backend/Postgresql"
	ut "awesomeProject/srcs/Backend/Utils"
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

func SigHandlerClose(channels *ut.Channels) {
	select {
	case <-channels.Interrupt:
		channels.StopQueueSelect <- true
		channels.StopMain <- true
	}
}

func main() {
	channels := ut.InitChan()
	config := ut.ParseArgs()
	backend := Backend.BackEnd(config, channels.StopQueueSelect)
	defer backend.Close()

	go SigHandlerClose(&channels)
	go producer("client-1", config.ClusterID)
	go DebugHandler(channels, backend)
	<-channels.StopMain
	fmt.Println("\rGood bye!")
}

func DebugHandler(channels ut.Channels, backend *Backend.CommonBackend) {
	var input string
	for {
		fmt.Scanln(&input)
		input = strings.ToLower(input)
		if strings.Compare(input, "get") == 0 {
			for i, v := range backend.JModelSlice.GetSlice() {
				fmt.Println(i, v.Locale)
			}
		} else if strings.Compare(input, "stop") == 0 {
			channels.Interrupt <- os.Interrupt
			return
		}
	}
}

func producer(clientID, clusterID string) {
	jsonByte, err := ioutil.ReadFile("model.json")
	if err != nil {
		log.Panic(err)
		return
	}
	connect, _ := stan.Connect(clusterID, clientID)
	for i := 0; i < 1000; i++ {
		//fmt.Println("inserting:", i)
		err = Postgresql.TryDoIt(time.Second, 10, func() (ok error) {
			ok = connect.Publish("jsonModel", jsonByte)
			return ok
		})
		if err != nil {
			log.Panic("producer err:", err)
			return
		}
	}
	fmt.Println("producers work is done")
}
