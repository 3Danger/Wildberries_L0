package Backend

import (
	"github.com/nats-io/stan.go"
	"log"
)

func ModelSubscribe(bk *CommonBackend, subject string) {
	bk.Connect.NewSubscribe(&subject, func(msg *stan.Msg) {
		_, ok := bk.DataBase.GetRaw().Exec("INSERT INTO models (model) VALUES ($1)", msg.Data)
		if ok != nil {
			log.Panic(ok)
			return
		}
		bk.JModelSlice.Lock()
		defer bk.JModelSlice.Unlock()
		ok = bk.JModelSlice.AddFromData(msg.Data)
		if ok != nil {
			log.Println(ok)
		}
	})
}
