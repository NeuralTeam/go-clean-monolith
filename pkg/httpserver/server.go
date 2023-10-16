package httpserver

import (
	"github.com/gin-gonic/gin"
)

type (
	// HandlerFunc wrapper for handler
	HandlerFunc func(*gin.Context) Response

	// HttpServer -> Gin Engine
	HttpServer struct {
		engine *gin.Engine
		router *gin.RouterGroup
	}
)

// New create HttpServer
func New() HttpServer {
	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	return HttpServer{
		engine: e,
		router: &e.RouterGroup,
	}
}

// Handler get Engine from server
func (s *HttpServer) Handler() *gin.Engine {
	return s.engine
}

// Routes gets all routes
func (s *HttpServer) Routes() gin.RoutesInfo {
	return s.engine.Routes()
}

// Use add middleware to Group of routes
func (s *HttpServer) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return s.router.Use(middleware...)
}

// EngineUse add middleware to all routes
func (s *HttpServer) EngineUse(middleware ...gin.HandlerFunc) gin.IRoutes {
	return s.engine.Use(middleware...)
}

// Group creates a new router group. You should add all the routes that have common middlewares or the same path prefix.
func (s *HttpServer) Group(relativePath string, handlers ...gin.HandlerFunc) *HttpServer {
	return &HttpServer{
		engine: s.engine,
		router: s.router.Group(relativePath, handlers...),
	}
}

// POST is a shortcut for router.Handle("POST", path, handlers).
func (s *HttpServer) POST(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return s.router.POST(relativePath, convertHandlers(handlers)...)
}

// GET is a shortcut for router.Handle("GET", path, handlers).
func (s *HttpServer) GET(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return s.router.GET(relativePath, convertHandlers(handlers)...)
}

// DELETE is a shortcut for router.Handle("DELETE", path, handlers).
func (s *HttpServer) DELETE(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return s.router.DELETE(relativePath, convertHandlers(handlers)...)
}

// PATCH is a shortcut for router.Handle("PATCH", path, handlers).
func (s *HttpServer) PATCH(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return s.router.PATCH(relativePath, convertHandlers(handlers)...)
}

// PUT is a shortcut for router.Handle("PUT", path, handlers).
func (s *HttpServer) PUT(relativePath string, handlers ...HandlerFunc) gin.IRoutes {
	return s.router.PUT(relativePath, convertHandlers(handlers)...)
}

func convertHandlers(handlers []HandlerFunc) []gin.HandlerFunc {
	var wrappedHandlers []gin.HandlerFunc
	for _, handler := range handlers {
		wrappedHandlers = append(wrappedHandlers, HandlerWrapper(handler))
	}
	return wrappedHandlers
}
