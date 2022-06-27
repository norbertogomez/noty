package controller

import (
	"net/http"
	"noty-users/internal/auth"
	"noty-users/internal/dto"
	"noty-users/internal/service"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService   auth.JWTService
}

func LoginHandler(
	loginService service.LoginService,
	jWtService auth.JWTService,
) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jWtService,
	}
}

func LoginUser(loginHandler LoginController) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := loginHandler.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credential dto.LoginCredentials
	err := ctx.ShouldBind(&credential)
	if err != nil {
		return "no data found"
	}
	isUserAuthenticated := controller.loginService.LoginUser(credential.Email, credential.Password)
	if isUserAuthenticated {
		return controller.jwtService.GenerateToken(credential.Email, true)

	}
	return ""
}
