package main

import (
	"GinCasbin/internal"
	"GinCasbin/middleware"
	"time"

	"github.com/gin-gonic/gin"
	selfLogger "github.com/thomascwei/golang_logger"
)

var (
	Logger = selfLogger.InitLogger("main")
)

func main() {
	Logger.Info("starting...")
	r := gin.Default()
	// 限流
	r.Use(middleware.RateLimitMiddleware(time.Second, 100, 100))

	r.GET("/", internal.HelloWorldHandler)

	users := r.Group("/v1/users")
	users.POST("/login", internal.LoginHandler)

	data := r.Group("/v1/data")
	data.GET("/:id", internal.JWTAuthMiddleware(), middleware.RBACAuthorizeMiddleware("data", "GET"), internal.GetDataHandler)
	data.POST("/", internal.JWTAuthMiddleware(), middleware.RBACAuthorizeMiddleware("data", "POST"), internal.POSTDataHandler)

	data.GET("/ABAC/:id", internal.JWTAuthMiddleware(), middleware.ABACAuthorizeMiddleware("data", "GET"), internal.GetDataHandler)
	data.POST("/ABAC/", internal.JWTAuthMiddleware(), middleware.ABACAuthorizeMiddleware("data", "POST"), internal.POSTDataHandler)

	r.Run(":9109")

}
