package BackEnd

import (
	"awesomeProject/srcs/BackEnd/Utils"
	"fmt"
	stan "github.com/nats-io/stan.go"
	"log"
)

type StanManager struct {
	connect  stan.Conn
	subjects map[string]stan.Subscription
}

func NewConnect(c *Utils.Configs, clientID string) (stanManager *StanManager) {
	var ok error
	stanManager = &StanManager{}
	stanManager.connect, ok = stan.Connect(c.ClusterID, clientID)
	if ok != nil {
		log.Panic(ok)
	}
	stanManager.subjects = make(map[string]stan.Subscription, 0)
	return stanManager
}

func (s *StanManager) NewSubscribe(subject *string, cb stan.MsgHandler, opts ...stan.SubscriptionOption) {
	subscribe, err := s.connect.Subscribe(*subject, cb, opts...)
	if err != nil {
		log.Panic(err)
	}
	s.subjects[*subject] = subscribe
	fmt.Println("subject:", *subject, "Subscribed")
}

func (s *StanManager) Unscribe(subject *string) {
	subscribe := s.subjects[*subject]
	if subscribe != nil {
		err := subscribe.Close()
		if err != nil {
			log.Print("unscribe err:", *subject, " ")
			log.Println(err)
		} else {
			fmt.Println("subject", *subject, "Unscribed")
		}
		delete(s.subjects, *subject)
	} else {
		fmt.Println("Subject", *subject, "not found")
	}
}

func (s *StanManager) UnscribeAll() {
	for k, v := range s.subjects {
		err := v.Close()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(k, "Unscribed")
	}
}
