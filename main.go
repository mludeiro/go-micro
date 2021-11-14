package main

import (
	"go-micro/controller"
	"go-micro/repository"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "go-micro ", log.LstdFlags)

	repository.Initialize()
	repository.Migrate()

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	cont := controller.ArticlesController{Logger: logger}

	sm.Methods(http.MethodGet).Subrouter().HandleFunc("/articles", cont.GetArticles)
	sm.Methods(http.MethodPost).Subrouter().HandleFunc("/articles", cont.PostArticle)

	s := http.Server{
		Addr:         ":9000",           // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     logger,            // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		logger.Println("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			logger.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	signal := <-c
	log.Println("Got signal:", signal)
}
