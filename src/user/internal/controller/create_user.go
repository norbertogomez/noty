package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"noty-users/internal/dto"
	"noty-users/internal/service"
)

func CreateUser(service service.UserService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user dto.UserCreationInfo
		if err := ctx.ShouldBind(&user); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid payload",
			})
		}

		ok, err := service.CreateUser(user)
		if ok {
			ctx.JSON(http.StatusCreated, gin.H{})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}
