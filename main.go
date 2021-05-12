package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	goHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/hashicorp/go-hclog"
	handelers "github.com/jaskeerat789/go-postgres-webserver/handlers"
	"github.com/jaskeerat789/go-postgres-webserver/model"
	"github.com/joho/godotenv"
)

func main() {

	// load env variables
	godotenv.Load()

	// initialize logger
	l := hclog.New(&hclog.LoggerOptions{
		Name:  "webserver",
		Level: hclog.LevelFromString("DEBUG"),
	})

	// DB Connection
	db := model.NewDBInstance(l)
	db.Connect()

	// configure handlers
	// bh := handelers.NewBookHandler(l, db)
	ph := handelers.NewPersonHandler(l, db)

	// create a new server mux
	sm := mux.NewRouter()
	getRouter := sm.Methods("GET").Subrouter()
	postRouter := sm.Methods("POST").Subrouter()
	// postRouter := sm.Methods("POST").Subrouter()

	// register handelers
	getRouter.HandleFunc("/people", ph.GetPeople)
	getRouter.HandleFunc("/person/{id:[0-9]+}", ph.GetPerson)

	postRouter.HandleFunc("/person", ph.CreatePerson)
	// CORS
	goHandlers.CORS()

	// create http server
	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		ErrorLog:     l.StandardLogger(&hclog.StandardLoggerOptions{}),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		l.Info("Starting server on Port 8080")
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Info("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
