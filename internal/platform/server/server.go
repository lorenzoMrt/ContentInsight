package server

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	cr "github.com/lorenzoMrt/ContentInsight/internal"
	"github.com/lorenzoMrt/ContentInsight/internal/platform/server/handler/contents"
	"github.com/lorenzoMrt/ContentInsight/internal/platform/server/handler/health"
)

type Server struct {
	httpAddr          string
	engine            *gin.Engine
	contentRepository cr.ContentRepository
}

func New(host string, port uint, contentRepository cr.ContentRepository) Server {
	srv := Server{
		engine:   gin.New(),
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
	s.engine.GET("/health", health.CheckHandler())
	s.engine.POST("/api/contents", contents.CreateHandler(s.contentRepository))
}
