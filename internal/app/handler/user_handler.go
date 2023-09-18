package handler

import (
	"fmt"
	"net/http"
	"rest_api/internal/app/dto"
	"rest_api/internal/app/service"
	utils "rest_api/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (u *UserHandler) CreateUser(c *gin.Context) {
	logger := utils.NewLogger(utils.Info)
	logger.Info("Connected To Database")
	var user dto.CreateUserDTO
	if err := c.ShouldBindJSON(&user); err != nil {
		message := fmt.Sprintf("Error creating user: %v", err)
		logger.Warning(message)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if err := u.userService.CreateUser(&user); err != nil {
		message := fmt.Sprintf("Error creating user: %v", err)
		logger.Error(message)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Created user successfully", "data": user})

}
