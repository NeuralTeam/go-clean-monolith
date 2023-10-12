package v1

import (
	"go-clean-monolith/controller/middleware"
	"go-clean-monolith/pkg/httpserver"
)

//----------------------------------------------------------------------------------------------------------------------

// Router struct
type Router struct {
	httpServer      httpserver.HttpServer
	middlewares     middleware.Middleware
	usersController IUsersController
}

// NewRouter description
func NewRouter(
	httpServer httpserver.HttpServer,
	middlewares middleware.Middleware,
	userController IUsersController,
) Router {
	return Router{
		httpServer:      httpServer,
		middlewares:     middlewares,
		usersController: userController,
	}
}

//----------------------------------------------------------------------------------------------------------------------

// Setup description
func (r Router) Setup() {
	api := r.httpServer.Group("/api/v1")

	users := api.Group("/users")
	{
		users.POST("/register", httpserver.HandlerWrapper(r.usersController.RegisterAnAccount))
		users.POST("/login", httpserver.HandlerWrapper(r.usersController.LoginInAccount))
		users.GET("/profile", httpserver.HandlerWrapper(r.usersController.GetProfile))
	}
}

//----------------------------------------------------------------------------------------------------------------------
