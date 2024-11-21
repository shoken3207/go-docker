package auth

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct{}

type User struct{
	Id string `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}
var users []User

type RegisterRequest struct{
	Name string `json:"name" binding:"required,min=3,max=100"`
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=50"`
}

type LoginRequest struct{
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6,max=50"`
}

type ErrorResponse struct{
	Error string `json:"error"`
	Message string `json:"message"`
}

func (h *AuthHandler) GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}


func (h *AuthHandler) Register(c *gin.Context) {
	var request RegisterRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid Request",
			Message: err.Error(),
		})
	}


	id := uuid.New()
	fmt.Print(request)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Server Error",
			Message: err.Error(),
		})
	}

	newUser := User{
		Id: id.String(),
		Name: request.Name,
		Email: request.Email,
		Password: string(hashedPassword),
	}



	users = append(users, newUser)
	c.JSON(http.StatusOK, newUser)
}

func findUserByEmail(email string) (*User, error) {
	for _, user := range users {
		if user.Email == email {
			return &user, nil
		}
	}

	return nil, fmt.Errorf("user not found")
}

func (h *AuthHandler)Login(c *gin.Context){
	var request LoginRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid Request",
			Message: err.Error(),
		})
		return;
	}

	user, err := findUserByEmail(request.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error: "Invalid credentials",
			Message: err.Error(),
		})
		return;
	}

	
	if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, ErrorResponse{
			Error: "Invalid credentials",
			Message: err.Error(),
		})
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	fmt.Println("token",token)

	err = godotenv.Load(".env")
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "server error",
			Message: err.Error(),
		})
		return;
	}
	secretKey := os.Getenv("SECRET_KEY")
	fmt.Print("key: ", secretKey)

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Failed to generate token",
			Message: err.Error(),
		})
		return;
	}
	fmt.Print("token",tokenString)
	c.SetCookie("token", tokenString, 3600 * 24, "/", "localhost", false, true)

	c.JSON(http.StatusOK, gin.H{"message": "Logged in successfully"})
}

// func authenticateJWT(c *gin.Context) {
// 	tokenString, err := c.Cookie("token")
// 	if err != nil {
// 		c.JSON(http.StatusUnauthorized, ErrorResponse{
// 			Error: "Authorization cookie not found",
// 			Message: err.Error(),
// 		})
// 	}

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHS256); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return 
// 	})
// }

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}