package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func NewProxy(target string) gin.HandlerFunc {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("[PROXY] Failed to parse target URL: %s, error: %v", target, err)
	}

	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("[PROXY] Error detail: %v", err)
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(`{"error":"upstream service unavailable"}`))
	}

	return func(c *gin.Context) {
		log.Printf("[PROXY] %s %s → %s", c.Request.Method, c.Request.URL.Path, target)
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}
