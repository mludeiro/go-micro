package main

import (
	"go-micro/tools"
	"os"
	"os/signal"
)

func main() {
	cont := NewContainer()

	cont.DataBase.InitializeSqlite().CreateSampleData()
	// 	cont.DataBase.InitializePostgress().CreateSampleData()

	go cont.WebServer.CreateServer()

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
