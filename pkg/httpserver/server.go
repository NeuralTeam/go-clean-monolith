package httpserver

import "github.com/gin-gonic/gin"

// HttpServer -> Gin Engine
type HttpServer struct {
	*gin.Engine
}

// New description
func New() HttpServer {
	gin.SetMode(gin.ReleaseMode)

	engine := gin.New()

	return HttpServer{
		engine,
	}
}
