package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	server "starter-go/api/rest"
	"starter-go/api/rest/example"
	exampleHandler "starter-go/api/rest/example"
	"starter-go/internal/pkg/app"
	"starter-go/internal/pkg/config"
	"starter-go/internal/pkg/driver/httpserver"
	"starter-go/internal/pkg/logger"
	exampleService "starter-go/internal/service/example"
)

// build vars
const (
	GitCommit  = "unknown"
	AppVersion = "unknown"
	AppName    = "starter-go"
)

func main() {
	cfgPath := flag.String("configpath", "./config/config.yaml", "path to config file")
	flag.Parse()

	err := config.Init(*cfgPath)
	if err != nil {
		panic(fmt.Errorf("error parsing config. %w", err))
	}

	buildInfo := map[string]string{
		"commit": GitCommit,
	}

	newLogger, err := logger.NewFromConfig(config.LoggerConfig())
	if err != nil {
		panic(fmt.Errorf("failed to create logger"))
	}
	defer newLogger.Stop()

	logger.SetDefaultLogger(newLogger.
		With("buildinfo", buildInfo).
		// With("country", "ID").
		With("version", AppVersion).
		With("service", AppName))

	srv := httpserver.NewServer()

	apps := []app.App{
		srv,
	}

	server.RegisterRoutes(srv.Engine())
	exSvc := exampleService.NewService(nil)
	exHdlr := exampleHandler.NewHandler(exSvc)
	example.RegisterRoutes(srv.Engine(), exHdlr)

	stopFn := app.AppController(apps...)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)

	sig := <-quit
	logger.Info(fmt.Sprintf("exiting. received signal: %s", sig.String()))

	stopFn(time.Duration(30) * time.Second)
}
