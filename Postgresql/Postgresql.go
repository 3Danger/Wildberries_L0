package Postgresql

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
)

type Postgresql struct {
	connStr string // := "postgresql://csamuro:PASSWORD@localhost/csamuro?sslmode=disable"
	open    *sql.DB
}

// Connect user : csamuro, address : localhost
// TODO check localhost:5432 in address
func (p *Postgresql) Connect(user, password, address string) error {
	var ok error
	p.connStr = "postgresql://" + user + ":" + password + "@" + address + "/" + user + "?sslmode=disable"
	p.open, ok = sql.Open("postgres", p.connStr)
	if ok != nil {
		return ok
	}
	return nil
}

func (p *Postgresql) Disconnect() error {
	return p.open.Close()
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
