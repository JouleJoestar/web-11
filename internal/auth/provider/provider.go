package provider

import (
	"database/sql"
	"fmt"
	"log"
)

type Provider struct {
	db *sql.DB
}

func NewProvider(host string, port int, user, password, dbName string) *Provider {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

	conn, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	return &Provider{db: conn}
}

func (p *Provider) GetUser(username string) (string, error) {
	var existingUser string
	err := p.db.QueryRow("SELECT username FROM users WHERE username = $1", username).Scan(&existingUser)
	if err != nil {
		return "", err
	}
	return existingUser, nil
}

func (p *Provider) CreateUser(username, password string) error {
	_, err := p.db.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", username, password)
	return err
}
