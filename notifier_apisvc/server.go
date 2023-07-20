package apisvc

import (
	"github.com/dhij/go-notifier"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

type Server struct {
	handler *Handler
}

func NewServer(client notifier.NotifierClient) *Server {
	handler := New(client)

	return &Server{
		handler: handler,
	}
}

func (s *Server) InitRouter() {
	r = gin.Default()

	r.POST("/notify", s.handler.Notify)
}

func (s *Server) Start(addr string) error {
	return r.Run(addr)
}
