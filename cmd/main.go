package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	server "starter-go/api/rest"
	"starter-go/api/rest/example"
	domainExample "starter-go/internal/domain/example"
	"starter-go/internal/pkg/app"
	"starter-go/internal/pkg/config"
	"starter-go/internal/pkg/driver/httpserver"
	"starter-go/internal/pkg/logger"
	exampleRepo "starter-go/internal/repository/example"
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
		With("version", AppVersion).
		With("service", AppName))

	// init HTTP Server
	srv := httpserver.NewServer()
	server.RegisterRoutes(srv.Engine())

	// init inmemory repository (for example only, can replace with actual db)
	exampleRepo := inMemoryInit()
	logger.Info("Starting application with inmemory repository")

	// initialize service and handler
	exSvc := exampleService.NewService(exampleRepo)
	exHdlr := example.NewHandler(exSvc)

	example.RegisterRoutes(srv.Engine(), exHdlr)

	apps := []app.App{
		srv,
	}

	stopFn := app.AppController(apps...)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)

	sig := <-quit
	logger.Info(fmt.Sprintf("exiting. received signal: %s", sig.String()))

	stopFn(time.Duration(30) * time.Second)
}

// just for example, should remove if not used
func inMemoryInit() *exampleRepo.MemoryRepository {
	exampleRepository := exampleRepo.NewMemoryRepository()

	// prepopulate with some examples for testing
	ctx := context.Background()
	err := exampleRepository.Preload(
		ctx,
		&domainExample.Example{Description: "Example 1"},
		&domainExample.Example{Description: "Example 2"},
	)
	if err != nil {
		logger.Error("Failed to preload examples", "error", err.Error())
	} else {
		logger.Info("Pre-populated repository with examples")
	}
	return exampleRepository
}
