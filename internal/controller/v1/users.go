package v1

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
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
		Q int `form:"q" binding:"required"`
	}

	var request requestQuery
	if err := ctx.ShouldBindQuery(&request); err != nil {
		var (
			details []*ValidationErrDetail
			vErrs   validator.ValidationErrors
		)
		if errors.As(err, &vErrs) {
			details = ValidationErrorDetails(&request, "form", vErrs)
		}
		for _, detail := range details {
			fmt.Println(detail.Message, detail.Value, detail.Field)
		}
		return Error(400, 5)
	}

	return SuccessJSON(200, gin.H{"username": 5 / (request.Q - 5)})
}
