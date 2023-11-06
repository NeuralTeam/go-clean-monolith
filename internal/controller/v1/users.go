package v1

import (
	"github.com/gin-gonic/gin"
	requests "go-clean-monolith/dto/requests"
	responses "go-clean-monolith/dto/responses"
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
	var request requests.UsersQueryRequest
	if err := BindQuery(ctx, &request); err != nil {
		return SuccessJSON(400, gin.H{"error": err.Error()})
	}
	return SuccessJSON(200, responses.UsersProfileResponse{Name: request.Name, Age: 4 / request.Age})
}
