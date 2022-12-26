package main

import (
	"home/bjit/Documents/Golang_Practice/3_Gin_Projects/YouTube_Pragmatic_Reviews/golang-gin-poc-jwt/controller"
	"home/bjit/Documents/Golang_Practice/3_Gin_Projects/YouTube_Pragmatic_Reviews/golang-gin-poc-jwt/middlewares"
	"home/bjit/Documents/Golang_Practice/3_Gin_Projects/YouTube_Pragmatic_Reviews/golang-gin-poc-jwt/service"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	videoService service.VideoService = service.New()
	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

func main() {

	server := gin.New()

	server.Use(gin.Recovery(), gin.Logger())

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	// Login Endpoint: Authentication + Token creation
	server.POST("/login", func(ctx *gin.Context) {
		token := loginController.Login(ctx)
		if token != "" {
			ctx.JSON(http.StatusOK, gin.H{
				"token": token,
			})
		} else {
			ctx.JSON(http.StatusUnauthorized, nil)
		}
	})

	// JWT Authorization Middleware applies to "/api" only.
	apiRoutes := server.Group("/api", middlewares.AuthorizeJWT())
	{
		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, videoController.FindAll())
		})

		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := videoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid!!"})
			}

		})
	}

	// The "/view" endpoints are public (no Authorization required)
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", videoController.ShowAll)
	}

	// We can setup this env variable from the EB console
	port := os.Getenv("PORT")
	// Elastic Beanstalk forwards requests to port 5000
	if port == "" {
		port = "5000"
	}
	server.Run(":" + port)
}

// go mod init home/bjit/Documents/Golang_Practice/3_Gin_Projects/YouTube_Pragmatic_Reviews/golang-gin-poc-jwt
// export GO111MODULE="on"
// go get github.com/gin-gonic/gin
// go get github.com/tpkeeper/gin-dump
// go get github.com/dgrijalva/jwt-go
// go get github.com/stretchr/testify/assert
// go get github.com/go-playground/validator
