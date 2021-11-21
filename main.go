package main

import (
	"go-micro/database"
	"go-micro/presentation"
	"go-micro/tools"
	"os"
	"os/signal"
)

func main() {
	database.InitializeSqlite()
	database.Migrate()
	database.CreateSampleData()

	go presentation.CreateServer()

	waitForInterruptSignal()
}

func waitForInterruptSignal() {
	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	signal := <-c
	tools.GetLogger().Println("Got signal:", signal)
}
