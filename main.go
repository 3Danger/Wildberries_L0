package main

import (
	"awesomeProject/srcs/Backend"
	"awesomeProject/srcs/Backend/Postgresql"
	"awesomeProject/srcs/Backend/Utils"
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"
)

type Channels struct {
	StopQueueSelect chan bool
	StopMain        chan bool
	Interrupt       chan os.Signal
}

func InitChan() Channels {
	var ch Channels
	ch.StopQueueSelect = make(chan bool)
	ch.StopMain = make(chan bool)
	ch.Interrupt = make(chan os.Signal)
	signal.Notify(ch.Interrupt, os.Interrupt)
	return ch
}

func SigHandlerClose(channels *Channels) {
	select {
	case <-channels.Interrupt:
		channels.StopQueueSelect <- true
		channels.StopMain <- true
	}
}

func main() {
	channels := InitChan()
	go SigHandlerClose(&channels)
	config := Utils.ParseArgs()
	backEnd := Backend.BackEnd(config, channels.StopQueueSelect)
	defer backEnd.Close()
	go producer("client-1")

	go func() {
		var input string
		for {
			fmt.Scanln(&input)
			input = strings.ToLower(input)
			if strings.Compare(input, "get") == 0 {
				for i, v := range backEnd.JModelSlice.GetSlice() {
					fmt.Println(i, v.Locale)
				}
			} else if strings.Compare(input, "stop") == 0 {
				channels.Interrupt <- os.Interrupt
				return
			}
		}
	}()
	<-channels.StopMain
	fmt.Println("\rGood bye!")
}

func producer(clientID string) {
	jsonByte, err := ioutil.ReadFile("model.json")
	if err != nil {
		log.Panic(err)
		return
	}
	connect, _ := stan.Connect("TEST-CLUSTER-ID", clientID)
	for i := 0; i < 1000; i++ {
		fmt.Println("inserting:", i)
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
