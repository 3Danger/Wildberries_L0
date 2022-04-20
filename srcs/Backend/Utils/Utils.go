package Utils

import (
	"flag"
)

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
	flag.StringVar(&c.ClusterID, "cid", "TEST-CLUSTER-ID", "cluster id of NATS-streaming")
	flag.StringVar(&c.ModelSubj, "sm", "jsonModel", "Subject of channel to getting json model")
	//flag.StringVar(&c.stopSubj, "ss", "stop", "Subject of channel to notify for stop listening")
	flag.Parse()
	return c
}
