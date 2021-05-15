package server

import (
	"net"
	"net/http"
	"os"

	"github.com/thanhptse/go-template/config"
	"github.com/thanhptse/go-template/handler"
	"github.com/thanhptse/go-template/pkg/contxt"
	"github.com/thanhptse/go-template/pkg/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Server struct {
	cfg        *config.AppConfig
	httpServer *http.Server
	router     *gin.Engine
}

func NewServer(cfg *config.AppConfig) (*Server, error) {
	router := gin.New()

	router.Use(middleware.SetRequestID())
	router.Use(contxt.SetupAppContext())
	router.Use(middleware.SetupLog())
	router.Use(gin.Recovery())

	s := &Server{
		router: router,
		cfg:    cfg,
	}

	return s, nil
}

func (s *Server) Init() {
	s.initRouters()
}

func (s *Server) ListenHTTP() error {
	address := ":" + os.Getenv("PORT")
	listen, err := net.Listen("tcp", address)
	if err != nil {
		panic(err)
	}

	s.httpServer = &http.Server{
		Handler: s.router,
		Addr:    address,
	}

	zap.S().Infof("Starting http server at port %s", address)

	return s.httpServer.Serve(listen)
}

func (s *Server) Close() {
	_ = s.httpServer.Close()
}

func (s *Server) initRouters() {
	hdl := handler.NewHandler(s.cfg, s.router)

	hdl.ConfigureRoute(s.router)
}
