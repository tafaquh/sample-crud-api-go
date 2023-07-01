package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	regisDomain "coba-go/repository/register"
	"coba-go/services"
	"coba-go/utils/errors"
)

func RegisterUser(c *gin.Context) {
	var requestBody regisDomain.RegisterUserRequest
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, &errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		})
		return
	}
	result, errInsert := services.RegisterService.CreateUser(&requestBody)
	if errInsert != nil {
		c.JSON(errInsert.Status, &errors.RestErr{
			Message: errInsert.Message,
			Status:  errInsert.Status,
			Error:   errInsert.Error,
		})
		return
	}
	c.JSON(http.StatusCreated, result)
}
