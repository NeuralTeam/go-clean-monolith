package dto

type UsersQueryRequest struct {
	Name string `query:"name" binding:"required"`
	Age  int8   `query:"age" binding:"required"`
}
