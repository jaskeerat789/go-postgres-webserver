package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

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
	host := os.Getenv("HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DBNAME")
	port := os.Getenv("PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata", host, username, password, DBName, port)
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
