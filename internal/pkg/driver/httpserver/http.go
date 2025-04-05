package httpserver

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	"starter-go/internal/pkg/config"
	mw "starter-go/internal/pkg/driver/httpserver/middleware"
)

type server struct {
	s *http.Server
	e *gin.Engine
	// port         string
	// readTimeout  time.Duration
	// writeTimeout time.Duration
}

func NewServer() server {
	router := gin.New()

	// recovery middleware
	router.Use(mw.CustomRecovery())

	// custom middlewares
	router.Use(mw.RequestTimer())
	router.Use(mw.CORS())
	router.Use(mw.Headers())
	router.Use(mw.ErrorHandler())

	s := &http.Server{
		Addr:         getPort(),
		Handler:      router,
		ReadTimeout:  time.Duration(config.Server().GetReadTimeout()) * time.Millisecond,
		WriteTimeout: time.Duration(config.Server().GetWriteTimeout()) * time.Millisecond,
		IdleTimeout:  time.Duration(config.Server().GetIdleTimeout()) * time.Millisecond,
	}

	srv := server{
		s: s,
		e: router,
	}

	return srv
}

func (srv server) Engine() *gin.Engine {
	return srv.e
}

func (srv server) Start() {
	if err := srv.s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(1)
	}
}

func (srv server) Stop() {
	ctx := context.Background()
	if err := srv.s.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // fallback default
	}
	return ":" + port
}
