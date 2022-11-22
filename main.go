package main

import (
	"GinCasbin/internal"
	"GinCasbin/middleware"

	"github.com/gin-gonic/gin"
	selfLogger "github.com/thomascwei/golang_logger"
)

var (
	// create main.log
	Logger = selfLogger.InitLogger("main")
)

func main() {
	Logger.Info("starting...")
	r := gin.Default()
	r.GET("/", internal.HellowWorldHandler)

	users := r.Group("/v1/users")
	users.POST("/login", internal.LoginHandler)

	data := r.Group("/v1/data")
	data.GET("/:id", internal.JWTAuthMiddleware(), middleware.AuthorizeMiddleware("data", "GET"), internal.GetDataHandler)
	r.Run(":9109")

}
