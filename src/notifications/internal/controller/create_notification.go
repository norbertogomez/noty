package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"noty-notifications/internal/dto"
	"noty-notifications/internal/service"
)

func CreateNotification(service service.NotificationService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var notification dto.NotificationCreationInfo
		if err := ctx.ShouldBind(&notification); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid payload",
			})
		}

		ok, err := service.CreateNotification(notification)
		if ok {
			ctx.JSON(http.StatusCreated, gin.H{})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}
