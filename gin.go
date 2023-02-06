package web

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type GinRouteFunc func(route *gin.Engine)

type GinService struct {
	handler *gin.Engine
	server  *http.Server
	opt     *Options
}

func NewGinService(handler *gin.Engine, opt *Options) *GinService {
	return &GinService{handler: handler, opt: opt}
}

func (s *GinService) makeServer() {
	s.server = &http.Server{
		Handler:      s.handler,
		Addr:         fmt.Sprintf("%s:%d", s.opt.Address, s.opt.Port),
		ReadTimeout:  time.Duration(s.opt.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(s.opt.WriteTimeout) * time.Second,
	}

	s.opt.BindRoutes(s.handler)
}

func (s *GinService) Boot() (err error) {
	s.makeServer()

	err = s.server.ListenAndServe()
	if err == http.ErrServerClosed {
		return nil
	}

	return
}

func (s *GinService) Routers(routeFunc GinRouteFunc) {
	routeFunc(s.handler)
}

func (s *GinService) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	if err := s.server.Shutdown(ctx); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}
}
