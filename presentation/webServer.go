package presentation

import (
	"fmt"
	"go-micro/tools"
	"net/http"
	"time"
)

type WebServer struct {
	Router WebRouter
	server http.Server
}

func (webServer *WebServer) CreateServer() {
	port := 80

	if !tools.IsRunningInDockerContainer() {
		port = 8080
	}

	webServer.server = http.Server{
		Addr:         fmt.Sprintf(":%d", port),     // configure the bind address
		Handler:      webServer.Router.GetRouter(), // set the default handler
		ErrorLog:     tools.GetLogger(),            // set the logger for the server
		ReadTimeout:  5 * time.Second,              // max time to read request from the client
		WriteTimeout: 10 * time.Second,             // max time to write response to the client
		IdleTimeout:  120 * time.Second,            // max time for connections using TCP Keep-Alive
	}

	tools.GetLogger().Println("Starting server on port ", port)

	err := webServer.server.ListenAndServe()
	if err != nil {
		tools.GetLogger().Printf("Error starting server: %s\n", err)
	}
}
