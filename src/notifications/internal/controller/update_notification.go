package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"noty-notifications/internal/dto"
	"noty-notifications/internal/service"
)

func UpdateNotification(service service.NotificationService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		notificationID := ctx.Param("id")
		var notification dto.NotificationUpdateInfo
		if err := ctx.ShouldBind(&notification); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Invalid payload",
			})
		}

		notification.ID = notificationID
		ok, err := service.UpdateNotification(notification)
		if ok {
			ctx.JSON(http.StatusOK, gin.H{})
			return
		}

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
}
