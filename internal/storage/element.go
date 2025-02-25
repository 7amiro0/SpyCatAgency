package storage

import (
	"fmt"
	"os"
)

type conf struct {
	user     string
	password string
	host     string
	port     string
	name     string
}

func initDB() conf {
	return conf{
		user:     os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		host:     os.Getenv("DB_HOST"),
		port:     os.Getenv("DB_PORT"),
		name:     os.Getenv("DB_NAME"),
	}
}

func (c conf) getLink() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.host,
		c.port,
		c.user,
		c.password,
		c.name,
	)
}

type Cat struct {
	id         int `json:id`
	Experience int	`json:experience`
	Salary     int	`json:salary`
	Name       string `json:name`
	Bread      string `json:bread`
}

type Target struct {
	Name     string
	Country  string
	Notes    string
	Complete bool
}

type Mision struct {
	id       int
	Cat      int
	Complete bool
	Targets  []Target
}
