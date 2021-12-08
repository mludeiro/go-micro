package main

import (
	"go-micro/container"
	"go-micro/tools"
	"os"
	"os/signal"
)

func main() {
	cont := container.NewContainer()

	cont.DataBase.InitializeSqlite().Migrate().CreateSampleData()
	// 	cont.DataBase.InitializePostgress().Migrate().CreateSampleData()

	go cont.WebServer.CreateServer()

	waitForInterruptSignal()
	 error de compilacion
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
