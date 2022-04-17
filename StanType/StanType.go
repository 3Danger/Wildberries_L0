package StanType

import (
	"errors"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

type StanType struct {
	connectionNats *nats.Conn
	connectionStan stan.Conn
}

func (s *StanType) Connect(address, port string) (ok error) {
	s.Disconnect()
	s.connectionNats, ok = nats.Connect("nats://" + address + ":" + port)
	return ok
}

// ConnectStanNats stanClusterID : foo & clientID : 999
// ConnectStanNats stanClusterID : test-cluster & clientID : client-1
func (s *StanType) ConnectStanNats(stanClusterID, clientID string) (ok error) {
	if s.connectionNats == nil {
		return errors.New("nats isn't connected")
	}
	s.connectionStan, ok = stan.Connect(
		stanClusterID, //# may be "test-cluster",
		clientID,      //# may be "client-1",
		stan.NatsConn(s.connectionNats))
	return ok
}

// ConnectStan stanClusterID : foo & clientID : 999
// ConnectStan stanClusterID : test-cluster & clientID : client-1
func (s *StanType) ConnectStan(stanClusterID, clientID string) (ok error) {
	s.connectionStan, ok = stan.Connect(
		stanClusterID, //# may be "test-cluster",
		clientID,      //# may be "client-1",
	)
	return ok
}

func (s *StanType) Disconnect() {
	if s.connectionNats != nil {
		s.connectionNats.Close()
	}
}

func (s *StanType) GetStan() stan.Conn { return s.connectionStan }
