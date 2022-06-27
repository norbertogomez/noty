package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"noty-common/pkg/middleware"
	"noty-notifications/internal/controller"
	"noty-notifications/internal/repository"
	"noty-notifications/internal/service"
)

type services struct {
	notificationSvc service.NotificationService
}

func main() {
	svcs := buildServices()

	server := gin.Default()
	server.SetTrustedProxies(nil)
	server.Use(middleware.Recovery())

	registerRoutes(server, svcs)
	port := "8082"
	if err := server.Run(":" + port); err != nil {
		panic(err)
	}
}

func registerRoutes(server *gin.Engine, svcs services) {
	noAuth := server.Group("/")
	noAuth.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	noAuth.POST("/notifications", controller.CreateNotification(svcs.notificationSvc))
}

func buildServices() services {
	notificationRepo := repository.NewNotificationRepository()
	var notificationService = service.NewNotificationService(notificationRepo)

	return services{
		notificationSvc: notificationService,
	}
}
