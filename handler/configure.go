package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/thanhptse/go-template/config"
)

type Handler struct {
	cfg    *config.AppConfig
	router *gin.Engine
}

func NewHandler(cfg *config.AppConfig, router *gin.Engine) *Handler {
	return &Handler{
		cfg:    cfg,
		router: router,
	}
}

func (h *Handler) ConfigureRoute(router *gin.Engine) {
	routers := router.Group("v1")

	routers.GET("/ping", h.ping())
}
