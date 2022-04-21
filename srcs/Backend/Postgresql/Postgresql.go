package Postgresql

import (
	"awesomeProject/srcs/Backend/Utils"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Postgresql struct {
	//sync.RWMutex
	connStr string // := "postgresql://csamuro:PASSWORD@localhost/csamuro?sslmode=disable"
	open    *sql.DB
}

// Connect user : csamuro, address : localhost
// TODO check localhost:5432 in address
func (p *Postgresql) Connect(cnf *Utils.Configs, s time.Duration) {
	var ok error
	p.connStr = "postgresql://" + cnf.UserDB + ":" + cnf.PassDB + "@" + cnf.AddrDB + "/" + cnf.NameDB + "?sslmode=disable"
	ok = Utils.TryDoIt(s, 10, func() error {
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

func (p *Postgresql) Disconnect() {
	err := p.open.Close()
	if err != nil {
		log.Panic(err)
	} else {
		fmt.Println("DataBase disconnected")
	}
}

func (p *Postgresql) GetRaw() *sql.DB {
	return p.open
}
