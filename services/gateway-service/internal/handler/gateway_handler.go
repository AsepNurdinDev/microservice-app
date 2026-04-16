package handler

import (
	"gateway-service/internal/config"
	"gateway-service/internal/proxy"

	"github.com/gin-gonic/gin"
)

type GatewayHandler struct {
	cfg config.Config
}

func NewGatewayHandler() *GatewayHandler {
	return &GatewayHandler{
		cfg: config.Load(),
	}
}

// AUTH (public)
func (h *GatewayHandler) AuthProxy() gin.HandlerFunc {
	return proxy.NewProxy(h.cfg.AuthService) 
}

// ARTICLE (protected)
func (h *GatewayHandler) ArticleProxy() gin.HandlerFunc {
	return proxy.NewProxy(h.cfg.ArticleService)
}