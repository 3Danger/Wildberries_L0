package main

import (
	ut "awesomeProject/srcs/Backend/Utils"
	"flag"
	"fmt"
	"github.com/nats-io/stan.go"
	"io/ioutil"
	"log"
	"time"
)

func ReadAll(path string) (res *[][]byte, ok error) {
	res = new([][]byte)
	dir, ok := ioutil.ReadDir(path)
	if ok != nil {
		return nil, ok
	}

	for _, v := range dir {
		if !v.IsDir() {
			file, err := ioutil.ReadFile(path + "/" + v.Name())
			if err == nil {
				*res = append(*res, file)
			}
		}
	}
	return res, ok
}

func main() {
	var clientID, clusterID, subject, jsonPath string
	flag.StringVar(&clientID, "c", "producer-1", "client name")
	flag.StringVar(&clusterID, "cid", "test-cluster", "cluster id for connect")
	flag.StringVar(&subject, "subj", "jsonModel", "client name")
	flag.StringVar(&jsonPath, "j", "./json", "folder with contains .json")
	flag.Parse()

	models, ok := ReadAll(jsonPath)
	if ok != nil {
		log.Println(ok)
		return
	}
	connect, _ := stan.Connect(clusterID, clientID)
	for i, v := range *models {
		ok = ut.TryDoIt(time.Second, 10, func() (ok error) {
			ok = connect.Publish(subject, v)
			return ok
		})
		if ok != nil {
			log.Panic("producer err:", ok)
			return
		}
		fmt.Print("\rPublished:", i)
		time.Sleep(time.Millisecond * 300)
	}
}
