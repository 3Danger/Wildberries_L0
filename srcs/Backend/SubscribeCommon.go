package Backend

import (
	JsonStruct2 "awesomeProject/srcs/Backend/JsonStruct"
	"awesomeProject/srcs/Backend/Postgresql"
	"awesomeProject/srcs/Backend/Utils"
	"fmt"
	"github.com/nats-io/stan.go"
	"time"
)

func ModelSubscribe(bk *CommonBackend, Configs *Utils.Configs, stop <-chan bool) {
	dataChannel := make(chan []byte, 1000)
	bk.Connect.NewSubscribe(&Configs.ModelSubj, func(msg *stan.Msg) { dataChannel <- msg.Data })
	go queueInserting(dataChannel, bk, stop)
}

func queueInserting(dataChan <-chan []byte, bk *CommonBackend, stop <-chan bool) {
	for {
		select {
		case data := <-dataChan:
			model, ok := JsonStruct2.ParseBytes(data)
			ok = Postgresql.TryDoIt(time.Second, 10, func() error {
				//bk.DataBase.Lock()
				//_, err := bk.DataBase.GetRaw().Query("INSERT INTO models (model) VALUES ($1)", model)
				rows, err := bk.DataBase.GetRaw().Query("INSERT INTO models (model) VALUES ($1)", model)
				rows.Close()
				//bk.DataBase.Unlock()
				return err
			})
			if ok != nil {
				fmt.Println("can't insert model from server")
				fmt.Println(ok)
				continue
			}
			bk.JModelSlice.Lock()
			bk.JModelSlice.Add(&model)
			bk.JModelSlice.Unlock()
		case <-stop:
			fmt.Println("\rGot signal from SELECT")
			return
		}
	}
}
