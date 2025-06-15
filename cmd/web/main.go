package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	port := flag.String("port", ":4000", "Network Address Port")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{}))

	app := &application{
		logger: logger,
	}

	logger.Info("Starting server", slog.String("Port", *port))
	err := http.ListenAndServe(*port, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
