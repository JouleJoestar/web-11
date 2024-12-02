package provider

import (
	"database/sql"
	"fmt"
	"log"
)

type Provider struct {
	conn *sql.DB
}

func NewProvider(host string, port int, user, password, dbName string) *Provider {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	return &Provider{conn: conn}
}

func (p *Provider) SelectUser(name string) (string, error) {
	var user string
	row := p.conn.QueryRow("SELECT name FROM mytable WHERE name = $1", name)
	err := row.Scan(&user)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}
	return user, nil
}

func (p *Provider) InsertUser(name string) error {
	_, err := p.conn.Exec("INSERT INTO mytable (name) VALUES ($1)", name)
	if err != nil {
		return err
	}
	return nil
}

func (p *Provider) Close() error {
	return p.conn.Close()
}
