package main

import (
	"github.com/gin-gonic/gin"
	"github.com/penglongli/gin-metrics/ginmetrics"
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
	// get global Monitor object
	m := ginmetrics.GetMonitor()

	// +optional set metric path, default /debug/metrics
	m.SetMetricPath("/metrics")
	// +optional set slow time, default 5s
	m.SetSlowTime(10)
	// +optional set request duration, default {0.1, 0.3, 1.2, 5, 10}
	// used to p95, p99
	m.SetDuration([]float64{0.1, 0.3, 1.2, 5, 10})

	// set middleware for gin
	m.Use(server)
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

	noAuth.POST("/notification", controller.CreateNotification(svcs.notificationSvc))
	noAuth.PUT("/notification/:id", controller.UpdateNotification(svcs.notificationSvc))
	noAuth.GET("/notification/:id", controller.GetNotification(svcs.notificationSvc))
}

func buildServices() services {
	notificationRepo := repository.NewNotificationRepository()
	var notificationService = service.NewNotificationService(notificationRepo)

	return services{
		notificationSvc: notificationService,
	}
}
