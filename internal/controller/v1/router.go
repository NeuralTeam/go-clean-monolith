package v1

import (
	"go-clean-monolith/internal/controller/middleware"
	"go-clean-monolith/pkg/httpserver"
)

//----------------------------------------------------------------------------------------------------------------------

// Router description
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
		users.POST("/register", r.usersController.RegisterAnAccount)
		users.POST("/login", r.usersController.LoginInAccount)
		users.GET("/profile", r.usersController.GetProfile)
	}
}

//----------------------------------------------------------------------------------------------------------------------
