package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	kitlog "github.com/go-kit/log"
	"github.com/lorenzoMrt/ContentInsight/internal/creating"
	"github.com/lorenzoMrt/ContentInsight/internal/health"
	"github.com/lorenzoMrt/ContentInsight/kit/command"
)

type ServerConfig struct {
	httpAddr        string
	shutdownTimeout time.Duration
}

type Server struct {
	config ServerConfig
	engine *gin.Engine
	logger kitlog.Logger
	// deps
	commandBus    command.Bus
	healthService health.Service
}

func New(ctx context.Context, host string, port uint, shutdownTimeout time.Duration, commandBus command.Bus, logger kitlog.Logger, healthService health.Service) (context.Context, Server) {
	cfg := ServerConfig{
		httpAddr: fmt.Sprintf("%s:%d", host, port),

		shutdownTimeout: shutdownTimeout,
	}
	srv := Server{
		config: cfg,
		engine: gin.New(),
		logger: logger,

		commandBus:    commandBus,
		healthService: healthService,
	}

	srv.registerRoutes()
	return serverContext(ctx), srv
}

func (s *Server) Run(ctx context.Context) error {
	s.logger.Log("Server running on", s.config.httpAddr)

	srv := &http.Server{
		Addr:    s.config.httpAddr,
		Handler: s.engine,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server shut down", err)
		}
	}()

	<-ctx.Done()
	ctxShutDown, cancel := context.WithTimeout(context.Background(), s.config.shutdownTimeout)
	defer cancel()

	return srv.Shutdown(ctxShutDown)
}

func (s *Server) registerRoutes() {
	httpLogger := kitlog.With(s.logger, "component", "http")
	s.engine.GET("/health", health.MakeHandler(s.healthService, httpLogger))
	s.engine.POST("/contents/v1/", creating.MakeHandler(s.commandBus, httpLogger))
}

func serverContext(ctx context.Context) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}
