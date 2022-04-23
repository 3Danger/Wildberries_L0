package Utils

import (
	"flag"
	"time"
)

type Configs struct {
	UserDB, PassDB, AddrDB, NameDB string
	ClusterID                      string
	ClientID                       string
	ModelSubj                      string
}

func ParseArgs() (c *Configs) {
	c = &Configs{}
	flag.StringVar(&c.UserDB, "u", "csamuro", "user name of database")
	flag.StringVar(&c.NameDB, "d", "csamuro", "name of database")
	flag.StringVar(&c.PassDB, "p", "csamuro", "password of database")
	flag.StringVar(&c.AddrDB, "h", "localhost", "host address of database")
	flag.StringVar(&c.ClusterID, "cid", "test-cluster", "cluster id of NATS-streaming")
	flag.StringVar(&c.ClientID, "cln", "server-1", "client name in NATS-connection")
	flag.StringVar(&c.ModelSubj, "sm", "jsonModel", "Subject of channel to getting json model")
	//flag.StringVar(&c.stopSubj, "ss", "stop", "Subject of channel to notify for stop listening")
	flag.Parse()
	return c
}

func TryDoIt(t time.Duration, attempts uint8, f func() error) (ok error) {
	ok = f()
	for ok != nil && attempts != 0 {
		time.Sleep(t)
		ok = f()
		attempts--
	}
	return ok
}
