package provider

import (
	"database/sql"
	"fmt"
	"log"
)

type provider struct {
	db *sql.DB
}

func NewProvider(host string, port int, user, password, dbName string) *provider {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	return &provider{db: conn}
}

func (p *provider) GetCounter() (int, error) {
	var counter int
	row := p.db.QueryRow("SELECT value FROM counter_table LIMIT 1")
	err := row.Scan(&counter)
	if err != nil {
		return 0, err
	}
	return counter, nil
}

func (p *provider) UpdateCounter(value int) error {
	_, err := p.db.Exec("UPDATE counter_table SET value = value + $1", value)
	return err
}
