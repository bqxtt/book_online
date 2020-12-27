package router

import (
	"github.com/bqxtt/book_online/api/handler"
	"github.com/bqxtt/book_online/api/router/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func Init() {
	router := gin.Default()
	router.Use(middleware.Cors())
	router.POST("/register", handler.Register)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("router run error: %v", err)
	}
}
