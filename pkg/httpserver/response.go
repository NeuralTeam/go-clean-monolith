package httpserver

import "github.com/gin-gonic/gin"

type Json map[string]any

type Response struct {
	HttpCode  int
	ErrorCode uint
	Data      interface{}
}

func Success(httpCode int) Response {
	return Response{httpCode, 0, nil}
}

func SuccessJSON(httpCode int, data interface{}) Response {
	return Response{httpCode, 0, data}
}

func SuccessPlain(httpCode int, text string) Response {
	return Response{httpCode, 0, text}
}

func Error(httpCode int, errorCode uint) Response {
	return Response{httpCode, errorCode, nil}
}

func HandlerWrapper(f func(c *gin.Context) Response) gin.HandlerFunc {
	return func(c *gin.Context) {
		response := f(c)

		// Error
		if response.ErrorCode != 0 {
			c.AbortWithStatusJSON(response.HttpCode, Json{"error": response.ErrorCode})
			return
		}

		// Success
		if response.Data == nil {
			c.Status(response.HttpCode)
			return
		}

		// SuccessPlain
		if str, ok := response.Data.(string); ok {
			c.String(response.HttpCode, str)
			return
		}

		// SuccessJSON
		c.JSON(response.HttpCode, response.Data)
		return
	}
}
