package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lorenzoMrt/ContentInsight/internal/creating"
	"github.com/lorenzoMrt/ContentInsight/internal/platform/server/handler/contents"
	"github.com/lorenzoMrt/ContentInsight/internal/platform/server/handler/health"
	"github.com/lorenzoMrt/ContentInsight/internal/platform/server/middleware/logging"
	"github.com/lorenzoMrt/ContentInsight/internal/platform/server/middleware/recovery"
)

type Server struct {
	httpAddr string
	engine   *gin.Engine

	//deps
	contentService creating.ContentService
}

func New(host string, port uint, contentService creating.ContentService) Server {
	srv := Server{
		engine:   gin.New(), //Can use gin.Default() for out of the box middlewares
		httpAddr: fmt.Sprintf("%s:%d", host, port),
	}

	srv.registerRoutes()
	return srv
}

func (s *Server) Run() error {
	log.Println("Server running on", s.httpAddr)
	return s.engine.Run(s.httpAddr)
}

func (s *Server) registerRoutes() {
	s.engine.Use(recovery.Middleware(), logging.Middleware())

	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/api/contents", contents.CreateHandler(s.contentService))
}
