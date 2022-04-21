package Utils

import (
	"flag"
	"os"
	"os/signal"
)

type Channels struct {
	StopQueueSelect chan bool
	StopMain        chan bool
	Interrupt       chan os.Signal
}

func SigHandlerClose(channels *Channels) {
	select {
	case <-channels.Interrupt:
		channels.StopQueueSelect <- true
		channels.StopMain <- true
	}
}

func InitChan() Channels {
	var ch Channels
	ch.StopQueueSelect = make(chan bool)
	ch.StopMain = make(chan bool)
	ch.Interrupt = make(chan os.Signal)
	signal.Notify(ch.Interrupt, os.Interrupt)
	return ch
}

type Configs struct {
	UserDB, PassDB, AddrDB, NameDB string
	ClusterID                      string
	ModelSubj                      string
}

func ParseArgs() (c *Configs) {
	c = &Configs{}
	flag.StringVar(&c.UserDB, "u", "csamuro", "user name of database")
	flag.StringVar(&c.NameDB, "d", "csamuro", "name of database")
	flag.StringVar(&c.PassDB, "p", "csamuro", "password of database")
	flag.StringVar(&c.AddrDB, "h", "localhost", "host address of database")
	flag.StringVar(&c.ClusterID, "cid", "test-cluster", "cluster id of NATS-streaming")
	flag.StringVar(&c.ModelSubj, "sm", "jsonModel", "Subject of channel to getting json model")
	//flag.StringVar(&c.stopSubj, "ss", "stop", "Subject of channel to notify for stop listening")
	flag.Parse()
	return c
}
