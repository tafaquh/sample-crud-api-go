package controller

import (
	"net/http"

	"coba-go/services"
	"coba-go/utils/errors"

	"github.com/gin-gonic/gin"

	loginDomain "coba-go/repository/login"
	"coba-go/utils/response"
)

func Login(c *gin.Context) {
	var requestBody loginDomain.LoginRequest
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, &errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		})
		return
	}
	result, errLogin := services.LoginService.Login(&requestBody)
	if errLogin != nil {
		c.JSON(errLogin.Status, &errors.RestErr{
			Message: errLogin.Message,
			Status:  errLogin.Status,
			Error:   errLogin.Error,
		})
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponseMap(result, "succesfully login"))
}
