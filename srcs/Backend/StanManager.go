package Backend

import (
	"awesomeProject/srcs/Backend/Utils"
	"fmt"
	stan "github.com/nats-io/stan.go"
	"log"
)

type StanManager struct {
	connect  stan.Conn
	subjects map[string]stan.Subscription
}

func NewConnect(c *Utils.Configs) (stanManager *StanManager) {
	var ok error
	stanManager = &StanManager{}
	stanManager.connect, ok = stan.Connect(c.ClusterID, c.ClientID)
	if ok != nil {
		log.Fatal("Не удалось подключиться к nats-streaming-server")
	}
	fmt.Println("\033[34m"+"clusterID:"+"\033[0m", c.ClusterID+",",
		"\033[34m"+"clientID:"+"\033[0m", c.ClientID,
		"\033[32m"+"connected!"+"\033[0m")
	stanManager.subjects = make(map[string]stan.Subscription, 0)
	return stanManager
}

func (s *StanManager) NewSubscribe(subject *string, cb stan.MsgHandler, opts ...stan.SubscriptionOption) {
	subscribe, err := s.connect.Subscribe(*subject, cb, opts...)
	if err != nil {
		log.Panic(err)
	}
	s.subjects[*subject] = subscribe
	fmt.Println("\033[34m"+"subject:"+"\033[0m", *subject, "\033[32m"+"Subscribed!"+"\033[0m")
}

func (s *StanManager) Unscribe(subject *string) {
	subscribe := s.subjects[*subject]
	if subscribe != nil {
		err := subscribe.Close()
		if err != nil {
			log.Print("\033[31m"+"Unscribe err:", *subject, "\033[0m")
			log.Println(err)
		} else {
			fmt.Println("\033[34m"+"Subject:"+"\033[0m", *subject, "\033[32m"+"Unscribed"+"\033[0m")
		}
		delete(s.subjects, *subject)
	} else {
		fmt.Println("Subject:", *subject, "not found")
	}
}

func (s *StanManager) UnscribeAll() {
	for k, v := range s.subjects {
		err := v.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("\033[34m"+"Subject:"+"\033[0m", k, "\033[32m"+"Unscribed"+"\033[0m")
	}
}
