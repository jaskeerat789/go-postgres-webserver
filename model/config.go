package model

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/hashicorp/go-hclog"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Log hclog.Logger
	DB  *gorm.DB
}

type DBEntity interface {
	ToJson(w io.Writer) error
}

func NewDBInstance(l hclog.Logger) *DB {
	return &DB{Log: l}
}

func NewPerson(r *http.Request) Person {
	var person Person
	json.NewDecoder(r.Body).Decode(&person)
	return person
}

func (db *DB) Connect() error {
	dsn := "host=localhost user=postgres password=abc123 dbname=microservice_1 port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	dbConnection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		db.Log.Error("Failed to connect with db")
		return err
	} else {
		db.DB = dbConnection
		db.Log.Info("Connected to DB successfully")
	}
	return nil
}

func (p *People) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Person) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}
