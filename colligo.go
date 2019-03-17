package main

import (
	"flag"
	"os"
	"os/signal"

	"github.com/PeterBooker/colligo/internal/config"
	"github.com/PeterBooker/colligo/internal/log"
	"github.com/PeterBooker/colligo/internal/server"
)

//go:generate go run -tags=dev embed_files.go

var (
	version string
	commit  string
	date    string
	env     string
)

func main() {

	flag.StringVar(&env, "env", "dev", "Enviroment, e.g. producton or development.")
	flag.Parse()

	// Create Logger
	l := log.New()

	l.Println("Starting Colligo...")

	// Create Config
	c := config.New(version, commit, date, env)

	// Create Server
	s := server.New(l, c)

	// Graceful Shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	// Setup HTTP server.
	s.Setup()

	<-stop

	l.Println("Stopping Colligo...")

}
