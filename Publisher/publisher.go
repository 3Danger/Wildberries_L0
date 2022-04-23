package main

import (
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
	var repeat, sleep int
	flag.StringVar(&clientID, "c", "producer-1", "client name")
	flag.StringVar(&clusterID, "cid", "test-cluster", "cluster id for connect")
	flag.StringVar(&subject, "subj", "jsonModel", "client name")
	flag.StringVar(&jsonPath, "j", "./json", "folder with contains .json")
	flag.IntVar(&repeat, "r", 1, "number of repetitions")
	flag.IntVar(&sleep, "ms", 0, "sleeping time (ms)")
	flag.Parse()

	models, ok := ReadAll(jsonPath)
	if ok != nil {
		log.Println(ok)
		return
	}
	connect, _ := stan.Connect(clusterID, clientID)
	for repeat > 0 {
		for i, v := range *models {
			ok = connect.Publish(subject, v)
			if ok != nil {
				log.Panic("producer err:", ok)
				return
			}
			fmt.Print("\rPublished:", i+1)
			if sleep > 0 {
				time.Sleep(time.Millisecond * time.Duration(sleep))
			}
		}
		repeat--
	}
}
