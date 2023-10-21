package v1

import (
	"github.com/gin-gonic/gin"
	"go-clean-monolith/internal/service"
	. "go-clean-monolith/pkg/httpserver"
	"go-clean-monolith/pkg/logger"
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
	type requestQuery struct {
		Q string `query:"q" binding:"default:neuralteam,min:10,max:12,required"`
	}

	var request requestQuery
	if err := BindQuery(ctx, &request); err != nil {
		return SuccessJSON(400, gin.H{"error": err.Error()})
	}

	return SuccessJSON(200, gin.H{"username": request.Q})
}
