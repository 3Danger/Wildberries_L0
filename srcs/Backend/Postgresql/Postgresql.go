package Postgresql

import (
	"awesomeProject/srcs/Backend/Utils"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"sync"
	"time"
)

type Postgresql struct {
	sync.RWMutex
	connStr string // := "postgresql://csamuro:PASSWORD@localhost/csamuro?sslmode=disable"
	open    *sql.DB
}

// Connect user : csamuro, address : localhost
// TODO check localhost:5432 in address
func (p *Postgresql) Connect(cnf *Utils.Configs, s time.Duration) {
	var ok error
	p.connStr = "postgresql://" + cnf.UserDB + ":" + cnf.PassDB + "@" + cnf.AddrDB + "/" + cnf.NameDB + "?sslmode=disable"
	ok = TryDoIt(s, 10, func() error {
		p.open, ok = sql.Open("postgres", p.connStr)
		return ok
	},
	)
	if ok == nil {
		fmt.Println("DataBase connected")
	} else {
		log.Panic(ok)
	}
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

func (p *Postgresql) Disconnect() {
	err := p.open.Close()
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println("DataBase disconnected")
	}
}

func (p *Postgresql) InsertModel(args ...any) error {
	if p.open == nil {
		return errors.New("not connected to sql")
	}
	_, err := p.open.Exec("INSERT INTO models (model) VALUES ($1)", args)
	return err
}

func (p *Postgresql) GetRaw() *sql.DB {
	return p.open
}
