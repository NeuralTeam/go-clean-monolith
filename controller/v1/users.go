package v1

import (
	"github.com/gin-gonic/gin"
	. "go-clean-monolith/pkg/httpserver"
	"go-clean-monolith/pkg/logger"
	"go-clean-monolith/service"
	"strconv"
)

var _ IUsersController = &UsersController{}

type (
	// UsersController description
	UsersController struct {
		service service.IUsersService
		logger  logger.Logger
	}

	// IUsersController description
	IUsersController interface {
		RegisterAnAccount(ctx *gin.Context) Response
		LoginInAccount(ctx *gin.Context) Response
		GetProfile(ctx *gin.Context) Response
	}
)

// NewUserController description
func NewUserController(service service.IUsersService, logger logger.Logger) IUsersController {
	return &UsersController{
		service: service,
		logger:  logger,
	}
}

// RegisterAnAccount description
func (c *UsersController) RegisterAnAccount(ctx *gin.Context) Response {
	return Error(400, 7787)
}

// LoginInAccount description
func (c *UsersController) LoginInAccount(ctx *gin.Context) Response {
	return Error(400, 7787)
}

// GetProfile description
func (c *UsersController) GetProfile(ctx *gin.Context) Response {
	a, _ := ctx.GetQuery("q")
	b, _ := strconv.Atoi(a)

	return SuccessJSON(200, gin.H{"username": 5 / b})
}
