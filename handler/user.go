package handler

import (
	"log"
	"net/http"
	"project/bwastartup/helper"
	"project/bwastartup/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	inputUser := user.RegisterUserInput{}

	err := c.ShouldBindJSON(&inputUser)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errMsg := gin.H{"errors": errors}

		log.Println(err)
		dataErr := helper.APIResponse("Something went wrong!", http.StatusUnprocessableEntity, "error", errMsg)
		c.JSON(http.StatusUnprocessableEntity, dataErr)
		return
	}

	data, err := h.userService.RegisterUser(inputUser)
	if err != nil {
		log.Println(err)
		errors := helper.APIResponse("Something went wrong!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, errors)
		return
	}

	formatter := user.FormatUser(data, "tokentoken")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}

func (h *userHandler) LoginUser(c *gin.Context) {
	input := user.LoginInput{}

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errMsg := gin.H{"errors": errors}

		log.Println(err)
		dataErr := helper.APIResponse("Something went wrong!", http.StatusUnprocessableEntity, "error", errMsg)
		c.JSON(http.StatusUnprocessableEntity, dataErr)
		return
	}

	data, err := h.userService.Login(input)
	if err != nil {
		log.Println(err)
		errMsg := gin.H{"errors": err.Error()}
		errors := helper.APIResponse("Something went wrong!", http.StatusBadRequest, "error", errMsg)
		c.JSON(http.StatusBadRequest, errors)
		return
	}

	formatter := user.FormatUser(data, "tokentoken")

	response := helper.APIResponse("Login Success", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)
}
