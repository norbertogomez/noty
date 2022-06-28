package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"noty-notifications/internal/service"
)

func GetNotification(service service.NotificationService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		notificationID := ctx.Param("id")
		if notificationID == "" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "Empty id",
			})
		}

		notification, err := service.FindByID(notificationID)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
		}

		if notification.ID == "" {
			ctx.JSON(http.StatusNotFound, gin.H{})
		}

		ctx.JSON(http.StatusOK, notification)
		return
	}
}
