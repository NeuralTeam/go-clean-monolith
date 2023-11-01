package v1

import (
	. "go-clean-monolith/__/dto/requests"
	"go-clean-monolith/internal/service"
	. "go-clean-monolith/pkg/httpserver"
	"go-clean-monolith/pkg/logger"
)

// UsersController description
// @RestController
type UsersController struct {
	usersService *service.UsersService
	logger       logger.Logger
}

type PathInt int
type QueryString string
type RequestSchema any

// RegisterAnAccount description
// @Http POST /register
func (c *UsersController) RegisterAnAccount(body UsersQueryRequest) Response {
	return Success(200)
}

// LoginInAccount description
// @Http POST /login
func (c *UsersController) LoginInAccount(name QueryString) Response {
	return Success(200)
}

// GetProfile description
// @Http GET /profile/{id}
func (c *UsersController) GetProfile(id PathInt) Response {
	return Success(200)
}
