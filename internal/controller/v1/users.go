package v1

import (
	"github.com/gin-gonic/gin"
	requests "go-clean-monolith/internal/dto/requests"
	responses "go-clean-monolith/internal/dto/responses"
	"go-clean-monolith/internal/service"
	. "go-clean-monolith/pkg/httpserver"
	"go-clean-monolith/pkg/logger"
)

// UsersController description
type UsersController struct {
	usersService *service.UsersService
	logger       logger.Logger
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
	if err := BindJSON(ctx, &request); err != nil {
		return SuccessJSON(400, gin.H{"error": err.Error()})
	}

	return SuccessJSON(200, responses.UsersProfileResponse{Name: request.Name, Age: request.Age})
}
