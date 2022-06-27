package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	common "noty-common/pkg/middleware"
	"noty-users/internal/auth"
	"noty-users/internal/controller"
	"noty-users/internal/middleware"
	"noty-users/internal/repository"
	"noty-users/internal/service"
)

type handlers struct {
	loginHandler controller.LoginController
}

type services struct {
	jwtSvc   auth.JWTService
	loginSvc service.LoginService
	userSvc  service.UserService
}

func main() {
	svcs := buildServices()
	h := buildHandlers(svcs)

	server := gin.Default()
	server.SetTrustedProxies(nil)

	registerNoAuthRoutes(server, h, svcs)
	registerAuthRoutes(server)
	server.Use(common.Recovery())
	port := "8080"
	if err := server.Run(":" + port); err != nil {
		panic(err)
	}
}

func buildHandlers(svcs services) handlers {
	var loginHandler = controller.LoginHandler(svcs.loginSvc, svcs.jwtSvc)

	return handlers{
		loginHandler: loginHandler,
	}
}

func buildServices() services {
	userRepo := repository.NewUserRepository()
	var UserService = service.NewUserService(userRepo)
	var loginService = service.NewLoginUserService(userRepo)
	var jwtService = auth.JWTAuthService()
	return services{
		jwtSvc:   jwtService,
		loginSvc: loginService,
		userSvc:  UserService,
	}
}

func registerAuthRoutes(server *gin.Engine) {
	authGroup := server.Group("/v1")
	authGroup.Use(middleware.AuthorizeJWT())
	authGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func registerNoAuthRoutes(server *gin.Engine, h handlers, s services) {
	noAuth := server.Group("/")
	noAuth.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong no auth",
		})
	})

	noAuth.POST("/users", controller.CreateUser(s.userSvc))
	noAuth.POST("/login", controller.LoginUser(h.loginHandler))
}
