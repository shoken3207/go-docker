package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type CreateUserRequest struct {
	Name  string `json:"name" binding:"required,min=3,max=100"`
	Email string `json:"email" binding:"required,email"`
	Age   int    `json:"age" binding:"required"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

var users = []User{
	{Id: 1, Name: "Alice", Email: "alice@example.com", Age: 25},
	{Id: 2, Name: "Bob", Email: "bob@example.com", Age: 30},
	{Id: 3, Name: "Charlie", Email: "charlie@example.com", Age: 35},
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "idが不正です", Message: "idは1以上の整数である必要があります。"})
	}

	for _, user := range users {
		if user.Id == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}

	c.JSON(http.StatusNotFound, ErrorResponse{Error: "ユーザーが見つかりません。", Message: "idが同じユーザーが存在しません。"})
}

func createUser(c *gin.Context) {
	var request CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid Request",
			Message: err.Error(),
		})
	}

	maxId := 0
	for _, user := range users {
		if user.Id > maxId {
			maxId = user.Id
		}
	}

	var newUser User
	c.ShouldBindJSON(&newUser)
	newUser.Id = maxId

	users = append(users, newUser)
	c.JSON(http.StatusOK, newUser)
}

func updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid Request", Message: "Idが不正な値です。"})
	}

	var request CreateUserRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error:   "Invalid Request",
			Message: err.Error(),
		})
	}

	var updatedUser User
	c.ShouldBindJSON(&updatedUser)
	updatedUser.Id = id

	for i, user := range users {
		if user.Id == id {
			users[i] = updatedUser
			c.JSON(http.StatusOK, updatedUser)
			return
		}
	}

	c.JSON(http.StatusNotFound, ErrorResponse{Error: "ユーザーが見つかりません。", Message: "idが同じユーザーが存在しません。"})
}

// func deleteUser(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))
// 	if err != nil || id <= 0 {
// 		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "invalid Request", Message: "Idが不正な値です。"})
// 	}

// 	for i, user := range users {
// 		if user.Id == id {
// 			users = append(users[:i], users[i+1:]...)

// 			users[i] = updatedUser
// 			c.JSON(http.StatusOK, updatedUser)
// 			return
// 		}
// 	}

// 	c.JSON(http.StatusNotFound, ErrorResponse{Error: "ユーザーが見つかりません。", Message: "idが同じユーザーが存在しません。"})
// }

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}
