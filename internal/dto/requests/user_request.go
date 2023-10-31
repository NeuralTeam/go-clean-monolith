package dto

type UsersQueryRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int8   `json:"age" binding:"required"`
}
