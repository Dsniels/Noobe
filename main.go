package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/dsniels/noobe/router"
)

func main() {
	slog.Info("Hi")

	port := flag.String("port", "80", "define the port")
	flag.Parse()

	listenPort := fmt.Sprintf(":%s", *port)

	routes := router.InitRoutes()
	server := &http.Server{
		Handler: routes,
		Addr:    listenPort,
	}
	slog.Info("server listening ", slog.String("PORT", *port))
	if err := server.ListenAndServe(); err != nil {
		slog.Error(err.Error())
	}

}
