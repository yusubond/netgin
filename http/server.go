package http

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type (
	Server struct {
		cfg *Config
		*zap.SugaredLogger
		stop chan struct{}

		*http.Server
		*gin.Engine
	}
)

func NewHttpServer() *Server {
	logger, _ := zap.NewDevelopment()
	svr := &Server{
		cfg:           NewDefaultConfig(),
		SugaredLogger: logger.Sugar(),
		stop:          make(chan struct{}),
		Server:        &http.Server{},
		Engine:        gin.New(),
	}
	return svr
}

func (s *Server) Init() error {
	s.Server.Addr = s.Addr()
	s.Server.Handler = s.Engine
	s.initHandlers()
	return nil
}

func (s *Server) Addr() string {
	return fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
}

func (s *Server) Start() {
	s.Infof("start http server")

	go func() {
		if err := s.ListenAndServe(); err != nil && s.IsStopped() {
			s.Fatal("http server start error", zap.Error(err))
		}
	}()

	s.Infof("http server listen at %s", s.Addr())
	return
}

func (s *Server) Stop() {
	s.Infof("stop http server")
	if s.Server != nil {
		ctx, cancel := context.WithTimeout(context.TODO(), 30*time.Second)
		defer cancel()

		if err := s.Server.Shutdown(ctx); err != nil {
			s.Fatal("http server shutdown error", zap.Error(err))
		}
		s.Infof("http server has quit.")
	}

	close(s.stop)
	return
}

func (s *Server) IsStopped() bool {
	select {
	case <-s.stop:
		return true
	default:
		return false
	}
}
