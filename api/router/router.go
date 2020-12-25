package router

import (
	"github.com/bqxtt/book_online/api/handler"
	"github.com/gin-gonic/gin"
	"log"
)

func Init() {
	router := gin.Default()
	router.POST("/register", handler.Register)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("router run error: %v", err)
	}
}
