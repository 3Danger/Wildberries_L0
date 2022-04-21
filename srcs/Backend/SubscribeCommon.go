package Backend

import (
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
)

func ModelSubscribe(bk *CommonBackend, subject string) {
	fmt.Println("\r Count elem in cache before:", len(bk.JModelSlice.GetSlice()), "\b")
	bk.Connect.NewSubscribe(&subject, func(msg *stan.Msg) {
		_, ok := bk.DataBase.GetRaw().Exec("INSERT INTO models (model) VALUES ($1)", msg.Data)
		if ok != nil {
			log.Panic(ok)
			return
		}
		bk.JModelSlice.Lock()
		defer bk.JModelSlice.Unlock()
		fmt.Print("\r Count elem in cache after: ", len(bk.JModelSlice.GetSlice()), "\b")
		ok = bk.JModelSlice.AddFromData(msg.Data)
		if ok != nil {
			log.Println(ok)
		}
	})
}
